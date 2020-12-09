package main

import (
	myhandlers "github.com/course_spec/data_control/autorization/handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"os"

	"log"
	"net/http"
)

const (
	connPort = "8080"
	connHost = "localhost"
)



func main(){
	router := mux.NewRouter()

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	router.Handle("/",handlers.LoggingHandler(logFile,myhandlers.LoginPageHandler)).Methods("GET")
	router.Handle("/home",handlers.LoggingHandler(logFile,myhandlers.HomePageHandler)).Methods("GET")
	router.Handle("/login",handlers.LoggingHandler(logFile,myhandlers.LoginHomePageHandler)).Methods("POST")
	router.Handle("/logout",handlers.LoggingHandler(logFile,http.HandlerFunc(myhandlers.LogoutFormPageHandler))).Methods("POST")

	//router.HandleFunc("/",myhandlers.LoginPageHandler)
	//router.HandleFunc("/home",myhandlers.HomePageHandler)
	//router.HandleFunc("/login",myhandlers.LoginHomePageHandler).Methods("POST")
	//router.HandleFunc("/logout",myhandlers.LogoutFormPageHandler).Methods("POST")	//router.HandleFunc("/",myhandlers.LoginPageHandler)
	//router.HandleFunc("/home",myhandlers.HomePageHandler)
	//router.HandleFunc("/login",myhandlers.LoginHomePageHandler).Methods("POST")
	//router.HandleFunc("/logout",myhandlers.LogoutFormPageHandler).Methods("POST")

	err = http.ListenAndServe(connHost+":"+connPort,router)
	if err != nil {
		log.Fatal("error starting server: ",err)
		return
	}

}






