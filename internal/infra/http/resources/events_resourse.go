package resources

import (
	"github.com/test_server/internal/domain"
	"time"
)

type EventsDto struct {
	Id               uint    `json:"id,omitempty"`
	Title            string  `json:"Title"`
	ShortDescription string  `json:"ShortDescription"`
	Description      string  `json:"Description"`
	Longitude        float64 `json:"Longitude"`
	Latitude         float64 `json:"Latitude"`
	Images           string  `json:"Images"`
	Preview          string  `json:"Preview"`
	Date             string  `json:"Date"`
}
type AllEventsDto struct {
	Id               uint      `json:"id,omitempty"`
	Title            string    `json:"Title"`
	ShortDescription string    `json:"ShortDescription"`
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

func MapDomainToAllEventsDto(events *domain.Events) *AllEventsDto {
	return &AllEventsDto{
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

func MapDomainToAllEventsDtoCollection(events []domain.Events) []AllEventsDto {
	var result []AllEventsDto
	for _, t := range events {
		dto := MapDomainToAllEventsDto(&t)
		result = append(result, *dto)
	}
	return result
}
