package checkr

import(
	"time"
	"encoding/json"
)

// MotorVehicleReport ...
// https://docs.checkr.com/#operation/getMotorVehicleReport
type MotorVehicleReport struct {
	ID                    string      `json:"id"`
	Object                string      `json:"object"`
	URI                   string      `json:"uri"`
	Status                string      `json:"status"`
	CreatedAt             time.Time   `json:"created_at"`
	CompletedAt           time.Time   `json:"completed_at"`
	TurnaroundTime        int         `json:"turnaround_time"`
	FullName              string      `json:"full_name"`
	LicenseNumber         string      `json:"license_number"`
	LicenseState          string      `json:"license_state"`
	PreviousLicenseNumber string      `json:"previous_license_number"`
	PreviousLicenseState  string      `json:"previous_license_state"`
	LicenseStatus         string      `json:"license_status"`
	LicenseType           string      `json:"license_type"`
	LicenseClass          string      `json:"license_class"`
	Dob                   string      `json:"dob"`
	ExpirationDate        string      `json:"expiration_date"`
	CovidExtension        bool        `json:"covid_extension"`
	IssuedDate            string      `json:"issued_date"`
	FirstIssuedDate       string      `json:"first_issued_date"`
	InferredIssuedDate    interface{} `json:"inferred_issued_date"`
	Restrictions          []string    `json:"restrictions"`
	CustomRules           []string    `json:"custom_rules"`
	NotFound              bool        `json:"not_found"`
	ExperienceFailed      bool        `json:"experience_failed"`
	PrivilegeToDrive      string      `json:"privilege_to_drive"`
	Accidents             []struct {
		AccidentDate          string      `json:"accident_date"`
		Description           string      `json:"description"`
		City                  interface{} `json:"city"`
		County                string      `json:"county"`
		State                 interface{} `json:"state"`
		OrderNumber           string      `json:"order_number"`
		Points                interface{} `json:"points"`
		VehicleSpeed          interface{} `json:"vehicle_speed"`
		ReinstatementDate     interface{} `json:"reinstatement_date"`
		ActionTaken           string      `json:"action_taken"`
		TicketNumber          interface{} `json:"ticket_number"`
		EnforcingAgency       string      `json:"enforcing_agency"`
		Jurisdiction          interface{} `json:"jurisdiction"`
		Severity              interface{} `json:"severity"`
		ViolationNumber       interface{} `json:"violation_number"`
		LicensePlate          string      `json:"license_plate"`
		FineAmount            interface{} `json:"fine_amount"`
		StateCode             interface{} `json:"state_code"`
		AcdCode               interface{} `json:"acd_code"`
		InjuryAccident        bool        `json:"injury_accident"`
		FatalityAccident      bool        `json:"fatality_accident"`
		FatalityCount         int         `json:"fatality_count"`
		InjuryCount           int         `json:"injury_count"`
		VehiclesInvolvedCount int         `json:"vehicles_involved_count"`
		ReportNumber          interface{} `json:"report_number"`
		PolicyNumber          interface{} `json:"policy_number"`
		Group                 string      `json:"group"`
	} `json:"accidents"`
	Violations []struct {
		Type           string      `json:"type"`
		IssuedDate     string      `json:"issued_date"`
		ConvictionDate string      `json:"conviction_date"`
		Description    string      `json:"description"`
		Points         int         `json:"points"`
		City           interface{} `json:"city"`
		County         string      `json:"county"`
		State          string      `json:"state"`
		TicketNumber   string      `json:"ticket_number"`
		Disposition    interface{} `json:"disposition"`
		Category       interface{} `json:"category"`
		CourtName      interface{} `json:"court_name"`
		AcdCode        interface{} `json:"acd_code"`
		StateCode      interface{} `json:"state_code"`
		Docket         interface{} `json:"docket"`
	} `json:"violations"`
	Suspensions []struct {
		Description string `json:"description"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		State       string `json:"state"`
	} `json:"suspensions"`
}

// Unmarshal ...
func (m *MotorVehicleReport) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &m)
}