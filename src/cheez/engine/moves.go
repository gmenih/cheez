package engine

func isInBounds(t Tile) bool {
	return t.X >= 0 && t.Y >= 0 && t.X <= 7 && t.Y <= 7
}

func filterInBounds(tiles ...Tile) (res []Tile) {
	for _, t := range tiles {
		if isInBounds(t) {
			res = append(res, t)
		}
	}

	return
}

func (e Engine) fillMoves(t Tile, mvFunc func(i int8) []Tile) []Tile {
	moves := make([]Tile, 14)
	blockedDirections := [4]bool{}
	sourceColor := e.GetTile(t).GetColor()
	m := 0

	for i := int8(1); i < 8; i++ {
		directions := mvFunc(i)

		for i, direction := range directions {
			if blockedDirections[i] {
				continue
			}

			if isInBounds(direction) {
				targetColor := e.GetTile(direction).GetColor()
				if targetColor != 0 {
					blockedDirections[i] = true

					if targetColor == sourceColor {
						continue
					}
				}

				moves[m] = direction
				m++
			}
		}
	}

	return moves[0:m]
}

func (e *Engine) movesInLine(t Tile) []Tile {
	return e.fillMoves(t, func(i int8) []Tile {
		return []Tile{
			t.Add(0, i),
			t.Add(0, -i),
			t.Add(-i, 0),
			t.Add(i, 0),
		}
	})
}

func (e *Engine) movesDiagonally(t Tile) []Tile {
	return e.fillMoves(t, func(i int8) []Tile {
		return []Tile{
			t.Add(i, i),
			t.Add(i, -i),
			t.Add(-i, i),
			t.Add(-i, -i),
		}
	})
}

func (e *Engine) movesInL(t Tile) []Tile {
	return filterInBounds(
		t.Add(-1, -2),
		t.Add(-1, 2),
		t.Add(-2, -1),
		t.Add(-2, 1),
		t.Add(1, -2),
		t.Add(1, 2),
		t.Add(2, -1),
		t.Add(2, 1),
	)
}
