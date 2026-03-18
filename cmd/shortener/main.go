package main

import (
	"errors"
	"net/http"

	"github.com/grizlaz/ya-shortener/internal/config"
	"github.com/grizlaz/ya-shortener/internal/handler"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
)

func main() {
	if err := logger.Initialize("info"); err != nil {
		panic(err)
	}
	config := config.Get()
	shorteningStorage := repository.NewInMemory()
	shortener := service.NewService(shorteningStorage)
	srv := handler.NewServer(shortener, config.BaseURL)
	if err := http.ListenAndServe(config.ServerAddress, srv); !errors.Is(err, http.ErrServerClosed) {
		logger.Log.Sugar().Fatalf("error running server: %v", err)
	}
}
