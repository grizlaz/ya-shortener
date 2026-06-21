package model

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	ErrNotFound         = errors.New("not found")                 // Ошибка в случае если ничего не найдено
	ErrIdentifierExists = errors.New("identifier already exists") // Ошибка конфликта создаваемых данных
	ErrInvalidURL       = errors.New("invalid url")               // Ошибка некорректной URL
	ErrConflict         = errors.New("url already exists")        // Ошибка при повторном сокращении ссылки
	ErrUnauthorized     = errors.New("unauthorized")              // Ошибка если пользователь не авторизован
	ErrURLDeleted       = errors.New("url deleted")               // Ошибка при переходе по уже удаленной ссылке
)

// Shortening предназначен для работы с сокращением ссылки
type Shortening struct {
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	ID          int       `json:"id"`
	IsDeleted   bool      `json:"is_deleted"`
	UserID      uuid.UUID `json:"-"`
}

// ShortenRequestBatch нужен массового сокращения ссылок
type ShortenRequestBatch struct {
	ID  string `json:"correlation_id"`
	URL string `json:"original_url"`
}

// User отвечает за работу jwt
type User struct {
	ID uuid.UUID `json:"id"`
}

// UserClaims отвечает за работу jwt
type UserClaims struct {
	jwt.RegisteredClaims
	User `json:"user"`
}

// DeleteUrls нужен для удаления пользовательских ссылок
type DeleteUrls struct {
	UserID uuid.UUID
	Urls   *[]string
}

// AuditMessage отвечает за структуру сообщений в аудит
type AuditMessage struct {
	TS     int64  `json:"ts"`
	Action string `json:"action"`
	UserID string `json:"user_id,omitempty"`
	URL    string `json:"url"`
}
