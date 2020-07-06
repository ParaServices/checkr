package checkr

import (
	"encoding/json"
	"time"
)

// SSNTrace ...
// https://docs.checkr.com/#tag/SSN-Trace
type SSNTrace struct {
	ID                  string    `json:"id"`
	Object              string    `json:"object"`
	URI                 string    `json:"uri"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
	CompletedAt         time.Time `json:"completed_at"`
	TurnaroundTime      int       `json:"turnaround_time"`
	Ssn                 string    `json:"ssn"`
	NoData              bool      `json:"no_data"`
	DobMismatch         bool      `json:"dob_mismatch"`
	NameMismatch        bool      `json:"name_mismatch"`
	DataMismatch        bool      `json:"data_mismatch"`
	ThinFile            bool      `json:"thin_file"`
	InvalidIssuanceYear bool      `json:"invalid_issuance_year"`
	DeathIndex          bool      `json:"death_index"`
	SsnAlreadyTaken     bool      `json:"ssn_already_taken"`
	IssuedYear          int       `json:"issued_year"`
	IssuedState         string    `json:"issued_state"`
	Addresses           []struct {
		Street   string `json:"street"`
		Unit     string `json:"unit"`
		City     string `json:"city"`
		State    string `json:"state"`
		Zipcode  string `json:"zipcode"`
		Country  string `json:"country"`
		FromDate string `json:"from_date"`
		ToDate   string `json:"to_date"`
	} `json:"addresses"`
	Aliases []struct {
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
	} `json:"aliases"`
}

// Unmarshal ...
func (s *SSNTrace) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &s)
}
