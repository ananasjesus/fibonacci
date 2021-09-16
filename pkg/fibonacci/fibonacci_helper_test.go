package fibonacci

import (
	"reflect"
	"testing"
)

type testPair struct {
	n                 int
	expectedFibonacci []int64
}

var tests = []testPair{
	{1, []int64{0}},
	{3, []int64{0, 1, 1}},
	{10, []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
}

func TestCalcFibonacci(t *testing.T) {
	for _, pair := range tests {
		calculatedFibonacci := CalcFibonacci(pair.n)
		if !reflect.DeepEqual(calculatedFibonacci, pair.expectedFibonacci) {
			t.Error(
				"For", pair.n,
				"expected", pair.expectedFibonacci,
				"got", calculatedFibonacci,
			)
		}
	}
}
