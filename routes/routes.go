package routes

import (
	"github.com/gorilla/mux"
	cards_controllers "github.com/italodavidb/goCrud/controllers"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/Cards", cards_controllers.CreateCard).Methods("POST")
	router.HandleFunc("/api/Cards", cards_controllers.FindAllCards).Methods("GET")
	router.HandleFunc("/api/Cards/Search", cards_controllers.FindCards).Methods("GET")
	router.HandleFunc("/api/Cards", cards_controllers.DeleteCards).Methods("DELETE")
	router.HandleFunc("/api/Cards/Update", cards_controllers.UpdateCard).Methods("PUT")
	// router.HandleFunc("/api/Register", users_controllers.CreateUser).Methods("POST")
	// router.HandleFunc("/api/Login", users_controllers.Login).Methods("POST")
	return router
}
