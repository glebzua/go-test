package database

import (
	"fmt"
	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
	"log"
)

type events struct {
	Id               uint    `db:"id,omitempty"`
	Title            string  `db:"Title"`
	ShortDescription string  `db:"Short Description"`
	Description      string  `db:"Description"`
	Longitude        float64 `db:"Longitude"`
	Latitude         float64 `db:"Latitude"`
	Images           string  `db:"Images"`
	Preview          string  `db:"Preview"`
	Date             string  `db:"Date"`
}

type EventsRepository interface {
	FindAll(page uint, pageSize uint, showDeleted bool) ([]domain.Events, error)
	FindOne(id uint64, showDeleted bool) (*domain.Events, error)
}

const EventsCount uint64 = 8

type eventsRepository struct {
	coll db.Collection
}

func NewEventsRepository(dbSession *db.Session) EventsRepository {
	return &eventsRepository{
		coll: (*dbSession).Collection("events"),
	}
}

func (r *eventsRepository) FindAll(page uint, pageSize uint, showDeleted bool) ([]domain.Events, error) {
	var clt []events
	err := r.coll.Find(softDelCond(nil, showDeleted)).Paginate(pageSize).Page(page).All(&clt)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapEventsDbModelToDomainCollection(clt), nil
}

func (r *eventsRepository) FindOne(id uint64, showDeleted bool) (*domain.Events, error) {
	var ctl events
	err := r.coll.Find(softDelCond(db.Cond{"id": id}, showDeleted)).One(&ctl)
	if err != nil {
		return nil, fmt.Errorf("repository FindOneCattle: %w", err)
	}

	return mapEventsDbModelToDomain(&ctl), nil
}

func mapEventsDbModelToDomainCollection(events []events) []domain.Events {
	var result []domain.Events

	for _, c := range events {
		newCtl := mapEventsDbModelToDomain(&c)
		result = append(result, *newCtl)
	}
	return result
}

func mapEventsDbModelToDomain(events *events) *domain.Events {
	return &domain.Events{
		Id:               events.Id,
		Title:            events.Title,
		ShortDescription: events.ShortDescription,
		Description:      events.Description,
		Longitude:        events.Longitude,
		Latitude:         events.Latitude,
		Images:           events.Images,
		Preview:          events.Preview,
		Date:             events.Date,
	}

}
