package checkr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetScreenings(t *testing.T) {
	candidate := createCandidate(t)
	reqPayload := &CreateReportRequest{
		CandidateID: candidate.ID,
		Package:     DriverPro.Code(),
	}
	client := newClient(t)
	rpt, err := client.CreateReport(reqPayload)
	require.NoError(t, err)
	require.NotNil(t, rpt)

	screenings := GetScreenings(rpt, client)
	require.NotEmpty(t, screenings)
}
