package psolver

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidState = errors.New("invalid state")
)

type NamedPiece string

const (
	PieceF NamedPiece = "F"
	PieceI NamedPiece = "I"
	PieceL NamedPiece = "L"
	PieceN NamedPiece = "N"
	PieceP NamedPiece = "P"
	PieceT NamedPiece = "T"
	PieceU NamedPiece = "U"
	PieceV NamedPiece = "V"
	PieceW NamedPiece = "W"
	PieceX NamedPiece = "X"
	PieceY NamedPiece = "Y"
	PieceZ NamedPiece = "Z"
)

type piecesI struct {
}

func (piecesI) Name() byte {
	return 'I'
}

func (piecesI) States() int {
	return 2
}

func (piecesI) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	case 0:
		return [5]Point{
			{x, y},
			{x + 1, y},
			{x + 2, y},
			{x + 3, y},
			{x + 4, y},
		}, nil
	case 1:
		return [5]Point{
			{x, y},
			{x, y + 1},
			{x, y + 2},
			{x, y + 3},
			{x, y + 4},
		}, nil

	default:
		return [5]Point{}, ErrInvalidState
	}
}

type piecesX struct {
}

func (piecesX) Name() byte {
	return 'X'
}

func (piecesX) States() int {
	return 1
}

func (piecesX) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	case 0:
		// . X . \
		// X X X \
		// . X .
		return [5]Point{
			{x, y},
			{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
			{x, y + 2},
		}, nil

	default:
		return [5]Point{}, ErrInvalidState
	}
}

type piecesL struct {
}

func (piecesL) Name() byte {
	return 'L'
}

func (piecesL) States() int {
	return 8
}

func (piecesL) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	// --- Rotational States (Standard 'L' shape) ---
	case 0:
		// 0-degrees (L . . . /
		// 			  L L L L) - 4x2 box
		return [5]Point{
			{x, y},
			{x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1}, {x + 3, y + 1},
		}, nil
	case 1:
		// 90-degrees (L L /
		// 			   L . /
		// 			   L . /
		// 			   L . ) - 2x4 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x, y + 1},
			{x, y + 2},
			{x, y + 3},
		}, nil
	case 2:
		// 180-degrees (L L L L /
		// 				. . . L) - 4x2 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y},
			{x + 3, y + 1},
		}, nil
	case 3:
		// 90-degrees (. L /
		// 			   . L /
		// 			   . L /
		// 			   L L ) - 2x4 box
		return [5]Point{
			{x, y},
			{x, y + 1},
			{x, y + 2},
			{x - 1, y + 3}, {x, y + 3},
		}, nil
	case 4:
		// 90-degrees (L . /
		// 			   L . /
		// 			   L . /
		// 			   L L ) - 2x4 box
		return [5]Point{
			{x, y},
			{x, y + 1},
			{x, y + 2},
			{x, y + 3}, {x + 1, y + 3},
		}, nil
	case 5:
		// 180-degrees (L L L L /
		// 				L . . .) - 4x2 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y},
			{x, y + 1},
		}, nil
	case 6:
		// 90-degrees (L L /
		// 			   . L /
		// 			   . L /
		// 			   . L ) - 2x4 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x + 1, y + 1},
			{x + 1, y + 2},
			{x + 1, y + 3},
		}, nil
	case 7:
		// 0-degrees (. . . L /
		// 			  L L L L) - 4x2 box
		return [5]Point{
			{x, y},
			{x - 3, y + 1}, {x + 2, y + 1}, {x + 1, y + 1}, {x, y + 1},
		}, nil

	default:
		return [5]Point{}, ErrInvalidState
	}
}

type piecesW struct {
}

func (piecesW) Name() byte {
	return 'W'
}

func (piecesW) States() int {
	return 8
}

