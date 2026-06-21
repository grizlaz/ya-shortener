package handler

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/grizlaz/ya-shortener/internal/audit"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/grizlaz/ya-shortener/internal/service"
)

type shortener interface {
	Shorten(context.Context, string, uuid.UUID) (*model.Shortening, error)
}

// Handler для обработки сокращения ссылки.
// Оригинальная ссылка ожидается в теле запроса в виде строки.
// Запрос:
//
//	http://example.com
//
// Ответ:
//
//	http://localhost:8080/b9j5za
func HandleShorten(shortener shortener, baseURL string, audit *audit.Audit) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer c.Request().Body.Close()

		contentType := c.Request().Header.Get("Content-Type")
		if contentType != "text/plain" {
			return echo.NewHTTPError(http.StatusBadRequest, "wrong content-type")
		}

		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		requestURL := string(body)
		if requestURL == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "empty body")
		}

		userID, err := getUserID(c)
		if err != nil {
			logger.Log.Sugar().Errorf("error get user id %q: %v", requestURL, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		returnCode := http.StatusCreated
		shortening, err := shortener.Shorten(c.Request().Context(), requestURL, userID)
		if err != nil {
			if !errors.Is(err, model.ErrConflict) {
				logger.Log.Sugar().Errorf("error shortening url %q: %v", requestURL, err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			returnCode = http.StatusConflict
		}

		shortURL, err := service.PrependBaseURL(baseURL, shortening.ShortURL)
		if err != nil {
			logger.Log.Sugar().Errorf("error generating full url for %q: %v", shortening.ShortURL, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		audit.Send(model.AuditMessage{
			TS:     time.Now().Unix(),
			Action: "shorten",
			UserID: userID.String(),
			URL:    requestURL,
		})
		return c.String(
			returnCode,
			shortURL,
		)
	}
}
