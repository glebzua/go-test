package http

import (
	"github.com/go-chi/chi"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/test_server/internal/infra/http/controllers"
)

func Router(eventController *controllers.EventsController) http.Handler {
	router := chi.NewRouter()

	// Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Use(middleware.RedirectSlashes)

		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes, cors.Handler(cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		}))
		apiRouter.Route("/v1", func(apiRouter chi.Router) {
			AddEventRoutes(&apiRouter, eventController)

			apiRouter.Handle("/*", NotFoundJSON())
		})
	})

	return router
}

func AddEventRoutes(router *chi.Router, eventController *controllers.EventsController) {
	(*router).Route("/events", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/all/{page}",
			eventController.FindAll(),
		)
		apiRouter.Get(
			"/upcoming/{page}",
			eventController.FindUpcoming(),
		)
		apiRouter.Get(
			"/{id}",
			eventController.FindOne(),
		)
		apiRouter.Post(
			"/",
			eventController.Add(),
		)

	})
}
