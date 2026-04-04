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

type apiShortenerBatch interface {
	ShortenBatch(context.Context, *[]model.ShortenRequestBatch) (*[]model.Shortening, error)
}

type batchResponse struct {
	ID  string `json:"correlation_id"`
	URL string `json:"short_url"`
}

func HandleAPIShortenBatch(shortener apiShortenerBatch, baseURL string) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer c.Request().Body.Close()
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		var request []model.ShortenRequestBatch
		err = json.Unmarshal(body, &request)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		for _, v := range request {
			if v.URL == "" {
				return echo.NewHTTPError(http.StatusBadRequest, "empty url")
			}
		}

		shortens, err := shortener.ShortenBatch(c.Request().Context(), &request)
		if err != nil {
			logger.Log.Sugar().Infof("error shortening batch urls: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if len(*shortens) != len(request) {
			logger.Log.Sugar().Infof("error shortening batch urls len(request) != len(result): %d != %d", len(*shortens), len(request))
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		response := make([]batchResponse, 0, len(*shortens))
		for i, v := range *shortens {
			shortURL, err := service.PrependBaseURL(baseURL, v.ShortURL)
			if err != nil {
				logger.Log.Sugar().Infof("error generating full url for %q: %v", v.ShortURL, err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			response = append(response, batchResponse{
				ID:  request[i].ID,
				URL: shortURL,
			})
		}

		return c.JSON(http.StatusCreated, response)
	}
}
