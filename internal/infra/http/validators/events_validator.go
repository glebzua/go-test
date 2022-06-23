package validators

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/test_server/internal/domain"
	"log"
	"net/http"
)

type EventsValidator struct {
	validator *validator.Validate
}

func NewEventsValidator() *EventsValidator {
	return &EventsValidator{
		validator: validator.New(),
	}
}

func (t EventsValidator) ValidateAndMap(r *http.Request) (*domain.Events, error) {
	var cattleResource eventsRequest
	err := json.NewDecoder(r.Body).Decode(&cattleResource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	err = t.validator.Struct(cattleResource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapEventsRequestToDomain(&cattleResource), nil
}
