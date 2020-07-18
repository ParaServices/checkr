package checkr

import (
	"encoding/json"
	"time"
)

// IdentityDocumentVerification ...
// https://docs.checkr.com/#tag/Identity-Document-Verification
type IdentityDocumentVerification struct {
	ID             string     `json:"id,omitempty"`
	Object         string     `json:"object,omitempty"`
	URI            string     `json:"uri,omitempty"`
	Status         string     `json:"status,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	CompletedAt    *time.Time `json:"completed_at,omitempty"`
	TurnaroundTime int        `json:"turnaround_time,omitempty"`
	Verified       bool       `json:"verified,omitempty"`
	CaptureURL     string     `json:"capture_url,omitempty"`
	Provider       string     `json:"provider,omitempty"`
	DocumentIds    []string   `json:"document_ids,omitempty"`
	Checks         struct {
		CandidateDataMatch []struct {
			Name   string `json:"name,omitempty"`
			Status string `json:"status,omitempty"`
		} `json:"candidate_data_match,omitempty"`
		FaceMatch []struct {
			Name   string `json:"name,omitempty"`
			Status string `json:"status,omitempty"`
		} `json:"face_match,omitempty"`
		DocumentValidation []struct {
			Name   string `json:"name,omitempty"`
			Status string `json:"status,omitempty"`
		} `json:"document_validation,omitempty"`
	} `json:"checks,omitempty"`
	ExtractedData struct {
		FirstName             string `json:"first_name,omitempty"`
		MiddleName            string `json:"middle_name,omitempty"`
		LastName              string `json:"last_name,omitempty"`
		FullName              string `json:"full_name,omitempty"`
		Dob                   string `json:"dob,omitempty"`
		Gender                string `json:"gender,omitempty"`
		DocumentType          string `json:"document_type,omitempty"`
		DocumentNumber        string `json:"document_number,omitempty"`
		DocumentExpiration    string `json:"document_expiration,omitempty"`
		DocumentIssuer        string `json:"document_issuer,omitempty"`
		DocumentIssuerRegion  string `json:"document_issuer_region,omitempty"`
		DocumentIssuerCountry string `json:"document_issuer_country,omitempty"`
	} `json:"extracted_data,omitempty"`
}

// Unmarshal ...
func (i *IdentityDocumentVerification) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &i)
}
