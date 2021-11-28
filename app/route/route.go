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
	router.HandleFunc("/recruitment", controller.RecruitmentFormPOST).Methods("POST") //validierung der text inputs verbessern (regex) -> attribute & gearscore = zahlen; name = Buchstaben&Zahlen&spaces

	router.HandleFunc("/recruitments", controller.Authenticate(controller.RecruitmentsGET)).Methods("GET")
	router.HandleFunc("/recruitments/{id}", controller.Authenticate(controller.RecruitmentEntryGET)).Methods("GET")
	router.HandleFunc("/recruitments/{id}/accepted", controller.Authenticate(controller.RecruitmentEntryAcceptedPOST)).Methods("POST")
	router.HandleFunc("/recruitments/{id}/rejected", controller.Authenticate(controller.RecruitmentEntryRejectedPOST)).Methods("POST")

	router.HandleFunc("/members", controller.Authenticate(controller.UsersGET)).Methods("GET")
	router.HandleFunc("/members/{id}", controller.Authenticate(controller.UserGET)).Methods("GET")
	router.HandleFunc("/members/{id}/changePassword", controller.Authenticate(controller.ChangePasswordGET)).Methods("GET")
	router.HandleFunc("/members/{id}/changePassword", controller.Authenticate(controller.ChangePasswordPOST)).Methods("POST") // fehlermeldung bei fehleingabe anzeigen

	router.HandleFunc("/imprint", controller.ImprintGET).Methods("GET")
	router.HandleFunc("/data_privacy", controller.DataPrivacyGET).Methods("GET")
	router.HandleFunc("/about_us", controller.AboutGET).Methods("GET")

	router.HandleFunc("/login", controller.Login).Methods("POST") // Bei login Fehler <div class="show"> zu modal & offcanvas hinzufügen, fehlermeldung im modal hinzufügen
	router.HandleFunc("/logout", controller.Authenticate(controller.Logout)).Methods("GET")

	// router.HandleFunc("/register", controller.RegisterGET).Methods("GET")
	// router.HandleFunc("/register", controller.Authenticate(controller.RegisterPOST)).Methods("POST")
	// TODO: sort options /recruitments, /members
	return router
}
