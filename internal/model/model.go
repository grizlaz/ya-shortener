package model

import (
	"errors"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrIdentifierExists = errors.New("identifier already exists")
	ErrInvalidURL       = errors.New("invalid url")
	ErrConflict         = errors.New("url already exists")
)

type Shortening struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	ID          int    `json:"id"`
}

type ShortenRequestBatch struct {
	ID  string `json:"correlation_id"`
	URL string `json:"original_url"`
}
