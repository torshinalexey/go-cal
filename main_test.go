package main

import (
	"bytes"
	"testing"
	"time"
)

func Test_printHeader(t *testing.T) {
	var (
		b    bytes.Buffer
		d    = time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
		want = `2000 February
Mo Tu We Th Fr Sa Su
--------------------
`
	)

	printHeader(&b, d)
	got := b.String()
	if got != want {
		t.Fatalf("\nwant: %q\ngot: %q\n", want, got)
	}
}

func Test_printMonth(t *testing.T) {
	var (
		b    bytes.Buffer
		d    = time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
		want = `   01*02 03 04 05 06
07 08 09 10 11 12 13
14 15 16 17 18 19 20
21 22 23 24 25 26 27
28 29 `
	)
	printMonth(&b, d)
	got := b.String()
	if got != want {
		t.Fatalf("\nwant: %q\n got: %q\n", want, got)
	}
}
