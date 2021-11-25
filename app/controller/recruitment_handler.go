package controller

import (
	"forgottennw/app/model"
	"html/template"
	"net/http"
	"time"
)

func RecruitmentGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/recruitment.html", "template/character_form.html", head, navigation, footer))
	tmpl.Execute(w, nil)
}

func RecruitmentPOST(w http.ResponseWriter, r *http.Request) {
	attributes := model.Attributes{}
	attributes.Strength = ParseInt(r.FormValue("strength"))
	attributes.Dexterity = ParseInt(r.FormValue("dexterity"))
	attributes.Intelligence = ParseInt(r.FormValue("intelligence"))
	attributes.Focus = ParseInt(r.FormValue("focus"))
	attributes.Constitution = ParseInt(r.FormValue("constitution"))

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
	character := model.Character{}
	character.Name = r.FormValue("charactername")
	character.Gearscore = ParseInt(r.FormValue("gearscore"))
	character.Attributes = attributes
	character.Roles = roles
	character.Weapons = weapons
	character.Old_Companies = r.FormValue("old-companies")

	date := model.Date{}
	date_Now := time.Now()
	date.Date = date_Now.Local()
	date.Second = date_Now.Second()
	date.Minute = date_Now.Minute()
	date.Hour = date_Now.Hour()
	date.Day = date_Now.Day()
	date.Month = int(date_Now.Month())
	date.Year = date_Now.Year()

	recruitment_entry := model.Recruitment_Entry{}
	recruitment_entry.Character = character
	recruitment_entry.Status = "pending"
	recruitment_entry.Date = date

	err := model.AddRecruitment_Entry(recruitment_entry)
	if err != nil {
		http.Redirect(w, r, "/you_fcked_up", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
