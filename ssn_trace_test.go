package checkr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const driverpackage = "driver_pro"

func TestClient_GetSSNTrace(t *testing.T) {
	candidate := createCandidate(t)
	reqPayload := &CreateReportRequest{
		CandidateID: candidate.ID,
		Package:     driverpackage,
	}
	client := newClient(t)
	rpt, err := client.CreateReport(reqPayload)
	require.NoError(t, err)
	require.NotNil(t, rpt)

	ssnTrace, err := rpt.GetSSNTrace(rpt.SSNTraceID, client)
	require.NoError(t, err)
	require.NotEmpty(t, ssnTrace)
}
