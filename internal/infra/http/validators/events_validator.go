package validators

import (
	"encoding/json"
	"fmt"
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
	var eventsResource eventsRequest
	err := json.NewDecoder(r.Body).Decode(&eventsResource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	err = t.validator.Struct(eventsResource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapEventsRequestToDomain(&eventsResource), nil
}

type EventsUpdateValidator struct {
	validator *validator.Validate
}

func NewEventsUpdateValidator() *EventsUpdateValidator {
	return &EventsUpdateValidator{
		validator: validator.New(),
	}
}

func (t EventsUpdateValidator) ValidateAndMap(r *http.Request) (*domain.Events, error) {
	var eventsResource eventsUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&eventsResource)
	if err != nil {
		return nil, fmt.Errorf("userValidator ValidateAndMap: %w", err)
	}
	err = t.validator.Struct(eventsResource)
	if err != nil {
		return nil, fmt.Errorf("userValidator ValidateAndMap: %w", err)
	}

	return mapEventsUpdateRequestToDomain(&eventsResource), nil
}
