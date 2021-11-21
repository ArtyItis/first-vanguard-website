package main

import (
	"forgottennw/app/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//exclude static prefix from url
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r := mux.NewRouter()
	http.Handle("/", r)
	// Create a file server which serves files out of the "../static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("../static/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	r.Handle("/static/", http.StripPrefix("/static", fileServer))
	r.HandleFunc("/", controller.Index).Methods("GET")
	r.HandleFunc("/recruitment", controller.RecruitmentGET).Methods("GET")
	r.HandleFunc("/recruitment", controller.RecruitmentPOST).Methods("POST")
	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}
