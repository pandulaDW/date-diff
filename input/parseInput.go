package input

import (
	"github.com/pandulaDW/date-diff/calc"
	"github.com/pandulaDW/date-diff/config"
)

// ParseDates takes the two date strings given by the user, validates them
// and return the two parsed dates.
func ParseDates(dStr1, dStr2 string) (*config.Date, *config.Date, error) {
	date1, err := calc.ParseDate(dStr1)
	if err != nil {
		return nil, nil, err
	}
	date2, err := calc.ParseDate(dStr2)
	if err != nil {
		return nil, nil, err
	}
	return date1, date2, nil
}
