package repository

import (
	"context"
	"database/sql"

	"github.com/grizlaz/ya-shortener/internal/model"
	_ "github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
)

type postgres struct {
	db *sql.DB
}

func NewPostgresDB(db *sql.DB) (*postgres, error) {
	pg := &postgres{db}
	migrationsDir := "migrations"
	if err := goose.Up(db, migrationsDir); err != nil {
		return nil, err
	}
	return pg, nil
}

func (p *postgres) Put(ctx context.Context, shortening model.Shortening) (*model.Shortening, error) {
	query := `INSERT INTO shortening (original_url, short_url)
			  VALUES ($1, $2)`
	_, err := p.db.ExecContext(ctx, query, shortening.OriginalURL, shortening.ShortURL)
	if err != nil {
		return nil, err
	}
	return &shortening, nil
}

func (p *postgres) Get(ctx context.Context, shortURL string) (*model.Shortening, error) {
	query := `SELECT s.id, s.original_url, s.short_url 
			  FROM shortening s
			  WHERE s.short_url = $1`
	var shortening model.Shortening
	row := p.db.QueryRowContext(ctx, query, shortURL)
	err := row.Scan(&shortening.ID, &shortening.OriginalURL, &shortening.ShortURL)
	if err != nil {
		return nil, err
	}
	return &shortening, nil
}
