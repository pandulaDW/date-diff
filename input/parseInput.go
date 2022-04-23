package input

import (
	"fmt"
	"github.com/pandulaDW/date-diff/calc"
	"github.com/pandulaDW/date-diff/config"
	"strconv"
	"strings"
)

// ParseDates takes the two date strings given by the user, validates them
// and return the two parsed dates.
func ParseDates(dStr1, dStr2 string) (*config.Date, *config.Date, error) {
	date1, err := parseDate(dStr1)
	if err != nil {
		return nil, nil, err
	}
	date2, err := parseDate(dStr2)
	if err != nil {
		return nil, nil, err
	}
	return date1, date2, nil
}

// parseDate is a helper function for ParseDates
func parseDate(dStr string) (*config.Date, error) {
	if !config.DateFormat.Match([]byte(dStr)) {
		return nil,
			fmt.Errorf("format of %s is incorrect. use the format (dd|d)/(mm|m)/yyyy.\nFor example: 02/04/2012", dStr)
	}

	split := strings.Split(dStr, "/")
	day, _ := strconv.ParseUint(split[0], 10, 64)
	month, _ := strconv.ParseUint(split[1], 10, 64)
	year, _ := strconv.ParseUint(split[2], 10, 64)

	if uint(year) < config.BaseYear || uint(year) > config.MaxYear {
		return nil, fmt.Errorf("year should be between 1990 and 2999 for %s", dStr)
	}

	if month < 1 || month > 12 {
		return nil, fmt.Errorf("month should be between 1 and 12 for %s", dStr)
	}

	numDays := calc.NumDaysPerMonth(uint8(month), uint(year))
	if day < 1 || uint8(day) > numDays {
		return nil, fmt.Errorf("number of days for month %d should be between 1 and %d for date %s", month, numDays, dStr)
	}

	date := config.NewDate(uint(year), uint8(month), uint8(day))
	return &date, nil
}