package checkr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
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

type Unmarshaler interface {
	Unmarshal(b []byte) error
}

// GetReport ...
func (w *WebhookEvent) GetReport(c *Client, u Unmarshaler) error {
	reportURI, err := w.ReportURI()
	if err != nil {
		return err
	}
	rel, err := url.Parse(reportURI.String())
	if err != nil {
		return err
	}

	baseURL := *c.BaseURL
	baseURL.Path = path.Join(baseURL.Path, rel.String())

	req, err := http.NewRequest(http.MethodPost, baseURL.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return NewError([]int{
			http.StatusCreated,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return u.Unmarshal(b)
}