func (piecesW) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	// --- Rotational States (Standard 'W' shape) ---
	case 0:
		// 0-degrees (W W . /
		// 			 . W W /
		// 			 . . W)
		return [5]Point{
			{x, y}, {x + 1, y},
			{x + 1, y + 1}, {x + 2, y + 1},
			{x + 2, y + 2},
		}, nil
	case 1:
		// 90-degrees ( . . W /
		// 			 	. W W /
		// 				W W . )
		return [5]Point{
			{x, y},
			{x - 1, y + 1}, {x, y + 1},
			{x - 2, y + 2}, {x - 1, y + 2},
		}, nil
	case 2:
		// 180-degrees ( W . . /
		// 				 W W . /
		// 				 . W W )
		return [5]Point{
			{x, y},
			{x, y + 1}, {x + 1, y + 1},
			{x + 1, y + 2}, {x + 2, y + 2},
		}, nil
	case 3:
		// 270-degrees ( . W W /
		// 				 W W . /
		// 				 W . . )
		return [5]Point{
			{x, y}, {x + 1, y},
			{x - 1, y + 1}, {x, y + 1},
			{x - 1, y + 2},
		}, nil

	// --- Reflected States (Flipped 'W' shape, W') ---
	case 4:
		// Flipped 0-degrees ( . W W /
		// 					   W W . /
		// 					   W . . )
		return [5]Point{
			{x, y}, {x + 1, y},
			{x - 1, y + 1}, {x, y + 1},
			{x - 1, y + 2},
		}, nil
	case 5:
		// Flipped 90-degrees ( W . . /
		// 						W W . /
		// 						. W W )
		return [5]Point{
			{x, y},
			{x, y + 1}, {x + 1, y + 1},
			{x + 1, y + 2}, {x + 2, y + 2},
		}, nil
	case 6:
		// Flipped 180-degrees ( . . W /
		// 						 . W W /
		// 						 W W . )
		return [5]Point{
			{x, y},
			{x - 1, y + 1}, {x, y + 1},
			{x - 2, y + 2}, {x - 1, y + 2},
		}, nil
	case 7:
		// Flipped 270-degrees ( W W . /
		// 						 . W W /
		// 						 . . W )
		return [5]Point{
			{x, y}, {x + 1, y},
			{x + 1, y + 1}, {x + 2, y + 1},
			{x + 2, y + 2},
		}, nil

	default:
		// You would need to define ErrInvalidState in your package
		return [5]Point{}, ErrInvalidState
	}
}

type piecesF struct {
}

func (piecesF) Name() byte {
	return 'F'
}

func (piecesF) States() int {
	return 8
}

func (piecesF) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	// --- Rotational States (Standard 'F' shape) ---
	case 0:
		// 0-degrees ( . F F /
		// 			   F F . /
		// 			   . F . ) - 3x3 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x - 1, y + 1}, {x, y + 1},
			{x, y + 2},
		}, nil
	case 1:
		// 90-degrees ( . F . /
		// 				F F F /
		// 				. . F ) - 3x3 box
		return [5]Point{
			{x, y},
			{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
			{x + 1, y + 2},
		}, nil
	case 2:
		// 180-degrees ( . F . /
		// 				 . F F /
		// 				 F F . ) - 3x3 box
		return [5]Point{
			{x, y},
			{x, y + 1}, {x + 1, y + 1},
			{x - 1, y + 2}, {x, y + 2},
		}, nil
	case 3:
		// 270-degrees ( F . . /
		// 				 F F F /
		// 				 . F . ) - 3x3 box
		return [5]Point{
			{x, y},
			{x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1},
			{x + 1, y + 2},
		}, nil

	// --- Reflected States (Flipped 'F' shape, F') ---
	case 4:
		// Flipped 0-degrees ( F F . /
		// 					   . F F /
		// 					   . F . ) - 3x3 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x + 1, y + 1}, {x + 2, y + 1},
			{x + 1, y + 2},
		}, nil
	case 5:
		// Flipped 90-degrees ( . . F /
		// 						F F F /
		// 						. F . ) - 3x3 box
		return [5]Point{
			{x, y},
			{x - 2, y + 1}, {x - 1, y + 1}, {x, y + 1},
			{x - 1, y + 2},
		}, nil
	case 6:
		// Flipped 180-degrees ( . F . /
		// 						 F F . /
		// 						 . F F ) - 3x3 box
		return [5]Point{
			{x, y},
			{x - 1, y + 1}, {x, y + 1},
			{x, y + 2}, {x + 1, y + 2},
		}, nil
	case 7:
		// 90-degrees ( . F . /
		// 				F F F /
		// 				F . . ) - 3x3 box
		return [5]Point{
			{x, y},
			{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
			{x - 1, y + 2},
		}, nil
	default:
		// You would need to define ErrInvalidState in your package
		return [5]Point{}, ErrInvalidState
	}
}

type piecesN struct {
}

func (piecesN) Name() byte {
	return 'N'
}

