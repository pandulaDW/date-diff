package time

import (
	"github.com/pandulaDW/date-diff/config"
	"testing"
)

func Test_getNumDaysSinceBaseDate(t *testing.T) {
	type test struct {
		input    config.Date
		expected uint
	}

	tests := []test{
		{input: config.NewDate(1903, 3, 24), expected: 1178},
		{input: config.NewDate(1904, 3, 7), expected: 1527},
		{input: config.NewDate(1905, 4, 12), expected: 1928},
	}

	for _, tc := range tests {
		actual := getNumDaysSinceBaseDate(&tc.input)
		if actual != tc.expected {
			t.Fatalf("expected: %d, got: %d", tc.expected, actual)
		}
	}
}

func TestDateDiff(t *testing.T) {
	type test struct {
		input1   config.Date
		input2   config.Date
		expected uint
	}

	tests := []test{
		{input1: config.NewDate(1900, 1, 20), input2: config.NewDate(1900, 3, 28), expected: 66},
		{input1: config.NewDate(2001, 1, 1), input2: config.NewDate(2001, 1, 3), expected: 1},
		{input1: config.NewDate(1983, 6, 2), input2: config.NewDate(1983, 6, 22), expected: 19},
		{input1: config.NewDate(1984, 7, 4), input2: config.NewDate(1984, 12, 25), expected: 173},
		{input1: config.NewDate(1989, 1, 3), input2: config.NewDate(1983, 8, 3), expected: 1979},
	}

	for _, tc := range tests {
		actual := DateDiff(&tc.input1, &tc.input2)
		if actual != tc.expected {
			t.Fatalf("expected: %d, got: %d", tc.expected, actual)
		}
	}
}
