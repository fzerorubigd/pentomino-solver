package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mshafiee/jalali"
)

func main() {
	// The daily job solves tomorrow, so the default offset is 1 and the output
	// keys keep their existing names — the workflow already consumes `today` and
	// `jtoday`, and renaming them here would break it for no gain.
	var offset int
	flag.IntVar(&offset, "days", 1, "Days from today to report. 0 is today, 1 is tomorrow")
	flag.Parse()

	now := time.Now().AddDate(0, 0, offset)

	// Gregorian
	fmt.Printf("today=%s\n", now.Format("2006-01-02"))

	// Jalali
	jDate := jalali.ToJalali(now)
	fmt.Printf("jtoday=%s\n", jDate.Format("%Y-%m-%d"))
}
