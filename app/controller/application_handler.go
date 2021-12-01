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
	roles_Form := r.Form["roles"]
	roles := []model.Role{}
	for _, r := range roles_Form {
		role, err := model.GetRoleById(r)
		if err == nil {
			roles = append(roles, role)
		}
	}

	weapons_Form := r.Form["weapons"]
	weapons := []model.Weapon{}
	for _, w := range weapons_Form {
		weapon, err := model.GetWeaponById(w)
		if err == nil {
			weapons = append(weapons, weapon)
		}
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
		Roles:            roles,
		Weapons:          weapons,
		Old_company:      r.FormValue("old-companies"),
		Crafting_jobs:    craftingJobs,
		Refining_jobs:    refiningJobs,
	}

	applicationEntry := model.ApplicationEntry{
		Character:   character,
		Status:      "offen",
		Date:        GetCurrentDate(),
		Discord_tag: r.FormValue("discord-tag"),
	}

	err := model.AddApplicationEntry(applicationEntry)
	if err != nil {
		http.Redirect(w, r, r.Header.Get("Referer")+"?application=failed", http.StatusFound)
	} else {
		http.Redirect(w, r, "/?application=success", http.StatusFound)
	}
}

func ApplicationsGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("applications.html").Funcs(template.FuncMap{"add": Add}).ParseFiles("template/applications.html", head, navigation, footer))
	applicationEntries, _ := model.GetAllApplicationEntriesOpen()
	data := Data{
		Session:            GetSessionInformation(r),
		ApplicationEntries: applicationEntries,
	}
	tmpl.Execute(w, data)
}

func ApplicationEntryGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("applicationEntry.html").
		Funcs(template.FuncMap{"containsWeapon": ContainsWeapon}).
		Funcs(template.FuncMap{"containsRole": ContainsRole}).
		ParseFiles("template/applicationEntry.html", head, navigation, footer))
	applicationEntryId := mux.Vars(r)["id"]
	applicationEntry, _ := model.GetApplicationEntryById(applicationEntryId)
	weapons, _ := model.GetAllWeapons()
	roles, _ := model.GetAllRoles()
	data := Data{
		Session:          GetSessionInformation(r),
		ApplicationEntry: applicationEntry,
		Weapons:          weapons,
		Roles:            roles,
	}
	tmpl.Execute(w, data)
}

func ApplicationEntryAcceptedPOST(w http.ResponseWriter, r *http.Request) {
	applicationEntry, _ := model.GetApplicationEntryById(mux.Vars(r)["id"])
	applicationEntry.Status = "angenommen"
	company := r.FormValue("company")
	permissionLevel := r.FormValue("permission-level")
	//create temporary password
	password := strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	user := model.User{
		Name:             strings.ToLower(strings.ReplaceAll(applicationEntry.Character.Name, " ", "")),
		Discord_tag:      applicationEntry.Discord_tag,
		Password_tmp:     password,
		Password:         b64HashedPwd,
		Company:          company,
		Permission_level: ParseInt(permissionLevel),
		Character:        applicationEntry.Character,
		Date:             GetCurrentDate(),
	}
	model.UpdateApplicationEntry(applicationEntry)
	err := model.AddUser(user)
	if err != nil {
		http.Redirect(w, r, GetPreviousRoute(r), http.StatusFound)
	} else {
		http.Redirect(w, r, "/applications", http.StatusFound)
	}
}

func ApplicationEntryRejectedPOST(w http.ResponseWriter, r *http.Request) {
	applicationEntry, _ := model.GetApplicationEntryById(mux.Vars(r)["id"])
	applicationEntry.Status = "abgelehnt"
	model.UpdateApplicationEntry(applicationEntry)
	http.Redirect(w, r, "/applications", http.StatusFound)
}
