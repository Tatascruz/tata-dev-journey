package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	DB, err = gorm.Open(sqlite.Open("produtos.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}
	log.Println("Banco de dados conectado com sucesso!")
}
