package checkr

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetCountryCriminalSearch(t *testing.T) {
	candidate := createCandidate(t)
	reqPayload := &CreateReportRequest{
		CandidateID: candidate.ID,
		Package:     DriverPro.Slug(),
	}
	client := newClient(t)
	rpt, err := client.CreateReport(reqPayload)
	require.NoError(t, err)
	require.NotNil(t, rpt)
	for _, c := range rpt.CountyCriminalSearchIDs {
		countryCriminalSearch, err := rpt.GetCountryCriminalSearch(c, client)
		require.Empty(t, err)
		require.NotEmpty(t, countryCriminalSearch)
		log.Println(countryCriminalSearch)
	}

}
