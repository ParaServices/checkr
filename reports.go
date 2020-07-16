package checkr

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
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
func (r *Report) GetSSNTrace(ssnTraceID string, c *Client) (*SSNTrace, *Error) {
	// since ID is empty we don't have anything to return. the caller should make sure if ssnTraceID is nil
	if ssnTraceID == "" {
		return nil, NewError(ssnTraceID, SsnTrace, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(ssnTracePath)
	if err != nil {
		return nil, NewError(ssnTraceID, SsnTrace, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), ssnTraceID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(ssnTraceID, SsnTrace, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(ssnTraceID, SsnTrace, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(ssnTraceID, SsnTrace, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(ssnTraceID, SsnTrace, err, http.StatusOK, resp)
	}

	getResp := &SSNTrace{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(ssnTraceID, SsnTrace, err, http.StatusOK, resp)
	}

	return getResp, nil
}

const sexOffenderSearchPath = "/v1/sex_offender_searches"

func (r *Report) GetSexOffenderSearch(sexOffenderSearchID string, c *Client) (*SexOffenderSearch, *Error) {

	if sexOffenderSearchID == "" {
		return nil, NewError(sexOffenderSearchID, Sexoffendersearch, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(sexOffenderSearchPath)
	if err != nil {
		return nil, NewError(sexOffenderSearchID, Sexoffendersearch, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), sexOffenderSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(sexOffenderSearchID, Sexoffendersearch, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(sexOffenderSearchID, Sexoffendersearch, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(sexOffenderSearchID, Sexoffendersearch, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(sexOffenderSearchID, Sexoffendersearch, err, http.StatusOK, resp)
	}

	getResp := &SexOffenderSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(sexOffenderSearchID, Sexoffendersearch, err, http.StatusOK, resp)
	}

	return getResp, nil
}

const globalWatchListSearchPath = "/v1/global_watchlist_searches"

func (r *Report) GetGlobalWatchListSearch(globalWatchlistSearchID string, c *Client) (*GlobalWatchListSearch, *Error) {

	if globalWatchlistSearchID == "" {
		return nil, NewError(globalWatchlistSearchID, GlobalWatchListsearch, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(globalWatchListSearchPath)
	if err != nil {
		return nil, NewError(globalWatchlistSearchID, GlobalWatchListsearch, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), globalWatchlistSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(globalWatchlistSearchID, GlobalWatchListsearch, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(globalWatchlistSearchID, GlobalWatchListsearch, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(globalWatchlistSearchID, GlobalWatchListsearch, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(globalWatchlistSearchID, GlobalWatchListsearch, err, http.StatusOK, resp)
	}

	getResp := &GlobalWatchListSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(globalWatchlistSearchID, GlobalWatchListsearch, err, http.StatusOK, resp)
	}

	return getResp, nil
}

const nationalCriminalSearchPath = "/v1/national_criminal_searches"

func (r *Report) GetNationalCriminalSearch(nationalCriminalSearchID string, c *Client) (*NationalCriminalSearch, *Error) {

	if nationalCriminalSearchID == "" {
		return nil, NewError(nationalCriminalSearchID, NationalCriminalsearch, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(nationalCriminalSearchPath)
	if err != nil {
		return nil, NewError(nationalCriminalSearchID, NationalCriminalsearch, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), nationalCriminalSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(nationalCriminalSearchID, NationalCriminalsearch, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(nationalCriminalSearchID, NationalCriminalsearch, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(nationalCriminalSearchID, NationalCriminalsearch, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(nationalCriminalSearchID, NationalCriminalsearch, err, http.StatusOK, resp)
	}

	getResp := &NationalCriminalSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(nationalCriminalSearchID, NationalCriminalsearch, err, http.StatusOK, resp)
	}

	return getResp, nil
}

const federalCriminalSearchPath = "/v1/federal_criminal_searches"

func (r *Report) GetFederalCriminalSearch(federalCrimeSearchID string, c *Client) (*FederalCriminalSearch, *Error) {

	if federalCrimeSearchID == "" {
		return nil, NewError(federalCrimeSearchID, FederalCriminalsearch, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(federalCriminalSearchPath)
	if err != nil {
		return nil, NewError(federalCrimeSearchID, FederalCriminalsearch, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), federalCrimeSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(federalCrimeSearchID, FederalCriminalsearch, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(federalCrimeSearchID, FederalCriminalsearch, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(federalCrimeSearchID, FederalCriminalsearch, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(federalCrimeSearchID, FederalCriminalsearch, err, http.StatusOK, resp)
	}

	getResp := &FederalCriminalSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(federalCrimeSearchID, FederalCriminalsearch, err, http.StatusOK, resp)
	}

	return getResp, nil
}

const countryCriminalSearchPath = "/v1/county_criminal_searches"

func (r *Report) GetCountryCriminalSearch(countryCriminalSearchID string, c *Client) (*CountryCriminalSearch, *Error) {

	if countryCriminalSearchID == "" {
		return nil, NewError(countryCriminalSearchID, CountryCriminalsearch, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(countryCriminalSearchPath)
	if err != nil {
		return nil, NewError(countryCriminalSearchID, CountryCriminalsearch, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), countryCriminalSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(countryCriminalSearchID, CountryCriminalsearch, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(countryCriminalSearchID, CountryCriminalsearch, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(countryCriminalSearchID, CountryCriminalsearch, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(countryCriminalSearchID, CountryCriminalsearch, err, http.StatusOK, resp)
	}

	getResp := &CountryCriminalSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(countryCriminalSearchID, CountryCriminalsearch, err, http.StatusOK, resp)
	}

	return getResp, nil
}

const stateCriminalSearchPath = "/v1/state_criminal_searches"

func (r *Report) GetStateCriminalSearch(stateCriminalSearchID string, c *Client) (*StateCriminalSearch, *Error) {

	if stateCriminalSearchID == "" {
		return nil, NewError(stateCriminalSearchID, StateCriminalsearch, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(stateCriminalSearchPath)
	if err != nil {
		return nil, NewError(stateCriminalSearchID, StateCriminalsearch, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), stateCriminalSearchID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(stateCriminalSearchID, StateCriminalsearch, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(stateCriminalSearchID, StateCriminalsearch, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(stateCriminalSearchID, StateCriminalsearch, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(stateCriminalSearchID, StateCriminalsearch, err, http.StatusOK, resp)
	}

	getResp := &StateCriminalSearch{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(stateCriminalSearchID, StateCriminalsearch, err, http.StatusOK, resp)
	}

	return getResp, nil
}

const motorVehicleReportSearchPath = "/v1/motor_vehicle_reports"

func (r *Report) GetMotorVehicleReportSearch(motorVehicleReportID string, c *Client) (*MotorVehicleReport, *Error) {

	if motorVehicleReportID == "" {
		return nil, NewError(motorVehicleReportID, MotorVehiclereport, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(motorVehicleReportSearchPath)
	if err != nil {
		return nil, NewError(motorVehicleReportID, MotorVehiclereport, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), motorVehicleReportID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(motorVehicleReportID, MotorVehiclereport, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(motorVehicleReportID, MotorVehiclereport, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(motorVehicleReportID, MotorVehiclereport, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(motorVehicleReportID, MotorVehiclereport, err, http.StatusOK, resp)
	}

	getResp := &MotorVehicleReport{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(motorVehicleReportID, MotorVehiclereport, err, http.StatusOK, resp)
	}
	return getResp, nil
}

const educationVerificationSearchPath = "/v1/education_verifications"

func (r *Report) GetEducationVerificationSearch(educationVerificationID string, c *Client) (*EducationVerification, *Error) {

	if educationVerificationID == "" {
		return nil, NewError(educationVerificationID, Educationverification, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(educationVerificationSearchPath)
	if err != nil {
		return nil, NewError(educationVerificationID, Educationverification, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), educationVerificationID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(educationVerificationID, Educationverification, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(educationVerificationID, Educationverification, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(educationVerificationID, Educationverification, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(educationVerificationID, Educationverification, err, http.StatusOK, resp)
	}

	getResp := &EducationVerification{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(educationVerificationID, Educationverification, err, http.StatusOK, resp)
	}

	return getResp, nil
}

const employmentVerificationSearchPath = "/v1/employment_verifications"

func (r *Report) GetEmploymentVerificationSearch(employmentVerificationID string, c *Client) (*EmploymentVerification, *Error) {

	if employmentVerificationID == "" {
		return nil, NewError(employmentVerificationID, Employmentverification, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(employmentVerificationSearchPath)
	if err != nil {
		return nil, NewError(employmentVerificationID, Employmentverification, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), employmentVerificationID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(employmentVerificationID, Employmentverification, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(employmentVerificationID, Employmentverification, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(employmentVerificationID, Employmentverification, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(employmentVerificationID, Employmentverification, err, http.StatusOK, resp)
	}

	getResp := &EmploymentVerification{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(employmentVerificationID, Employmentverification, err, http.StatusOK, resp)
	}

	return getResp, nil
}

const identityDocumentVerificationSearchPath = "/v1/identity_document_verifications"

func (r *Report) GetIdentityDocumentSearch(identityDocumentVerificationID string, c *Client) (*IdentityDocumentVerification, *Error) {

	if identityDocumentVerificationID == "" {
		return nil, NewError(identityDocumentVerificationID, IdentityDocumentverification, errors.New("empty ID"), 0, nil)
	}

	rel, err := url.Parse(identityDocumentVerificationSearchPath)
	if err != nil {
		return nil, NewError(identityDocumentVerificationID, IdentityDocumentverification, err, 0, nil)
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String(), identityDocumentVerificationID)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, NewError(identityDocumentVerificationID, IdentityDocumentverification, err, 0, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, NewError(identityDocumentVerificationID, IdentityDocumentverification, err, 0, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(identityDocumentVerificationID, IdentityDocumentverification, err, http.StatusOK, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(identityDocumentVerificationID, IdentityDocumentverification, err, http.StatusOK, resp)
	}

	getResp := &IdentityDocumentVerification{}
	err = json.Unmarshal(b, getResp)
	if err != nil {
		return nil, NewError(identityDocumentVerificationID, IdentityDocumentverification, err, http.StatusOK, resp)
	}

	return getResp, nil
}

// GetScreenings returns all the screenings for a report
func (r *Report) GetScreenings(c *Client) (*Screenings, error) {
	cs := &Screenings{}
	var err *Error
	var wg sync.WaitGroup

	screenIds := make(map[ScreenType]string)
	screenIds[SsnTrace] = r.SSNTraceID
	screenIds[Sexoffendersearch] = r.SexOffenderSearchID
	screenIds[GlobalWatchListsearch] = r.GlobalWatchlistSearchID
	screenIds[NationalCriminalsearch] = r.NationalCriminalSearchID
	screenIds[FederalCriminalsearch] = r.FederalCrimeSearchID
	screenIds[MotorVehiclereport] = r.MotorVehicleReportID
	if len(r.CountyCriminalSearchIDs) > 0 {
		screenIds[CountryCriminalsearch] = r.CountyCriminalSearchIDs[0]
	}
	if len(r.StateCriminalSearchIDs) > 0 {
		screenIds[StateCriminalsearch] = r.StateCriminalSearchIDs[0]
	}
	//change these two if we know how to get the ids of these to pass to functions
	screenIds[Educationverification] = ""
	screenIds[Employmentverification] = ""
	screenIds[IdentityDocumentverification] = r.IdentityDocumentVerificationID

	screenErrorsChan := make(chan *Error, len(screenIds)+len(r.CountyCriminalSearchIDs)+len(r.StateCriminalSearchIDs))
	for screenType, id := range screenIds {
		if id == "" {
			continue
		}
		wg.Add(1)
		go func(t ScreenType, errChan chan<- *Error) {
			defer wg.Done()
			switch t {
			case SsnTrace:
				cs.SSNTrace, err = r.GetSSNTrace(r.SSNTraceID, c)
				if err != nil {
					errChan <- err
					log.Println("error getting ssn trace ", err)
				}
			case Sexoffendersearch:
				cs.SexOffenderSearch, err = r.GetSexOffenderSearch(r.SexOffenderSearchID, c)
				if err != nil {
					errChan <- err
					log.Println("error getting sex offender searches", err)
				}
			case GlobalWatchListsearch:
				cs.GlobalWatchListSearch, err = r.GetGlobalWatchListSearch(r.GlobalWatchlistSearchID, c)
				if err != nil {
					errChan <- err
					log.Println("error getting global watch list searches", err)
				}
			case NationalCriminalsearch:
				cs.NationalCriminalSearch, err = r.GetNationalCriminalSearch(r.NationalCriminalSearchID, c)
				if err != nil {
					errChan <- err
					log.Println("error getting national criminal searches", err)
				}
			case FederalCriminalsearch:
				cs.FederalCriminalSearch, err = r.GetFederalCriminalSearch(r.FederalCrimeSearchID, c)
				if err != nil {
					errChan <- err
					log.Println("error getting federal criminal search", err)
				}
			case CountryCriminalsearch:
				for _, countyCriminalSearchID := range r.CountyCriminalSearchIDs {
					ccs, err := r.GetCountryCriminalSearch(countyCriminalSearchID, c)
					if err != nil {
						errChan <- err
						log.Println("error getting country criminal searches", err)
						continue
					}
					cs.CountryCriminalSearches = append(cs.CountryCriminalSearches, *ccs)
				}
			case StateCriminalsearch:
				for _, stateCriminalSearchID := range r.StateCriminalSearchIDs {
					scs, err := r.GetStateCriminalSearch(stateCriminalSearchID, c)
					if err != nil {
						errChan <- err
						log.Println("error getting state criminal searches", err)
						continue
					}
					cs.StateCriminalSearch = append(cs.StateCriminalSearch, *scs)
				}
			case MotorVehiclereport:
				cs.MotorVehicleReport, err = r.GetMotorVehicleReportSearch(r.MotorVehicleReportID, c)
				if err != nil {
					errChan <- err
					log.Println("error getting motor vehicle report", err)
				}
			case Educationverification:
				// cs.EducationVerification, err = r.GetEducationVerificationSearch(r.EducationVerificationSearchID, c)
				// if err != nil {
				// errChan <- err
				// 	log.Println("error getting education verification", err)
				// }
			case Employmentverification:
				// cs.EmploymentVerification, err = r.GetEmploymentVerificationSearch(r.EmploymentVerificationSearchID, c)
				// if err != nil {
				// errChan <- err
				// 	log.Println("error getting employment verification", err)
				// }
			case IdentityDocumentverification:
				cs.IdentityDocumentVerification, err = r.GetIdentityDocumentSearch(r.IdentityDocumentVerificationID, c)
				if err != nil {
					errChan <- err
					log.Println("error getting identity document search", err)
				}
			}
		}(screenType, screenErrorsChan)

	}
	wg.Wait()
	close(screenErrorsChan)
	sr := &ScreeningError{}
	for err := range screenErrorsChan {
		sr.errMap[err.screenType] = append(sr.errMap[err.screenType], err)
	}
	return cs, sr
}

const createReportPath = "/v1/reports"

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
