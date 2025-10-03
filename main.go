package main

import (
	"fmt"
	"sync"
)

type Matrix struct {
	Width  int
	Height int

	data []byte
}

type Point struct {
	X, Y int
}

type Piece interface {
	Name() byte
	States() int
	Position(ref Point, state int) ([5]Point, error)
}

func NewMatrix(w, h int) *Matrix {
	m := Matrix{
		Width:  w,
		Height: h,
		data:   make([]byte, w*h),
	}

	return &m
}

func (m *Matrix) String() string {
	res := ""
	for j := 0; j < m.Height; j++ {
		for i := 0; i < m.Width; i++ {
			if m.data[j*m.Width+i] == 0 {
				res += "."
			} else {
				res += string(m.data[j*m.Width+i])
			}
		}
		res += "\n"
	}

	return res
}

func (m *Matrix) isFull() bool {
	for i := range m.data {
		if m.data[i] == 0 {
			return false
		}
	}
	return true
}

func (m *Matrix) findFirstEmpty() (Point, bool) {
	for j := 0; j < m.Height; j++ {
		for i := 0; i < m.Width; i++ {
			if m.data[j*m.Width+i] == 0 {
				return Point{X: i, Y: j}, true
			}
		}
	}
	return Point{}, false
}

func (m *Matrix) canPlace(p Piece, pos Point, state int) bool {
	points, err := p.Position(pos, state)
	if err != nil {
		return false
	}

	for _, pt := range points {
		if pt.X < 0 || pt.X >= m.Width || pt.Y < 0 || pt.Y >= m.Height {
			return false
		}
		if m.data[pt.Y*m.Width+pt.X] != 0 {
			return false
		}
	}
	return true
}

func (m *Matrix) place(p Piece, pos Point, state int) error {
	if !m.canPlace(p, pos, state) {
		return fmt.Errorf("can not place %s in %d : (%d, %d)", string(p.Name()), state, pos.X, pos.Y)
	}

	points, err := p.Position(pos, state)
	if err != nil {
		return nil
	}

	for _, pt := range points {
		m.data[pt.Y*m.Width+pt.X] = p.Name()
	}

	return nil
}

func (m *Matrix) remove(p Piece) {
	name := p.Name()
	for j := 0; j < m.Height; j++ {
		for i := 0; i < m.Width; i++ {
			if m.data[j*m.Width+i] == name {
				m.data[j*m.Width+i] = 0
			}
		}
	}
}

func minus(p []Piece, idx int) []Piece {
	if len(p) < 4 {
		//fmt.Println(".")
	}
	nP := make([]Piece, 0, len(p)-1)
	for i := range p {
		if idx != i {
			nP = append(nP, p[i])
		}
	}

	return nP
}

func solve(m *Matrix, p []Piece, ans chan string) {
	// return condition
	if len(p) == 0 {
		if m.isFull() {
			ans <- m.String()
		}
	}

	empty, has := m.findFirstEmpty()
	if !has {
		return
	}
	for idx, current := range p {
		rest := minus(p, idx)
		states := current.States()
		for s := range states {
			if err := m.place(current, empty, s); err == nil {
				//fmt.Println(m.String())
				solve(m, rest, ans)
				m.remove(current)
			}
		}
	}
}

func main() {
	pieces := []Piece{
		piecesI{},
		piecesL{},
		piecesP{},
		piecesT{},
		piecesU{},
		piecesV{},
		piecesW{},
		piecesX{},
		piecesY{},
		piecesZ{},
		piecesF{},
		piecesN{},
	}
	wg := sync.WaitGroup{}

	resp := make(chan string, 10)
	for i, current := range pieces {
		wg.Add(1)
		rest := minus(pieces, i)
		go func(ans chan string, main Piece, rest []Piece, w, h int) {
			defer wg.Done()
			m := NewMatrix(w, h)
			states := main.States()
			for s := range states {
				if m.canPlace(main, Point{X: 0, Y: 0}, s) {
					m.place(main, Point{0, 0}, s)
					solve(m, rest, ans)
					m.remove(main)
				}
			}

		}(resp, current, rest, 10, 6)

	}

	go func() {
		wg.Wait()
		close(resp)
	}()
	// m := NewMatrix(10, 6)
	// m.place(pieces[0], Point{0, 0}, 1)
	// go func() {
	// 	solve(m, pieces[1:], resp)
	// 	close(resp)
	// }()

	i := 1
	for r := range resp {
		fmt.Println(i, "===>")
		fmt.Println(r)
		i += 1
	}

	// m.place(piecesI{}, Point{0, 1}, 1)
	// m.place(piecesZ{}, Point{0, 0}, 0)
	// m.place(piecesU{}, Point{2, 0}, 2)
	// m.place(piecesL{}, Point{5, 0}, 5)
	// m.place(piecesT{}, Point{9, 0}, 1)
	// m.place(piecesW{}, Point{3, 1}, 2)
	// m.place(piecesX{}, Point{6, 1}, 0)
	// m.place(piecesF{}, Point{8, 2}, 1)
	// m.place(piecesV{}, Point{1, 3}, 0)
	// m.place(piecesP{}, Point{2, 3}, 1)
	// m.place(piecesY{}, Point{5, 4}, 0)
	// m.place(piecesN{}, Point{6, 4}, 6)

	// fmt.Println(m.isFull())
	// fmt.Println(m.String())
}
