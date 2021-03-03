package engine

func isInBounds(t Tile) bool {
	return t.X >= 0 && t.Y >= 0 && t.X <= 7 && t.Y <= 7
}

type predicateFn func(t Tile, i int8) []Tile

func linearPredicate(t Tile, i int8) []Tile {
	return []Tile{
		t.Add(0, i),
		t.Add(0, -i),
		t.Add(-i, 0),
		t.Add(i, 0),
	}
}

func diagonalPredicate(t Tile, i int8) []Tile {
	return []Tile{
		t.Add(i, i),
		t.Add(i, -i),
		t.Add(-i, i),
		t.Add(-i, -i),
	}
}

func knightPredicate(t Tile, _ int8) []Tile {
	return []Tile{
		t.Add(-1, -2),
		t.Add(-1, 2),
		t.Add(-2, -1),
		t.Add(-2, 1),
		t.Add(1, -2),
		t.Add(1, 2),
		t.Add(2, -1),
		t.Add(2, 1),
	}
}

func joinPredicate(predicates ...predicateFn) predicateFn {
	return func(t Tile, i int8) []Tile {
		moves := []Tile{}
		for _, p := range predicates {
			moves = append(moves, p(t, i)...)
		}
		return moves
	}
}

func (e Engine) getMoves(t Tile, maxDistance int8, predicate predicateFn) []Tile {
	moves := []Tile{}
	blockedDirections := map[int]bool{}
	sourceColor := e.GetTile(t).GetColor()

	for i := int8(1); i < maxDistance; i++ {
		directions := predicate(t, i)

		for i, direction := range directions {
			if blockedDirections[i] == true {
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

				moves = append(moves, direction)
			}
		}
	}

	return moves
}

func (e Engine) getMovesOnTile(tile Tile) []Tile {
	figure := e.GetTile(tile)

	switch figure.GetPlain() {
	case Knight:
		return e.getMoves(tile, 2, knightPredicate)
	case Rook:
		return e.getMoves(tile, 8, linearPredicate)
	case Bishop:
		return e.getMoves(tile, 8, diagonalPredicate)
	case King:
		return e.getMoves(tile, 2, joinPredicate(linearPredicate, diagonalPredicate))
	case Queen:
		return e.getMoves(tile, 8, joinPredicate(linearPredicate, diagonalPredicate))
	}

	return []Tile{}
}