func (piecesN) States() int {
	return 8
}

func (piecesN) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	// --- Rotational States (Standard 'N' shape) ---
	case 0:
		// 0-degrees ( . . N N /
		// 			   N N N . ) - 4x2 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x - 2, y + 1}, {x - 1, y + 1}, {x, y + 1},
		}, nil
	case 1:
		// 90-degrees ( N . /
		// 				N N /
		// 				. N /
		// 				. N ) - 2x4 box
		return [5]Point{
			{x, y},
			{x, y + 1}, {x + 1, y + 1},
			{x + 1, y + 2},
			{x + 1, y + 3},
		}, nil
	case 2:
		// 180-degrees ( . N N N /
		// 				 N N . . ) - 4x2 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y},
			{x - 1, y + 1}, {x, y + 1},
		}, nil
	case 3:
		// 270-degrees ( N . /
		// 				 N . /
		// 				 N N /
		// 				 . N ) - 2x4 box
		return [5]Point{
			{x, y},
			{x, y + 1},
			{x, y + 2}, {x + 1, y + 2},
			{x + 1, y + 3},
		}, nil

	// --- Reflected States (Flipped 'N' shape, N') ---
	case 4:
		// Flipped 0-degrees ( N N . . /
		// 					   . N N N ) - 4x2 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x + 1, y + 1}, {x + 2, y + 1}, {x + 3, y + 1},
		}, nil
	case 5:
		// Flipped 90-degrees ( . N /
		// 						. N /
		// 						N N /
		// 						N . ) - 2x4 box
		return [5]Point{
			{x, y},
			{x, y + 1},
			{x - 1, y + 2}, {x, y + 2},
			{x - 1, y + 3},
		}, nil
	case 6:
		// Flipped 180-degrees ( N N N . /
		// 						 . . N N ) - 4x2 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y},
			{x + 2, y + 1}, {x + 3, y + 1},
		}, nil
	case 7:
		// Flipped 270-degrees ( . N /
		// 					     N N /
		// 						 N . /
		// 						 N . ) - 2x4 box
		return [5]Point{
			{x, y},
			{x - 1, y + 1}, {x, y + 1},
			{x - 1, y + 2},
			{x - 1, y + 3},
		}, nil

	default:
		return [5]Point{}, ErrInvalidState
	}
}

type piecesP struct {
}

func (piecesP) Name() byte {
	return 'P'
}

func (piecesP) States() int {
	return 8
}

func (piecesP) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	// --- Rotational States (Standard 'P' shape) ---
	case 0:
		// 0-degrees (P P /
		// 			  P P /
		// 			  P . ) - 2x3 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x, y + 1}, {x + 1, y + 1},
			{x, y + 2},
		}, nil
	case 1:
		// 90-degrees (P P . /
		// 			   P P P ) - 3x2 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1},
		}, nil
	case 2:
		// 180-degrees ( . P /
		// 				 P P /
		// 				 P P) - 2x3 box
		return [5]Point{
			{x, y},
			{x - 1, y + 1}, {x, y + 1},
			{x - 1, y + 2}, {x, y + 2},
		}, nil
	case 3:
		// 270-degrees ( P P P /
		// 				 . P P) - 3x2 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y},
			{x + 1, y + 1}, {x + 2, y + 1},
		}, nil

	// --- Reflected States (Flipped 'P' shape, P') ---
	case 4:
		// Flipped 0-degrees ( P P /
		// 					   P P /
		// 					   . P ) - 2x3 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x, y + 1}, {x + 1, y + 1},
			{x + 1, y + 2},
		}, nil
	case 5:
		// Flipped 90-degrees ( . P P /
		// 						P P P ) - 3x2 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
		}, nil
	case 6:
		// Flipped 180-degrees ( P . /
		// 						 P P /
		// 						 P P) - 2x3 box
		return [5]Point{
			{x, y},
			{x, y + 1}, {x + 1, y + 1},
			{x, y + 2}, {x + 1, y + 2},
		}, nil
	case 7:
		// Flipped 270-degrees ( P P P /
		// 						 P P .) - 3x2 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y},
			{x, y + 1}, {x + 1, y + 1},
		}, nil

	default:
		return [5]Point{}, ErrInvalidState
	}
}

type piecesT struct {
}

func (piecesT) Name() byte {
	return 'T'
}

func (piecesT) States() int {
	return 4
}

