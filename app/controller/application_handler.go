package controller

import (
	"encoding/base64"
	"forgottennw/app/model"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func ApplicationFormGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/applicationForm.html", head, navigation, footer))
	weapons, _ := model.GetAllWeapons()
	roles, _ := model.GetAllRoles()
	data := Data{
		Session: GetSessionInformation(r),
		Weapons: weapons,
		Roles:   roles,
	}
	tmpl.Execute(w, data)
}

func ApplicationFormPOST(w http.ResponseWriter, r *http.Request) {
	// //reCaptcha start
	// rec, errRecaptcha := recaptcha.New(model.RECAPTCHA_SECRET)
	// if errRecaptcha != nil {
	// 	panic(errRecaptcha)
	// }
	// response, errResponse := rec.GetRecaptchaToken(r)
	// if errResponse != nil {
	// 	panic(errResponse)
	// }

	// if errResponse = rec.Verify(response); errResponse != nil {
	// 	panic(errResponse)
	// }
	// //reCaptcha end
	attributes := model.Attributes{
		Strength:     ParseInt(r.FormValue("strength")),
		Dexterity:    ParseInt(r.FormValue("dexterity")),
		Intelligence: ParseInt(r.FormValue("intelligence")),
		Focus:        ParseInt(r.FormValue("focus")),
		Constitution: ParseInt(r.FormValue("constitution")),
	}

	craftingJobs_Form := r.Form["crafting-jobs"]
	craftingJobs := model.CraftingJobs{}
	for _, craftingJob := range craftingJobs_Form {
		switch craftingJob {
		case "armoring":
			craftingJobs.Armoring = true
		case "weaponsmithing":
			craftingJobs.Weaponsmithing = true
		case "engineering":
			craftingJobs.Engineering = true
		case "jewelcrafting":
			craftingJobs.Jewelcrafting = true
		case "arcana":
			craftingJobs.Arcana = true
		case "cooking":
			craftingJobs.Cooking = true
		case "furnishing":
			craftingJobs.Furnishing = true
		}
	}

	refiningJobs_Form := r.Form["refining-jobs"]
	refiningJobs := model.RefiningJobs{}
	for _, refiningJob := range refiningJobs_Form {
		switch refiningJob {
		case "smelting":
			refiningJobs.Smelting = true
		case "stonecutting":
			refiningJobs.Stonecutting = true
		case "tanning":
			refiningJobs.Tanning = true
		case "weaving":
			refiningJobs.Weaving = true
		case "woodworking":
			refiningJobs.Woodworking = true
		}
	}

	character := model.Character{
		Name:             r.FormValue("charactername"),
		Gearscore:        ParseInt(r.FormValue("gearscore")),
		Equipment_weight: r.FormValue("armor-weight"),
		Attributes:       attributes,
		Roles:            r.Form["roles"],
		Weapons:          r.Form["weapons"],
		Old_company:      r.FormValue("old-companies"),
		Crafting_jobs:    craftingJobs,
		Refining_jobs:    refiningJobs,
	}

	application := model.Application{
		Character:   character,
		Status:      "offen",
		Date:        GetCurrentDate(),
		Discord_tag: r.FormValue("discord-tag"),
	}

	err := model.AddApplication(application)
	if err != nil {
		http.Redirect(w, r, r.Header.Get("Referer")+"?application=failed", http.StatusFound)
	} else {
		http.Redirect(w, r, "/?application=success", http.StatusFound)
	}
}

func ApplicationsGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("applications.html").
		Funcs(template.FuncMap{"add": Add}).
		Funcs(template.FuncMap{"getRoleName": GetRoleName}).
		ParseFiles("template/applications.html", head, navigation, footer))
	applications, _ := model.GetAllApplicationsOpen()
	data := Data{
		Session:      GetSessionInformation(r),
		Applications: applications,
	}
	tmpl.Execute(w, data)
}

func ApplicationGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("application.html").
		Funcs(template.FuncMap{"getWeaponName": GetWeaponName}).
		Funcs(template.FuncMap{"getWeaponByType": GetWeaponByType}).
		Funcs(template.FuncMap{"getRoleName": GetRoleName}).
		ParseFiles("template/application.html", attributes, jobs, roles, weapons, head, navigation, footer))
	applicationId := mux.Vars(r)["id"]
	application, _ := model.GetApplicationById(applicationId)
	weapons, _ := model.GetAllWeapons()
	roles, _ := model.GetAllRoles()
	data := Data{
		Session:     GetSessionInformation(r),
		Application: application,
		Weapons:     weapons,
		Roles:       roles,
	}
	tmpl.Execute(w, data)
}

func ApplicationAcceptedPOST(w http.ResponseWriter, r *http.Request) {
	application, _ := model.GetApplicationById(mux.Vars(r)["id"])
	application.Status = "angenommen"
	company := r.FormValue("company")
	permissionLevel := r.FormValue("permission-level")
	//create temporary password
	password := strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	_, week := time.Now().ISOWeek()
	taxes := model.Taxes{}
	taxes.Previous_week.Week = week - 1
	taxes.Current_week.Week = week
	taxes.Next_week.Week = week + 1
	taxes.Second_next_week.Week = week + 2

	user := model.User{
		Name:             strings.ToLower(strings.ReplaceAll(application.Character.Name, " ", "")),
		Discord_tag:      application.Discord_tag,
		Password_tmp:     password,
		Password:         b64HashedPwd,
		Company:          company,
		Permission_level: ParseInt(permissionLevel),
		Character:        application.Character,
		Date:             GetCurrentDate(),
		Taxes:            taxes,
	}
	model.UpdateApplication(application)
	err := model.AddUser(user)
	if err != nil {
		http.Redirect(w, r, GetPreviousRoute(r), http.StatusFound)
	} else {
		http.Redirect(w, r, "/applications", http.StatusFound)
	}
}

func ApplicationRejectedPOST(w http.ResponseWriter, r *http.Request) {
	application, _ := model.GetApplicationById(mux.Vars(r)["id"])
	application.Status = "abgelehnt"
	model.UpdateApplication(application)
	http.Redirect(w, r, "/applications", http.StatusFound)
}
