// Модуль handler обрабатывает все входящие запросы к сервису.
package handler

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/pprof"
	"slices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/grizlaz/ya-shortener/internal/audit"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/service"
)

type Server struct {
	e         *echo.Echo
	shortener *service.Service
	baseURL   string
	db        *sql.DB
	audit     *audit.Audit
}

func NewServer(shortener *service.Service, baseURL string, db *sql.DB, audit *audit.Audit) *Server {
	s := &Server{
		shortener: shortener,
		baseURL:   baseURL,
		db:        db,
		audit:     audit,
	}
	s.setupRouter()

	return s
}

func (s *Server) setupRouter() {
	s.e = echo.New()
	s.e.HideBanner = true

	s.e.Pre(middleware.RemoveTrailingSlash())
	s.e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Skipper: skipper}))
	s.e.Use(middleware.DecompressWithConfig(middleware.DecompressConfig{Skipper: skipper}))
	s.e.Use(logger.WithLogging())
	// s.e.Use(audit.WithAudit(s.audit))
	s.e.Use(WithJWT())

	s.e.POST("/", HandleShorten(s.shortener, s.baseURL, s.audit))
	s.e.POST("/api/shorten", HandleAPIShorten(s.shortener, s.baseURL, s.audit))
	s.e.POST("/api/shorten/batch", HandleAPIShortenBatch(s.shortener, s.baseURL))
	s.e.GET("/api/user/urls", HandleUserUrls(s.shortener, s.baseURL))
	s.e.DELETE("/api/user/urls", HandleDeleteUserUrls(s.shortener))
	s.e.GET("/:identifier", HandleRedirect(s.shortener, s.audit))
	s.e.GET("/ping", HandlePing(context.TODO(), s.db))
	s.e.Any("/debug/pprof", func(c echo.Context) error {
		pprof.Index(c.Response().Writer, c.Request())
		return nil
	})
	s.e.Any("/debug/pprof/cmdline", func(c echo.Context) error {
		pprof.Cmdline(c.Response().Writer, c.Request())
		return nil
	})
	s.e.Any("/debug/pprof/profile", func(c echo.Context) error {
		pprof.Profile(c.Response().Writer, c.Request())
		return nil
	})
	s.e.Any("/debug/pprof/symbol", func(c echo.Context) error {
		pprof.Symbol(c.Response().Writer, c.Request())
		return nil
	})
	s.e.Any("/debug/pprof/trace", func(c echo.Context) error {
		pprof.Trace(c.Response().Writer, c.Request())
		return nil
	})
	s.e.Any("/*", func(c echo.Context) error {
		return c.String(http.StatusBadRequest, "wrong url")
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.e.ServeHTTP(w, r)
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.e.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func skipper(c echo.Context) bool {
	//не убрал text/plain т.к. он передается в тестах для 8 инкремента
	return !slices.Contains([]string{"application/json", "text/plain", "text/html"}, c.Request().Header.Get(echo.HeaderContentType))
}
