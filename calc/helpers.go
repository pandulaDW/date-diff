package calc

import "github.com/pandulaDW/date-diff/config"

// isLeapYear returns true if the given year is a leap year.
// Returns false otherwise.
func isLeapYear(year uint) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// numDaysPerMonth returns the number of days for a given month after
// checking if the year is a leap year
func numDaysPerMonth(month uint8, year uint) uint8 {
	if month == 2 && isLeapYear(year) {
		return 29
	}
	return config.MonthDayMap[month]
}
