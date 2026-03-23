package logger

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var Log *zap.Logger = zap.NewNop()

func Initialize(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}
	cfg := zap.NewProductionConfig()

	cfg.Level = lvl

	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	Log = zl
	return nil
}

func WithLogging() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			r := c.Request()

			err := next(c)

			duration := time.Since(start)
			res := c.Response()

			Log.Info(
				"HTTP request",
				zap.String("URI", r.RequestURI),
				zap.String("method", r.Method),
				zap.Duration("duration", duration),
				zap.Int("statusCode", res.Status),
				zap.Int("size", int(res.Size)),
			)
			if err != nil {
				Log.Error("HTTP request error", zap.Error(err))
			}
			return err
		}
	}
}
