package engine

// Piece is a single game piece
type Piece uint8

// All pieces. To get white or black, we have to do Light | Pawn, Dark | Pawn, etc...
// the pieces should also roughly represent their value
const (
	Pawn   Piece = 1
	Bishop Piece = 2
	Knight Piece = 3
	Rook   Piece = 4
	Queen  Piece = 6
	King   Piece = 8

	Dark  Piece = 0x20
	Light Piece = 0x10
)

// GetPlain removes the color from the piece, only returning the first 3 bits
func (p Piece) GetPlain() Piece {
	return p & 0b111
}

// GetColor returns the color of the piece
func (p Piece) GetColor() Piece {
	return p >> 3
}
