package checkr

import (
	"bytes"
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

// NewErrorX wraps all details we get during screening or checkr requests
func NewErrorX(id string, st ScreenType, err error, expectedRespCode int, resp *http.Response) *ErrorX {
	return &ErrorX{
		id:                   id,
		expectedResponseCode: expectedRespCode,
		response:             resp,
	}
}

// ErrorX ...
type ErrorX struct {
	expectedResponseCode int
	response             *http.Response
	id                   string
	err                  error
	screenType           ScreenType
}

func (e *ErrorX) Error() string {
	if e.response != nil {
		b, err := ioutil.ReadAll(e.response.Body)
		if err != nil {
			return fmt.Sprintf(
				"id: %v,  error: %v",
				e.id,
				e.err,
			)
		}
		defer e.response.Body.Close()
		return fmt.Sprintf(
			"id: %v,  Expected response code: %v, actual response code: %v, url: %v, method: %v, resp_body: %v, error: %v",
			e.id,
			e.expectedResponseCode,
			e.response.StatusCode,
			e.response.Request.URL.String(),
			e.response.Request.Method,
			string(b),
			e.err,
		)
	}
	return fmt.Sprintf(
		"id: %v,  error: %v",
		e.id,
		e.err,
	)
}

// ErrorsX ...
type ErrorsX []ErrorX

func (e ErrorsX) String() string {
	buf := bytes.Buffer{}
	for _, err := range e {
		buf.WriteString(err.Error() + "\n")
	}
	return buf.String()
}

// ScreeningError ...
type ScreeningError struct {
	errMap map[ScreenType][]*ErrorX
}

func (s *ScreeningError) Error() string {
	buf := bytes.Buffer{}
	for k, v := range s.errMap {
		buf.WriteString(fmt.Sprintf("screenType: %v,  Errors: %v \n", k, v))
	}
	return buf.String()
}
