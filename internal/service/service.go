package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/model"
)

type Storage interface {
	Put(ctx context.Context, shortering model.Shortening) (*model.Shortening, error)
	Get(ctx context.Context, identifier string) (*model.Shortening, error)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Shorten(ctx context.Context, input string) (*model.Shortening, error) {
	shortURL := Shorten(uuid.New().ID())

	inputShorterin := model.Shortening{
		ShortURL:    shortURL,
		OriginalURL: input,
	}

	shortering, err := s.storage.Put(ctx, inputShorterin)
	if err != nil {
		return nil, err
	}

	return shortering, nil
}

func (s *Service) Get(ctx context.Context, shortURL string) (*model.Shortening, error) {
	return s.storage.Get(ctx, shortURL)
}

func (s *Service) Redirect(ctx context.Context, shortURL string) (string, error) {
	shortering, err := s.storage.Get(ctx, shortURL)
	if err != nil {
		return "", err
	}

	return shortering.OriginalURL, nil
}
