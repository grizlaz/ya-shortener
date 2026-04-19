package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/grizlaz/ya-shortener/internal/model"
)

type inFileStorage struct {
	inMemory *inMemory
	file     *os.File
	encoder  *json.Encoder
	decoder  *json.Decoder
	lastID   int
}

func NewInFile(filename string) (*inFileStorage, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	storage := &inFileStorage{
		inMemory: NewInMemory(),
		file:     file,
		encoder:  json.NewEncoder(file),
		decoder:  json.NewDecoder(file),
		lastID:   0,
	}
	err = storage.loadFromFile()
	if err != nil {
		return nil, err
	}
	return storage, nil
}

func (s *inFileStorage) Close() error {
	return s.file.Close()
}

func (s *inFileStorage) loadFromFile() error {
	shortening, err := s.readFromFile()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	for err == nil {
		shortening, err = s.inMemory.Put(context.Background(), *shortening)
		if err != nil {
			return fmt.Errorf("error put in inMemory storage: %w", err)
		}
		s.lastID = shortening.ID
		shortening, err = s.readFromFile()
	}
	if err != io.EOF {
		return err
	}
	return nil
}

func (s *inFileStorage) writeToFile(_ context.Context, shortening model.Shortening) error {
	return s.encoder.Encode(&shortening)
}

func (s *inFileStorage) writeToFileBatch(_ context.Context, shortens *[]model.Shortening) error {
	var err error
	for _, v := range *shortens {
		err = s.encoder.Encode(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *inFileStorage) readFromFile() (*model.Shortening, error) {
	var shortening *model.Shortening
	if err := s.decoder.Decode(&shortening); err != nil {
		return nil, err
	}
	return shortening, nil
}

func (s *inFileStorage) Get(ctx context.Context, shortURL string) (*model.Shortening, error) {
	return s.inMemory.Get(ctx, shortURL)
}

func (s *inFileStorage) Put(ctx context.Context, shortening model.Shortening) (*model.Shortening, error) {
	s.lastID++
	shortening.ID = s.lastID
	_, err := s.inMemory.Put(ctx, shortening)
	if err != nil {
		return nil, err
	}
	err = s.writeToFile(ctx, shortening)
	if err != nil {
		return nil, err
	}
	return &shortening, nil
}

func (s *inFileStorage) PutBatch(ctx context.Context, shortens *[]model.Shortening) (int64, error) {
	count := int64(0)
	for _, v := range *shortens {
		s.lastID++
		v.ID = s.lastID
		_, err := s.inMemory.Put(ctx, v)
		if err != nil {
			return 0, err
		}
		count++
	}
	err := s.writeToFileBatch(ctx, shortens)
	if err != nil {
		return count, err //TODO нужно ли тут возвращать количество сохраненных строк?
	}
	return count, nil
}

func (s *inFileStorage) GetUserUrls(ctx context.Context, userID uuid.UUID) (*[]model.Shortening, error) {
	return s.inMemory.GetUserUrls(ctx, userID)
}

func (s *inFileStorage) DeleteUserUrls(ctx context.Context, deleteUrls ...model.DeleteUrls) error {
	err := s.inMemory.DeleteUserUrls(ctx, deleteUrls...)
	if err != nil {
		return err
	}
	//TODO не придумал пока как реализовать обновлять отдельные записи, просто - через перезапись всего файла, красиво пока не разобрался как
	return nil
}
