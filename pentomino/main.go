package main

import (
	"flag"
	"fmt"
	"os"

	psolver "github.com/fzerorubigd/pentomino-solver"
)

func main() {
	var w, h, count int
	flag.IntVar(&count, "count", -1, "The count of the solution to show before exit, -1 to show all")
	flag.IntVar(&w, "width", 10, "Width of the puzzle")
	flag.IntVar(&h, "height", 6, "Width of the puzzle")
	flag.Parse()
	if w*h != 60 {
		fmt.Println("The size should be 60")
		os.Exit(-1)
	}

	puzzle := psolver.NewMatrix(w, h)
	pie := psolver.New12()
	resp := make(chan *psolver.Matrix, 10)
	psolver.Solve(puzzle, pie, resp)

	i := 1
	for r := range resp {
		fmt.Println(i, "===>")
		fmt.Println(r)
		if count > 0 && i == count {
			os.Exit(0)
		}
		i += 1
	}
}
