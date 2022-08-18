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
	Create(task *domain.Events) (*domain.Events, error)
	Delete(eventId int64) error
	Update(user *domain.Events) (*domain.Events, error)
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
	//ct := fmt.Sprint(time.Now())
	err := r.coll.Find(isDeleted(db.Cond{"Date >": time.Now()}, showDeleted)).Paginate(pageSize).Page(page).All(&event)
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
func (r *eventsRepository) Create(event *domain.Events) (*domain.Events, error) {
	nEvent := mapDomainToEventDbModel(event)
	err := r.coll.InsertReturning(nEvent)
	if err != nil {
		log.Print("InsertReturning err", err)
		return nil, err
	}
	return mapEventsDbModelToDomain(nEvent), nil
}
func (r *eventsRepository) Update(event *domain.Events) (*domain.Events, error) {
	eventToUpdate := mapDomainToEventDbModel(event)

	err := r.coll.UpdateReturning(eventToUpdate)
	if err != nil {
		return nil, fmt.Errorf("userRepository UpdateUser: %w", err)
	}

	return mapEventsDbModelToDomain(eventToUpdate), nil
}

func (r *eventsRepository) Delete(eventId int64) error {
	var event events
	err := r.coll.Find(isDeleted(db.Cond{"id": eventId}, false)).One(&event)
	if err != nil {
		return fmt.Errorf("eventsRepository Delete: %w", err)
	}
	err = r.coll.Find(eventId).Update(map[string]interface{}{"deletedDate": time.Now()})
	if err != nil {
		return fmt.Errorf("eventsRepository DeleteEvent: %w", err)
	}

	return nil
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
