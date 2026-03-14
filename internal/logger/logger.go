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

// func WithLogging(h http.HandlerFunc) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()

// 		responseData := &responseData{0, 0}
// 		lw := &loggingResponseWriter{
// 			ResponseWriter: w,
// 			responseData:   responseData,
// 		}

// 		h.ServeHTTP(lw, r)

// 		duration := time.Since(start)
// 		Log.Info(
// 			"HTTP request",
// 			zap.String("URI", r.RequestURI),
// 			zap.String("method", r.Method),
// 			zap.Duration("duration", duration),
// 			zap.Int("statusCode", responseData.status),
// 			zap.Int("size", responseData.size),
// 		)
// 	})
// }

func WithLogging() echo.MiddlewareFunc {
	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	start := time.Now()

	// 	responseData := &responseData{0, 0}
	// 	lw := &loggingResponseWriter{
	// 		ResponseWriter: w,
	// 		responseData:   responseData,
	// 	}

	// 	h.ServeHTTP(lw, r)

	// 	duration := time.Since(start)
	// 	Log.Info(
	// 		"HTTP request",
	// 		zap.String("URI", r.RequestURI),
	// 		zap.String("method", r.Method),
	// 		zap.Duration("duration", duration),
	// 		zap.Int("statusCode", responseData.status),
	// 		zap.Int("size", responseData.size),
	// 	)
	// })
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

// type responseData struct {
// 	status int
// 	size   int
// }

// type loggingResponseWriter struct {
// 	http.ResponseWriter
// 	responseData *responseData
// }

// func (w *loggingResponseWriter) Write(b []byte) (int, error) {
// 	size, err := w.ResponseWriter.Write(b)
// 	w.responseData.size += size
// 	return size, err
// }

// func (w *loggingResponseWriter) WriteHeader(statusCode int) {
// 	w.ResponseWriter.WriteHeader(statusCode)
// 	w.responseData.status = statusCode
// }
