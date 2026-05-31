package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
)

type apiShortener interface {
	Shorten(context.Context, string, uuid.UUID) (*model.Shortening, error)
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Result string `json:"result"`
}

func HandleAPIShorten(shortener apiShortener, baseURL string) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer c.Request().Body.Close()
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		var request shortenRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if request.URL == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "empty url")
		}

		userID, err := getUserID(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		returnCode := http.StatusCreated
		shortening, err := shortener.Shorten(c.Request().Context(), request.URL, userID)
		if err != nil {
			if !errors.Is(err, model.ErrConflict) {
				logger.Log.Sugar().Infof("error shortening url %q: %v", request.URL, err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			returnCode = http.StatusConflict
		}

		shortURL, err := service.PrependBaseURL(baseURL, shortening.ShortURL)
		if err != nil {
			logger.Log.Sugar().Infof("error generating full url for %q: %v", shortening.ShortURL, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(returnCode, shortenResponse{
			Result: shortURL,
		})
	}
}
