// Go supports time formatting and parsing via pattern-based
// layouts.
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Here’s a basic example of formatting a time according to
	// RFC3339, using the corresponding layout constant.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Time parsing uses the same layout values as Format.
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)
	p(t1)

	// Format and Parse use example-based layouts. Usually you’ll
	// use a constant from time for these layouts, but you can also
	// supply custom layouts. Layouts must use the reference time Mon
	// Jan 2 15:04:05 MST 2006 to show the pattern with which to
	// format/parse a given time/string. The example time must be
	// exactly as shown: the year 2006, 15 for the hour, Monday for
	// the day of the week, etc.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01002T15:04:05.999999-07:00"))
	from := "3 04 PM"
	t2, _ := time.Parse(from, "8 41 PM")
	p(t2)

	// For purely numeric representations you can also use standard
	// string formatting with the extracted components of the time
	// value.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Parse will return an error on malformed input explaining the
	// parsing problem.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")

	p(e)

}

// $ go run time_formatting_parsing/time_formation_parsing.go
// 2023-05-29T17:03:28+03:30
// 2012-11-01 22:08:41 +0000 +0000
// 5:03PM
// Mon May 29 17:03:28 2023
// 2023-05149T17:03:28.885935+03:30
// 0000-01-01 20:41:00 +0000 UTC
// 2023-05-29T17:03:28-00:00
// parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
