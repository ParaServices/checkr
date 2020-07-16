package checkr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetSexOffenderListSearch(t *testing.T) {
	candidate := createCandidate(t, false)
	reqPayload := &CreateReportRequest{
		CandidateID: candidate.ID,
		Package:     DriverPro.Slug(),
	}
	client := newClient(t)
	rpt, err := client.CreateReport(reqPayload)
	require.NoError(t, err)
	require.NotNil(t, rpt)
	sexOffenderSearch, err := rpt.GetSexOffenderSearch(rpt.SexOffenderSearchID, client)
	require.Empty(t, err)
	require.NotEmpty(t, sexOffenderSearch)

}
