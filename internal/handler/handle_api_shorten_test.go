package handler_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/grizlaz/ya-shortener/internal/handler"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleShorten(t *testing.T) {
	t.Run("get short url", func(t *testing.T) {
		baseURL := "http://localhost:8080"
		url := "https://practicum.yandex.ru"
		path := "/api/shorten"
		body := strings.NewReader(fmt.Sprintf(`{"url":"%s"}`, url))

		shorten := service.NewService(context.Background(), repository.NewInMemory())
		handler := handler.HandleAPIShorten(shorten, baseURL)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, path, body)
		e := echo.New()
		c := e.NewContext(request, recorder)

		c.SetPath(path)

		require.NoError(t, handler(c))

		responseBody, err := io.ReadAll(recorder.Result().Body) //nolint:bodyclose
		recorder.Result().Body.Close()
		require.NoError(t, err)

		assert.Contains(t, string(responseBody), baseURL)
		assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)
	})
}
