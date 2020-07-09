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

// Candidate ...
// https://docs.checkr.com/#section/Getting-Started/Create-a-Candidate
type Candidate struct {
	ID                          string      `json:"id,omitempty"`
	CustomID                    string      `json:"custom_id,omitempty"`
	Object                      string      `json:"object,omitempty"`
	URI                         string      `json:"uri,omitempty"`
	CreatedAt                   *time.Time `json:"created_at,omitempty"`
	FirstName                   string      `json:"first_name,omitempty"`
	LastName                    string      `json:"last_name,omitempty"`
	MiddleName                  string      `json:"middle_name,omitempty"`
	MotherMaidenName            string      `json:"mother_maiden_name,omitempty"`
	DOB                         *time.Time  `json:"dob,omitempty"`
	SSN                         string      `json:"ssn,omitempty"`
	Email                       string      `json:"email,omitempty"`
	Zipcode                     string      `json:"zipcode,omitempty"`
	Phone                       string      `json:"phone,omitempty"`
	DriverLicenseState          string      `json:"driver_license_state,omitempty"`
	DriverLicenseNumber         string      `json:"driver_license_number,omitempty"`
	CopyRequested               bool        `json:"copy_requested,omitempty"`
	PreviousDriverLicenseState  string      `json:"previous_driver_license_state,omitempty"`
	PreviousDriverLicenseNumber string      `json:"previous_driver_license_number,omitempty"`
	Adjudication                string      `json:"adjudication,omitempty"`
	NoMiddleName                bool        `json:"no_middle_name,omitempty"`
	ReportIDs                   []string    `json:"report_ids,omitempty"`
	GeoIDs                      []string    `json:"geo_ids,omitempty"`
}

// Unmarshal ...
func (c *Candidate) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &c)
}

type CreateCandidateRequest struct {
	CustomID                    string     `json:"custom_id,omitempty"`
	LastName                    string     `json:"last_name,omitempty"`
	FirstName                   string     `json:"first_name,omitempty"`
	MiddleName                  string     `json:"middle_name,omitempty"`
	MotherMaidenName            string     `json:"mother_maiden_name,omitempty"`
	NoMiddleName                bool       `json:"no_middle_name,omitempty"`
	Email                       string     `json:"email,omitempty"`
	Phone                       string     `json:"phone,omitempty"`
	ZipCode                     string     `json:"zipcode,omitempty"`
	DOB                         *time.Time `json:"dob,omitempty"`
	SSN                         string     `json:"ssn,omitempty"`
	DriverLicenseNumber         string     `json:"driver_license_number,omitempty"`
	DriverLicenseState          string     `json:"driver_license_state,omitempty"`
	PreviousDriverLicenseNumber string     `json:"previous_driver_license_number,omitempty"`
	PreviousDriverLicenseState  string     `json:"previous_driver_license_state,omitempty"`
	CopyRequested               bool       `json:"copy_requested,omitempty"`
	Adjucation                  string     `json:"adjudication,omitempty"`
}

const createCandidatePath = "/v1/candidates"

func (c *Client) CreateCandidate(reqPayload *CreateCandidateRequest) (*Candidate, error) {
	rel, err := url.Parse(createCandidatePath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String())

	if reqPayload.MiddleName == "" && reqPayload.MotherMaidenName == "" {
		reqPayload.NoMiddleName = true
	}

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

	createResp := &Candidate{}
	err = json.Unmarshal(b, createResp)
	if err != nil {
		return nil, err
	}

	return createResp, nil
}
