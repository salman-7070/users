package router

import (
	"USER_TEST/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUserRouter(router *mux.Router) *mux.Router {
	myControllers := &controller.UserController{}

	router.Handle("/students/create", http.HandlerFunc(myControllers.Create)).Methods("POST")
	router.Handle("/students/update", http.HandlerFunc(myControllers.Update)).Methods("POST")
	router.Handle("/students/getbyname", http.HandlerFunc(myControllers.GetAllUserData)).Methods("POST")
	router.Handle("/students/getbyaddress", http.HandlerFunc(myControllers.Delete)).Methods("POST")

	return router
}
