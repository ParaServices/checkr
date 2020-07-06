package checkr

import (
	"encoding/json"
	"time"
)

// EducationVerification ...
// https://docs.checkr.com/#operation/getEducationVerification
type EducationVerification struct {
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
			Verified bool `json:"verified"`
		} `json:"result"`
		School struct {
			ID          string `json:"id"`
			Object      string `json:"object"`
			URI         string `json:"uri"`
			CandidateID string `json:"candidate_id"`
			Name        string `json:"name"`
			Degree      string `json:"degree"`
			YearAwarded int    `json:"year_awarded"`
			Major       string `json:"major"`
			Phone       string `json:"phone"`
			Minor       string `json:"minor"`
			StartDate   string `json:"start_date"`
			EndDate     string `json:"end_date"`
			Current     bool   `json:"current"`
			SchoolURL   string `json:"school_url"`
			Address     struct {
				Street  string `json:"street"`
				Unit    int    `json:"unit"`
				City    string `json:"city"`
				State   string `json:"state"`
				Zipcode string `json:"zipcode"`
				Country string `json:"country"`
			} `json:"address"`
		} `json:"school"`
		Events []struct {
			Text      string    `json:"text"`
			CreatedAt time.Time `json:"created_at"`
			Type      string    `json:"type"`
		} `json:"events"`
		Status string `json:"status"`
	} `json:"records"`
}

// Unmarshal ...
func (e *EducationVerification) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &e)
}
