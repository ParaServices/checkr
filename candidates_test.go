package checkr

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/magicalbanana/tg"
	"github.com/stretchr/testify/require"
)

func TestClient_CreateCandidate(t *testing.T) {
	opts := ClientOpts{}
	opts.APIKey = testAPIKey(t)
	client, err := NewClient(&opts)
	require.NoError(t, err)
	customID, err := tg.RandGen(15, tg.LowerUpperDigit, "", "")
	require.NoError(t, err)
	reqPayload := CreateCandidateRequest{}
	reqPayload.CustomID = customID
	reqPayload.LastName = fake.LastName()
	reqPayload.FirstName = fake.FirstName()
	reqPayload.Email = fake.EmailAddress()
	resp, err := client.CreateCandidate(&reqPayload)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
