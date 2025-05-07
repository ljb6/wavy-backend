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
	user, err := s.repo.GetTokenByEmail(email)
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

func (s *Service) ChangePassword(id, password, newPassword string) error {

	user, err := s.repo.GetUserDataByID(id)
	if err != nil {
		return err
	}

	// checking password
	if (!security.CheckPassword(user.Password, password)) {
		return err
	}

	newHashedPassword, err := security.HashPassword(newPassword)
	if err != nil {
		return err
	}
	
	err = s.repo.ChangePassword(id, newHashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateUserSettings(req UserSettings) error {

	encryptedKey, err := security.Encrypt(req.SMTP_KEY)
	if err != nil {
		return err
	}

	req.SMTP_KEY = encryptedKey

	err = s.repo.CreateUserSettings(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUserSettings(userID string) (*UserSettings, error) {
	settings, err := s.repo.GetUserSettings(userID)
	if err != nil {
		return nil, err
	}
	return settings, nil
}