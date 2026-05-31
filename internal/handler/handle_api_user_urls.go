package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
)

type shortenerUserUrls interface {
	GetUserUrls(context.Context, uuid.UUID) (*[]model.Shortening, error)
}

type userURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func HandleUserUrls(shortener shortenerUserUrls, baseURL string) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := getUserID(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		shortenings, err := shortener.GetUserUrls(c.Request().Context(), userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if len(*shortenings) == 0 {
			return c.NoContent(http.StatusNoContent)
		}

		response := make([]userURL, 0, len(*shortenings))
		for _, v := range *shortenings {
			shortURL, err := service.PrependBaseURL(baseURL, v.ShortURL)
			if err != nil {
				logger.Log.Sugar().Infof("error generating full url for %q: %v", v.ShortURL, err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			response = append(response, userURL{
				OriginalURL: v.OriginalURL,
				ShortURL:    shortURL,
			})
		}
		return c.JSON(http.StatusOK, response)
	}
}
