package main

import (
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
		"Slack Email": {input: "Welcome to Slack!", expected: map[string]int{
			"welcome": 1,
			"to":      1,
			"slack":   1,
		}},
		"Mixed-Case": {input: "Danny DAnNy dAnNy Yo YO yOO", expected: map[string]int{
			"danny": 3,
			"yo":    2,
			"yoo":   1,
		}},
		"Mixed-Unicode": {input: "chinese unit test:单元测试 japanese test こんにちは korean unit test::::안녕하세요", expected: map[string]int{
			"chinese":  1,
			"japanese": 1,
			"korean":   1,
			"unit":     2,
			"test":     3,
		}},
	}
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := computeFrequency(test.input)
			expected := test.expected
			assert.Equalf(t, expected, actual, "expected %v, actual %v", expected, actual)
		})
	}
}
