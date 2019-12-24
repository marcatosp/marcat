package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Guest"
	}

	log.Printf("Received request for %s\n", name)
	w.Write([]byte (fmt.Sprintf("Hello, %s\n", name)))
}

func main() {
	// Create Server and Route Handlers

	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		TLSConfig:    nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Config logging
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")

	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500,
			MaxAge:     28,
			MaxBackups: 3,
			Compress:   true,
		})
	}

	// Start server
	go func() {
		log.Println("Starting server")

		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interrupChan := make(chan os.Signal, 1)
	signal.Notify(interrupChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal
	<-interrupChan

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
