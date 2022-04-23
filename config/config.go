package config

import "regexp"

// DateFormat specifies the regex syntax of date format expected by the
// user's date input
var DateFormat *regexp.Regexp

// MonthDayMap gives the number of days for a given month.
//
// Number of days for February can be 29, based on the year.
var MonthDayMap map[uint8]uint8

func init() {
	MonthDayMap = map[uint8]uint8{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	DateFormat = regexp.MustCompile("^\\d{1,2}/\\d{1,2}/\\d{4}$")
}

// BaseYear represents the lowest year that the application allows
const BaseYear uint = 1900

// MaxYear represents the highest year that the application allows
const MaxYear uint = 2999

// Date structure represents an any given date
type Date struct {
	Year  uint
	Month uint8
	Day   uint8
}

// NewDate returns a new Date instance based on the given year, month and day
func NewDate(year uint, month, day uint8) Date {
	return Date{Year: year, Month: month, Day: day}
}
