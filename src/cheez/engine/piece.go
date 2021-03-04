package engine

import "unicode"

// Piece is a single game piece
type Piece uint8

// All pieces. To get white or black, we have to do Light | Pawn, Dark | Pawn, etc...
// the pieces should also roughly represent their value
const (
	Pawn   Piece = 1 // 1
	Bishop Piece = 2 // 01
	Knight Piece = 3 // 11
	Rook   Piece = 4 // 001
	Queen  Piece = 6 // 011
	King   Piece = 8 // 111

	Dark  Piece = 0x20
	Light Piece = 0x10
)

// GetPlain removes the color from the piece, only returning the first 3 bits
func (p Piece) GetPlain() Piece {
	return p & 0xf
}

// GetColor returns the color of the piece
func (p Piece) GetColor() Piece {
	return p & 0xf0
}

func (p Piece) SameColor(p2 Piece) bool {
	return p.GetColor() == p2.GetColor()
}

func (p Piece) IsDark() bool {
	return p.GetColor() == Dark
}

func (p Piece) IsLight() bool {
	return p.GetColor() == Light
}

func (p Piece) ToRune() rune {
	v := '.'

	switch p.GetPlain() {
	case King:
		v = 'k'
		break
	case Queen:
		v = 'q'
		break
	case Pawn:
		v = 'p'
		break
	case Bishop:
		v = 'b'
		break
	case Rook:
		v = 'r'
		break
	case Knight:
		v = 'n'
		break
	}

	if p.GetColor() == Light {
		v = unicode.ToUpper(v)
	}

	return v
}
