package checkr

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sync"
	"time"
)

// ScreenType ...
type ScreenType int

// all screentypes
const (
	SsnTrace ScreenType = iota
	Sexoffendersearch
	GlobalWatchListsearch
	NationalCriminalsearch
	FederalCriminalsearch
	CountryCriminalsearch
	StateCriminalsearch
	MotorVehiclereport
	Educationverification
	Employmentverification
	IdentityDocumentverification
)

// CreateReportRequest ...
type CreateReportRequest struct {
	Package     string `json:"package,omitempty"`
	CandidateID string `json:"candidate_id,omitempty"`
}

// Report ...
type Report struct {
	ID                               string     `json:"id,omitempty"`
	Object                           string     `json:"object,omitempty"`
	URI                              string     `json:"uri,omitempty"`
	Status                           string     `json:"status,omitempty"`
	CreatedAt                        *time.Time `json:"created_at,omitempty"`
	CompletedAt                      string     `json:"completed_at,omitempty"`
	RevisedAt                        string     `json:"revised_at,omitempty"`
	UpgradedAt                       string     `json:"upgraded_at,omitempty"`
	TurnaroundTime                   int        `json:"turnaround_time,omitempty"`
	DueTime                          *time.Time `json:"due_time,omitempty"`
	Adjudication                     string     `json:"adjudication,omitempty"`
	Package                          string     `json:"package,omitempty"`
	Source                           string     `json:"source,omitempty"`
	CandidateID                      string     `json:"candidate_id,omitempty"`
	SSNTraceID                       string     `json:"ssn_trace_id,omitempty"`
	ArrestSearchID                   string     `json:"arrest_search_id,omitempty"`
	FACISSearchID                    string     `json:"facis_search_id,omitempty"`
	FederalCrimeSearchID             string     `json:"federal_crime_search_id,omitempty"`
	GlobalWatchlistSearchID          string     `json:"global_watchlist_search_id,omitempty"`
	SexOffenderSearchID              string     `json:"sex_offender_search_id,omitempty"`
	NationalCriminalSearchID         string     `json:"national_criminal_search_id,omitempty"`
	MotorVehicleReportID             string     `json:"motor_vehicle_report_id,omitempty"`
	ProgramID                        string     `json:"program_id,omitempty"`
	CountyCriminalSearchIDs          []string   `json:"country_criminal_search_ids,omitempty"`
	PersonalReferenceVerificationIDs []string   `json:"personal_reference_verification_ids,omitempty"`
	StateCriminalSearchIDs           []string   `json:"state_criminal_search_ids,omitempty"`
	DocumentIDs                      []string   `json:"document_ids,omitempty"`
	GeoIDs                           []string   `json:"geo_ids,omitempty"`
	IdentityDocumentVerificationID   string     `json:"identity_document_verification_id"`
}

// Unmarshal ...
func (r *Report) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &r)
}

const ssnTracePath = "/v1/ssn_traces"

