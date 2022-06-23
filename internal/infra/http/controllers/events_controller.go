package controllers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/http/resources"
	"github.com/test_server/internal/infra/http/validators"
	"log"
	"net/http"
	"strconv"
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

		err = success(w, resources.MapDomainToEventsDtoCollection(events))
		if err != nil {
			log.Print(err)

		}
	}
}

func (c *EventsController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := (*c.eventsService).FindOne(uint64(id))
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, resources.MapDomainToEventsDto(event))
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}
