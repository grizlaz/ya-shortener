package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
)

type apiShortener interface {
	Shorten(context.Context, string) (*model.Shortening, error)
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

		shortening, err := shortener.Shorten(c.Request().Context(), request.URL)
		if err != nil {
			logger.Log.Sugar().Infof("error shortening url %q: %v", request.URL, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		shortURL, err := service.PrependBaseURL(baseURL, shortening.ShortURL)
		if err != nil {
			logger.Log.Sugar().Infof("error generating full url for %q: %v", shortening.ShortURL, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusCreated, shortenResponse{
			Result: shortURL,
		})
	}
}
