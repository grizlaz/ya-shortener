package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/grizlaz/ya-shortener/internal/audit"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleShorten(t *testing.T) {
	initCap := 100
	t.Run("success", func(t *testing.T) {
		baseURL := "http://localhost:8080"
		url := "https://practicum.yandex.ru"
		path := "/"
		body := strings.NewReader(url)

		shorten := service.NewService(context.Background(), repository.NewInMemory(initCap))
		audit := audit.NewAudit()
		handler := HandleShorten(shorten, baseURL, audit)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, path, body)
		request.Header.Add("Content-Type", "text/plain")
		e := echo.New()
		c := e.NewContext(request, recorder)

		c.SetPath(path)

		require.NoError(t, handler(c))
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}
