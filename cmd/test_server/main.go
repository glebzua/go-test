package main

import (
	"context"
	"fmt"
	"github.com/test_server/config"
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/database"
	"github.com/test_server/internal/infra/http"
	"github.com/test_server/internal/infra/http/controllers"
	"github.com/test_server/internal/infra/http/middlewares"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

func main() {

	var conf = config.GetConfiguration()

	sess, err := postgresql.Open(
		postgresql.ConnectionURL{
			User:     conf.DatabaseUser,
			Host:     conf.DatabaseHost,
			Password: conf.DatabasePassword,
			Database: conf.DatabaseName,
		})
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer func(sess db.Session) {
		err := sess.Close()
		if err != nil {
			log.Print("defer func sess.Close error:", err)
		}
	}(sess)
	log.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), conf.DatabaseHost)

	exitCode := 0
	ctx, cancel := context.WithCancel(context.Background())

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The system panicked!: %v\n", r)
			fmt.Printf("Stack trace form panic: %s\n", string(debug.Stack()))
			exitCode = 1
		}
		os.Exit(exitCode)
	}()

	// Signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("Received signal '%s', stopping... \n", sig.String())
		cancel()
		fmt.Printf("Sent cancel to all threads...")
	}()

	userRepository := database.NewUserRepository(&sess)
	userService := app.NewUserService(&userRepository)
	tokenService := app.NewTokenService([]byte(conf.AuthAccessKeySecret))
	userController := controllers.NewUserController(&userService, &tokenService)

	eventsRepository := database.NewEventsRepository(&sess)
	eventService := app.NewEventsService(&eventsRepository)
	eventController := controllers.NewEventsController(&eventService)
	authMiddleware := middlewares.AuthMiddleware(tokenService)
	// HTTP Server
	errHttp := http.Server(
		ctx,
		http.Router(
			userController,
			eventController,
			authMiddleware,
		),
	)

	if errHttp != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}

}
