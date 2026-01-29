package main

import (
	"fmt"
	"time"

	"github.com/mshafiee/jalali"
)

func main() {
	now := time.Now().AddDate(0, 0, 1)

	// Gregorian
	fmt.Printf("today=%s\n", now.Format("2006-01-02"))

	// Jalali
	jDate := jalali.ToJalali(now)
	fmt.Printf("jtoday=%s\n", jDate.Format("%Y-%m-%d"))
}
