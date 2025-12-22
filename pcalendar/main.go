package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
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
		printColoredMatrix(r)
		if count > 0 && i == count {
			os.Exit(0)
		}
		i += 1
	}
}


func printColoredMatrix(m *psolver.Matrix) {
	for j := 0; j < m.Height; j++ {
		for i := 0; i < m.Width; i++ {
			char := m.At(i, j) 
			
			style, ok := colorMap[char]
			if !ok {
				style = lipgloss.NewStyle() // Default white
			}
			
			fmt.Print(style.Render("██"))
		}
		fmt.Println()
	}
}

var colorMap = map[byte]lipgloss.Style{
	'I': lipgloss.NewStyle().Foreground(lipgloss.Color("#e6194B")), // Rose Red
	'L': lipgloss.NewStyle().Foreground(lipgloss.Color("#3cb44b")), // Green
	'P': lipgloss.NewStyle().Foreground(lipgloss.Color("#ffe119")), // Yellow
	'T': lipgloss.NewStyle().Foreground(lipgloss.Color("#4363d8")), // Blue
	'U': lipgloss.NewStyle().Foreground(lipgloss.Color("#f58231")), // Orange
	'V': lipgloss.NewStyle().Foreground(lipgloss.Color("#911eb4")), // Purple
	'W': lipgloss.NewStyle().Foreground(lipgloss.Color("#42d4f4")), // Cyan
	'X': lipgloss.NewStyle().Foreground(lipgloss.Color("#f032e6")), // Magenta
	'Y': lipgloss.NewStyle().Foreground(lipgloss.Color("#bfef45")), // Lime
	'Z': lipgloss.NewStyle().Foreground(lipgloss.Color("#fabed4")), // Pink
	'F': lipgloss.NewStyle().Foreground(lipgloss.Color("#469990")), // Teal
	'N': lipgloss.NewStyle().Foreground(lipgloss.Color("#dcbeff")), // Lavender
	'O': lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")), // Dark Gray for Date cells
	'.': lipgloss.NewStyle().Foreground(lipgloss.Color("#222222")), // Near black for empty
}

