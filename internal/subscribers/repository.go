package subscribers

import (
	"database/sql"
)

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

func (r *Repository) GetSubscribers(userID string) ([]Subscriber, error) {
	query := `SELECT id, name, email FROM subscribers WHERE user_id = $1`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscribers []Subscriber

	for rows.Next() {
		var sub Subscriber
		err := rows.Scan(&sub.ID, &sub.Name, &sub.Email)
		if err != nil {
			return nil, err
		}
		subscribers = append(subscribers, sub)
	}

	return subscribers, nil
}

func (r *Repository) ClearSubscribersFromID(userID string) error {
	query := `DELETE FROM subscribers WHERE user_id = $1`
   	_, err := r.DB.Exec(query, userID)
    return err
}