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
	type args struct {
		d time.Time
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "01.02.2000",
			args: args{
				d: time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
			},
			wantW: `   01*02 03 04 05 06
07 08 09 10 11 12 13
14 15 16 17 18 19 20
21 22 23 24 25 26 27
28 29
`,
		},
		{
			name: "07.03.2018",
			args: args{
				d: time.Date(2018, 7, 3, 12, 30, 0, 0, time.UTC),
			},
			wantW: `                  01
02 03*04 05 06 07 08
09 10 11 12 13 14 15
16 17 18 19 20 21 22
23 24 25 26 27 28 29
30 31
`,
		},
		{
			name: "03.05.2006",
			args: args{
				d: time.Date(2006, 5, 3, 12, 30, 0, 0, time.UTC),
			},
			wantW: `01 02 03*04 05 06 07
08 09 10 11 12 13 14
15 16 17 18 19 20 21
22 23 24 25 26 27 28
29 30 31
`,
		},
		{
			name: "02.08.2010",
			args: args{
				d: time.Date(2010, 8, 2, 12, 30, 0, 0, time.UTC),
			},
			wantW: `                  01
02*03 04 05 06 07 08
09 10 11 12 13 14 15
16 17 18 19 20 21 22
23 24 25 26 27 28 29
30 31
`,
		},
		{
			name: "19.09.2010",
			args: args{
				d: time.Date(2010, 9, 19, 12, 30, 0, 0, time.UTC),
			},
			wantW: `      01 02 03 04 05
06 07 08 09 10 11 12
13 14 15 16 17 18 19*
20 21 22 23 24 25 26
27 28 29 30
`,
		},
		{
			name: "29.12.2013",
			args: args{
				d: time.Date(2013, 12, 29, 12, 30, 0, 0, time.UTC),
			},
			wantW: `                  01
02 03 04 05 06 07 08
09 10 11 12 13 14 15
16 17 18 19 20 21 22
23 24 25 26 27 28 29*
30 31
`,
		},
		{
			name: "01.10.2014",
			args: args{
				d: time.Date(2014, 9, 1, 12, 30, 0, 0, time.UTC),
			},
			wantW: `01*02 03 04 05 06 07
08 09 10 11 12 13 14
15 16 17 18 19 20 21
22 23 24 25 26 27 28
29 30
`,
		},
		{
			name: "30.04.2023",
			args: args{
				d: time.Date(2023, 4, 30, 12, 30, 0, 0, time.UTC),
			},
			wantW: `               01 02
03 04 05 06 07 08 09
10 11 12 13 14 15 16
17 18 19 20 21 22 23
24 25 26 27 28 29 30*
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			printMonth(w, tt.args.d)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("\n got: %q\nwant: %q", gotW, tt.wantW)
			}
		})
	}
}
