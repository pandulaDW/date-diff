package input

import (
	"bufio"
	"fmt"
	"github.com/pandulaDW/date-diff/time"
	"os"
	"strings"
)

// ParseUserInput reads the two dates from the standard input,
// parses the dates, calculates the whole day difference and sends the output to the stdout.
//
// If an error encounters in parsing the dates, it will display the error in stderr.
func ParseUserInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter date 1: ")
	inputStr1, _ := reader.ReadString('\n')
	d1, err := time.ParseDate(strings.TrimSpace(inputStr1))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error in parsing date:\n\t%s", err.Error())
		os.Exit(1)
	}

	fmt.Print("Enter date 2: ")
	inputStr2, _ := reader.ReadString('\n')
	d2, err := time.ParseDate(strings.TrimSpace(inputStr2))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error in parsing date:\n\t%s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("The date difference is %d days\n", time.DateDiff(d1, d2))
}
