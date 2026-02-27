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
	Identifier  string
	OriginalURL string
}