func (piecesT) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	// --- Rotational States (Standard 'T' shape) ---
	case 0:
		// 0-degrees (T T T /
		// 			  . T . /
		// 	          . T . ) - 3x3 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y},
			{x + 1, y + 1},
			{x + 1, y + 2},
		}, nil
	case 1:
		// 90-degrees ( . . T /
		// 				T T T /
		// 				. . T ) - 3x3 box
		return [5]Point{
			{x, y},
			{x - 2, y + 1}, {x - 1, y + 1}, {x, y + 1},
			{x, y + 2},
		}, nil
	case 2:
		// 180-degrees (. T . /
		// 			    . T . /
		// 				T T T ) - 3x3 box
		return [5]Point{
			{x, y},
			{x, y + 1},
			{x - 1, y + 2}, {x, y + 2}, {x + 1, y + 2},
		}, nil
	case 3:
		// 270-degrees ( T . ./
		// 				 T T T /
		// 				 T . . ) - 3x3 box
		return [5]Point{
			{x, y},
			{x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1},
			{x, y + 2},
		}, nil

	default:
		// You would need to define ErrInvalidState in your package
		return [5]Point{}, ErrInvalidState
	}
}

type piecesU struct {
}

func (piecesU) Name() byte {
	return 'U'
}

func (piecesU) States() int {
	return 4
}

func (piecesU) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	// --- Rotational States (Standard 'U' shape) ---
	case 0:
		// 0-degrees (U . U /
		// 			  U U U) - 3x2 box
		return [5]Point{
			{x, y}, {x + 2, y},
			{x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1},
		}, nil
	case 1:
		// 90-degrees (U U /
		// 			   U . /
		// 			   U U) - 2x3 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x, y + 1},
			{x, y + 2}, {x + 1, y + 2},
		}, nil
	case 2:
		// 180-degrees (U U U /
		// 				U . U) - 3x2 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y},
			{x, y + 1}, {x + 2, y + 1},
		}, nil
	case 3:
		// 270-degrees (U U /
		// 				. U /
		// 				U U) - 2x3 box
		return [5]Point{
			{x, y}, {x + 1, y},
			{x + 1, y + 1},
			{x, y + 2}, {x + 1, y + 2},
		}, nil

	default:
		// You would need to define ErrInvalidState in your package
		return [5]Point{}, ErrInvalidState
	}
}

type piecesV struct {
}

func (piecesV) Name() byte {
	return 'V'
}

func (piecesV) States() int {
	return 4
}

func (piecesV) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	// --- Rotational States (Standard 'V' shape) ---
	case 0:
		// 0-degrees (V . . /
		// 			  V . . /
		// 			  V V V ) - 3x3 box
		return [5]Point{
			{x, y},
			{x, y + 1},
			{x, y + 2}, {x + 1, y + 2}, {x + 2, y + 2},
		}, nil
	case 1:
		// 90-degrees (V V V /
		// 			   V . . /
		// 			   V . . ) - 3x3 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y},
			{x, y + 1},
			{x, y + 2},
		}, nil
	case 2:
		// 180-degrees (V V V /
		// 				. . V /
		// 				. . V) - 3x3 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y},
			{x + 2, y + 1},
			{x + 2, y + 2},
		}, nil
	case 3:
		// 270-degrees ( . . V /
		// 				 . . V /
		// 				 V V V) - 3x3 box
		return [5]Point{
			{x, y},
			{x, y + 1},
			{x - 2, y + 2}, {x - 1, y + 2}, {x, y + 2},
		}, nil

	default:
		// You would need to define ErrInvalidState in your package
		return [5]Point{}, ErrInvalidState
	}
}

type piecesY struct {
}

func (piecesY) Name() byte {
	return 'Y'
}

func (piecesY) States() int {
	return 8
}

