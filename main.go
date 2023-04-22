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

// printHeader prints a line with the year, month, and abbreviated weekday names,
// followed by a line of dashes with the same length as the weekday names.
func printHeader(w io.Writer, d time.Time) {
	weekdays := "Mo Tu We Th Fr Sa Su"
	fmt.Fprintf(w, "%d %s\n%s\n%s\n",
		d.Year(), d.Month(), weekdays, strings.Repeat("-", len(weekdays)),
	)
}

// printMonth prints a calendar month to the given writer for the given date.
func printMonth(w io.Writer, d time.Time) {
	var monthRepr strings.Builder
	monthRepr.Grow(100)
	for day := d.AddDate(0, 0, -d.Day()+1); day.Month() == d.Month(); day = day.AddDate(0, 0, 1) {
		dayOfMonth := day.Day()
		weekday := day.Weekday()
		if day.AddDate(0, 0, -1).Month() != day.Month() &&
			weekday != time.Monday {
			n := int(weekday) - 1
			if n < 0 {
				n = 6
			}
			monthRepr.WriteString(strings.Repeat("   ", n))
		}
		var dayRepr strings.Builder
		dayRepr.Grow(3)
		dayRepr.WriteString(fmt.Sprintf("%.2d", dayOfMonth))
		switch {
		case dayOfMonth == d.Day() && weekday == time.Sunday:
			dayRepr.WriteString("*\n")
		case dayOfMonth == d.Day():
			dayRepr.WriteString("*")
		case weekday == time.Sunday,
			day.AddDate(0, 0, 1).Month() != day.Month():
			dayRepr.WriteRune('\n')
		default:
			dayRepr.WriteRune(' ')
		}
		monthRepr.WriteString(dayRepr.String())
	}
	fmt.Fprint(w, monthRepr.String())
}
