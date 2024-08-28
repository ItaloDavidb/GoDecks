package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/italodavidb/goCrud/database"
	"github.com/italodavidb/goCrud/routes"
)

func main() {
	fmt.Println("Server on")
	database.ConnectToDb()
	router := routes.LoadRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
