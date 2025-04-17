package main

import (
	"github.com/ljb6/wavy-backend/database"
	router "github.com/ljb6/wavy-backend/routes"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()
	
	router.InitializeServer()
}