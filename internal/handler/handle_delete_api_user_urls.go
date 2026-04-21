package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/labstack/echo/v4"
)

type shortenerDeleteUserUrls interface {
	DeleteUserUrls(context.Context, model.DeleteUrls) error
}

func HandleDeleteUserUrls(shortener shortenerDeleteUserUrls) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer c.Request().Body.Close()
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		var urls []string
		err = json.Unmarshal(body, &urls)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		userID, err := getUserID(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		err = shortener.DeleteUserUrls(c.Request().Context(), model.DeleteUrls{
			UserID: userID,
			Urls:   &urls})
		if err != nil {
			logger.Log.Sugar().Errorf("error delete user urls for userID %s: %v", userID, err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return c.NoContent(http.StatusAccepted)
	}
}
