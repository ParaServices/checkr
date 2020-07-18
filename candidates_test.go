package checkr

import (
	"strings"
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
	digit, err := tg.RandGen(5, tg.Digit, "", "")
	require.NoError(t, err)
	reqPayload.Email = strings.Join([]string{"testdev", digit, "@joinpara.com"}, "")
	resp, err := client.CreateCandidate(&reqPayload)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
