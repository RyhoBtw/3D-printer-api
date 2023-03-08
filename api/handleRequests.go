package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RyhoBtw/3D-printer-api/api/printer"
	"github.com/RyhoBtw/3D-printer-api/api/db"
	"github.com/gorilla/mux"
)

// handle api requests
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/print", HomePage)
	router.HandleFunc("/print/status", printer.GetStatus)
	router.HandleFunc("/print/login{info}", db.login)
	router.HandleFunc("/print/Gcode", printer.PostGcode).Methods("POST")
	log.Log().Fatal(http.ListenAndServe(":8000", router))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Print API")
	fmt.Println("Endpoint Hit: homePage")
}
