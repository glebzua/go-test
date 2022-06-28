package database

import (
	"fmt"
	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
	"log"
	"time"
)

type events struct {
	Id               uint      `db:"id,omitempty"`
	Title            string    `db:"Title"`
	ShortDescription string    `db:"ShortDescription"`
	Description      string    `db:"Description"`
	Longitude        float64   `db:"Longitude"`
	Latitude         float64   `db:"Latitude"`
	Images           string    `db:"Images"`
	Preview          string    `db:"Preview"`
	Date             string    `db:"Date"`
	IsEnded          bool      `db:"isEnded"`
	DeletedDate      time.Time `db:"deletedDate,omitempty"`
}

type EventsRepository interface {
	FindAll(page uint, pageSize uint) ([]domain.Events, error)
	FindUpcoming(page uint, pageSize uint, showDeleted bool) ([]domain.Events, error)
	FindOne(id uint64, showDeleted bool) (*domain.Events, error)
	AddEvent(task *domain.Events) (*domain.Events, error)
}

type eventsRepository struct {
	coll db.Collection
}

func NewEventsRepository(dbSession *db.Session) EventsRepository {
	return &eventsRepository{
		coll: (*dbSession).Collection("events"),
	}
}

func (r *eventsRepository) FindAll(page uint, pageSize uint) ([]domain.Events, error) {
	var event []events
	err := r.coll.Find().Paginate(pageSize).Page(page).All(&event)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return mapEventsDbModelToDomainCollection(event), nil
}
func (r *eventsRepository) FindUpcoming(page uint, pageSize uint, showDeleted bool) ([]domain.Events, error) {
	var event []events
	err := r.coll.Find(isDeleted(db.Cond{"isEnded": false}, showDeleted)).Paginate(pageSize).Page(page).All(&event)
	if err != nil {

		log.Print(err)
		return nil, err
	}
	return mapAllEventsDbModelToDomainCollection(event), nil
}
func (r *eventsRepository) FindOne(id uint64, showDeleted bool) (*domain.Events, error) {
	var event events
	err := r.coll.Find(isDeleted(db.Cond{"id": id}, showDeleted)).One(&event)
	if err != nil {
		return nil, fmt.Errorf("repository FindOneEvent: %w", err)
	}

	return mapEventsDbModelToDomain(&event), nil
}
func (r *eventsRepository) AddEvent(event *domain.Events) (*domain.Events, error) {
	nEvent := mapDomainToEventDbModel(event)
	err := r.coll.InsertReturning(nEvent)
	if err != nil {
		log.Print("InsertReturning err", err)
		return nil, err
	}
	return mapEventsDbModelToDomain(nEvent), nil
}

func mapEventsDbModelToDomainCollection(events []events) []domain.Events {
	var result []domain.Events
	for _, e := range events {
		newEvent := mapEventsDbModelToDomain(&e)
		result = append(result, *newEvent)
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
		IsEnded:          events.IsEnded,
	}

}
func mapAllEventsDbModelToDomainCollection(events []events) []domain.Events {
	var result []domain.Events
	for _, e := range events {
		newEvent := mapAllEventsDbModelToDomain(&e)
		result = append(result, *newEvent)
	}
	return result
}

func mapAllEventsDbModelToDomain(events *events) *domain.Events {
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
		IsEnded:          events.IsEnded,
		DeletedDate:      events.DeletedDate,
	}

}

func mapDomainToEventDbModel(nEvent *domain.Events) *events {
	return &events{
		Id:               nEvent.Id,
		Title:            nEvent.Title,
		ShortDescription: nEvent.ShortDescription,
		Description:      nEvent.Description,
		Longitude:        nEvent.Longitude,
		Latitude:         nEvent.Latitude,
		Images:           nEvent.Images,
		Preview:          nEvent.Preview,
		Date:             nEvent.Date,
		IsEnded:          nEvent.IsEnded,
		DeletedDate:      nEvent.DeletedDate,
	}
}
