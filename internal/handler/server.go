package handler

import (
	"context"
	"net/http"
	"slices"

	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CloseFunc func(context.Context) error

type Server struct {
	e         *echo.Echo
	shortener *service.Service
	baseURL   string
}

func NewServer(shortener *service.Service, baseURL string) *Server {
	s := &Server{
		shortener: shortener,
		baseURL:   baseURL,
	}
	s.setupRouter()

	return s
}

func (s *Server) setupRouter() {
	s.e = echo.New()
	s.e.HideBanner = true

	s.e.Pre(middleware.RemoveTrailingSlash())
	s.e.Use(logger.WithLogging())
	s.e.Use(middleware.GzipWithConfig(makeGzipConfig()))

	s.e.POST("/", HandleShorten(s.shortener, s.baseURL))
	s.e.POST("/api/shorten", HandleAPIShorten(s.shortener, s.baseURL))
	s.e.GET("/:identifier", HandleRedirect(s.shortener))
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

func makeGzipConfig() middleware.GzipConfig {
	return middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return !slices.Contains([]string{"application/json", "text/html"}, c.Request().Header.Get(echo.HeaderContentType))
		},
	}
}
