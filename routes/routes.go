package routes

import (
	"github.com/gorilla/mux"
	cards_controllers "github.com/italodavidb/goCrud/controllers"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/Cards", cards_controllers.CreateCard).Methods("POST")
	router.HandleFunc("/api/Cards", cards_controllers.FindAllCards).Methods("GET")
	router.HandleFunc("/api/Cards/search", cards_controllers.FindCards).Methods("GET")
	router.HandleFunc("/api/Cards", cards_controllers.DeleteCards).Methods("DELETE")
	return router
}
