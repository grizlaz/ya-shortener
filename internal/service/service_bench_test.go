package service_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/google/uuid"

	"github.com/grizlaz/ya-shortener/internal/model"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
)

func BenchmarkShortenBatch(b *testing.B) {
	urlsCount := b.N
	userID := uuid.New()
	service := service.NewService(context.Background(), repository.NewInMemory(urlsCount))
	urls := make([]model.ShortenRequestBatch, 0, urlsCount)
	for i := range urlsCount {
		urls = append(urls, model.ShortenRequestBatch{
			ID:  strconv.Itoa(i),
			URL: fmt.Sprintf("https://practicum.yandex.ru/?%d", i),
		})
	}

	b.ResetTimer()

	for b.Loop() {
		service.ShortenBatch(context.Background(), &urls, userID)
	}
}
