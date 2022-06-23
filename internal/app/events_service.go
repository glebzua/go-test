package app

import (
	"github.com/test_server/internal/domain"
	"github.com/test_server/internal/infra/database"
)

type Service interface {
	FindAll() ([]domain.Event, error)
	FindOne(id uint64) (*domain.Event, error)
}

type service struct {
	repo *database.Repository
}

func NewService(r *database.Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]domain.Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id uint64) (*domain.Event, error) {
	return (*s.repo).FindOne(id)
}
