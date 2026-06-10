package audit

import (
	"go.uber.org/zap"

	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/model"
)

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

//Хотел через middleware сделать, но что-то муторно получать action, url
// func WithAudit(auditSrv *Audit) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			err := next(c)
// 			auditMessage := getAuditMessage(c)
// 			if auditMessage != nil {
// 				auditSrv.Send(*auditMessage)
// 			}
// 			return err
// 		}
// 	}
// }

// func getAuditMessage(c echo.Context) *model.AuditMessage {
// 	action := ""
// 	isShorten := c.Request().Method == "POST" && c.Request().URL.Path == ""
// 	isApiShorten := c.Request().Method == "POST" && c.Request().URL.Path == "api/shorten"
// 	if isShorten || isApiShorten {
// 		action = "shorten"
// 	}
// 	isRedirect := !(isShorten || isApiShorten) && c.Request().Method == "GET" && c.Param("identifier") != ""
// 	if isRedirect {
// 		action = "follow"
// 	}
// 	if action == "" {
// 		return nil
// 	}
// 	message := &model.AuditMessage{
// 		Ts:     time.Now().Unix(),
// 		Action: action,
// 	}
// 	return message
// }
