package checkr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetNationalCriminalListSearch(t *testing.T) {
	candidate := createCandidate(t)
	reqPayload := &CreateReportRequest{
		CandidateID: candidate.ID,
		Package:     DriverPro.Slug(),
	}
	client := newClient(t)
	rpt, err := client.CreateReport(reqPayload)
	require.NoError(t, err)
	require.NotNil(t, rpt)
	nationalCriminalSearch, err := rpt.GetNationalCriminalSearch(rpt.NationalCriminalSearchID, client)
	require.Empty(t, err)
	require.NotEmpty(t, nationalCriminalSearch)

}
