package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	psolver "github.com/fzerorubigd/pentomino-solver"
	"github.com/mshafiee/jalali"
)

func getTomorrow(j bool) (int, int, int, int) {
	date := time.Now().AddDate(0, 0, 1)
	if j {
		jDate := jalali.ToJalali(date)

		wd := int(jDate.Weekday()) + 2
		if wd > 7 {
			wd -= 7
		}
		return wd, jDate.Day(), int(jDate.Month()), jDate.Year() - 1403
	}
	return int(date.Weekday()) + 1, date.Day(), int(date.Month()), date.Year() - 2024

}

func main() {
	var W, D, M, Y, count int
	var color, svg, tomorrow, jalaliDate bool
	var outputDir string
	flag.IntVar(&W, "weekday", 1, "The weekday, 1 for the first day, 7 for 7th day. (Shanbe is first for Persian, Monday for Gregorian)")
	flag.IntVar(&D, "day", 1, "The day of the month, 1 to 31")
	flag.IntVar(&M, "month", 1, "The month, 1 to 12")
	flag.IntVar(&Y, "year", 1, "The year of the calendar, for Persian 1=1404 and for Gregorian 1=2025, max 10")
	flag.IntVar(&count, "count", -1, "The count of the solution to show before exit, -1 to show all")
	flag.BoolVar(&color, "color", true, "Use color output")
	flag.BoolVar(&svg, "svg", false, "Output SVG files (1.svg, 2.svg, ...)")
	flag.StringVar(&outputDir, "output-dir", "", "Output directory for SVG files")
	flag.BoolVar(&tomorrow, "tomorrow", false, "Output tomorrow's calendar, ignore all other date related flags")

	flag.BoolVar(&jalaliDate, "jalali", false, "Use jalali calendar")
	flag.Parse()

	var exporter psolver.Exporter
	if svg {
		exporter = psolver.NewSVGExporter()
	} else if color {
		exporter = psolver.NewColorStringExporter()
	} else {
		exporter = &psolver.StringExporter{}
	}

	if tomorrow {
		W, D, M, Y = getTomorrow(jalaliDate)
	}

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

		if svg {
			fileName := filepath.Join(outputDir, fmt.Sprintf("%d.svg", i))
			f, err := os.Create(fileName)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", fileName, err)
				continue
			}
			if err := exporter.Export(r, f); err != nil {
				fmt.Printf("Error exporting to %s: %v\n", fileName, err)
			}
			f.Close()
			fmt.Printf("Exported %s\n", fileName)
		} else {
			fmt.Println(i, "===>")
			if err := exporter.Export(r, os.Stdout); err != nil {
				fmt.Println("Error exporting:", err)
			}
		}

		if count > 0 && i == count {
			os.Exit(0)
		}
		i += 1
	}
}
