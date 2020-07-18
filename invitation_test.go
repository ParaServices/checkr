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
	candidate := createCandidate(t, false)
	reqPayload := CreateInvitationRequest{}
	reqPayload.CandidateID = candidate.ID
	reqPayload.Package = "driver_pro"
	resp, err := client.CreateInvitation(&reqPayload)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestClient_GetInvitiation(t *testing.T) {
	opts := ClientOpts{}
	opts.APIKey = testAPIKey(t)
	client, err := NewClient(&opts)
	require.NoError(t, err)

	// create a candidate first
	candidate := createCandidate(t, false)
	reqPayload := CreateInvitationRequest{}
	reqPayload.CandidateID = candidate.ID
	reqPayload.Package = "driver_pro"
	createResp, err := client.CreateInvitation(&reqPayload)
	require.NoError(t, err)
	require.NotNil(t, createResp)

	getResp, err := client.GetInvitation(createResp.ID)
	require.NoError(t, err)
	require.NotNil(t, getResp)
	require.Equal(t, createResp.ID, getResp.ID)
}