// GetSSNTrace return the ssn trace screening report
func (r *Report) GetSSNTrace(ssnTraceID string, c *Client) (*SSNTrace, error) {
	// since ID is empty we don't have anything to return. the caller should make sure if ssnTraceID is nil
	if ssnTraceID == "" {
		return nil, errors.New("ssnTraceID is empty")
	}

	rel, err := url.Parse(ssnTracePath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), ssnTraceID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &SSNTrace{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

const sexOffenderSearchPath = "/v1/sex_offender_searches"

// GetSexOffenderSearch ...
func (r *Report) GetSexOffenderSearch(sexOffenderSearchID string, c *Client) (*SexOffenderSearch, error) {

	if sexOffenderSearchID == "" {
		return nil, errors.New("sexOffenderSearchID is empty")
	}

	rel, err := url.Parse(sexOffenderSearchPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), sexOffenderSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &SexOffenderSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

const globalWatchListSearchPath = "/v1/global_watchlist_searches"

// GetGlobalWatchListSearch ...
func (r *Report) GetGlobalWatchListSearch(globalWatchlistSearchID string, c *Client) (*GlobalWatchListSearch, error) {

	if globalWatchlistSearchID == "" {
		return nil, errors.New("globalWatchlistSearchID is empty")
	}

	rel, err := url.Parse(globalWatchListSearchPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), globalWatchlistSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &GlobalWatchListSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

const nationalCriminalSearchPath = "/v1/national_criminal_searches"

// GetNationalCriminalSearch ...
func (r *Report) GetNationalCriminalSearch(nationalCriminalSearchID string, c *Client) (*NationalCriminalSearch, error) {

	if nationalCriminalSearchID == "" {
		return nil, errors.New("nationalCriminalSearchID is empty")
	}

	rel, err := url.Parse(nationalCriminalSearchPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), nationalCriminalSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &NationalCriminalSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

const federalCriminalSearchPath = "/v1/federal_criminal_searches"

// GetFederalCriminalSearch ...
func (r *Report) GetFederalCriminalSearch(federalCrimeSearchID string, c *Client) (*FederalCriminalSearch, error) {

	if federalCrimeSearchID == "" {
		return nil, errors.New("federalCrimeSearchID is empty")
	}

	rel, err := url.Parse(federalCriminalSearchPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), federalCrimeSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &FederalCriminalSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

const countryCriminalSearchPath = "/v1/county_criminal_searches"

// GetCountryCriminalSearch ...
func (r *Report) GetCountryCriminalSearch(countryCriminalSearchID string, c *Client) (*CountryCriminalSearch, error) {

	if countryCriminalSearchID == "" {
		return nil, errors.New("countryCriminalSearchID is empty")
	}

	rel, err := url.Parse(countryCriminalSearchPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), countryCriminalSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &CountryCriminalSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

const stateCriminalSearchPath = "/v1/state_criminal_searches"

// GetStateCriminalSearch ...
func (r *Report) GetStateCriminalSearch(stateCriminalSearchID string, c *Client) (*StateCriminalSearch, error) {

	if stateCriminalSearchID == "" {
		return nil, errors.New("stateCriminalSearchID is empty")
	}

	rel, err := url.Parse(stateCriminalSearchPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), stateCriminalSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &StateCriminalSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

const motorVehicleReportSearchPath = "/v1/motor_vehicle_reports"

// GetMotorVehicleReportSearch ...
func (r *Report) GetMotorVehicleReportSearch(motorVehicleReportID string, c *Client) (*MotorVehicleReport, error) {

	if motorVehicleReportID == "" {
		return nil, errors.New("motorVehicleReportID is empty")
	}

	rel, err := url.Parse(motorVehicleReportSearchPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), motorVehicleReportID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &MotorVehicleReport{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

const educationVerificationSearchPath = "/v1/education_verifications"

// GetEducationVerificationSearch ...
func (r *Report) GetEducationVerificationSearch(educationVerificationID string, c *Client) (*EducationVerification, error) {

	if educationVerificationID == "" {
		return nil, errors.New("educationVerificationID is empty")
	}

	rel, err := url.Parse(educationVerificationSearchPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), educationVerificationID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &EducationVerification{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

const employmentVerificationSearchPath = "/v1/employment_verifications"

// GetEmploymentVerificationSearch ...
func (r *Report) GetEmploymentVerificationSearch(employmentVerificationID string, c *Client) (*EmploymentVerification, error) {

	if employmentVerificationID == "" {
		return nil, errors.New("employmentVerificationID is empty")
	}

	rel, err := url.Parse(employmentVerificationSearchPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), employmentVerificationID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &EmploymentVerification{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

const identityDocumentVerificationSearchPath = "/v1/identity_document_verifications"

// GetIdentityDocumentSearch ...
func (r *Report) GetIdentityDocumentSearch(identityDocumentVerificationID string, c *Client) (*IdentityDocumentVerification, error) {

	if identityDocumentVerificationID == "" {
		return nil, errors.New("identityDocumentVerificationID is empty")
	}

	rel, err := url.Parse(identityDocumentVerificationSearchPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), identityDocumentVerificationID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			resp.StatusCode,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getResp := &IdentityDocumentVerification{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, err
	}

	return getResp, nil
}

// GetScreenings returns all the screenings for a report
func (r *Report) GetScreenings(c *Client) (*Screenings, error) {
	cs := &Screenings{}
	var err error
	var wg sync.WaitGroup

	errorChan := make(chan error, 9+len(r.CountyCriminalSearchIDs)+len(r.StateCriminalSearchIDs))

	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.SSNTrace, err = r.GetSSNTrace(r.SSNTraceID, c)
		if err != nil {
			errorChan <- err
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.SexOffenderSearch, err = r.GetSexOffenderSearch(r.SexOffenderSearchID, c)
		if err != nil {
			errorChan <- err
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.GlobalWatchListSearch, err = r.GetGlobalWatchListSearch(r.GlobalWatchlistSearchID, c)
		if err != nil {
			errorChan <- err
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.NationalCriminalSearch, err = r.GetNationalCriminalSearch(r.NationalCriminalSearchID, c)
		if err != nil {
			errorChan <- err
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.FederalCriminalSearch, err = r.GetFederalCriminalSearch(r.FederalCrimeSearchID, c)
		if err != nil {
			errorChan <- err
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, countyCriminalSearchID := range r.CountyCriminalSearchIDs {
			ccs, err := r.GetCountryCriminalSearch(countyCriminalSearchID, c)
			if err != nil {
				errorChan <- err
			}
			cs.CountryCriminalSearches = append(cs.CountryCriminalSearches, *ccs)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, stateCriminalSearchID := range r.StateCriminalSearchIDs {
			scs, err := r.GetStateCriminalSearch(stateCriminalSearchID, c)
			if err != nil {
				errorChan <- err
			}
			cs.StateCriminalSearch = append(cs.StateCriminalSearch, *scs)
		}

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.MotorVehicleReport, err = r.GetMotorVehicleReportSearch(r.MotorVehicleReportID, c)
		if err != nil {
			errorChan <- err
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.IdentityDocumentVerification, err = r.GetIdentityDocumentSearch(r.IdentityDocumentVerificationID, c)
		if err != nil {
			errorChan <- err
		}
	}()

	wg.Wait()
	close(errorChan)
	sr := &ScreeningErrors{}
	for err := range errorChan {
		sr.errors = append(sr.errors, err)
	}
	return cs, sr
}

const createReportPath = "/v1/reports"

// CreateReport ...
func (c *Client) CreateReport(reqPayload *CreateReportRequest) (*Report, error) {
	rel, err := url.Parse(createReportPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String())

	b, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Conent-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, NewResponseError([]int{
			http.StatusCreated,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	createResp := &Report{}
	err = json.Unmarshal(b, createResp)
	if err != nil {
		return nil, err
	}

	return createResp, nil
}

const getReport = "/v1/reports"

// GetReport Retrieves a report
func (c *Client) GetReport(reportID string) (*Report, error) {
	rel, err := url.Parse(getReport)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), reportID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewResponseError([]int{
			http.StatusCreated,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	report := &Report{}
	err = json.Unmarshal(b, report)
	if err != nil {
		return nil, err
	}

	return report, nil
}
