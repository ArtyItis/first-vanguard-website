package route

import (
	"first-vanguard/app/controller"
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

	router.HandleFunc("/companies/{company}/taxes", controller.Authenticate(controller.TaxesGET)).Methods("GET")
	router.HandleFunc("/companies/{company}/taxes", controller.Authenticate(controller.TaxesPOST)).Methods("POST")
	router.HandleFunc("/companies/{company}/members", controller.Authenticate(controller.UsersGET)).Methods("GET")
	router.HandleFunc("/companies/{company}/members/{id}", controller.Authenticate(controller.UserGET)).Methods("GET")
	router.HandleFunc("/companies/{company}/members/{id}/changePassword", controller.Authenticate(controller.ChangePasswordGET)).Methods("GET")
	router.HandleFunc("/companies/{company}/members/{id}/changePassword", controller.Authenticate(controller.ChangePasswordPOST)).Methods("POST")
	router.HandleFunc("/companies/{company}/members/{id}/delete", controller.Authenticate(controller.UserDeleteGET)).Methods("GET")
	router.HandleFunc("/companies/{company}/members/{id}/edit", controller.Authenticate(controller.UserEditGet)).Methods("GET")
	router.HandleFunc("/companies/{company}/members/{id}/edit", controller.Authenticate(controller.UserEditPost)).Methods("POST")

	router.HandleFunc("/roles", controller.Authenticate(controller.RolesGET)).Methods("GET")
	router.HandleFunc("/roles/new", controller.Authenticate(controller.RoleFormGET)).Methods("GET")
	router.HandleFunc("/roles/new", controller.Authenticate(controller.RoleFormPOST)).Methods("POST")
	router.HandleFunc("/roles/{id}", controller.Authenticate(controller.RoleGET)).Methods("GET")
	router.HandleFunc("/roles/{id}/edit", controller.Authenticate(controller.RoleEditGET)).Methods("GET")
	router.HandleFunc("/roles/{id}/edit", controller.Authenticate(controller.RoleEditPOST)).Methods("POST")

	router.HandleFunc("/refining", controller.Authenticate(controller.RefiningGET)).Methods("GET")
	router.HandleFunc("/crafting", controller.Authenticate(controller.CraftingGET)).Methods("GET")

	router.HandleFunc("/imprint", controller.ImprintGET).Methods("GET")

	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Authenticate(controller.Logout)).Methods("GET")

	// router.HandleFunc("/register", controller.RegisterGET).Methods("GET")
	// router.HandleFunc("/register", controller.RegisterPOST).Methods("POST")

	return router
}
