// Package discord contains the Discord-specific helpers and tools.
// See https://discord.com/developers/docs/.
package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Webhook represents a Discord Webhook.
// See https://discord.com/developers/docs/resources/webhook.
type Webhook struct {
	ID    string
	Token string
}

// https://discord.com/developers/docs/resources/webhook#execute-webhook-jsonform-params
// Currently only use a subset that we're actually interested in sending
type WebhookBody struct {
	Content string `json:"content"`
}

// URL returns the fully-formatted webhook URL.
func (w *Webhook) URL() string {
	return fmt.Sprintf("https://discord.com/api/webhooks/%s/%s", w.ID, w.Token)
}

// Excute implements the Exectute verb on a webhook.
// https://discord.com/developers/docs/resources/webhook#execute-webhook
func (w *Webhook) Execute(body WebhookBody) error {
	bodySer, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resp, err := http.Post(w.URL(), "application/json", bytes.NewBuffer(bodySer))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("request returned %d: %v", resp.StatusCode, resp)
	}
	return nil
}
