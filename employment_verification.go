package checkr

import (
	"encoding/json"
	"time"
)

// EmploymentVerification ...
// https://docs.checkr.com/#operation/getEmploymentVerification
type EmploymentVerification struct {
	ID             string    `json:"id,omitempty"`
	Object         string    `json:"object,omitempty"`
	URI            string    `json:"uri,omitempty"`
	Status         string    `json:"status,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	CompletedAt    time.Time `json:"completed_at,omitempty"`
	TurnaroundTime int       `json:"turnaround_time,omitempty"`
	Records        []struct {
		ID     string `json:"id,omitempty"`
		Result struct {
			StartDate struct {
				Verified bool   `json:"verified,omitempty"`
				Comments string `json:"comments,omitempty"`
				Ignored  string `json:"ignored,omitempty"`
			} `json:"start_date,omitempty"`
			EndDate struct {
				Verified bool   `json:"verified,omitempty"`
				Comments string `json:"comments,omitempty"`
				Ignored  string `json:"ignored,omitempty"`
			} `json:"end_date,omitempty"`
			Position struct {
				Verified bool   `json:"verified,omitempty"`
				Comments string `json:"comments,omitempty"`
				Ignored  string `json:"ignored,omitempty"`
			} `json:"position,omitempty"`
			ContractType struct {
				Verified bool   `json:"verified,omitempty"`
				Comments string `json:"comments,omitempty"`
				Ignored  string `json:"ignored,omitempty"`
			} `json:"contract_type,omitempty"`
			Salary struct {
				Verified bool   `json:"verified,omitempty"`
				Comments string `json:"comments,omitempty"`
				Ignored  string `json:"ignored,omitempty"`
			} `json:"salary,omitempty"`
			Questions []struct {
				SortNumber int    `json:"sort_number,omitempty"`
				Text       string `json:"text,omitempty"`
				Response   string `json:"response,omitempty"`
			} `json:"questions,omitempty"`
		} `json:"result,omitempty"`
		Employer struct {
			ID           string `json:"id,omitempty"`
			Object       string `json:"object,omitempty"`
			URI          string `json:"uri,omitempty"`
			CandidateID  string `json:"candidate_id,omitempty"`
			Name         string `json:"name,omitempty"`
			Position     string `json:"position,omitempty"`
			Salary       int    `json:"salary,omitempty"`
			ContractType string `json:"contract_type,omitempty"`
			DoNotContact bool   `json:"do_not_contact,omitempty"`
			StartDate    string `json:"start_date,omitempty"`
			EndDate      string `json:"end_date,omitempty"`
			EmployerURL  string `json:"employer_url,omitempty"`
			Address      struct {
				Street  string `json:"street,omitempty"`
				Unit    int    `json:"unit,omitempty"`
				City    string `json:"city,omitempty"`
				State   string `json:"state,omitempty"`
				Zipcode string `json:"zipcode,omitempty"`
				Country string `json:"country,omitempty"`
			} `json:"address,omitempty"`
			Manager struct {
				Name  string `json:"name,omitempty"`
				Title string `json:"title,omitempty"`
				Email string `json:"email,omitempty"`
				Phone string `json:"phone,omitempty"`
			} `json:"manager,omitempty"`
		} `json:"employer,omitempty"`
		Events []struct {
			Text      string    `json:"text,omitempty"`
			CreatedAt time.Time `json:"created_at,omitempty"`
			Type      string    `json:"type,omitempty"`
		} `json:"events,omitempty"`
		Status string `json:"status,omitempty"`
	} `json:"records,omitempty"`
}

// Unmarshal ...
func (e *EmploymentVerification) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &e)
}
