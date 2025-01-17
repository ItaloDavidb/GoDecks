package routes

import (
	"github.com/gorilla/mux"
	"github.com/italodavidb/goCrud/internal/controllers/cardsController"
	"github.com/italodavidb/goCrud/internal/controllers/loginController"
	"github.com/italodavidb/goCrud/internal/controllers/usersController"
	"github.com/italodavidb/goCrud/internal/middleware"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/Login", loginController.Login).Methods("POST")
	router.HandleFunc("/api/Users", usersController.CreateUser).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.JwtMiddleware)
	api.HandleFunc("/Cards", cardsController.CreateCard).Methods("POST")
	api.HandleFunc("/Cards", cardsController.FindAllCards).Methods("GET")
	api.HandleFunc("/Cards/Search", cardsController.FindCards).Methods("GET")
	api.HandleFunc("/Cards", cardsController.DeleteCards).Methods("DELETE")
	api.HandleFunc("/Cards/Update", cardsController.UpdateCard).Methods("PUT")
	api.HandleFunc("/Users", usersController.FindAllUsers).Methods("GET")
	api.HandleFunc("/Users/Search", usersController.FindUser).Methods("GET")
	api.HandleFunc("/Users", usersController.DeleteUser).Methods("DELETE")
	api.HandleFunc("/Users/Update", usersController.UpdateUser).Methods("PUT")
	api.HandleFunc("/Auth", loginController.Auth).Methods("POST")

	return router
}
