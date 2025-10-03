package psolver

import "fmt"

type PersianCalendar struct {
	Matrix
}

var mapping = []string{
	"W1", "W7", "D1", "D2", "D3", "D4", "D5", "D6", "M1", "M2",
	"W2", "D7", "D8", "D9", "D10", "D11", "D12", "D13", "M3", "M4",
	"W3", "D14", "D15", "D16", "D17", "D18", "D19", "D20", "M5", "M6",
	"W4", "D21", "D22", "D23", "D24", "D25", "D26", "D27", "M7", "M8",
	"W5", "D28", "D29", "D30", "D31", "Y1", "Y2", "Y3", "M9", "M10",
	"W6", "Y4", "Y5", "Y6", "Y7", "Y8", "Y9", "Y10", "M11", "M12",
	"E1", "E2", "E3", "E4",
}

func (p *PersianCalendar) SetDate(WD, D, M, Y int) error {
	w := fmt.Sprintf("W%d", WD)
	m := fmt.Sprintf("M%d", M)
	d := fmt.Sprintf("D%d", D)
	y := fmt.Sprintf("Y%d", Y-1403)
	data := [4]int{-1, -1, -1, -1}
	for i := range mapping {
		if mapping[i] == w {
			data[0] = i
		}

		if mapping[i] == m {
			data[1] = i
		}

		if mapping[i] == d {
			data[2] = i
		}

		if mapping[i] == y {
			data[3] = i
		}
	}

	for i := range data {
		if data[i] < 0 {
			return fmt.Errorf("invalid input")
		}
	}

	for i := range p.Matrix.data {
		p.Matrix.data[i] = 0
	}
	p.Matrix.data[data[0]] = 'O'
	p.Matrix.data[data[1]] = 'O'
	p.Matrix.data[data[2]] = 'O'
	p.Matrix.data[data[3]] = 'O'

	p.Matrix.data[64] = 'O'
	p.Matrix.data[65] = 'O'
	p.Matrix.data[66] = 'O'
	p.Matrix.data[67] = 'O'
	p.Matrix.data[68] = 'O'
	p.Matrix.data[69] = 'O'

	return nil
}

func NewPersianCalendar() *PersianCalendar {
	m := NewMatrix(10, 7)
	p := &PersianCalendar{
		Matrix: *m,
	}

	return p
}
