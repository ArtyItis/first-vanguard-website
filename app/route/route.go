package route

import (
	"forgottennw/app/controller"
	"net/http"

	"github.com/gorilla/mux" //go get -u -v github.com/gorilla/mux
)

func MapURLs() {
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
	r.HandleFunc("/imprint", controller.ImprintGET).Methods("GET")
	r.HandleFunc("/imprint", controller.ImprintPOST).Methods("POST")
	r.HandleFunc("/news", controller.NewsGET).Methods("GET")
	r.HandleFunc("/news", controller.NewsPOST).Methods("POST")
	r.HandleFunc("/about", controller.AboutGET).Methods("GET")
	r.HandleFunc("/about", controller.AboutPOST).Methods("POST")
}
