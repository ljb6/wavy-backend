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

type UserSettings struct {
	ID       string `jsond:"id"`
	User_ID  string `json:"user_id"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	SMTP_KEY string `jsond:"smtp_key"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
