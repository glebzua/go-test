package validators

import (
	"github.com/test_server/internal/domain"
	"time"
)

type eventsRequest struct {
	Id               uint      `json:"id,omitempty"`
	Title            string    `json:"Title"`
	ShortDescription string    `json:"Short Description"`
	Description      string    `json:"Description"`
	Longitude        float64   `json:"Longitude"`
	Latitude         float64   `json:"Latitude"`
	Images           string    `json:"Images"`
	Preview          string    `json:"Preview"`
	Date             string    `json:"Date"`
	IsEnded          bool      `json:"isEnded"`
	DeletedDate      time.Time `json:"DeletedDate,omitempty"`
}

func mapEventsRequestToDomain(eventsRequest *eventsRequest) *domain.Events {
	var events domain.Events
	events.Id = eventsRequest.Id
	events.Title = eventsRequest.Title
	events.ShortDescription = eventsRequest.ShortDescription
	events.Description = eventsRequest.Description
	events.Longitude = eventsRequest.Longitude
	events.Latitude = eventsRequest.Latitude
	events.Images = eventsRequest.Images
	events.Preview = eventsRequest.Preview
	events.Date = eventsRequest.Date
	events.IsEnded = eventsRequest.IsEnded
	events.DeletedDate = eventsRequest.DeletedDate
	return &events
}
