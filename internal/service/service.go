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
	identifier := Shorten(uuid.New().ID())

	inputShorterin := model.Shortening{
		Identifier:  identifier,
		OriginalURL: input,
	}

	shortering, err := s.storage.Put(ctx, inputShorterin)
	if err != nil {
		return nil, err
	}

	return shortering, nil
}

func (s *Service) Get(ctx context.Context, identifier string) (*model.Shortening, error) {
	return s.storage.Get(ctx, identifier)
}

func (s *Service) Redirect(ctx context.Context, identifier string) (string, error) {
	shortering, err := s.storage.Get(ctx, identifier)
	if err != nil {
		return "", err
	}

	return shortering.OriginalURL, nil
}
