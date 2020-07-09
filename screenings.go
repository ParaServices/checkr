package checkr

import (
	"log"
	"sync"
)

// Screenings ...
type Screenings struct {
	SSNTrace                     *SSNTrace
	SexOffenderSearch            *SexOffenderSearch
	GlobalWatchListSearch        *GlobalWatchListSearch
	NationalCriminalSearch       *NationalCriminalSearch
	FederalCriminalSearch        *FederalCriminalSearch
	CountryCriminalSearches      []*CountryCriminalSearch
	StateCriminalSearch          []*StateCriminalSearch
	MotorVehicleReport           *MotorVehicleReport
	EducationVerification        *EducationVerification
	EmploymentVerification       *EmploymentVerification
	IdentityDocumentVerification *IdentityDocumentVerification
}

// GetScreenings returns all the screenings for a report
func GetScreenings(report *Report, c *Client) *Screenings {
	cs := &Screenings{}
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.SSNTrace, err = report.GetSSNTrace(report.SSNTraceID, c)
		if err != nil {
			log.Println("error getting ssn trace ", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.SexOffenderSearch, err = report.GetSexOffenderSearch(report.SexOffenderSearchID, c)
		if err != nil {
			log.Println("error getting sex offender searches", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.GlobalWatchListSearch, err = report.GetGlobalWatchListSearch(report.GlobalWatchlistSearchID, c)
		if err != nil {
			log.Println("error getting global watch list searches", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.NationalCriminalSearch, err = report.GetNationalCriminalSearch(report.NationalCriminalSearchID, c)
		if err != nil {
			log.Println("error getting national criminal searches", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.FederalCriminalSearch, err = report.GetFederalCriminalSearch(report.FederalCrimeSearchID, c)
		if err != nil {
			log.Println("error getting federal criminal search", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, countyCriminalSearchID := range report.CountyCriminalSearchIDs {
			ccs, err := report.GetCountryCriminalSearch(countyCriminalSearchID, c)
			if err != nil {
				log.Println("error getting country criminal searches", err)
				continue
			}
			cs.CountryCriminalSearches = append(cs.CountryCriminalSearches, ccs)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, stateCriminalSearchID := range report.StateCriminalSearchIDs {
			scs, err := report.GetStateCriminalSearch(stateCriminalSearchID, c)
			if err != nil {
				log.Println("error getting state criminal searches", err)
				continue
			}
			cs.StateCriminalSearch = append(cs.StateCriminalSearch, scs)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.MotorVehicleReport, err = report.GetMotorVehicleReportSearch(report.MotorVehicleReportID, c)
		if err != nil {
			log.Println("error getting motor vehicle report", err)
		}
	}()

	//todo the fields EducationVerificationSearchID and EmploymentVerificationSearchID are not present
	// not sure how to get EducationVerificationSearch and EmploymentVerificationSearch
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	cs.EducationVerification, err = report.GetEducationVerificationSearch(report.EducationVerificationSearchID, c)
	// 	if err != nil {
	// 		log.Println("error getting education verification", err)
	// 	}
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	cs.EmploymentVerification, err = report.GetEmploymentVerificationSearch(report.EmploymentVerificationSearchID, c)
	// 	if err != nil {
	// 		log.Println("error getting employment verification", err)
	// 	}
	// }()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.IdentityDocumentVerification, err = report.GetIdentityDocumentSearch(report.IdentityDocumentVerificationID, c)
		if err != nil {
			log.Println("error getting identity document search", err)
		}
	}()

	wg.Wait()

	return cs

}
