package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/config"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/labstack/echo/v4"
)

const cookieName = "Authorization"

func WithJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, err := getUserID(c)
			if err != nil && errors.Is(err, model.ErrUnauthorized) && strings.Contains(c.Request().URL.Path, "api/user/urls") {
				return c.NoContent(http.StatusUnauthorized)
			}
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err)
			}
			token, err := makeJWT(userID)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err)
			}
			cookie := &http.Cookie{
				Name:     cookieName,
				Value:    token,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
			}
			c.Request().AddCookie(cookie)

			c.SetCookie(cookie)

			return next(c)
		}
	}
}

func makeJWT(ID uuid.UUID) (string, error) {
	cfg := config.Get()
	claims := model.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.TokenExp)),
		},
		User: model.User{
			ID: ID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(cfg.SecretKey)
}

func getUserID(c echo.Context) (uuid.UUID, error) {
	userCookie, err := c.Cookie(cookieName)
	if err != nil && err == http.ErrNoCookie {
		if err == http.ErrNoCookie {
			return uuid.New(), nil
		}
		return uuid.Nil, err
	}
	tokenString := userCookie.Value
	cfg := config.Get()

	claims := &model.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return cfg.SecretKey, nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	if claims.User.ID == uuid.Nil {
		return uuid.Nil, model.ErrUnauthorized
	}

	if !token.Valid {
		logger.Log.Info("invalid token")
		return uuid.Nil, err
	}

	return claims.User.ID, nil
}
