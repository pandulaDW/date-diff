package input

import (
	"github.com/pandulaDW/date-diff/config"
	"testing"
)

func Test_parseDate(t *testing.T) {
	type test struct {
		input            string
		expectedDate     *config.Date
		expectedErrorMsg string
	}

	tests := []test{
		{"02-12-2023", nil, "format of 02-12-2023 is incorrect. use the format (dd|d)/(mm|m)/yyyy.\nFor example: 02/04/2012"},
		{"02/2/", nil, "format of 02/2/ is incorrect. use the format (dd|d)/(mm|m)/yyyy.\nFor example: 02/04/2012"},
		{"31/6/1883", nil, "year should be between 1990 and 2999 for 31/6/1883"},
		{"31/6/3083", nil, "year should be between 1990 and 2999 for 31/6/3083"},
		{"31/13/1983", nil, "month should be between 1 and 12 for 31/13/1983"},
		{"31/0/1983", nil, "month should be between 1 and 12 for 31/0/1983"},
		{"31/6/1983", nil, "number of days for month 6 should be between 1 and 30 for date 31/6/1983"},
		{"0/6/1983", nil, "number of days for month 6 should be between 1 and 30 for date 0/6/1983"},
	}

	for _, tc := range tests {
		actualDate, actualError := parseDate(tc.input)
		if actualDate != tc.expectedDate {
			t.Fatalf("expected date: %v, got: %v", tc.expectedDate, actualDate)
		}
		if actualError.Error() != tc.expectedErrorMsg {
			t.Fatalf("expected date: %s, got: %s", tc.expectedErrorMsg, actualError.Error())
		}
	}

	actual, err := parseDate("5/06/1983")
	expected := config.NewDate(1983, 6, 5)
	if *actual != expected {
		t.Fatalf("expected: %v, got: %v", expected, actual)
	}
	if err != nil {
		t.Fatalf("expected error to be nil")
	}
}
