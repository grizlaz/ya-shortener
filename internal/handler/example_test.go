package handler_test

import (
	"github.com/grizlaz/ya-shortener/internal/audit"
	"github.com/grizlaz/ya-shortener/internal/handler"
	"github.com/grizlaz/ya-shortener/internal/service"
	"github.com/labstack/echo/v4"
)

func Example() {
	shorten := &service.Service{}
	audit := &audit.Audit{}
	baseURL := "baseURL"
	// Инициализируем http framework
	e := echo.New()

	// Добавляем handler обработки запросов на получение сокращенной ссылки текстовой ссылки в теле запроса.
	e.POST("/", handler.HandleShorten(shorten, baseURL, nil))
	// Добавляем handler обработки запросов на получение сокращенной ссылки с json в теле запроса.
	e.POST("/api/shorten", handler.HandleAPIShorten(shorten, baseURL, audit))
	// Добавляем handler с пакетной загрузкой ссылок для сокращения.
	e.POST("/api/shorten/batch", handler.HandleAPIShortenBatch(shorten, baseURL))
	// Добавляем handler получения ссылок конкретного пользователя.
	e.GET("/api/user/urls", handler.HandleUserUrls(shorten, baseURL))
	// Добавляем handler удаления сокращенной ссылки пользователя.
	e.DELETE("/api/user/urls", handler.HandleDeleteUserUrls(shorten))
	// Добавляем handler редирект с сокращенной ссылки на оригинальную.
	e.GET("/:identifier", handler.HandleRedirect(shorten, audit))
}

// func ExampleShorten() {

// }

// func ExampleShortenBatch() {

// }

// func ExampleUserUrls() {

// }

// func ExampleDeleteUserUrls() {

// }

// func ExampleRedirect() {

// }

// func ExamplePing() {

// }
