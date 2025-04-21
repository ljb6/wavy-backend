package user

import "time"

type User struct {
	ID        int       `jsond:"id"`
	Name      string    `jsond:"name"`
	Email     string    `jsond:"email"`
	Password  string    `jsond:"-"`
	Plan      string    `json:"plan"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
