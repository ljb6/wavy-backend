package email

import (
	"fmt"
	"log"

	"github.com/ljb6/wavy-backend/internal/security"
	"github.com/ljb6/wavy-backend/internal/subscribers"
	"github.com/ljb6/wavy-backend/internal/user"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	userRepo *user.Repository
	subscriberRepo *subscribers.Repository
}

func NewEmailService(urepo *user.Repository, srepo *subscribers.Repository) *EmailService{
	return &EmailService{userRepo: urepo, subscriberRepo: srepo }
}

func (s *EmailService) SendEmail(req EmailReq, userID string) error {

	settings, err := s.userRepo.GetUserSettings(userID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("chegou")

	decrypted_key, err := security.Decrypt(settings.SMTP_KEY)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(settings.Host, settings.Port, settings.Username, decrypted_key)
	sd, err := d.Dial()
	if err != nil {
		panic(err)
	}

	recipients, err := s.subscriberRepo.GetSubscribers(userID)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	for _, r := range recipients {
		m.SetHeader("From", "luccajbecker@gmail.com")
		m.SetAddressHeader("To", r.Email, r.Name)
		m.SetHeader("Subject", "Newsletter #1")
		m.SetBody("text/html", fmt.Sprintf("Hello %s!", r.Name))

		if err := gomail.Send(sd, m); err != nil {
			log.Printf("Could not send email to %q: %v", r.Email, err)
		}
		m.Reset()
	}

	return nil
}
