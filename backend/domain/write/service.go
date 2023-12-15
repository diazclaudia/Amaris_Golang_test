package domain

import "backend/config"

type Service interface {
	Update(id, points string) (*string, error)
}

type Repository interface {
	Update(id, points, key, secret string) (*string, error)
}

type service struct {
	repo   Repository
	config *config.Config
}

func (s *service) Update(id, points string) (*string, error) {
	return s.repo.Update(id, points, s.config.Aws.Key, s.config.Aws.Secret)
}

func NewPointService(repo Repository, config *config.Config) Service {
	return &service{
		repo:   repo,
		config: config,
	}
}
