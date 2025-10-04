package main

import (
	"flag"
	"fmt"
	"os"

	psolver "github.com/fzerorubigd/pentomino-solver"
)

func main() {
	var W, D, M, Y, count int
	flag.IntVar(&W, "weekday", 1, "The weekday, 1 for the first day, 7 for 7th day. (Shanbe is first for Persian, Monday for Gregorian)")
	flag.IntVar(&D, "day", 1, "The day of the month, 1 to 31")
	flag.IntVar(&M, "month", 1, "The month, 1 to 12")
	flag.IntVar(&Y, "year", 1, "The year of the calendar, for Persian 1=1404 and for Gregorian 1=2025, max 10")
	flag.IntVar(&count, "count", -1, "The count of the solution to show before exit, -1 to show all")
	flag.Parse()

	cal := psolver.NewPersianCalendar()
	pie := psolver.New12()
	cal.SetDate(W, D, M, Y+1403)
	resp := make(chan *psolver.Matrix, 10)
	psolver.Solve(&cal.Matrix, pie, resp)

	mm := map[string]struct{}{}
	i := 1
	for r := range resp {
		if _, ok := mm[r.Hash()]; ok {
			continue
		}
		mm[r.Hash()] = struct{}{}
		fmt.Println(i, "===>")
		fmt.Println(r)
		if count > 0 && i == count {
			os.Exit(0)
		}
		i += 1
	}
}
