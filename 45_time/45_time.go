// 45_time.go
// Topic: time package — Time, Duration, format, parse, timers
//
// CORE TYPES
//   time.Time      — instant in time
//   time.Duration  — int64 nanoseconds (with units: time.Second, time.Millisecond)
//   time.Location  — timezone
//
// CONSTRUCTION
//   t := time.Now()
//   t := time.Date(2025, time.March, 14, 13, 37, 0, 0, time.UTC)
//   t, err := time.Parse(layout, "2025-03-14")
//
// FORMAT / PARSE — uses REFERENCE TIME, not strftime
//   Reference: Mon Jan 2 15:04:05 MST 2006   (= 01/02 03:04:05PM '06 -0700)
//   Memorize: 1 2 3 4 5 6 7 -> Jan 2 03:04:05 2006 (MST is -7)
//
// PRESETS
//   time.RFC3339      "2006-01-02T15:04:05Z07:00"
//   time.RFC822, time.Kitchen, time.Stamp, ...
//   time.DateOnly = "2006-01-02"   (Go 1.20+)
//   time.TimeOnly = "15:04:05"
//
// DURATIONS
//   d := 3 * time.Second
//   d.Seconds(), d.Milliseconds()
//   time.Sleep(d)
//
// ARITHMETIC
//   t.Add(d)            Time + Duration -> Time
//   t.Sub(t2)           Time - Time -> Duration
//   t.AddDate(y, m, d)  calendar
//   t.Before(t2), After, Equal
//
// TIMEZONE
//   loc, _ := time.LoadLocation("America/New_York")
//   t.In(loc)
//
// TIMERS / TICKERS
//   <-time.After(d)
//   timer := time.NewTimer(d); <-timer.C; timer.Stop()
//   ticker := time.NewTicker(d); <-ticker.C; defer ticker.Stop()
//
// MEASURE
//   start := time.Now()
//   ... work ...
//   elapsed := time.Since(start)
//
// MONOTONIC vs WALL
//   time.Now() includes monotonic reading. Sub() uses it for accuracy across
//   wall-clock changes. Strip with t.Round(0) when comparing across processes.
//
// Run: go run 45_time.go

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now:", now)
	fmt.Println("unix:", now.Unix(), "unix nano:", now.UnixNano())

	// Build a time
	t := time.Date(2025, time.March, 14, 13, 37, 0, 0, time.UTC)
	fmt.Println("t:", t)

	// Format
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("Mon, 02 Jan 06 15:04 MST"))

	// Parse
	parsed, err := time.Parse("2006-01-02", "2025-12-31")
	fmt.Println(parsed, err)

	// Duration
	d := 2*time.Hour + 30*time.Minute
	fmt.Println("duration:", d, d.Minutes())

	// Arithmetic
	tomorrow := now.Add(24 * time.Hour)
	diff := tomorrow.Sub(now)
	fmt.Println("diff:", diff)

	nextMonth := now.AddDate(0, 1, 0)
	fmt.Println("next month:", nextMonth)

	// Before / After
	fmt.Println(now.Before(tomorrow))

	// Timezone
	if loc, err := time.LoadLocation("Asia/Tokyo"); err == nil {
		fmt.Println("Tokyo:", now.In(loc))
	}

	// Sleep
	time.Sleep(50 * time.Millisecond)

	// Measure
	start := time.Now()
	time.Sleep(20 * time.Millisecond)
	fmt.Println("elapsed:", time.Since(start))

	// Timer + Ticker (brief demo)
	timer := time.NewTimer(30 * time.Millisecond)
	<-timer.C
	fmt.Println("timer fired")

	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < 3; i++ {
		<-ticker.C
		fmt.Println("tick", i)
	}
}
