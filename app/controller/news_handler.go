package controller

import (
	"html/template"
	"net/http"
)

func NewsGET(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/news.html", head, navigation, footer))
	tmpl.Execute(w, nil)
}
func NewsPOST(w http.ResponseWriter, r *http.Request) {

}
