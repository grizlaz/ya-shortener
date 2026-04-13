package model

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrIdentifierExists = errors.New("identifier already exists")
	ErrInvalidURL       = errors.New("invalid url")
	ErrConflict         = errors.New("url already exists")
	ErrUnauthorized     = errors.New("unauthorized")
)

type Shortening struct {
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	ID          int       `json:"id"`
	UserID      uuid.UUID `json:"-"`
}

type ShortenRequestBatch struct {
	ID  string `json:"correlation_id"`
	URL string `json:"original_url"`
}

type User struct {
	ID uuid.UUID `json:"id"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	User `json:"user"`
}
