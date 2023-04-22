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
	var b strings.Builder
	b.Grow(100)
	// Iterate over each day in the month.
	for day := d.AddDate(0, 0, -d.Day()+1); day.Month() == d.Month(); day = day.AddDate(0, 0, 1) {
		curDay := day.Day()
		// If this is the first day of the month and it is not a Monday, add padding to align the first week.
		if day.AddDate(0, 0, -1).Month() != day.Month() && day.Weekday() != time.Monday {
			b.WriteString(strings.Repeat("   ", int(day.Weekday())-1))
		}
		// Print the day number with an asterisk if it is the current day.
		if curDay == d.Day() {
			b.WriteString(fmt.Sprintf("%.2d*", curDay))
			continue
		}
		// Print the day number with line break if it is Sunday.
		if day.Weekday() == time.Sunday {
			b.WriteString(fmt.Sprintf("%.2d\n", curDay))
			continue
		}
		// Print the day number with a trailing space.
		b.WriteString(fmt.Sprintf("%.2d ", curDay))
	}
	fmt.Fprint(w, b.String())
}
