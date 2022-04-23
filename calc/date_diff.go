package calc

import (
	"github.com/pandulaDW/date-diff/config"
)

// getNumDaysSinceBaseDate returns the number of days passed since the base-date.
//
// This function assumes that the base-date is set to be {config.BaseYear}/01/01.
func getNumDaysSinceBaseDate(date *config.Date) uint {
	var numDays uint

	// calculate days until the current year
	currentYear := config.BaseYear
	for currentYear < date.Year {
		if isLeapYear(currentYear) {
			numDays += 366
		} else {
			numDays += 365
		}
		currentYear++
	}

	// calculate days until the current month
	var currentMonth uint8 = 1
	for currentMonth < date.Month {
		numDays += uint(numDaysPerMonth(currentMonth, date.Year))
		currentMonth += 1
	}

	return numDays + uint(date.Day)
}

// DateDiff returns the whole days between two dates counting only
// the days in between those dates
func DateDiff(date1 *config.Date, date2 *config.Date) uint {
	diff1FromBase := getNumDaysSinceBaseDate(date1)
	diff2FromBase := getNumDaysSinceBaseDate(date2)

	if diff1FromBase < diff2FromBase {
		return diff2FromBase - diff1FromBase - 1
	}
	return diff1FromBase - diff2FromBase - 1
}
