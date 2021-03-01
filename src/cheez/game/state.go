package game

type tileInt uint8

func getTileInt(x, y uint8) tileInt {
	_x := x & 0xf
	_y := (y & 0xf) << 4

	return tileInt(_x | _y)
}

func reverseTileInt(ti tileInt) (uint8, uint8) {
	x := ti & 0xf
	y := ti >> 4

	return uint8(x), uint8(y)
}

type gameState struct {
	draggingTile tileInt
	hoveredTile  tileInt
}

func newState() *gameState {
	return &gameState{
		draggingTile: 0xff,
		hoveredTile:  0xff,
	}
}

func (s *gameState) setDragging(x, y uint8) {
	s.draggingTile = getTileInt(x, y)
}

func (s *gameState) resetDragging() (uint8, uint8) {
	x, y := reverseTileInt(s.draggingTile)
	s.draggingTile = 255

	return x, y
}

func (s *gameState) isDragging(x, y uint8) bool {
	return s.draggingTile == getTileInt(x, y)
}

func (s *gameState) highlightTile(x, y uint8) {
	s.hoveredTile = getTileInt(x, y)
}
