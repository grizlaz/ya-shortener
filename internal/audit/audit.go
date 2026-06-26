package audit

import (
	"go.uber.org/zap"

	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
)

// generate:reset
type Audit struct {
	observers map[string]*Observer
}

func NewAudit() *Audit {
	return &Audit{
		observers: make(map[string]*Observer),
	}
}

func (a *Audit) Send(message model.AuditMessage) {
	var err error
	if a == nil {
		return
	}
	for _, o := range a.observers {
		if err = o.SendAuditMessage(message); err != nil {
			logger.Log.Error("error send audit message", zap.String("observerID", o.GetID()), zap.Error(err))
		}
	}
}

func (a *Audit) Register(o *Observer) {
	a.observers[o.GetID()] = o
}

func (a *Audit) Deregister(o Observer) {
	delete(a.observers, o.GetID())
}
