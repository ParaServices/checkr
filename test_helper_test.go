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

func createCandidate(t *testing.T) *Candidate {
	client := newClient(t)
	customID, err := tg.RandGen(15, tg.LowerUpperDigit, "", "")
	require.NoError(t, err)
	reqPayload := CreateCandidateRequest{}
	reqPayload.CustomID = customID
	reqPayload.LastName = fake.LastName()
	reqPayload.FirstName = fake.LastName()
	digit, err := tg.RandGen(5, tg.Digit, "", "")
	require.NoError(t, err)
	reqPayload.Email = strings.Join([]string{"testdev", digit, "@joinpara.com"}, "")
	reqPayload.DOB = "1990-02-14"
	reqPayload.SSN = randata.RandomSSN(false, 1000)
	reqPayload.ZipCode = "60616"
	reqPayload.DriverLicenseNumber = "Y2367382"
	reqPayload.DriverLicenseState = "CA"
	resp, err := client.CreateCandidate(&reqPayload)
	require.NoError(t, err)
	require.NotNil(t, resp)
	return resp
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
