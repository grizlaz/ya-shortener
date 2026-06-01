package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grizlaz/ya-shortener/internal/model"
)

type auditClient struct {
	client  *http.Client
	address string
}

func NewAuditClient(address string) *auditClient {
	return &auditClient{
		client:  &http.Client{},
		address: address,
	}
}

func (a *auditClient) SendAuditMessage(message model.AuditMessage) error {
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshal audit message: %w", err)
	}
	resp, err := a.client.Post(a.address, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("error send audit message: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error send audit message: %d", resp.StatusCode)
	}
	return nil
}
