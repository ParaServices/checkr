package checkr

import (
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/icrowley/fake"
	"github.com/magicalbanana/randata"
	"github.com/magicalbanana/tg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testAPIKey(t *testing.T) string {
	v := os.Getenv("CHECKR_TEST_API_KEY")
	if v == "" {
		assert.FailNow(t, "CHECKR_TEST_API_KEY must be set")
	}
	return v
}

func newClient(t *testing.T) *Client {
	opts := ClientOpts{}
	opts.APIKey = testAPIKey(t)
	client, err := NewClient(&opts)
	require.NoError(t, err)
	return client
}

var testCandidate *Candidate

func createCandidate(t *testing.T, createNew bool) *Candidate {
	if createNew || testCandidate == nil {
		retries := 50
		errs := make([]error, 0)
		var candidate *Candidate
		for i := 0; i < retries; i++ {
			client := newClient(t)
			customID, err := tg.RandGen(15, tg.LowerUpperDigit, "", "")
			if err != nil {
				errs = append(errs, err)
				continue
			}
			reqPayload := CreateCandidateRequest{}
			reqPayload.CustomID = customID
			reqPayload.LastName = fake.LastName()
			reqPayload.FirstName = fake.LastName()
			digit, err := tg.RandGen(5, tg.Digit, "", "")
			if err != nil {
				errs = append(errs, err)
				continue
			}
			reqPayload.Email = strings.Join([]string{"testdev", digit, "@joinpara.com"}, "")
			reqPayload.DOB = "1990-02-14"
			reqPayload.SSN = randata.RandomSSN(false, 1000)
			reqPayload.ZipCode = "60616"
			reqPayload.DriverLicenseNumber = "Y2367382"
			reqPayload.DriverLicenseState = "CA"
			resp, err := client.CreateCandidate(&reqPayload)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			if resp != nil {
				candidate = resp
			}
		}
		if candidate == nil {
			for i := range errs {
				assert.NoError(t, errs[i])
			}
			require.FailNow(t, "Couldn't generate a valid candidate")
		}
		testCandidate = candidate

	}
	// we will still assert here so it fails
	return testCandidate
}

func listPackages(t *testing.T) *ListPackagesResponse {
	client := newClient(t)
	resp, err := client.ListPackages()
	require.NoError(t, err)
	return resp
}

func randomPackage(t *testing.T) *Package {
	listPackages := listPackages(t)
	rand.Seed(time.Now().UnixNano())

	return &listPackages.Data[rand.Intn(len(listPackages.Data))]
}
