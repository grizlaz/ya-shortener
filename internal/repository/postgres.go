package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/jackc/pgerrcode"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pressly/goose/v3"
)

type postgres struct {
	db *sql.DB
}

const insertQuery = "INSERT INTO shortening (original_url, short_url) VALUES ($1, $2)"

func NewPostgresDB(db *sql.DB) (*postgres, error) {
	pg := &postgres{db}
	migrationsDir := "migrations"
	if err := goose.Up(db, migrationsDir); err != nil {
		return nil, err
	}
	return pg, nil
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

func (p *postgres) GetByOriginalURL(ctx context.Context, originalURL string) (*model.Shortening, error) {
	query := `SELECT s.id, s.original_url, s.short_url
			  FROM shortening s
			  WHERE s.original_url = $1`
	var shortening model.Shortening
	row := p.db.QueryRowContext(ctx, query, originalURL)
	err := row.Scan(&shortening.ID, &shortening.OriginalURL, &shortening.ShortURL)
	if err != nil {
		return nil, err
	}
	return &shortening, nil
}

func (p *postgres) Put(ctx context.Context, shortening model.Shortening) (*model.Shortening, error) {
	_, err := p.db.ExecContext(ctx, insertQuery, shortening.OriginalURL, shortening.ShortURL)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			err = model.ErrConflict
			conflictShorten, errGet := p.GetByOriginalURL(ctx, shortening.OriginalURL)
			if errGet != nil {
				return nil, errGet
			}
			//на сколько это нормально вообще/для данной задачи?
			return conflictShorten, err
		}
		return nil, err
	}
	return &shortening, nil
}

func (p *postgres) PutBatch(ctx context.Context, shortens *[]model.Shortening) (int64, error) {
	inserts := int64(0)
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	stmt, err := p.db.PrepareContext(ctx, insertQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	for _, v := range *shortens {
		res, err := stmt.ExecContext(ctx, v.OriginalURL, v.ShortURL)
		if err != nil {
			return inserts, err //TODO нужно ли тут возвращать количество сохраненных строк?
		}
		if ins, err := res.RowsAffected(); err == nil {
			inserts += ins
		}
	}

	err = tx.Commit()
	if err != nil {
		return inserts, err
	}
	return inserts, nil
}
