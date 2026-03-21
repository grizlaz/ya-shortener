package repository

import (
	"context"

	"github.com/grizlaz/ya-shortener/internal/model"
)

type inMemory struct {
	m map[string]*model.Shortening
}

func NewInMemory() *inMemory {
	return &inMemory{m: make(map[string]*model.Shortening)}
}

func (i *inMemory) Put(_ context.Context, shortening model.Shortening) (*model.Shortening, error) {
	if _, exist := i.m[shortening.ShortURL]; exist {
		return nil, model.ErrIdentifierExists
	}

	i.m[shortening.ShortURL] = &shortening

	return &shortening, nil
}

func (i *inMemory) Get(_ context.Context, shortURL string) (*model.Shortening, error) {
	v, ok := i.m[shortURL]
	if !ok {
		return nil, model.ErrNotFound
	}

	return v, nil
}
