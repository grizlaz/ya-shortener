package handler

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
)

type shortener interface {
	Shorten(context.Context, string, uuid.UUID) (*model.Shortening, error)
}

func HandleShorten(shortener shortener, baseURL string) echo.HandlerFunc {
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
		return c.String(
			returnCode,
			shortURL,
		)
	}
}
