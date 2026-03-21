package model

import (
	"errors"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrIdentifierExists = errors.New("identifier already exists")
	ErrInvalidURL       = errors.New("invalid url")
)

type Shortening struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	ID          int    `json:"id"`
}
