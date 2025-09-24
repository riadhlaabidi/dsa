package fractionrecurring

import (
	"dsa/internal/testutil"
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	var tests = []struct {
		name        string
		numerator   int
		denominator int
		want        string
	}{
		{"Not recurring", 1, 2, "0.5"},
		{"Integral", 2, 1, "2"},
		{"Recurring sequence of three", 4, 333, "0.(012)"},
		{"Recurring sequence of one", 1, 6, "0.1(6)"},
		{"Negative not recurring", -50, 8, "-6.25"},
		{"Negative recurring", 7, -12, "-0.58(3)"},
		{"Both negative", -3, -4, "0.75"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testutil.AssertEquals(t, test.want, fractionToDecimal(test.numerator, test.denominator))
		})
	}
}
