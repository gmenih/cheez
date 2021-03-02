package engine

// Tile represents X and Y on the board
type Tile struct {
	X uint8
	Y uint8
}

// T is a quick-help function to make a Tile
func T(x, y uint8) Tile {
	return Tile{x, y}
}

// Equals compares if two tiles are equal
func (t Tile) Equals(t2 Tile) bool {
	return t.X == t2.X && t.Y == t2.Y
}

func (t Tile) Add(dX, dY int8) Tile {
	return Tile{
		uint8(int8(t.X) + dX),
		uint8(int8(t.Y) + dY),
	}
}
