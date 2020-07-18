package checkr

import (
	"encoding/json"
	"time"
)

// CountryCriminalSearch
// https://docs.checkr.com/#tag/County-Criminal-Search
type CountryCriminalSearch struct {
	ID                      string     `json:"id,omitempty"`
	Object                  string     `json:"object,omitempty"`
	URI                     string     `json:"uri,omitempty"`
	Status                  string     `json:"status,omitempty"`
	CreatedAt               *time.Time `json:"created_at,omitempty"`
	CompletedAt             *time.Time `json:"completed_at,omitempty"`
	TurnaroundTime          int        `json:"turnaround_time,omitempty"`
	EstimatedCompletionTime *time.Time `json:"estimated_completion_time,omitempty"`
	EstimatedCompletionDate string     `json:"estimated_completion_date,omitempty"`
	County                  string     `json:"county,omitempty"`
	State                   string     `json:"state,omitempty"`
	Records                 []struct {
		ID                string     `json:"id,omitempty"`
		CaseNumber        string     `json:"case_number,omitempty"`
		FileDate          *time.Time `json:"file_date,omitempty"`
		ArrestingAgency   string     `json:"arresting_agency,omitempty"`
		CourtJurisdiction string     `json:"court_jurisdiction,omitempty"`
		CourtOfRecord     string     `json:"court_of_record,omitempty"`
		FullName          string     `json:"full_name,omitempty"`
		Dob               *time.Time `json:"dob,omitempty"`
		Yob               int        `json:"yob,omitempty"`
		County            string     `json:"county,omitempty"`
		State             string     `json:"state,omitempty"`
		Address           struct {
			Street  string `json:"street,omitempty"`
			Unit    int    `json:"unit,omitempty"`
			City    string `json:"city,omitempty"`
			State   string `json:"state,omitempty"`
			Zipcode string `json:"zipcode,omitempty"`
			Country string `json:"country,omitempty"`
		} `json:"address,omitempty"`
		Charges []struct {
			Charge          string     `json:"charge,omitempty"`
			ChargeType      string     `json:"charge_type,omitempty"`
			ChargeID        string     `json:"charge_id,omitempty"`
			Classification  string     `json:"classification,omitempty"`
			Deposition      string     `json:"deposition,omitempty"`
			Defendant       string     `json:"defendant,omitempty"`
			Plaintiff       string     `json:"plaintiff,omitempty"`
			Sentence        string     `json:"sentence,omitempty"`
			Disposition     string     `json:"disposition,omitempty"`
			ProbationStatus string     `json:"probation_status,omitempty"`
			OffenseDate     *time.Time `json:"offense_date,omitempty"`
			DepositionDate  string     `json:"deposition_date,omitempty"`
			ArrestDate      string     `json:"arrest_date,omitempty"`
			ChargeDate      string     `json:"charge_date,omitempty"`
			SentenceDate    string     `json:"sentence_date,omitempty"`
			DispositionDate string     `json:"disposition_date,omitempty"`
			ConvictionDate  string     `json:"conviction_date,omitempty"`
			ReleaseDate     *time.Time `json:"release_date,omitempty"`
			NextCourtDate   string     `json:"next_court_date,omitempty"`
			Court           string     `json:"court,omitempty"`
			Plea            string     `json:"plea,omitempty"`
			Assessment      string     `json:"assessment,omitempty"`
			PrisonTime      string     `json:"prison_time,omitempty"`
			JailTime        string     `json:"jail_time,omitempty"`
			ProbationTime   string     `json:"probation_time,omitempty"`
			Restitution     string     `json:"restitution,omitempty"`
		} `json:"charges,omitempty"`
	} `json:"records,omitempty"`
	FilteredByPositiveAdjudicationRecords []struct {
		ID                string     `json:"id,omitempty"`
		CaseNumber        string     `json:"case_number,omitempty"`
		FileDate          *time.Time `json:"file_date,omitempty"`
		ArrestingAgency   string     `json:"arresting_agency,omitempty"`
		CourtJurisdiction string     `json:"court_jurisdiction,omitempty"`
		CourtOfRecord     string     `json:"court_of_record,omitempty"`
		FullName          string     `json:"full_name,omitempty"`
		Dob               *time.Time `json:"dob,omitempty"`
		Yob               int        `json:"yob,omitempty"`
		County            string     `json:"county,omitempty"`
		State             string     `json:"state,omitempty"`
		Address           struct {
			Street  string `json:"street,omitempty"`
			Unit    int    `json:"unit,omitempty"`
			City    string `json:"city,omitempty"`
			State   string `json:"state,omitempty"`
			Zipcode string `json:"zipcode,omitempty"`
			Country string `json:"country,omitempty"`
		} `json:"address,omitempty"`
		FilteredByPositiveAdjudicationCharges []struct {
			Charge          string     `json:"charge,omitempty"`
			ChargeType      string     `json:"charge_type,omitempty"`
			ChargeID        string     `json:"charge_id,omitempty"`
			Classification  string     `json:"classification,omitempty"`
			Deposition      string     `json:"deposition,omitempty"`
			Defendant       string     `json:"defendant,omitempty"`
			Plaintiff       string     `json:"plaintiff,omitempty"`
			Sentence        string     `json:"sentence,omitempty"`
			Disposition     string     `json:"disposition,omitempty"`
			ProbationStatus string     `json:"probation_status,omitempty"`
			OffenseDate     *time.Time `json:"offense_date,omitempty"`
			DepositionDate  string     `json:"deposition_date,omitempty"`
			ArrestDate      string     `json:"arrest_date,omitempty"`
			ChargeDate      string     `json:"charge_date,omitempty"`
			SentenceDate    string     `json:"sentence_date,omitempty"`
			DispositionDate string     `json:"disposition_date,omitempty"`
			ConvictionDate  string     `json:"conviction_date,omitempty"`
			ReleaseDate     *time.Time `json:"release_date,omitempty"`
			NextCourtDate   string     `json:"next_court_date,omitempty"`
			Court           string     `json:"court,omitempty"`
			Plea            string     `json:"plea,omitempty"`
			Assessment      string     `json:"assessment,omitempty"`
			PrisonTime      string     `json:"prison_time,omitempty"`
			JailTime        string     `json:"jail_time,omitempty"`
			ProbationTime   string     `json:"probation_time,omitempty"`
			Restitution     string     `json:"restitution,omitempty"`
		} `json:"filtered_by_positive_adjudication_charges,omitempty"`
	} `json:"filtered_by_positive_adjudication_records,omitempty"`
}

// Unmarshal ...
func (c *CountryCriminalSearch) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &c)
}
