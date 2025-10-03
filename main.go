package main

import (
	"fmt"

	"github.com/fzerorubigd/pentomino-solver/psolver"
)

func main() {
	cal := psolver.NewPersianCalendar()
	//mm := psolver.NewMatrix(10, 6)
	pie := psolver.New12()

	cal.SetDate(7, 11, 7, 1404)
	resp := make(chan *psolver.Matrix, 10)
	// go func() {
	// 	psolver.SolveSingle(&cal.Matrix, pie, resp)
	// 	close(resp)
	// }()
	psolver.Solve(&cal.Matrix, pie, resp)
	i := 1
	for r := range resp {
		fmt.Println(i, "===>")
		fmt.Println(r)
		i += 1
	}
}
