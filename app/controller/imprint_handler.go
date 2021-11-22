package controller

import (
	"html/template"
	"net/http"
)

func ImprintGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/imprint.html", head, navigation, footer))
	tmpl.Execute(w, nil)
}
func ImprintPOST(w http.ResponseWriter, r *http.Request) {

}
