package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

func printMonth(w io.Writer, target time.Time) {
	firstDayOfMonthWeekday := target.AddDate(0, 0, -target.Day()+1).Weekday()
	lastDayOfMonth := target.AddDate(0, 1, -target.Day())
	var monthBuilder strings.Builder
	monthBuilder.Grow(100)
	if firstDayOfMonthWeekday != time.Monday {
		wDay := firstDayOfMonthWeekday
		if wDay == time.Sunday {
			wDay = 7
		}
		for i := 0; i < (int(wDay)-1)*3; i++ {
			monthBuilder.WriteRune(' ')
		}
	}
	curWeekDay := firstDayOfMonthWeekday
	for day := 1; day <= lastDayOfMonth.Day(); day++ {
		if day < 10 {
			monthBuilder.WriteRune('0')
		}
		monthBuilder.WriteString(strconv.Itoa(day))
		if day == target.Day() {
			monthBuilder.WriteRune('*')
		}
		if curWeekDay == time.Sunday ||
			day == lastDayOfMonth.Day() {
			monthBuilder.WriteRune('\n')
		} else if day != target.Day() {
			monthBuilder.WriteRune(' ')
		}
		if curWeekDay == time.Saturday {
			curWeekDay = time.Sunday
			continue
		}
		curWeekDay++
	}
	// monthBuilder.WriteRune('\n')
	if _, err := w.Write([]byte(monthBuilder.String())); err != nil {
		log.Fatal(err)
	}
}
