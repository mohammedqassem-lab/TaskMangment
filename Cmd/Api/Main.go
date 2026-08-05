// @title Task Management API
// @version 1.0
// @description Task Management System API
// @termsOfService http://swagger.io/terms/

// @contact.name Mohammed
// @contact.email mohammed.qassem@baly.iq

// @license.name MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	app "TaskMangment/Internal/App"
	logfile "TaskMangment/Internal/LogFile"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "TaskMangment/docs"
)

func main() {
	logfileCleanup := logfile.InitLogger()
	defer logfileCleanup()
	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: application.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Error: %v", err)
	}

	if err := application.Shutdown(); err != nil {
		log.Printf("Application Shutdown Error: %v", err)
	}

	log.Println("Application Stopped Successfully")
}
