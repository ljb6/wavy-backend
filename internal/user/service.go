package user

import (
	"errors"

	"github.com/ljb6/wavy-backend/security"
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

func (s *Service) Login(email, password string) (*User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !security.CheckPassword(user.Password, password) {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}