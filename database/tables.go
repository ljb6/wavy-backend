package database

import (
	"database/sql"
	"log"

	"github.com/ljb6/wavy-backend/models"
	"github.com/ljb6/wavy-backend/security"
)

func CreateTables(db *sql.DB) {
	createUsersTableqQuery := `
		CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		plan VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`

	_, err := db.Exec(createUsersTableqQuery)
	if err != nil {
		log.Fatal("Error while creating 'users' table:", err)
	}
}

func CreateUser(db *sql.DB, user models.User) (int, error) {
	query := "INSERT INTO users (name, email, password, plan) VALUES ($1, $2, $3, $4) RETURNING id"

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	var userID int
	err = db.QueryRow(query, user.Name, user.Email, hashedPassword, user.Plan).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}