package checkr

import (
	"encoding/json"
	"time"
)

// IdentityDocumentVerification ...
// https://docs.checkr.com/#tag/Identity-Document-Verification
type IdentityDocumentVerification struct {
	ID             string     `json:"id"`
	Object         string     `json:"object"`
	URI            string     `json:"uri"`
	Status         string     `json:"status"`
	CreatedAt      *time.Time `json:"created_at"`
	CompletedAt    *time.Time `json:"completed_at"`
	TurnaroundTime int        `json:"turnaround_time"`
	Verified       bool       `json:"verified"`
	CaptureURL     string     `json:"capture_url"`
	Provider       string     `json:"provider"`
	DocumentIds    []string   `json:"document_ids"`
	Checks         struct {
		CandidateDataMatch []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"candidate_data_match"`
		FaceMatch []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"face_match"`
		DocumentValidation []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"document_validation"`
	} `json:"checks"`
	ExtractedData struct {
		FirstName             string `json:"first_name"`
		MiddleName            string `json:"middle_name"`
		LastName              string `json:"last_name"`
		FullName              string `json:"full_name"`
		Dob                   string `json:"dob"`
		Gender                string `json:"gender"`
		DocumentType          string `json:"document_type"`
		DocumentNumber        string `json:"document_number"`
		DocumentExpiration    string `json:"document_expiration"`
		DocumentIssuer        string `json:"document_issuer"`
		DocumentIssuerRegion  string `json:"document_issuer_region"`
		DocumentIssuerCountry string `json:"document_issuer_country"`
	} `json:"extracted_data"`
}

// Unmarshal ...
func (i *IdentityDocumentVerification) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &i)
}