func (piecesY) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	// --- Rotational States (Standard 'Y' shape) ---
	case 0:
		// 0-degrees ( . Y . . /
		// 			   Y Y Y Y) - 4x2 box
		return [5]Point{
			{x, y},
			{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1},
		}, nil
	case 1:
		// 90-degrees (Y . /
		// 			   Y Y /
		// 			   Y . /
		// 			   Y . ) - 2x4 box
		return [5]Point{
			{x, y},
			{x, y + 1}, {x + 1, y + 1},
			{x, y + 2},
			{x, y + 3},
		}, nil
	case 2:
		// 180-degrees (Y Y Y Y /
		// 				. . Y . ) - 4x2 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y},
			{x + 2, y + 1},
		}, nil
	case 3:
		// 270-degrees ( . Y /
		//				 . Y /
		// 				 Y Y /
		// 				 . Y) - 2x4 box
		return [5]Point{
			{x, y},
			{x, y + 1},
			{x - 1, y + 2}, {x, y + 2},
			{x, y + 3},
		}, nil

	// --- Reflected States (Flipped 'Y' shape, Y') ---
	case 4:
		// Flipped 0-degrees ( . . Y . /
		// 					   Y Y Y Y) - 4x2 box
		return [5]Point{
			{x, y},
			{x - 2, y + 1}, {x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
		}, nil
	case 5:
		// Flipped 90-degrees ( . Y /
		// 						Y Y /
		// 						. Y /
		// 						. Y) - 2x4 box
		return [5]Point{
			{x, y},
			{x - 1, y + 1}, {x, y + 1},
			{x, y + 2},
			{x, y + 3},
		}, nil
	case 6:
		// Flipped 180-degrees (Y Y Y Y /
		// 						. Y . . ) - 4x2 box
		return [5]Point{
			{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y},
			{x + 1, y + 1},
		}, nil
	case 7:
		// Flipped 270-degrees (Y . /
		// 						Y . /
		// 						Y Y /
		// 						Y . ) - 2x4 box
		return [5]Point{
			{x, y},
			{x, y + 1},
			{x, y + 2}, {x + 1, y + 1},
			{x, y + 3},
		}, nil

	default:
		// You would need to define ErrInvalidState in your package
		return [5]Point{}, ErrInvalidState
	}
}

type piecesZ struct {
}

func (piecesZ) Name() byte {
	return 'Z'
}

func (piecesZ) States() int {
	return 8
}

func (piecesZ) Position(ref Point, state int) ([5]Point, error) {
	x := ref.X
	y := ref.Y
	switch state {
	// --- Rotational States (Standard 'Z' shape) ---
	case 0:
		// 0-degrees (	Z Z . /
		// 				. Z . /
		// 				. Z Z ) - 3x3 box
		return [5]Point{
			{X: x, Y: y}, {X: x + 1, Y: y},
			{X: x + 1, Y: y + 1},
			{X: x + 1, Y: y + 2}, {X: x + 2, Y: y + 2},
		}, nil
	case 1:
		// 90-degrees ( . . Z /
		// 				Z Z Z /
		// 				Z . . ) - 3x3 box
		return [5]Point{
			{x, y},
			{x - 2, y + 1}, {x - 1, y + 1}, {x, y + 1},
			{x - 2, y + 2},
		}, nil
	case 2:
		// 180-degrees (. Z Z /
		// 				. Z . /
		// 				Z Z .) - 3x3 box
		return [5]Point{
			{X: x, Y: y}, {X: x + 1, Y: y},
			{X: x, Y: y + 1},
			{X: x - 1, Y: y + 2}, {X: x, Y: y + 2},
		}, nil
	case 3:
		// 270-degrees ( Z . . /
		// 				 Z Z Z /
		// 				 . . Z ) - 3x3 box
		return [5]Point{
			{X: x, Y: y},
			{X: x, Y: y + 1}, {X: x + 1, Y: y + 1}, {X: x + 2, Y: y + 1},
			{X: x + 2, Y: y + 2},
		}, nil

	default:
		// Assuming ErrInvalidState is defined elsewhere
		return [5]Point{}, ErrInvalidState
	}
}

func NewNamePiece(p NamedPiece) (Piece, error) {
	switch p {
	case PieceF:
		return piecesF{}, nil
	case PieceI:
		return piecesI{}, nil
	case PieceL:
		return piecesL{}, nil
	case PieceN:
		return piecesN{}, nil
	case PieceP:
		return piecesP{}, nil
	case PieceT:
		return piecesT{}, nil
	case PieceU:
		return piecesU{}, nil
	case PieceV:
		return piecesV{}, nil
	case PieceW:
		return piecesW{}, nil
	case PieceX:
		return piecesX{}, nil
	case PieceY:
		return piecesY{}, nil
	case PieceZ:
		return piecesZ{}, nil
	default:
		return nil, fmt.Errorf("%q is invalid name", p)
	}
}

func New12() []Piece {
	return []Piece{
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
}
