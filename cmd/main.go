package main

import (
	"api/internal/application"
	"api/internal/infrastructure/repository"
	"api/internal/transport/http"
	"api/pkg/configuration"
	"api/pkg/postgresql"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := configuration.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgresql.Connect(&cfg.Postgresql)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.New(db)
	useCase := application.New(repository)

	httpServer := http.New(useCase)

	go func() {
		if err := httpServer.Run(cfg.HttpSocket); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	_, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
}
