package controller

import (
	"encoding/base64"
	"forgottennw/app/model"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func RecruitmentFormGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/recruitmentForm.html", head, navigation, footer))
	data := Data{
		Session: GetSessionInformation(r),
	}
	tmpl.Execute(w, data)
}

func RecruitmentFormPOST(w http.ResponseWriter, r *http.Request) {
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
	roles := model.Roles{}
	for _, role := range roles_Form {
		switch role {
		case "dd":
			roles.DD = true
		case "tank":
			roles.Tank = true
		case "heal":
			roles.Heal = true
		case "siege":
			roles.Siege = true
		}
	}

	weapons_Form := r.Form["weapons"]
	weapons := model.Weapons{}
	for _, weapon := range weapons_Form {
		switch weapon {
		case "sword":
			weapons.Sword = true
		case "shield":
			weapons.Shield = true
		case "rapier":
			weapons.Rapier = true
		case "hatchet":
			weapons.Hatchet = true
		case "spear":
			weapons.Spear = true
		case "great-axe":
			weapons.Great_Axe = true
		case "warhammer":
			weapons.Warhammer = true
		case "bow":
			weapons.Bow = true
		case "musket":
			weapons.Musket = true
		case "fire-staff":
			weapons.Firestaff = true
		case "life-staff":
			weapons.Lifestaff = true
		case "ice-gauntlet":
			weapons.Ice_Gauntlet = true
		case "void-gauntlet":
			weapons.Void_Gauntlet = true
		}
	}
	character := model.Character{
		Name:       r.FormValue("charactername"),
		Gearscore:  ParseInt(r.FormValue("gearscore")),
		Attributes: attributes,
		// Roles:       roles,
		// Weapons:     weapons,
		Old_Company: r.FormValue("old-companies"),
	}

	recruitmentEntry := model.RecruitmentEntry{
		Character: character,
		Status:    "offen",
		Date:      GetCurrentDate(),
	}

	err := model.AddRecruitmentEntry(recruitmentEntry)
	if err != nil {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func RecruitmentsGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("recruitments.html").Funcs(template.FuncMap{"add": Add}).ParseFiles("template/recruitments.html", head, navigation, footer))
	recruitmentEntries, _ := model.GetAllRecruitmentEntries()
	data := Data{
		Session:            GetSessionInformation(r),
		RecruitmentEntries: recruitmentEntries,
	}
	tmpl.Execute(w, data)
}

func RecruitmentEntryGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/recruitmentEntry.html", head, navigation, footer))
	recruitmentErtyId := mux.Vars(r)["id"]
	recruitmentEntriy, _ := model.GetRecruitmentEntryById(recruitmentErtyId)
	data := Data{
		Session:          GetSessionInformation(r),
		RecruitmentEntry: recruitmentEntriy,
	}
	tmpl.Execute(w, data)
}

func RecruitmentEntryAcceptedPOST(w http.ResponseWriter, r *http.Request) {
	recruitmentEntry, _ := model.GetRecruitmentEntryById(mux.Vars(r)["id"])
	recruitmentEntry.Status = "angenommen"
	company := r.FormValue("company")
	permissionLevel := r.FormValue("permission-level")
	//create temporary password
	password := strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	user := model.User{
		Name:             recruitmentEntry.Character.Name,
		Password_Tmp:     password,
		Password:         b64HashedPwd,
		Company:          company,
		Permission_Level: ParseInt(permissionLevel),
		Character:        recruitmentEntry.Character,
		Date:             GetCurrentDate(),
	}
	model.UpdateRecruitmentEntry(recruitmentEntry)
	err := model.AddUser(user)
	if err != nil {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		http.Redirect(w, r, "/recruitments", http.StatusFound)
	}
}

func RecruitmentEntryRejectedPOST(w http.ResponseWriter, r *http.Request) {
	recruitmentEntry, _ := model.GetRecruitmentEntryById(mux.Vars(r)["id"])
	recruitmentEntry.Status = "abgelehnt"
	model.UpdateRecruitmentEntry(recruitmentEntry)
	http.Redirect(w, r, "/recruitments", http.StatusFound)
}
