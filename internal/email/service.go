package email

import (
	"fmt"

	"github.com/ljb6/wavy-backend/internal/user"
)

type EmailService struct {
	userRepo *user.Repository
}

func NewEmailService(repo *user.Repository) *EmailService {
	return &EmailService{userRepo: repo}
}

func (s *EmailService) SendEmail(req EmailReq, ID string) error {

	fmt.Println(req.Subject, req.Body, ID)

	return nil
}
