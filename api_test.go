package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestComputeFrequency(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected map[string]int
	}{
		"Empty": {input: "", expected: map[string]int{}},
		"One word #1": {input: "hello", expected: map[string]int{
			"hello": 1,
		}},
		"One word #2": {input: " !@#!#@! hello hello hello123123!@#$!#%!@#!@#!   12312 ", expected: map[string]int{
			"hello": 3,
		}},
		"From Slack Email": {input: "Welcome to Slack!", expected: map[string]int{
			"welcome": 1,
			"to":      1,
			"slack":   1,
		}},
		"Mixed-Case": {input: "Danny DAnNy dAnNy Yo YO yOO", expected: map[string]int{
			"danny": 3,
			"yo":    2,
			"yoo":   1,
		}},
		"Mixed-Unicode": {input: "chinese unit test:单元测试 japanese test こんにちは korean unit test::안녕하세요", expected: map[string]int{
			"chinese":  1,
			"japanese": 1,
			"korean":   1,
			"unit":     2,
			"test":     3,
		}},
	}
	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := computeFrequency(test.input)
			assert.Equalf(t, test.expected, actual, "expected %v, actual %v", test.expected, actual)
		})
	}
}

func TestHandleGetWordFrequency(t *testing.T) {
	testCases := map[string]struct {
		input              string
		expectedStatusCode int
		expectedResponse   string
	}{
		"Invalid JSON": {
			input:              ``,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   ``,
		},
		"Empty JSON": {
			input:              `{}`,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"count":0,"words":{}}`,
		},
		"From Slack Email": {
			input:              `{"input": "Welcome to Slack!"}`,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"count":3,"words":{"slack":1,"to":1,"welcome":1}}`,
		},
		"Mixed-Unicode": {
			input:              `{"input":"chinese unit test:单元测试 japanese test こんにちは korean unit test::안녕하세요"}`,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"count":5,"words":{"chinese":1,"japanese":1,"korean":1,"test":3,"unit":2}}`,
		},
	}
	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "http://localhost:8080/getwordfreq", strings.NewReader(test.input))
			if err != nil {
				t.Fatalf("error calling /getwordfreq: %s\n", err.Error())
			}

			rr := httptest.NewRecorder()
			handleGetWordFrequency(rr, req)

			actualResponse := rr.Body.String()
			actualStatusCode := rr.Code

			assert.Equalf(t, test.expectedStatusCode, actualStatusCode, "expected status code %v, actual status code %v", test.expectedStatusCode, actualStatusCode)
			assert.Equalf(t, test.expectedResponse, actualResponse, "expected %v, actual %v", test.expectedResponse, actualResponse)
		})
	}
}

func TestHandleGetWordFrequencyGET(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/getwordfreq", strings.NewReader(""))
	if err != nil {
		t.Fatalf("error calling /getwordfreq: %s\n", err.Error())
	}

	rr := httptest.NewRecorder()
	handleGetWordFrequency(rr, req)

	actualResponse := rr.Body.String()
	actualStatusCode := rr.Code

	assert.Equal(t, "", actualResponse, "expected nothing, actual %v", actualResponse)
	assert.Equal(t, http.StatusNotFound, actualStatusCode, "expected status code 404, actual status code %v", actualStatusCode)

}
