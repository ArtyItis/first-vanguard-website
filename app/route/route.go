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
	router.HandleFunc("/applications/{id}", controller.Authenticate(controller.ApplicationGET)).Methods("GET")
	router.HandleFunc("/applications/{id}/accepted", controller.Authenticate(controller.ApplicationAcceptedPOST)).Methods("POST")
	router.HandleFunc("/applications/{id}/rejected", controller.Authenticate(controller.ApplicationRejectedPOST)).Methods("POST")

	router.HandleFunc("/members/{company}", controller.Authenticate(controller.UsersGET)).Methods("GET")
	router.HandleFunc("/members/{company}/taxes", controller.Authenticate(controller.TaxesGET)).Methods("GET")
	router.HandleFunc("/members/{company}/taxes", controller.Authenticate(controller.TaxesPOST)).Methods("POST")
	router.HandleFunc("/members/{id}", controller.Authenticate(controller.UserGET)).Methods("GET")
	router.HandleFunc("/members/{id}/changePassword", controller.Authenticate(controller.ChangePasswordGET)).Methods("GET")
	router.HandleFunc("/members/{id}/changePassword", controller.Authenticate(controller.ChangePasswordPOST)).Methods("POST")
	router.HandleFunc("/members/{id}/delete", controller.Authenticate(controller.UserDeleteGET)).Methods("GET")
	router.HandleFunc("/members/{id}/edit", controller.Authenticate(controller.UserEditGet)).Methods("GET")
	router.HandleFunc("/members/{id}/edit", controller.Authenticate(controller.UserEditPost)).Methods("POST")

	router.HandleFunc("/builds", controller.Authenticate(controller.BuildsGET)).Methods("GET")

	router.HandleFunc("/refining", controller.Authenticate(controller.RefiningGET)).Methods("GET")
	router.HandleFunc("/crafting", controller.Authenticate(controller.CraftingGET)).Methods("GET")

	router.HandleFunc("/imprint", controller.ImprintGET).Methods("GET")

	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Authenticate(controller.Logout)).Methods("GET")

	// router.HandleFunc("/register", controller.RegisterGET).Methods("GET")
	// router.HandleFunc("/register", controller.RegisterPOST).Methods("POST")

	return router
}
