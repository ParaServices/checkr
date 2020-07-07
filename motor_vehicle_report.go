package checkr

import(
	"time"
	"encoding/json"
)

// MotorVehicleReport ...
// https://docs.checkr.com/#operation/getMotorVehicleReport
type MotorVehicleReport struct {
	ID                    string      `json:"id,omitempty"`
	Object                string      `json:"object,omitempty"`
	URI                   string      `json:"uri,omitempty"`
	Status                string      `json:"status,omitempty"`
	CreatedAt             time.Time   `json:"created_at,omitempty"`
	CompletedAt           time.Time   `json:"completed_at,omitempty"`
	TurnaroundTime        int         `json:"turnaround_time,omitempty"`
	FullName              string      `json:"full_name,omitempty"`
	LicenseNumber         string      `json:"license_number,omitempty"`
	LicenseState          string      `json:"license_state,omitempty"`
	PreviousLicenseNumber string      `json:"previous_license_number,omitempty"`
	PreviousLicenseState  string      `json:"previous_license_state,omitempty"`
	LicenseStatus         string      `json:"license_status,omitempty"`
	LicenseType           string      `json:"license_type,omitempty"`
	LicenseClass          string      `json:"license_class,omitempty"`
	Dob                   string      `json:"dob,omitempty"`
	ExpirationDate        string      `json:"expiration_date,omitempty"`
	CovidExtension        bool        `json:"covid_extension,omitempty"`
	IssuedDate            string      `json:"issued_date,omitempty"`
	FirstIssuedDate       string      `json:"first_issued_date,omitempty"`
	InferredIssuedDate    string `json:"inferred_issued_date,omitempty"`
	Restrictions          []string    `json:"restrictions,omitempty"`
	CustomRules           []string    `json:"custom_rules,omitempty"`
	NotFound              bool        `json:"not_found,omitempty"`
	ExperienceFailed      bool        `json:"experience_failed,omitempty"`
	PrivilegeToDrive      string      `json:"privilege_to_drive,omitempty"`
	Accidents             []struct {
		AccidentDate          string      `json:"accident_date,omitempty"`
		Description           string      `json:"description,omitempty"`
		City                  string `json:"city,omitempty"`
		County                string      `json:"county,omitempty"`
		State                 string `json:"state,omitempty"`
		OrderNumber           string      `json:"order_number,omitempty"`
		Points                string `json:"points,omitempty"`
		VehicleSpeed          string `json:"vehicle_speed,omitempty"`
		ReinstatementDate     string `json:"reinstatement_date,omitempty"`
		ActionTaken           string      `json:"action_taken,omitempty"`
		TicketNumber          string `json:"ticket_number,omitempty"`
		EnforcingAgency       string      `json:"enforcing_agency,omitempty"`
		Jurisdiction          string `json:"jurisdiction,omitempty"`
		Severity              string `json:"severity,omitempty"`
		ViolationNumber       string `json:"violation_number,omitempty"`
		LicensePlate          string      `json:"license_plate,omitempty"`
		FineAmount            string `json:"fine_amount,omitempty"`
		StateCode             string `json:"state_code,omitempty"`
		AcdCode               string `json:"acd_code,omitempty"`
		InjuryAccident        bool        `json:"injury_accident,omitempty"`
		FatalityAccident      bool        `json:"fatality_accident,omitempty"`
		FatalityCount         int         `json:"fatality_count,omitempty"`
		InjuryCount           int         `json:"injury_count,omitempty"`
		VehiclesInvolvedCount int         `json:"vehicles_involved_count,omitempty"`
		ReportNumber          string `json:"report_number,omitempty"`
		PolicyNumber          string `json:"policy_number,omitempty"`
		Group                 string      `json:"group,omitempty"`
	} `json:"accidents,omitempty"`
	Violations []struct {
		Type           string      `json:"type,omitempty"`
		IssuedDate     string      `json:"issued_date,omitempty"`
		ConvictionDate string      `json:"conviction_date,omitempty"`
		Description    string      `json:"description,omitempty"`
		Points         int         `json:"points,omitempty"`
		City           string `json:"city,omitempty"`
		County         string      `json:"county,omitempty"`
		State          string      `json:"state,omitempty"`
		TicketNumber   string      `json:"ticket_number,omitempty"`
		Disposition    string `json:"disposition,omitempty"`
		Category       string `json:"category,omitempty"`
		CourtName      string `json:"court_name,omitempty"`
		AcdCode        string `json:"acd_code,omitempty"`
		StateCode      string `json:"state_code,omitempty"`
		Docket         string `json:"docket,omitempty"`
	} `json:"violations,omitempty"`
	Suspensions []struct {
		Description string `json:"description,omitempty"`
		StartDate   string `json:"start_date,omitempty"`
		EndDate     string `json:"end_date,omitempty"`
		State       string `json:"state,omitempty"`
	} `json:"suspensions,omitempty"`
}

// Unmarshal ...
func (m *MotorVehicleReport) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &m)
}