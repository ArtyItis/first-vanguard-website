package controller

import (
	"html/template"
	"net/http"
)

func AboutGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/about_us.html", head, navigation, footer))
	tmpl.Execute(w, nil)
}
func AboutPOST(w http.ResponseWriter, r *http.Request) {

}
