package resources

import (
	"github.com/test_server/internal/domain"
	"time"
)

type EventsDto struct {
	Id               uint      `json:"id,omitempty"`
	Title            string    `json:"Title"`
	ShortDescription string    `json:"Short Description"`
	Description      string    `json:"Description"`
	Longitude        float64   `json:"Longitude"`
	Latitude         float64   `json:"Latitude"`
	Images           string    `json:"Images"`
	Preview          string    `json:"Preview"`
	Date             string    `json:"Date"`
	IsEnded          bool      `json:"IsEnded"`
	DeletedDate      time.Time `json:"DeletedDate"`
}

func MapDomainToEventsDto(events *domain.Events) *EventsDto {
	return &EventsDto{
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

func MapDomainToEventsDtoCollection(events []domain.Events) []EventsDto {
	var result []EventsDto
	for _, t := range events {
		dto := MapDomainToEventsDto(&t)
		result = append(result, *dto)
	}
	return result
}
