package controller

import (
	"forgottennw/app/model"
	"html/template"
	"net/http"
)

func RolesGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("roles.html").
		Funcs(template.FuncMap{"modulo": Modulo}).
		ParseFiles("template/roles.html", head, navigation, footer))
	roles, _ := model.GetAllRoles()
	data := Data{
		Session: GetSessionInformation(r),
		Roles:   roles,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RoleGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/role.html", head, navigation, footer))
	data := Data{
		Session: GetSessionInformation(r),
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RoleEditGET(w http.ResponseWriter, r *http.Request) {

}

func RoleEditPOST(w http.ResponseWriter, r *http.Request) {

}

func RoleFormGET(w http.ResponseWriter, r *http.Request) {

}

func RoleFormPOST(w http.ResponseWriter, r *http.Request) {

}
