package user

import (
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
