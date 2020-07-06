package checkr

import (
	"encoding/json"
	"time"
)

// SexOffenderSearch ...
// https://docs.checkr.com/#tag/Sex-Offender-Registry-Search
type SexOffenderSearch struct {
	ID             string    `json:"id"`
	Object         string    `json:"object"`
	URI            string    `json:"uri"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	CompletedAt    time.Time `json:"completed_at"`
	TurnaroundTime int       `json:"turnaround_time"`
	Records        []struct {
		Registry          string `json:"registry"`
		FullName          string `json:"full_name"`
		Age               int    `json:"age"`
		Dob               string `json:"dob"`
		RegistrationStart string `json:"registration_start"`
		RegistrationEnd   string `json:"registration_end"`
		State             string `json:"state"`
	} `json:"records"`
}

// Unmarshal ...
func (s *SexOffenderSearch) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &s)
}
