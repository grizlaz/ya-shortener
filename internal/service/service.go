package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
	"go.uber.org/zap"
)

type Storage interface {
	Get(ctx context.Context, identifier string) (*model.Shortening, error)
	GetUserUrls(ctx context.Context, userID uuid.UUID) (*[]model.Shortening, error)
	Put(ctx context.Context, shortering model.Shortening) (*model.Shortening, error)
	PutBatch(ctx context.Context, shortering *[]model.Shortening) (int64, error)
	DeleteUserUrls(ctx context.Context, deleteUrls ...model.DeleteUrls) error
}

type Service struct {
	storage  Storage
	deleteCh chan model.DeleteUrls
}

func NewService(ctx context.Context, storage Storage) *Service {
	service := &Service{
		storage:  storage,
		deleteCh: make(chan model.DeleteUrls, 100),
	}
	go service.flushDeleteUrls(ctx)
	return service
}

func (s *Service) Shorten(ctx context.Context, input string, userID uuid.UUID) (*model.Shortening, error) {
	shortURL := Shorten(uuid.New().ID())

	inputShorterin := model.Shortening{
		ShortURL:    shortURL,
		OriginalURL: input,
		UserID:      userID,
		IsDeleted:   false,
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
			IsDeleted:   false,
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
	if shortering.IsDeleted {
		return "", model.ErrURLDeleted
	}

	return shortering.OriginalURL, nil
}

func (s *Service) GetUserUrls(ctx context.Context, userID uuid.UUID) (*[]model.Shortening, error) {
	return s.storage.GetUserUrls(ctx, userID)
}

func (s *Service) DeleteUserUrls(_ context.Context, deleteUrls model.DeleteUrls) error {
	s.deleteCh <- deleteUrls
	// return s.storage.DeleteUserUrls(ctx, deleteUrls)
	return nil
}

func (s *Service) flushDeleteUrls(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	// Норм канал здесь закрывать? Пишется он в DeleteUserUrls, но объект тот же и слежу за контекстом процесса здесь
	// defer close(s.deleteCh)

	var queue []model.DeleteUrls

	for {
		select {
		case delUrls := <-s.deleteCh:
			queue = append(queue, delUrls)
		case <-ticker.C:
			if len(queue) == 0 {
				continue
			}
			err := s.storage.DeleteUserUrls(ctx, queue...)
			if err != nil {
				logger.Log.Debug("cannot delete urls", zap.Error(err))
				continue
			}
			queue = nil
		case <-ctx.Done():
			return
		}
	}
}
