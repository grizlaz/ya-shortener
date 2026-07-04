package audit

import (
	"github.com/grizlaz/ya-shortener/internal/model"
)

type auditStorage interface {
	SendAuditMessage(message model.AuditMessage) error
}

// generate:reset
type Observer struct {
	id      string
	storage auditStorage
}

func NewObserver(id string, storage auditStorage) *Observer {
	return &Observer{
		id:      id,
		storage: storage,
	}
}

func (o *Observer) GetID() string {
	return o.id
}

func (o *Observer) SendAuditMessage(message model.AuditMessage) error {
	return o.storage.SendAuditMessage(message)
}
