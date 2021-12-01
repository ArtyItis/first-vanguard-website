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

	router.HandleFunc("/applications", controller.ApplicationsGET).Methods("GET")
	router.HandleFunc("/applications/new", controller.ApplicationFormGET).Methods("GET")
	router.HandleFunc("/applications/new", controller.ApplicationFormPOST).Methods("POST")
	router.HandleFunc("/applications/{id}", controller.Authenticate(controller.ApplicationEntryGET)).Methods("GET")
	router.HandleFunc("/applications/{id}/accepted", controller.Authenticate(controller.ApplicationEntryAcceptedPOST)).Methods("POST")
	router.HandleFunc("/applications/{id}/rejected", controller.Authenticate(controller.ApplicationEntryRejectedPOST)).Methods("POST")

	router.HandleFunc("/members", controller.Authenticate(controller.UsersGET)).Methods("GET")
	router.HandleFunc("/members/{id}", controller.Authenticate(controller.UserGET)).Methods("GET")
	router.HandleFunc("/members/{id}/changePassword", controller.Authenticate(controller.ChangePasswordGET)).Methods("GET")
	router.HandleFunc("/members/{id}/changePassword", controller.Authenticate(controller.ChangePasswordPOST)).Methods("POST") // fehlermeldung bei fehleingabe anzeigen

	router.HandleFunc("/imprint", controller.ImprintGET).Methods("GET")
	router.HandleFunc("/data_privacy", controller.DataPrivacyGET).Methods("GET")
	router.HandleFunc("/about_us", controller.AboutGET).Methods("GET")

	router.HandleFunc("/login", controller.Login).Methods("POST") // Bei login Fehler <div class="show"> zu modal & offcanvas hinzufügen, fehlermeldung im modal hinzufügen
	router.HandleFunc("/logout", controller.Authenticate(controller.Logout)).Methods("GET")

	router.HandleFunc("/register", controller.RegisterGET).Methods("GET")
	router.HandleFunc("/register", controller.RegisterPOST).Methods("POST")

	return router
}
