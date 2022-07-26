package http

import (
	"github.com/go-chi/chi"
	"github.com/test_server/internal/infra/http/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/test_server/internal/infra/http/controllers"
)

type HandlerMiddleware func(http.Handler) http.Handler

func Router(
	userController *controllers.UserController,
	eventController *controllers.EventsController,
	authMiddleware HandlerMiddleware,
) http.Handler {
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
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		}))
		apiRouter.Route("/a", func(apiRouter chi.Router) {
			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouter.Use(authMiddleware)
				AdminEventRoutes(&apiRouter, eventController)
				UserRouter(&apiRouter, userController)
				apiRouter.Handle("/*", NotFoundJSON())
			})
			apiRouter.Post(
				"/user/login",
				userController.LogIn(),
			)
			apiRouter.Handle("/*", NotFoundJSON())
		})
	})
	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes, cors.Handler(cors.Options{
			AllowedMethods: []string{"GET"},
		}))
		apiRouter.Route("/g", func(apiRouter chi.Router) {
			apiRouter.Group(func(apiRouter chi.Router) {
				GuestEventRoutes(&apiRouter, eventController)
				apiRouter.Handle("/*", NotFoundJSON())
			})

			apiRouter.Handle("/*", NotFoundJSON())
		})
	})
	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes, cors.Handler(cors.Options{
			AllowedMethods: []string{"GET"},
		}))
		apiRouter.Route("/m", func(apiRouter chi.Router) {
			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouterModeratorOnly := apiRouter.With(middlewares.ModeratorOnly)
				apiRouterModeratorOnly.Get(
					"/all/{page}",
					eventController.FindAll(),
				)
				apiRouterModeratorOnly.Get(
					"/upcoming/{page}",
					eventController.FindUpcoming(),
				)
				apiRouterModeratorOnly.Get(
					"/{id}",
					eventController.FindOne(),
				)
				apiRouterModeratorOnly.Post(
					"/",
					eventController.Create(),
				)
				apiRouter.Handle("/*", NotFoundJSON())
			})

			apiRouter.Handle("/*", NotFoundJSON())
		})
	})
	return router
}

func AdminEventRoutes(router *chi.Router, eventController *controllers.EventsController) {
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
			eventController.Create(),
		)

	})
}
func GuestEventRoutes(router *chi.Router, eventController *controllers.EventsController) {
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

	})
}

func UserRouter(router *chi.Router, userController *controllers.UserController) {
	(*router).Route("/user", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			userController.PaginateAll(),
		)
		apiRouter.Get(
			"/profile",
			userController.FindOne(),
		)
		apiRouter.Get(
			"/checkauth",
			userController.CheckAuth(),
		)
		apiRouterAdminOnly := apiRouter.With(middlewares.AdminOnly)
		apiRouterAdminOnly.Post(
			"/",
			userController.Create(),
		)
		apiRouterAdminOnly.Put(
			"/{id}",
			userController.Update(),
		)
		apiRouterAdminOnly.Delete(
			"/{id}",
			userController.Delete(),
		)
	})
}
