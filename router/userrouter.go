package router

import (
	"USER_TEST/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUserRouter(router *mux.Router) *mux.Router {
	myControllers := &controller.UserController{}

	router.Handle("/user/create", http.HandlerFunc(myControllers.Create)).Methods("POST")
	router.Handle("/user/update", http.HandlerFunc(myControllers.Update)).Methods("POST")
	router.Handle("/user/getbyname", http.HandlerFunc(myControllers.GetAllUserData)).Methods("POST")
	router.Handle("/user/getbyaddress", http.HandlerFunc(myControllers.Delete)).Methods("POST")

	return router
}
