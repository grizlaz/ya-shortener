package main

import (
	"database/sql"
	"errors"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/grizlaz/ya-shortener/internal/config"
	"github.com/grizlaz/ya-shortener/internal/handler"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
)

func main() {
	var err error
	if err = logger.Initialize("info"); err != nil {
		panic(err)
	}
	config := config.Get()

	var shorteningStorage service.Storage
	var db *sql.DB
	if config.DatabaseDSN != "" {
		db, err = sql.Open("pgx", config.DatabaseDSN)
		if err != nil {
			logger.Log.Sugar().Fatalf("error init db: %v", err)
		}
		defer db.Close()
		shorteningStorage, err = repository.NewPostgresDB(db)
	} else {
		// shorteningStorage := repository.NewInMemory()
		shorteningStorage, err = repository.NewInFile(config.FileStoragePath)
	}
	if err != nil {
		logger.Log.Sugar().Fatalf("error init file storage: %v", err)
	}

	shortener := service.NewService(shorteningStorage)
	srv := handler.NewServer(shortener, config.BaseURL, db)
	if err := http.ListenAndServe(config.ServerAddress, srv); !errors.Is(err, http.ErrServerClosed) {
		logger.Log.Sugar().Fatalf("error running server: %v", err)
	}
}
