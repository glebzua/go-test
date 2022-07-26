package controllers

import (
	"github.com/go-chi/chi"
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/http/resources"
	"github.com/test_server/internal/infra/http/validators"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type EventsController struct {
	eventsService *app.EventsService
	validator     *validators.EventsValidator
}

func NewEventsController(s *app.EventsService) *EventsController {
	return &EventsController{
		eventsService: s,
		validator:     validators.NewEventsValidator(),
	}
}

func (c *EventsController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(chi.URLParam(r, "page"))
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		events, err := (*c.eventsService).FindAll(uint(page), 20)
		if err != nil {
			log.Printf("EventController.FindAll(): %s", err)
			return
		}

		success(w, resources.MapDomainToAllEventsDtoCollection(events))

	}
}

func (c *EventsController) FindUpcoming() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(chi.URLParam(r, "page"))
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		events, err := (*c.eventsService).FindUpcoming(uint(page), 20)
		if err != nil {
			log.Printf("EventController.FindUpcoming(): %s", err)
			return
		}

		success(w, resources.MapDomainToEventsDtoCollection(events))

	}
}

func (c *EventsController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Print("eventsController FindOne ParseInt", err)
			badRequest(w, err)
			return
		}
		event, err := (*c.eventsService).FindOne(uint64(id))
		if err != nil {
			if strings.HasSuffix(err.Error(), "upper: no more rows in this result set") {
				notFound(w, err)
			} else {
				internalServerError(w, err)
			}
			log.Print("eventsController FindOne error:", err)
			return
		}

		success(w, resources.MapDomainToEventsDto(event))

	}
}

func (c *EventsController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		event, err := c.validator.ValidateAndMap(r)
		if err != nil {
			log.Print(err)
			badRequest(w, err)
			return
		}

		addedEvent, err := (*c.eventsService).Create(event)
		if err != nil {
			log.Print(err)
			internalServerError(w, err)
			return
		}

		success(w, resources.MapDomainToEventsDto(addedEvent))

	}
}
