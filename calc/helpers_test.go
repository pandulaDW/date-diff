package calc

import "testing"

func Test_isLeapYear(t *testing.T) {
	type test struct {
		input    uint
		expected bool
	}

	tests := []test{
		{input: 1992, expected: true},
		{input: 2003, expected: false},
		{input: 2016, expected: true},
		{input: 2000, expected: true},
		{input: 1900, expected: false},
		{input: 1904, expected: true},
		{input: 2300, expected: false},
		{input: 2400, expected: true},
	}

	for _, tc := range tests {
		actual := isLeapYear(tc.input)
		if actual != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, actual)
		}
	}
}

func Test_numDaysPerMonth(t *testing.T) {
	type test struct {
		month    uint8
		year     uint
		expected uint8
	}

	tests := []test{
		{month: 12, year: 2009, expected: 31},
		{month: 2, year: 2013, expected: 28},
		{month: 2, year: 2016, expected: 29},
	}

	for _, tc := range tests {
		actual := numDaysPerMonth(tc.month, tc.year)
		if actual != tc.expected {
			t.Fatalf("expected: %d, got: %d", tc.expected, actual)
		}
	}
}
