package engine

// Board is the game board type
type Board [8][8]Piece

func (b Board) ToString() string {
	s := ""

	for x := uint8(0); x < 8; x++ {
		for y := uint8(0); y < 8; y++ {
			s += string(b[x][y].ToRune())
		}

		s += "\n"
	}

	return s
}
