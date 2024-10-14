package main

import (
	"USER_TEST/router"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/goinggo/tracelog"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func readServerSettings() (string, string) {
	return "localhost", "8080"
}

func main() {
	tracelog.Start(tracelog.LevelTrace)

	myCore := cors.New(cors.Options{

		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"POST", "GET"},
	})

	// Call Starting

	log.Println("initialising Routes")

	var myRouter *mux.Router

	myRouter = router.InitRoutes()

	n := negroni.Classic()
	n.Use(myCore)
	n.UseHandler(myRouter)

	myListenerIPAddress, myListenerPort := readServerSettings()

	server := &http.Server{
		Addr:    myListenerIPAddress + ":" + myListenerPort,
		Handler: n,
	}
	log.Println("Listening On :", server.Addr)
	server.ListenAndServe()
	log.Println("Exit")

}
