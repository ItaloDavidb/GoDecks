package routes

import (
	"github.com/gorilla/mux"
	"github.com/italodavidb/goCrud/controllers"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/Cards", controllers.CreateCard).Methods("POST")
	router.HandleFunc("/api/Cards", controllers.FindAllCards).Methods("GET")
	router.HandleFunc("/api/Cards/search", controllers.FindCards).Methods("GET")
	return router
}
