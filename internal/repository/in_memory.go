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

func (i *inMemory) Get(_ context.Context, shortURL string) (*model.Shortening, error) {
	v, ok := i.m[shortURL]
	if !ok {
		return nil, model.ErrNotFound
	}

	return v, nil
}

func (i *inMemory) Put(_ context.Context, shortening model.Shortening) (*model.Shortening, error) {
	if _, exist := i.m[shortening.ShortURL]; exist {
		return nil, model.ErrIdentifierExists
	}

	i.m[shortening.ShortURL] = &shortening

	return &shortening, nil
}

func (i *inMemory) PutBatch(ctx context.Context, shortens *[]model.Shortening) (int64, error) {
	count := int64(0)
	for _, v := range *shortens {
		_, err := i.Put(ctx, v)
		if err != nil {
			return count, err //TODO нужно ли тут возвращать количество сохраненных строк?
		}
		count++
	}
	return count, nil
}
