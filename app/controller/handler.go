package controller

import (
	"forgottennw/app/model"
	"html/template"
	"net/http"
)

func ExecuteTemplate(w http.ResponseWriter, r *http.Request, fileName string) {
	tmpl := template.Must(template.ParseFiles("template/"+fileName+".html", head, navigation, footer))
	data := Data{
		Session: GetSessionInformation(r),
	}
	tmpl.Execute(w, data)
}

func IndexGET(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(w, r, "index")
}

func ImprintGET(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(w, r, "imprint")
}

func BuildsGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/builds.html", head, navigation, footer))
	weapons, _ := model.GetAllWeapons()
	roles, _ := model.GetAllRoles()
	data := Data{
		Session: GetSessionInformation(r),
		Weapons: weapons,
		Roles:   roles,
	}
	tmpl.Execute(w, data)
	// ExecuteTemplate(w, r, "builds")
}
