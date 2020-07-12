package checkr

import (
	"encoding/json"
	"time"
)

// EducationVerification ...
// https://docs.checkr.com/#operation/getEducationVerification
type EducationVerification struct {
	ID             string     `json:"id,omitempty"`
	Object         string     `json:"object,omitempty"`
	URI            string     `json:"uri,omitempty"`
	Status         string     `json:"status,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	CompletedAt    *time.Time `json:"completed_at,omitempty"`
	TurnaroundTime int        `json:"turnaround_time,omitempty"`
	Records        []struct {
		ID     string `json:"id,omitempty"`
		Result struct {
			Verified bool `json:"verified,omitempty"`
		} `json:"result,omitempty"`
		School struct {
			ID          string     `json:"id,omitempty"`
			Object      string     `json:"object,omitempty"`
			URI         string     `json:"uri,omitempty"`
			CandidateID string     `json:"candidate_id,omitempty"`
			Name        string     `json:"name,omitempty"`
			Degree      string     `json:"degree,omitempty"`
			YearAwarded int        `json:"year_awarded,omitempty"`
			Major       string     `json:"major,omitempty"`
			Phone       string     `json:"phone,omitempty"`
			Minor       string     `json:"minor,omitempty"`
			StartDate   string     `json:"start_date,omitempty"`
			EndDate     *time.Time `json:"end_date,omitempty"`
			Current     bool       `json:"current,omitempty"`
			SchoolURL   string     `json:"school_url,omitempty"`
			Address     struct {
				Street  string `json:"street,omitempty"`
				Unit    int    `json:"unit,omitempty"`
				City    string `json:"city,omitempty"`
				State   string `json:"state,omitempty"`
				Zipcode string `json:"zipcode,omitempty"`
				Country string `json:"country,omitempty"`
			} `json:"address,omitempty"`
		} `json:"school,omitempty"`
		Events []struct {
			Text      string     `json:"text,omitempty"`
			CreatedAt *time.Time `json:"created_at,omitempty"`
			Type      string     `json:"type,omitempty"`
		} `json:"events,omitempty"`
		Status string `json:"status,omitempty"`
	} `json:"records,omitempty"`
}

// Unmarshal ...
func (e *EducationVerification) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &e)
}
