package controller

import (
	"forgottennw/app/model"
	"html/template"
	"net/http"
)

func RecruitmentGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/recruitment.html", "template/character_form.html", head, navigation, footer))
	data := Data{
		Session: GetSessionInformation(r),
	}
	tmpl.Execute(w, data)
}

func RecruitmentPOST(w http.ResponseWriter, r *http.Request) {
	attributes := model.Attributes{
		Strength:     ParseInt(r.FormValue("strength")),
		Dexterity:    ParseInt(r.FormValue("strength")),
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
		Name:          r.FormValue("charactername"),
		Gearscore:     ParseInt(r.FormValue("gearscore")),
		Attributes:    attributes,
		Roles:         roles,
		Weapons:       weapons,
		Old_Companies: r.FormValue("old-companies"),
	}

	recruitment_entry := model.Recruitment_Entry{
		Character: character,
		Status:    "pending",
		Date:      GetCurrentDate(),
	}

	err := model.AddRecruitment_Entry(recruitment_entry)
	if err != nil {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func RecruitmentsGET(w http.ResponseWriter, r *http.Request) {

}
