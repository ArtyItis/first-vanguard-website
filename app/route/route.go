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

	router.HandleFunc("/recruitment", controller.RecruitmentFormGET).Methods("GET")
	router.HandleFunc("/recruitment", controller.RecruitmentFormPOST).Methods("POST")

	router.HandleFunc("/recruitments", controller.Authenticate(controller.RecruitmentsGET)).Methods("GET")
	router.HandleFunc("/recruitments/{id}", controller.Authenticate(controller.RecruitmentEntryGET)).Methods("GET")
	router.HandleFunc("/recruitments/{id}/accepted", controller.Authenticate(controller.RecruitmentEntryAcceptedPOST)).Methods("POST")
	router.HandleFunc("/recruitments/{id}/rejected", controller.Authenticate(controller.RecruitmentEntryRejectedPOST)).Methods("POST")

	router.HandleFunc("/members", controller.UsersGET).Methods("GET")
	router.HandleFunc("/members/{id}", controller.UserGET).Methods("GET")
	router.HandleFunc("/members/{id}/changePassword", controller.Authenticate(controller.ChangePasswordGET)).Methods("GET")
	router.HandleFunc("/members/{id}/changePassword", controller.Authenticate(controller.ChangePasswordPOST)).Methods("POST")

	router.HandleFunc("/imprint", controller.ImprintGET).Methods("GET")
	router.HandleFunc("/data_privacy", controller.DataPrivacyGET).Methods("GET")
	router.HandleFunc("/about_us", controller.AboutGET).Methods("GET")

	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	router.HandleFunc("/register", controller.RegisterGET).Methods("GET")
	router.HandleFunc("/register", controller.RegisterPOST).Methods("POST")
	// TODO:  /members, /members/{id} -> change password
	// TODO: sort options /recruitments, /members
	return router
}
