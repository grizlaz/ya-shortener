package repository

import (
	"context"
	"slices"

	"github.com/google/uuid"

	"github.com/grizlaz/ya-shortener/internal/model"
)

type inMemory struct {
	m map[string]*model.Shortening
}

func NewInMemory(cap int) *inMemory {
	return &inMemory{m: make(map[string]*model.Shortening, cap)}
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

func (i *inMemory) GetUserUrls(ctx context.Context, userID uuid.UUID) (*[]model.Shortening, error) {
	shortenings := make([]model.Shortening, 0)
	for _, s := range i.m {
		if s.UserID == userID {
			shortenings = append(shortenings, *s)
		}
	}
	return &shortenings, nil
}

func (i *inMemory) DeleteUserUrls(ctx context.Context, deleteUrls ...model.DeleteUrls) error {
	for _, delUrls := range deleteUrls {
		for _, s := range i.m {
			if s.UserID == delUrls.UserID && slices.Contains(*delUrls.Urls, s.ShortURL) {
				s.IsDeleted = true
				i.m[s.ShortURL] = s
			}
		}
	}
	return nil
}

func (i *inMemory) GetUrlsCount(ctx context.Context) (int, error) {
	// TODO если нужны только актиывные, то логика сложнее будет
	return len(i.m), nil
}

func (i *inMemory) GetUsersCount(ctx context.Context) (int, error) {
	users := make(map[uuid.UUID]struct{})
	for _, s := range i.m {
		users[s.UserID] = struct{}{}
	}
	return len(users), nil
}
