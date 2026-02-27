package repository

import (
	"context"
	"sync"

	"github.com/grizlaz/ya-shortener/internal/model"
)

type inMemory struct {
	m sync.Map
}

func NewInMemory() *inMemory {
	return &inMemory{}
}

func (i *inMemory) Put(_ context.Context, shortening model.Shortening) (*model.Shortening, error) {
	if _, exist := i.m.Load(shortening.Identifier); exist {
		return nil, model.ErrIdentifierExists
	}

	i.m.Store(shortening.Identifier, shortening)

	return &shortening, nil
}

func (i *inMemory) Get(_ context.Context, identifier string) (*model.Shortening, error) {
	v, ok := i.m.Load(identifier)
	if !ok {
		return nil, model.ErrNotFound
	}

	shortening := v.(model.Shortening)

	return &shortening, nil
}
