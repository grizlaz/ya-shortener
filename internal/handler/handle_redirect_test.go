package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/handler"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleRedirect(t *testing.T) {
	t.Run("redirects to original URL", func(t *testing.T) {
		url := "https://practicum.yandex.ru"

		redirecter := service.NewService(repository.NewInMemory())
		handler := handler.HandleRedirect(redirecter)

		shortening, err := redirecter.Shorten(context.Background(), url, uuid.New())
		require.NoError(t, err)
		identifier := shortening.ShortURL

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/"+identifier, nil)
		e := echo.New()
		c := e.NewContext(request, recorder)

		c.SetPath("/:identifier")
		c.SetParamNames("identifier")
		c.SetParamValues(identifier)

		require.NoError(t, handler(c))
		assert.Equal(t, http.StatusTemporaryRedirect, recorder.Code)
		assert.Equal(t, url, recorder.Header().Get("Location"))
	})

	t.Run("returns 404 if identifier is not found", func(t *testing.T) {
		identifier := "ya"
		redirecter := service.NewService(repository.NewInMemory())
		handler := handler.HandleRedirect(redirecter)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/"+identifier, nil)
		e := echo.New()
		c := e.NewContext(request, recorder)

		c.SetPath("/:identifier")
		c.SetParamNames("identifier")
		c.SetParamValues(identifier)

		require.Error(t, handler(c))
	})
}
