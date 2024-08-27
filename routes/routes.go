package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func LoadRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", nil).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
