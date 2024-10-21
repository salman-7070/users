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
	router.Handle("/user/getalldata", http.HandlerFunc(myControllers.GetAllUserData)).Methods("GET")
	router.Handle("/user/delete", http.HandlerFunc(myControllers.Delete)).Methods("POST")
	router.Handle("/user/login", http.HandlerFunc(myControllers.Login)).Methods("POST")

	return router
}
