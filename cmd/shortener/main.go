package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/grizlaz/ya-shortener/internal/handler"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
)

func main() {
	shorteningStorage := repository.NewInMemory()
	shortener := service.NewService(shorteningStorage)
	srv := handler.NewServer(shortener)
	if err := http.ListenAndServe("localhost:8080", srv); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("error running server: %v", err)
	}
}
