package checkr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestClient_CreateAdverseAction(t *testing.T) {
// 	opts := ClientOpts{}
// 	opts.APIKey = testAPIKey(t)
// 	client, err := NewClient(&opts)
// 	require.NoError(t, err)

// 	// create a candidate first
// 	candidate := createCandidate(t)
// 	reportReqPayload := &CreateReportRequest{
// 		CandidateID: candidate.ID,
// 		Package:     randomPackage(t).Slug,
// 	}
// 	log.Println(candidate)
// 	reportResp, err := client.CreateReport(reportReqPayload)
// 	require.NoError(t, err)
// 	require.NotNil(t, reportResp)
// 	log.Println(reportResp)

// 	reqPayload := CreateAdverseActionRequest{}
// 	reqPayload.PostNoticeScheduledAt = time.Now().Add(time.Hour * 1)
// 	reqPayload.AdverseItemIds = []string{"57ed4ce3057e0b002adc6d90", "57ed4ce3057e0b002adc6d90"}
// 	resp, err := client.CreateAdverseActionRequest(reportResp.ID, &reqPayload)
// 	require.NoError(t, err)
// 	require.NotNil(t, resp)
// }

func TestClient_GetAdverseAction(t *testing.T) {
	opts := ClientOpts{}
	opts.APIKey = testAPIKey(t)
	client, err := NewClient(&opts)
	require.NoError(t, err)

	candidate := createCandidate(t)
	reportReqPayload := &CreateReportRequest{
		CandidateID: candidate.ID,
		Package:     randomPackage(t).Slug,
	}
	reportResp, err := client.CreateReport(reportReqPayload)
	require.NoError(t, err)
	require.NotNil(t, reportResp)

	// reqPayload := CreateAdverseActionRequest{}
	// reqPayload.PostNoticeScheduledAt = time.Now().Add(time.Hour * 1)
	// reqPayload.AdverseItemIds = []string{"e44aa283528e6fde7d542194"}
	// resp, err := client.CreateAdverseActionRequest(reportResp.ID, &reqPayload)
	// require.NoError(t, err)
	// require.NotNil(t, resp)

	// getResp, err := client.GetAdverseAction("e44aa283528e6fde7d542194")
	// require.NoError(t, err)
	// require.NotNil(t, getResp)
	// require.Equal(t, resp.ID, getResp.ID)
}
