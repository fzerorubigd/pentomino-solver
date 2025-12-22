package psolver

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

type Matrix struct {
	Width  int
	Height int

	data   []byte
	pieces map[byte]int
}

func (m *Matrix) At(x, y int) byte {
	return m.data[y*m.Width+x]
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
		pieces: make(map[byte]int, 12),
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

func (m *Matrix) Hash() string {
	hash := sha1.New()
	return fmt.Sprintf("%x", hash.Sum(m.data))
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
	m.pieces[p.Name()] = state

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

func (m *Matrix) duplicate() *Matrix {
	n := NewMatrix(m.Width, m.Height)
	copy(n.data, m.data)
	for i := range m.pieces {
		n.pieces[i] = m.pieces[i]
	}

	return n
}

func (m *Matrix) Pieces() map[byte]int {
	return m.pieces
}

func Minus(p []Piece, idx int) []Piece {
	nP := make([]Piece, 0, len(p)-1)
	for i := range p {
		if idx != i {
			nP = append(nP, p[i])
		}
	}

	return nP
}

func SolveSingle(m *Matrix, p []Piece, ans chan *Matrix) {
	// return condition
	if len(p) == 0 {
		if m.isFull() {
			ans <- m.duplicate()
		}
	}

	empty, has := m.findFirstEmpty()
	if !has {
		return
	}
	for idx, current := range p {
		rest := Minus(p, idx)
		states := current.States()
		for s := range states {
			if err := m.place(current, empty, s); err == nil {
				//fmt.Println(m.String())
				SolveSingle(m, rest, ans)
				m.remove(current)
			}
		}
	}
}

func Solve(m *Matrix, pieces []Piece, ans chan *Matrix) {
	wg := sync.WaitGroup{}
	wg.Add(len(pieces))
	for i, current := range pieces {
		rest := Minus(pieces, i)
		go func(ans chan *Matrix, main Piece, rest []Piece) {
			defer func() {
				wg.Done()
			}()
			fs := m.duplicate()
			start, _ := fs.findFirstEmpty()
			states := main.States()
			for s := range states {
				if fs.canPlace(main, start, s) {
					fs.place(main, start, s)
					SolveSingle(fs, rest, ans)
					fs.remove(main)
				}
			}

		}(ans, current, rest)

	}

	go func() {
		wg.Wait()
		close(ans)
	}()
}
