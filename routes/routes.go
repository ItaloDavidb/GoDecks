package routes

import (
	"github.com/gorilla/mux"
	cardsController "github.com/italodavidb/goCrud/cardsController"
	"github.com/italodavidb/goCrud/usersController"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/Cards", cardsController.CreateCard).Methods("POST")
	router.HandleFunc("/api/Cards", cardsController.FindAllCards).Methods("GET")
	router.HandleFunc("/api/Cards/Search", cardsController.FindCards).Methods("GET")
	router.HandleFunc("/api/Cards", cardsController.DeleteCards).Methods("DELETE")
	router.HandleFunc("/api/Cards/Update", cardsController.UpdateCard).Methods("PUT")
	router.HandleFunc("/api/Users", usersController.CreateUser).Methods("POST")
	router.HandleFunc("/api/Users", usersController.FindAllUsers).Methods("GET")
	router.HandleFunc("/api/Users/Search", usersController.FindUser).Methods("GET")
	// router.HandleFunc("/api/Register", users_controllers.CreateUser).Methods("POST")
	// router.HandleFunc("/api/Login", users_controllers.Login).Methods("POST")
	return router
}
