package checkr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

// AdverseAction ...
type AdverseAction struct {
	ID                    string     `json:"id,omitempty"`
	Object                string     `json:"object,omitempty"`
	URI                   string     `json:"uri,omitempty"`
	CreatedAt             *time.Time `json:"created_at,omitempty"`
	Status                string     `json:"status,omitempty"`
	ReportID              string     `json:"report_id,omitempty"`
	PostNoticeScheduledAt *time.Time `json:"post_notice_scheduled_at,omitempty"`
	PostNoticeReadyAt     *time.Time `json:"post_notice_ready_at,omitempty"`
	CanceledAt            *time.Time `json:"canceled_at,omitempty"`
	AdverseItems          []struct {
		ID     string `json:"id,omitempty"`
		Object string `json:"object,omitempty"`
		Text   string `json:"text,omitempty"`
	} `json:"adverse_items,omitempty"`
	IndividualizedAssessmentEngaged bool `json:"individualized_assessment_engaged,omitempty"`
}

// Unmarshal ...
func (a *AdverseAction) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &a)
}

// CreateAdverseActionRequest ...
type CreateAdverseActionRequest struct {
	PostNoticeScheduledAt *time.Time `json:"post_notice_scheduled_at,omitempty"`
	AdverseItemIds        []string   `json:"adverse_item_ids,omitempty"`
}

const createAdverseAction = "/v1/reports/%s/adverse_actions"

// CreateAdverseActionRequest ...
func (c *Client) CreateAdverseActionRequest(reportID string, reqPayload *CreateAdverseActionRequest) (*AdverseAction, error) {
	rel, err := url.Parse(fmt.Sprintf(createAdverseAction, reportID))
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String())

	b, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, NewError([]int{
			http.StatusCreated,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	createResp := &AdverseAction{}
	err = json.Unmarshal(b, createResp)
	if err != nil {
		return nil, err
	}

	return createResp, nil
}

const getAdverseAction = "/v1/adverse_actions"

// GetAdverseAction ...
func (c *Client) GetAdverseAction(adverseActionID string) (*AdverseAction, error) {
	rel, err := url.Parse(getAdverseAction)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), adverseActionID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError([]int{
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
		return nil, err
	}

	getResp := &AdverseAction{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}
