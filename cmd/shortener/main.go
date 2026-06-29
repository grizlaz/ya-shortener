package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/grizlaz/ya-shortener/internal/audit"
	"github.com/grizlaz/ya-shortener/internal/config"
	"github.com/grizlaz/ya-shortener/internal/handler"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func main() {
	setDefaultBuildInfo()
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	var err error
	if err = logger.Initialize("info"); err != nil {
		panic(err)
	}
	cfg := config.Get()

	var shorteningStorage service.Storage
	var db *sql.DB
	if cfg.DatabaseDSN != "" {
		db, err = sql.Open("pgx", cfg.DatabaseDSN)
		if err != nil {
			logger.Log.Sugar().Fatalf("error init db: %w", err)
		}
		defer func() {
			if derr := db.Close(); derr != nil {
				logger.Log.Sugar().Errorf("error defer db.Close(): %v", derr)
			}
		}()
		shorteningStorage, err = repository.NewPostgresDB(db)
	} else {
		shorteningStorage, err = repository.NewInFile(cfg.FileStoragePath, false)
	}
	if err != nil {
		logger.Log.Sugar().Fatalf("error init storage: %w", err)
	}
	auditService, err := initAudit()
	if err != nil {
		logger.Log.Sugar().Fatalf("error init audit service: %w", err)
	}

	shortener := service.NewService(context.Background(), shorteningStorage)
	srv := handler.NewServer(shortener, cfg.BaseURL, db, auditService)
	if err := http.ListenAndServe(cfg.ServerAddress, srv); !errors.Is(err, http.ErrServerClosed) {
		logger.Log.Sugar().Fatalf("error running server: %w", err)
	}
}

func setDefaultBuildInfo() {
	if buildVersion == "" {
		buildVersion = "N/A"
	}
	if buildDate == "" {
		buildDate = "N/A"
	}
	if buildCommit == "" {
		buildCommit = "N/A"
	}
}

func initAudit() (*audit.Audit, error) {
	cfg := config.Get()
	auditService := audit.NewAudit()
	if cfg.AuditFilePath != "" {
		fileStorage, err := repository.NewInFile(cfg.AuditFilePath, true)
		if err != nil {
			return nil, err
		}
		auditInFile := audit.NewObserver("file", fileStorage)
		auditService.Register(auditInFile)
	}
	if cfg.AuditURL != "" {
		auditClient := repository.NewAuditClient(cfg.AuditURL)
		httpAudit := audit.NewObserver("http", auditClient)
		auditService.Register(httpAudit)
	}
	return auditService, nil
}
