package main

import (
	"log"
	"net/http"

	"github.com/italodavidb/goCrud/internal/database"
	"github.com/italodavidb/goCrud/internal/routes"
)

func main() {
	database.ConnectToDb()
	router := routes.LoadRoutes()

	log.Println("Iniciando o servidor na porta 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
