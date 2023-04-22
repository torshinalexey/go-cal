package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	today := time.Now()
	printHeader(os.Stdout, today)
	printMonth(os.Stdout, today)
}

func printHeader(w io.Writer, d time.Time) {
	weekdays := "Mo Tu We Th Fr Sa Su"
	fmt.Fprintf(w, "%d %s\n%s\n%s\n",
		d.Year(), d.Month(), weekdays, strings.Repeat("-", len(weekdays)),
	)
}

// printMonth prints a calendar month to the given writer for the given date.
func printMonth(w io.Writer, d time.Time) {
	// Iterate over each day in the month.
	for day := d.AddDate(0, 0, -d.Day()+1); day.Month() == d.Month(); day = day.AddDate(0, 0, 1) {
		// If this is the first day of the month and it is not a Monday, add padding to align the first week.
		if day.AddDate(0, 0, -1).Month() != day.Month() && day.Weekday() != time.Monday {
			fmt.Fprint(w, strings.Repeat("   ", int(day.Weekday())-1))
		}
		// Print the day number with an asterisk if it is the current day.
		if day.Day() == d.Day() {
			fmt.Fprintf(w, "%.2d*", day.Day())
			// Print the day number with linebreak if it is Sunday.
		} else if day.Weekday() == time.Sunday {
			fmt.Fprintf(w, "%.2d\n", day.Day())
			// Print the day number with a trailing space.
		} else {
			fmt.Fprintf(w, "%.2d ", day.Day())
		}
	}
}
