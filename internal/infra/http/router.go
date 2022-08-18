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
				AdminUserRouter(&apiRouter, userController)
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
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		}))
		apiRouter.Route("/m", func(apiRouter chi.Router) {
			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouterModeratorOnly := apiRouter.With(middlewares.ModeratorOnly)
				apiRouterModeratorOnly.Get(
					"/events/all/{page}",
					eventController.FindAll(),
				)
				apiRouterModeratorOnly.Get(
					"/events/upcoming/{page}",
					eventController.FindUpcoming(),
				)
				apiRouterModeratorOnly.Get(
					"/events/{id}",
					eventController.FindOne(),
				)
				apiRouterModeratorOnly.Post(
					"/events",
					eventController.Create(),
				)
				apiRouterModeratorOnly.Put(
					"/events/{id}",
					eventController.Update(),
				)
				apiRouterModeratorOnly.Delete(
					"/events/{id}",
					eventController.Delete(),
				)

				apiRouter.Handle("/*", NotFoundJSON())
			})
			apiRouter.Post(
				"/user/login",
				userController.LogIn(),
			)
			apiRouter.Handle("/*", NotFoundJSON())
		})
	})
	return router
}

func AdminEventRoutes(router *chi.Router, eventController *controllers.EventsController) {
	(*router).Route("/events", func(apiRouter chi.Router) {
		apiRouterAdminOnly := apiRouter.With(middlewares.AdminOnly)
		apiRouterAdminOnly.Get(
			"/all/{page}",
			eventController.FindAll(),
		)
		apiRouterAdminOnly.Get(
			"/upcoming/{page}",
			eventController.FindUpcoming(),
		)
		apiRouterAdminOnly.Get(
			"/{id}",
			eventController.FindOne(),
		)
		apiRouterAdminOnly.Put(
			"/{id}",
			eventController.Update(),
		)
		apiRouterAdminOnly.Post(
			"/",
			eventController.Create(),
		)

		apiRouterAdminOnly.Delete(
			"/",
			eventController.Delete(),
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

func AdminUserRouter(router *chi.Router, userController *controllers.UserController) {
	(*router).Route("/users", func(apiRouter chi.Router) {
		apiRouterAdminOnly := apiRouter.With(middlewares.AdminOnly)
		apiRouterAdminOnly.Get(
			"/{id}",
			userController.FindOneById(),
		)
		apiRouterAdminOnly.Put(
			"/{id}",
			userController.Update(),
		)
		apiRouterAdminOnly.Delete(
			"/{id}",
			userController.Delete(),
		)
		apiRouterAdminOnly.Post(
			"/",
			userController.Create(),
		)
		apiRouterAdminOnly.Get(
			"/",
			userController.PaginateAll(),
		)
		apiRouterAdminOnly.Get(
			"/profile/{id}",
			userController.FindOne(),
		)
		apiRouterAdminOnly.Get(
			"/checkauth",
			userController.CheckAuth(),
		)
	})
}
