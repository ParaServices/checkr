package checkr

// Screenings ...
type Screenings struct {
	SSNTrace                     *SSNTrace                     `json:"ssn_trace,omitempty"`
	SexOffenderSearch            *SexOffenderSearch            `json:"sex_offender_searches,omitempty"`
	GlobalWatchListSearch        *GlobalWatchListSearch        `json:"global_watchlist_searches,omitempty"`
	NationalCriminalSearch       *NationalCriminalSearch       `json:"national_criminal_searches,omitempty"`
	FederalCriminalSearch        *FederalCriminalSearch        `json:"federal_criminal_searches,omitempty"`
	CountryCriminalSearches      []CountryCriminalSearch       `json:"country_criminal_searches,omitempty"`
	StateCriminalSearch          []StateCriminalSearch         `json:"state_criminal_searches,omitempty"`
	MotorVehicleReport           *MotorVehicleReport           `json:"motor_vehicle_report,omitempty"`
	EducationVerification        *EducationVerification        `json:"education_verification,omitempty"`
	EmploymentVerification       *EmploymentVerification       `json:"employment_verification,omitempty"`
	IdentityDocumentVerification *IdentityDocumentVerification `json:"identity_document_verification,omitempty"`
}
