package controller

import (
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

func DataPrivacyGET(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(w, r, "data_privacy")
}

func NewsGET(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(w, r, "news")
}

func AboutGET(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(w, r, "about")
}
