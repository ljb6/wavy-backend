package database

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) {
	createUsersTableqQuery := `
		CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		plan VARCHAR(100) NOT NULL
		)`

	_, err := db.Exec(createUsersTableqQuery)
	if err != nil {
		log.Fatal("Error while creating 'users' table:", err)
	}
}