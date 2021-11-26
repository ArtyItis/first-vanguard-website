package route

import (
	"forgottennw/app/controller"
	"net/http"

	"github.com/gorilla/mux" //go get -u -v github.com/gorilla/mux
)

func NewRouter() *mux.Router {
	staticDir := "/static/"
	router := mux.NewRouter()

	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	router.HandleFunc("/", controller.IndexGET).Methods("GET")
	router.HandleFunc("/recruitment", controller.RecruitmentGET).Methods("GET")
	router.HandleFunc("/recruitment", controller.RecruitmentPOST).Methods("POST")

	router.HandleFunc("/news", controller.NewsGET).Methods("GET")

	router.HandleFunc("/imprint", controller.ImprintGET).Methods("GET")
	router.HandleFunc("/data_privacy", controller.DataPrivacyGET).Methods("GET")
	router.HandleFunc("/about_us", controller.AboutGET).Methods("GET")

	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	router.HandleFunc("/register", controller.RegisterGET).Methods("GET")
	router.HandleFunc("/register", controller.RegisterPOST).Methods("POST")
	return router
}
