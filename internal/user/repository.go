package user

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(user User) error {
    query := `INSERT INTO users (name, email, password, plan) VALUES ($1, $2, $3, $4)`
   	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.Plan)
    return err
}
