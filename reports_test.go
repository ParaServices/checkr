package checkr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_CreateReport(t *testing.T) {
	candidate := createCandidate(t)
	reqPayload := &CreateReportRequest{
		CandidateID: candidate.ID,
		Package:     randomPackage(t).Slug,
	}
	client := newClient(t)
	resp, err := client.CreateReport(reqPayload)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
