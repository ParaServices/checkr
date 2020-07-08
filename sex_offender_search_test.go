package checkr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetSexOffenderListSearch(t *testing.T) {
	candidate := createCandidate(t)
	reqPayload := &CreateReportRequest{
		CandidateID: candidate.ID,
		Package:     driverpackage,
	}
	client := newClient(t)
	rpt, err := client.CreateReport(reqPayload)
	require.NoError(t, err)
	require.NotNil(t, rpt)
	ssnTrace, err := rpt.GetSexOffenderSearch(rpt.SexOffenderSearchID, client)
	require.NoError(t, err)
	require.NotEmpty(t, ssnTrace)

}
