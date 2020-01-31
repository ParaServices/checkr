package checkr

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Invitiation ...
type Invitation struct {
	ID            string     `json:"id,omitempty"`
	Object        string     `json:"object,omitempty"`
	URI           string     `json:"uri,omitempty"`
	InvitationURL string     `json:"invitation_url,omitempty"`
	Status        string     `json:"status,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	ExpiresAt     *time.Time `json:"expires_at,omitempty"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	Package       string     `json:"package,omitempty"`
	CandidateID   string     `json:"candidate_id,omitempty"`
	ReportID      string     `json:"report_id,omitempty"`
}

// CreateInvitationRequest ...
type CreateInvitationRequest struct {
	CandidateID string `json:"candidate_id,omitempty"`
	Package     string `json:"package,omitempty"`
}

const createInvitiation = "/invitations"

// CreateInvitation ...
func (c *Client) CreateInvitation(reqPayload *CreateInvitationRequest) (*Invitation, error) {
	rel, err := url.Parse(createInvitiation)
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
	req.Header.Add("Conent-Type", "application/json")
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

	createResp := &Invitation{}
	err = json.Unmarshal(b, createResp)
	if err != nil {
		return nil, err
	}

	return createResp, nil
}
