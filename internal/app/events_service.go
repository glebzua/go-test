package app

import (
	"github.com/test_server/internal/domain"
	"github.com/test_server/internal/infra/database"
)

type EventsService interface {
	FindAll(page uint, pageSize uint) ([]domain.Events, error)
	FindUpcoming(page uint, pageSize uint) ([]domain.Events, error)
	FindOne(id uint64) (*domain.Events, error)
	Create(task *domain.Events) (*domain.Events, error)
}

type eventsService struct {
	repo *database.EventsRepository
}

func NewEventsService(r *database.EventsRepository) EventsService {
	return &eventsService{
		repo: r,
	}
}

func (s *eventsService) FindAll(page uint, pageSize uint) ([]domain.Events, error) {
	return (*s.repo).FindAll(page, pageSize)
}

func (s *eventsService) FindUpcoming(page uint, pageSize uint) ([]domain.Events, error) {
	return (*s.repo).FindUpcoming(page, pageSize, false)
}

func (s *eventsService) FindOne(id uint64) (*domain.Events, error) {
	return (*s.repo).FindOne(id, false)
}

func (s *eventsService) Create(event *domain.Events) (*domain.Events, error) {
	return (*s.repo).Create(event)
}
