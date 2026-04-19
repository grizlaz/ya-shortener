package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Shorten(t *testing.T) {
	t.Run("generates shortening for a given URL", func(t *testing.T) {
		svc := service.NewService(context.TODO(), repository.NewInMemory())
		input := "https://practicum.yandex.ru/"

		shortening, err := svc.Shorten(context.Background(), input, uuid.New())
		require.NoError(t, err)

		assert.NotEmpty(t, shortening.ShortURL)
		assert.Equal(t, input, shortening.OriginalURL)
	})
}

func TestService_Redirect(t *testing.T) {
	t.Run("returns redirect URL for a given identifier", func(t *testing.T) {
		inMemoryStorage := repository.NewInMemory()
		svc := service.NewService(context.TODO(), inMemoryStorage)
		input := "https://practicum.yandex.ru/"

		shortening, err := svc.Shorten(context.Background(), input, uuid.New())
		require.NoError(t, err)

		redirectURL, err := svc.Redirect(context.Background(), shortening.ShortURL)
		require.NoError(t, err)

		assert.Equal(t, input, redirectURL)
	})

	t.Run("returns error if identifier is not found", func(t *testing.T) {
		var svc = service.NewService(context.TODO(), repository.NewInMemory())

		_, err := svc.Redirect(context.Background(), "yandex")
		assert.ErrorIs(t, err, model.ErrNotFound)
	})
}
