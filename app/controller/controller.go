package controller

import (
	"html/template"
	"net/http"
	"strconv"
)

var (
	tmpl       *template.Template
	head       = "template/head.html"
	navigation = "template/navigation.html"
	footer     = "template/footer.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html", head, navigation, footer))
	tmpl.Execute(w, nil)
}

func ParseInt(str string) (result int) {
	result, _ = strconv.Atoi(str)
	return result
}
