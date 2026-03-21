package repository

import (
	"context"
	"encoding/json"
	"io"
	"os"

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
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	for err == nil {
		shortening, err = s.inMemory.Put(context.Background(), *shortening)
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

func (s *inFileStorage) readFromFile() (*model.Shortening, error) {
	var shortening *model.Shortening
	if err := s.decoder.Decode(&shortening); err != nil {
		return nil, err
	}
	return shortening, nil
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

func (s *inFileStorage) Get(ctx context.Context, shortURL string) (*model.Shortening, error) {
	return s.inMemory.Get(ctx, shortURL)
}
