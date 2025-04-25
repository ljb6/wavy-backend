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

func (r *Repository) GetTokenByEmail(email string) (*User, error) {
	var user User
	row := r.DB.QueryRow(`SELECT id, name, email, password, plan, created_at FROM users WHERE email = $1`, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Plan, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserDataByID(id string) (*User, error) {
	var user User
	row := r.DB.QueryRow(`SELECT id, name, email, password, plan, created_at FROM users WHERE id = $1`, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Plan, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) ChangePassword(id, newPassword string) error {
	_, err := r.DB.Exec(`UPDATE users SET password = $1 WHERE id = $2`, newPassword, id)
	if err != nil {
		return err
	}
	return nil
}