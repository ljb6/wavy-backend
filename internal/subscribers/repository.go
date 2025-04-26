package subscribers

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) AddSubscriber(req SubRequest, userID string) error {
	query := `INSERT INTO subscribers (user_id, name, email) VALUES ($1, $2, $3)`
   	_, err := r.DB.Exec(query, userID, req.Name, req.Email)
    return err
}