package checkr

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Error ...
type Error interface {
	Error() string
	Response() *http.Response
}

type errResponse struct {
	expectedResponseCode []int
	response             *http.Response
}

func (e *errResponse) Error() string {
	b, err := ioutil.ReadAll(e.response.Body)
	if err != nil {
		return err.Error()
	}
	defer e.response.Body.Close()
	return fmt.Sprintf(
		"Expected response code: %v, actual response code: %v, url: %v, method: %v, resp_body: %v",
		e.expectedResponseCode,
		e.response.StatusCode,
		e.response.Request.URL.String(),
		e.response.Request.Method,
		string(b),
	)
}

// ResponseCode ...
func (e *errResponse) Response() *http.Response {
	return e.response
}

// NewError ...
func NewError(expectedRespCode []int, resp *http.Response) Error {
	return &errResponse{
		expectedResponseCode: expectedRespCode,
		response:             resp,
	}
}
