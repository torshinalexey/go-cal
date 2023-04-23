package main

import (
	"bytes"
	"testing"
	"time"
)

func Benchmark_printHeader(b *testing.B) {
	var buf bytes.Buffer
	d := time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)
	for n := 0; n < b.N; n++ {
		buf.Reset()
		b.ReportAllocs()
		printHeader(&buf, d)
	}
}

func Benchmark_printMonth(b *testing.B) {
	var buf bytes.Buffer
	d := time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)
	for n := 0; n < b.N; n++ {
		buf.Reset()
		b.ReportAllocs()
		printMonth(&buf, d)
	}
}
