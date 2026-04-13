package handler

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/labstack/echo/v4"
)

func HandlePing(ctx context.Context, db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer c.Request().Body.Close()
		err := db.PingContext(ctx)
		if err != nil {
			logger.Log.Sugar().Infof("error ping db: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusOK)
	}
}
