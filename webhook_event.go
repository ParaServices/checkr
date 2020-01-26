package checkr

import (
	"encoding/json"
	"net/url"
)

// WebhookEvent represents the payload form a webhook event
type WebhookEvent struct {
	ID         string           `json:"id,omitempty"`
	Object     string           `json:"object,omitempty"`
	Type       string           `json:"type,omitempty"`
	CreatedAt  string           `json:"created_at,omitempty"`
	WebhookURL string           `json:"webhook_url,omitempty"`
	Data       *json.RawMessage `json:"data,omitempty"`
}

// Unmarshal ...
func (w *WebhookEvent) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &w)
}

func (w *WebhookEvent) ReportURI() (*url.URL, error) {
	type Object struct {
		Object struct {
			ID     string `json:"id"`
			Object string `json:"candidate"`
			URI    string `json:"uri"`
		} `json:"object"`
	}

	o := Object{}
	err := json.Unmarshal([]byte(*w.Data), &o)
	if err != nil {
		return nil, err
	}

	return url.Parse(o.Object.URI)
}
