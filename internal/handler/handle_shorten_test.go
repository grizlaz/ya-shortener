package handler_test

import (
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
	t.Run("returns shortened URL for a given URL", func(t *testing.T) {
		var (
			shortener = service.NewService(repository.NewInMemory())
			handler   = handler.HandleShorten(shortener)
			recorder  = httptest.NewRecorder()
			request   = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://practicum.yandex.ru/"))
			e         = echo.New()
			c         = e.NewContext(request, recorder)
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)

		require.NoError(t, handler(c))
		assert.Equal(t, http.StatusCreated, recorder.Code)

		defer recorder.Result().Body.Close()
		body, err := io.ReadAll(recorder.Result().Body)

		require.NoError(t, err)
		assert.NotEmpty(t, body)
	})
}
