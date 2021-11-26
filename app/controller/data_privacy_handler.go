package controller

import (
	"html/template"
	"net/http"
)

func DataPrivacyGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/data_privacy.html", head, navigation, footer))
	tmpl.Execute(w, nil)
}
func DataPrivacyPOST(w http.ResponseWriter, r *http.Request) {

}
