package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/model"
)

type Storage interface {
	Get(ctx context.Context, identifier string) (*model.Shortening, error)
	GetUserUrls(ctx context.Context, userID uuid.UUID) (*[]model.Shortening, error)
	Put(ctx context.Context, shortering model.Shortening) (*model.Shortening, error)
	PutBatch(ctx context.Context, shortering *[]model.Shortening) (int64, error)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Shorten(ctx context.Context, input string, userID uuid.UUID) (*model.Shortening, error) {
	shortURL := Shorten(uuid.New().ID())

	inputShorterin := model.Shortening{
		ShortURL:    shortURL,
		OriginalURL: input,
		UserID:      userID,
	}

	shortering, err := s.storage.Put(ctx, inputShorterin)
	if err != nil {
		return shortering, err
	}

	return shortering, nil
}

func (s *Service) ShortenBatch(ctx context.Context, inputs *[]model.ShortenRequestBatch, userID uuid.UUID) (*[]model.Shortening, error) {
	result := make([]model.Shortening, 0, len(*inputs))
	for _, v := range *inputs {
		shortURL := Shorten(uuid.New().ID())
		result = append(result, model.Shortening{
			ShortURL:    shortURL,
			OriginalURL: v.URL,
			UserID:      userID,
		})
	}
	_, err := s.storage.PutBatch(ctx, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
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

func (s *Service) GetUserUrls(ctx context.Context, userID uuid.UUID) (*[]model.Shortening, error) {
	return s.storage.GetUserUrls(ctx, userID)
}
