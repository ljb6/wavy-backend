package email

import "github.com/ljb6/wavy-backend/internal/user"

type EmailService struct {
	userRepo *user.Repository
}

func NewEmailService(repo *user.Repository) *EmailService {
	return &EmailService{userRepo: repo}
}