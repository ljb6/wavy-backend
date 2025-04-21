package user

import (
	"errors"
	"strconv"

	"github.com/ljb6/wavy-backend/internal/security"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(user User) error {

    hashedPassword, err := security.HashPassword(user.Password)
    if err != nil {
        return err
    }

    user.Password = hashedPassword
    return s.repo.Create(user)
}

func (s *Service) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !security.CheckPassword(user.Password, password) {
		return "", errors.New("incorrect password")
	}

	token, err := security.GenerateJWT(strconv.Itoa(user.ID))
	if err != nil {
		return "", errors.New("errow while generating JWT token")
	}

	return token, nil
}