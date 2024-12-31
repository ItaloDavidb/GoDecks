package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDb() {
	dsn := "host=localhost user=meu_usuario password=meu_password dbname=meu_database port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados:", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
}
