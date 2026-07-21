package handler

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grizlaz/ya-shortener/internal/config"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type shortenerGetStats interface {
	GetUrlsCount(ctx context.Context) (int, error)
	GetUsersCount(ctx context.Context) (int, error)
}

type statsResponse struct {
	Urls  int `json:"urls"`
	Users int `json:"users"`
}

func HandleGetStats(shortener shortenerGetStats) echo.HandlerFunc {
	return func(c echo.Context) error {
		cfg := config.Get()
		if cfg.TrustedSubnet == "" || cfg.Subnet == nil {
			logger.Log.Info("empty cfg.TrustedSubnet")
			return c.NoContent(http.StatusForbidden)
		}
		userIP := net.ParseIP(c.RealIP())
		if correctSubnet := cfg.Subnet.Contains(userIP); !correctSubnet {
			logger.Log.Info("wrong usersubnet", zap.String("userIP", userIP.String()))
			return c.NoContent(http.StatusForbidden)
		}
		urlsCount, err := shortener.GetUrlsCount(c.Request().Context())
		if err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		usersCount, err := shortener.GetUsersCount(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		response := statsResponse{
			Urls:  urlsCount,
			Users: usersCount,
		}
		return c.JSON(http.StatusOK, response)
	}
}

// func getRequestIP(c echo.Context) (net.IP, error) {

// }
