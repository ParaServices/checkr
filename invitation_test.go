package checkr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_CreateInvitiation(t *testing.T) {
	opts := ClientOpts{}
	opts.APIKey = testAPIKey(t)
	client, err := NewClient(&opts)
	require.NoError(t, err)

	// create a candidate first
	candidate := createCandidate(t)
	reqPayload := CreateInvitationRequest{}
	reqPayload.CandidateID = candidate.ID
	reqPayload.Package = "driver_pro"
	resp, err := client.CreateInvitation(&reqPayload)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
