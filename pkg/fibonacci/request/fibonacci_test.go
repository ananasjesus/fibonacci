package fibonacci

import (
	"testing"
)

type testPair struct {
	From               string
	To                 string
	ExpectedValidation bool
}

var tests = []testPair{
	{"1", "5", true},
	{"5", "90", true},
	{"5", "5", true},
	{"0", "5", false},
	{"5", "4", false},
	{"5", "91", false},
}

func TestValidate(t *testing.T) {
	for _, pair := range tests {
		validated := (&FibonacciRequest{pair.From, pair.To}).Validate()
		if validated != pair.ExpectedValidation {
			t.Error(
				"For from:", pair.From, "to:", pair.To,
				"expected", pair.ExpectedValidation,
				"got", validated,
			)
		}
	}
}
