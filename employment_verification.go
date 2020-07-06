package checkr

import (
	"encoding/json"
	"time"
)

// EmploymentVerification ...
// https://docs.checkr.com/#operation/getEmploymentVerification
type EmploymentVerification struct {
	ID             string    `json:"id"`
	Object         string    `json:"object"`
	URI            string    `json:"uri"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	CompletedAt    time.Time `json:"completed_at"`
	TurnaroundTime int       `json:"turnaround_time"`
	Records        []struct {
		ID     string `json:"id"`
		Result struct {
			StartDate struct {
				Verified bool        `json:"verified"`
				Comments string      `json:"comments"`
				Ignored  interface{} `json:"ignored"`
			} `json:"start_date"`
			EndDate struct {
				Verified bool        `json:"verified"`
				Comments string      `json:"comments"`
				Ignored  interface{} `json:"ignored"`
			} `json:"end_date"`
			Position struct {
				Verified bool   `json:"verified"`
				Comments string `json:"comments"`
				Ignored  string `json:"ignored"`
			} `json:"position"`
			ContractType struct {
				Verified bool   `json:"verified"`
				Comments string `json:"comments"`
				Ignored  string `json:"ignored"`
			} `json:"contract_type"`
			Salary struct {
				Verified bool        `json:"verified"`
				Comments string      `json:"comments"`
				Ignored  interface{} `json:"ignored"`
			} `json:"salary"`
			Questions []struct {
				SortNumber int    `json:"sort_number"`
				Text       string `json:"text"`
				Response   string `json:"response"`
			} `json:"questions"`
		} `json:"result"`
		Employer struct {
			ID           string `json:"id"`
			Object       string `json:"object"`
			URI          string `json:"uri"`
			CandidateID  string `json:"candidate_id"`
			Name         string `json:"name"`
			Position     string `json:"position"`
			Salary       int    `json:"salary"`
			ContractType string `json:"contract_type"`
			DoNotContact bool   `json:"do_not_contact"`
			StartDate    string `json:"start_date"`
			EndDate      string `json:"end_date"`
			EmployerURL  string `json:"employer_url"`
			Address      struct {
				Street  string `json:"street"`
				Unit    int    `json:"unit"`
				City    string `json:"city"`
				State   string `json:"state"`
				Zipcode string `json:"zipcode"`
				Country string `json:"country"`
			} `json:"address"`
			Manager struct {
				Name  string `json:"name"`
				Title string `json:"title"`
				Email string `json:"email"`
				Phone string `json:"phone"`
			} `json:"manager"`
		} `json:"employer"`
		Events []struct {
			Text      string    `json:"text"`
			CreatedAt time.Time `json:"created_at"`
			Type      string    `json:"type"`
		} `json:"events"`
		Status string `json:"status"`
	} `json:"records"`
}

// Unmarshal ...
func (e *EmploymentVerification) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &e)
}
