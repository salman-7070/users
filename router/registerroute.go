package router

import (
	"log"

	"github.com/gorilla/mux"
)

type RouterManager struct{}

//this method is for registering depended routes

func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router = SetUserRouter(router)

	log.Println("Router Registered")
	return router

}
