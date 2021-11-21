package controller

import (
	"html/template"
	"net/http"
)

func RecruitmentGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/recruitment.html", head, navigation, footer))
	tmpl.Execute(w, nil)
}
func RecruitmentPOST(w http.ResponseWriter, r *http.Request) {

}
