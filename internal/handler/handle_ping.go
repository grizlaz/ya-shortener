package handler

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/grizlaz/ya-shortener/internal/logger"
)

// Handler для проверки связи с БД.
func HandlePing(ctx context.Context, db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if derr := c.Request().Body.Close(); derr != nil {
				logger.Log.Sugar().Errorf("error defer c.Request().Body.Close(): %v", derr)
			}
		}()
		err := db.PingContext(ctx)
		if err != nil {
			logger.Log.Sugar().Infof("error ping db: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusOK)
	}
}
