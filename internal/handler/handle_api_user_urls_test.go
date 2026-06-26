package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	// "your-project/internal/models" // содержит структуры, например Shortening и userURL
	// содержит PrependBaseURL
)

func TestHandleUserUrls(t *testing.T) {
	initCap := 100
	baseURL := "http://localhost:8080"
	t.Run("empty userID", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
		recorder := httptest.NewRecorder()
		c := e.NewContext(request, recorder)
		shorten := service.NewService(context.Background(), repository.NewInMemory(initCap))
		handler := HandleUserUrls(shorten, baseURL)
		err := handler(c)

		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})

	t.Run("success", func(t *testing.T) {
		userID := uuid.New()
		userURLs := []model.Shortening{
			{
				ShortURL:    "123",
				OriginalURL: "https://practicum.yandex.ru",
				ID:          1,
				IsDeleted:   false,
				UserID:      userID,
			},
		}
		jwt, err := makeJWT(userID)
		require.NoError(t, err)

		request := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
		cookie := &http.Cookie{
			Name:  "Authorization",
			Value: jwt,
		}
		request.AddCookie(cookie)
		recorder := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(request, recorder)

		repository := repository.NewInMemory(initCap)
		repository.PutBatch(t.Context(), &userURLs)
		shorten := service.NewService(context.Background(), repository)
		handler := HandleUserUrls(shorten, baseURL)
		err = handler(c)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
