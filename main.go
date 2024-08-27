package main

import (
	"fmt"

	"github.com/italodavidb/goCrud/database"
	"github.com/italodavidb/goCrud/routes"
)

func main() {
	fmt.Println("Server on")
	database.ConnectToDb()
	routes.LoadRoutes()
}
