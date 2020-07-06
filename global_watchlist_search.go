package checkr

import(
	"time"
	"encoding/json"
)
// GlobalWatchListSearch ...
// https://docs.checkr.com/#tag/Global-Watchlist-Search
type GlobalWatchListSearch struct {
	ID             string    `json:"id"`
	Object         string    `json:"object"`
	URI            string    `json:"uri"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	CompletedAt    time.Time `json:"completed_at"`
	TurnaroundTime int       `json:"turnaround_time"`
	Records        []struct {
		ID                string      `json:"id"`
		CaseNumber        string      `json:"case_number"`
		FileDate          interface{} `json:"file_date"`
		ArrestingAgency   string      `json:"arresting_agency"`
		CourtJurisdiction interface{} `json:"court_jurisdiction"`
		CourtOfRecord     interface{} `json:"court_of_record"`
		FullName          string      `json:"full_name"`
		Dob               string      `json:"dob"`
		Yob               int         `json:"yob"`
		County            string      `json:"county"`
		State             string      `json:"state"`
		Address           struct {
			Street  string `json:"street"`
			Unit    int    `json:"unit"`
			City    string `json:"city"`
			State   string `json:"state"`
			Zipcode string `json:"zipcode"`
			Country string `json:"country"`
		} `json:"address"`
		Charges []struct {
			Charge          string      `json:"charge"`
			ChargeType      interface{} `json:"charge_type"`
			ChargeID        interface{} `json:"charge_id"`
			Classification  string      `json:"classification"`
			Deposition      interface{} `json:"deposition"`
			Defendant       interface{} `json:"defendant"`
			Plaintiff       interface{} `json:"plaintiff"`
			Sentence        string      `json:"sentence"`
			Disposition     string      `json:"disposition"`
			ProbationStatus interface{} `json:"probation_status"`
			OffenseDate     string      `json:"offense_date"`
			DepositionDate  string      `json:"deposition_date"`
			ArrestDate      interface{} `json:"arrest_date"`
			ChargeDate      interface{} `json:"charge_date"`
			SentenceDate    interface{} `json:"sentence_date"`
			DispositionDate string      `json:"disposition_date"`
			ConvictionDate  string      `json:"conviction_date"`
			ReleaseDate     string      `json:"release_date"`
			NextCourtDate   string      `json:"next_court_date"`
			Court           string      `json:"court"`
			Plea            interface{} `json:"plea"`
			Assessment      interface{} `json:"assessment"`
			PrisonTime      string      `json:"prison_time"`
			JailTime        string      `json:"jail_time"`
			ProbationTime   interface{} `json:"probation_time"`
			Restitution     string      `json:"restitution"`
		} `json:"charges"`
	} `json:"records"`
}


// Unmarshal ...
func (g *GlobalWatchListSearch) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &g)
}