package domain

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Service interface {
	Find(id string) (map[string]types.AttributeValue, error)
}

type Repository interface {
	Find(key string) (map[string]types.AttributeValue, error)
}

type service struct {
	repo Repository
}

func NewPointServiceRead(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Find(id string) (map[string]types.AttributeValue, error) {
	return s.repo.Find(id)
}
