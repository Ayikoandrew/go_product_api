package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"main.go/handlers"
	"main.go/middleware"
)

func main() {
	lg := log.New(os.Stdout, "product-api ", log.LstdFlags)

	hp := handlers.NewProduct(lg)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", hp.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", hp.AddProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", hp.UpdateProduct)

	sm.Use(middleware.LogReport(lg, hp))

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ErrorLog:     lg,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		lg.Printf("Starting server on port 8080")
		if err := s.ListenAndServe(); err != nil {
			log.Println("Error", err)
			return
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	lg.Printf("Received terminate signal: %v\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		lg.Printf("Server forced to shutdown: %v", err)
		os.Exit(1)
	}

	lg.Printf("Server stopped gracefully")
}
