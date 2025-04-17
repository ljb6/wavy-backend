package models

type User struct {
	ID       int    `jsond:"id"`
	Name     string `jsond:"name"`
	Email    string `jsond:"email"`
	Password string `jsond:"-"`
	Plan     string `jsond:"plan"`
}
