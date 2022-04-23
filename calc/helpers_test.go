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
