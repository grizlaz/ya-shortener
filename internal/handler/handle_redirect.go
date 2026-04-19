package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/labstack/echo/v4"
)

type redirecter interface {
	Redirect(ctx context.Context, identifier string) (string, error)
}

func HandleRedirect(redirecter redirecter) echo.HandlerFunc {
	return func(c echo.Context) error {
		identifier := c.Param("identifier")

		redirectURL, err := redirecter.Redirect(c.Request().Context(), identifier)

		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		if errors.Is(err, model.ErrURLDeleted) {
			return c.NoContent(http.StatusGone)
		}
		if err != nil {
			logger.Log.Sugar().Infof("error getting redirect url for %q: %v", identifier, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}
