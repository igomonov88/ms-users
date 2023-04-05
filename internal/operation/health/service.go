package health

import "context"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CheckHealth(ctx context.Context) error {
	return nil
}
