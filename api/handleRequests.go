package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// handle api requests
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/print", homePage)
	router.HandleFunc("/print/status")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Print API")
	fmt.Println("Endpoint Hit: homePage")
}
