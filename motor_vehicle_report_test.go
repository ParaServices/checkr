package checkr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetMotorVehicleReport(t *testing.T) {
	candidate := createCandidate(t)
	reqPayload := &CreateReportRequest{
		CandidateID: candidate.ID,
		Package:     driverPro,
	}
	client := newClient(t)
	rpt, err := client.CreateReport(reqPayload)
	require.NoError(t, err)
	require.NotNil(t, rpt)
	mvReport, err := rpt.GetMotorVehicleReportSearch(rpt.MotorVehicleReportID, client)
	require.NoError(t, err)
	require.NotEmpty(t, mvReport)

}
