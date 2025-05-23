package subscribers

import (
	"encoding/json"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddSubscriber(req SubRequest, userID string) error {
	err := s.repo.AddSubscriber(req, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetSubscribers(userID string) ([]byte, error) {
	subs, err := s.repo.GetSubscribers(userID)
	if err != nil {
		return nil, err
	}

	jsonSubs, err := json.Marshal(subs)
	if err != nil {
		return nil, err
	}

	return jsonSubs, nil
}

func (s *Service) ClearSubscribers(userID string) error {
	err := s.repo.ClearSubscribersFromID(userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DownloadData(userID string) ([][]string, error) {
	subscribers, err := s.repo.GetSubscribers(userID)
	if err != nil {
		return nil, err
	}

	data := [][]string{
		{"name", "email"},
	}

	for _, sub := range subscribers {
		data = append(data, []string{sub.Name, sub.Email})
	}

	return data, nil
}