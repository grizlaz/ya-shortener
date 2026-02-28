package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/grizlaz/ya-shortener/internal/config"
	"github.com/grizlaz/ya-shortener/internal/handler"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
)

func main() {
	config := config.Get()
	shorteningStorage := repository.NewInMemory()
	shortener := service.NewService(shorteningStorage)
	srv := handler.NewServer(shortener)
	if err := http.ListenAndServe(config.A, srv); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("error running server: %v", err)
	}
}
