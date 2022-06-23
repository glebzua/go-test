package main

import (
	"context"
	"fmt"
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/database"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/test_server/internal/infra/http"
	"github.com/test_server/internal/infra/http/controllers"
)

//database
var settings = postgresql.ConnectionURL{
	Database: `training`,
	Host:     `localhost:54322`,
	User:     `postgres`,
	Password: `root`,
}

// @title                       Test Server
// @version                     0.1.0
// @description                 Test Server boilerplate
func main() {

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

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

	// Event

	eventRepository := database.NewRepository()
	eventService := app.NewService(&eventRepository)
	eventController := controllers.NewEventController(&eventService)

	// HTTP Server
	errHttp := http.Server(
		ctx,
		http.Router(
			eventController,
		),
	)

	if errHttp != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
}
