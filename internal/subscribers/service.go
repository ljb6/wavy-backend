package subscribers

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