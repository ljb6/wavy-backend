package database

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) {
	createUsersTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		plan VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`

	createUserSettingsTable := `
		CREATE TABLE IF NOT EXISTS user_settings (
		id SERIAL PRIMARY KEY,
		user_id VARCHAR(100) NOT NULL,
		host VARCHAR(100) NOT NULL,
		port int NOT NULL,
		username VARCHAR(100) NOT NULL,
		smtp_key TEXT
		)`

	createSubscribersTableQuery := `
		CREATE TABLE IF NOT EXISTS subscribers (
		id SERIAL PRIMARY KEY, 
		user_id VARCHAR(100) NOT NULL,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`

	_, err := db.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal("Error while creating 'users' table:", err)
	}

	_, err = db.Exec(createSubscribersTableQuery)
	if err != nil {
		log.Fatal("Error while creating 'subscribers' table:", err)
	}

	_, err = db.Exec(createUserSettingsTable)
	if err != nil {
		log.Fatal("Error while creating 'user_settings' table:", err)
	}
}