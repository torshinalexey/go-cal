package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	today := time.Now()
	printHeader(today)
	printMonth(today)
}

func printHeader(today time.Time) {
	weekdays := "Mo Tu We Th Fr Sa Su"
	fmt.Printf(
		"%d %s\n%s\n%s\n",
		today.Year(), today.Month(), weekdays, strings.Repeat("-", len(weekdays)),
	)
}

func printMonth(today time.Time) {
	for day := today.AddDate(0, 0, -today.Day()+1); day.Month() == today.Month(); day = day.AddDate(0, 0, 1) {
		if day.AddDate(0, 0, -1).Month() != day.Month() && day.Weekday() != time.Monday {
			fmt.Print(strings.Repeat("   ", int(day.Weekday())-1))
		}
		if day.Day() == today.Day() {
			fmt.Printf("\033[46m%.2d\033[0m ", day.Day())
		} else {
			fmt.Printf("%.2d ", day.Day())
		}
		if day.Weekday() == time.Sunday {
			fmt.Println()
		}
		if day.AddDate(0, 0, 1).Month() != day.Month() && day.Weekday() != time.Sunday {
			fmt.Println(strings.Repeat("   ", 7-int(day.Weekday())))
		}
	}
}
