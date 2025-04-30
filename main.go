package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ljb6/wavy-backend/database"
	"github.com/ljb6/wavy-backend/internal/security"
	router "github.com/ljb6/wavy-backend/routes"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Erro ao carregar .env")
    }

	security.InitEncryptionKey()

	db := database.ConnectDB()
	defer db.Close()

	database.CreateTables(db)

	router.InitializeServer(db)
}