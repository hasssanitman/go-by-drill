// Go offers extensive support for times and durations; here are
// some examples.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// We’ll start by getting the current time.
	now := time.Now()
	p(now)

	// You can build a time struct by providing the year, month, day,
	// etc. Times are always associated with a Location, i.e. time
	// zone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	p(then)

	// You can extract the various components of the time value as
	// expected.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// The Monday-Sunday Weekday is also available.
	p(then.Weekday())

	// These methods compare two times, testing if the first occurs
	// before, after, or at the same time as the second, respectively.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// The Sub methods returns a Duration representing the interval
	// between two times.
	diff := now.Sub(then)
	p(diff)

	// We can compute the length of the duration in various units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// You can use Add to advance a time by a given duration, or with
	// a - to move backwards by a duration.
	p(then.Add(diff))
	p(then.Add(-diff))

}

// $ go run time/time.go
// 2023-05-29 16:40:03.873805766 +0330 +0330 m=+0.000020172
// 2009-11-17 20:34:58.651387237 +0000 UTC
// 2009
// November
// 17
// 20
// 34
// 58
// 651387237
// UTC
// Tuesday
// true
// false
// false
// 118576h35m5.222418529s
// 118576.58478400514
// 7.114595087040309e+06
// 4.2687570522241855e+08
// 426875705222418529
// 2023-05-29 13:10:03.873805766 +0000 UTC
// 1996-05-09 03:59:53.428968708 +0000 UTC

// Next we’ll look at the related idea of time relative to the Unix epoch.
