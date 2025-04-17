package models

import "time"

type User struct {
	ID        int       `jsond:"id"`
	Name      string    `jsond:"name"`
	Email     string    `jsond:"email"`
	Password  string    `jsond:"-"`
	Plan      string    `jsond:"plan"`
	CreatedAt time.Time `json:"created_at"`
}
