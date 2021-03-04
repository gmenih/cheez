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

func makeForwardPredicate(p Piece) predicateFn {
	color := p.GetColor()
	dir := int8(1)
	if color == Dark {
		dir = -1
	}

	return func(t Tile, _ int8) []Tile {
		moves := []Tile{
			t.Add(0, dir),
		}

		if dir == 1 && t.Y == 1 || dir == -1 && t.Y == 6 {
			moves = append(moves, t.Add(0, dir*2))
		}

		return moves
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

	for i := int8(1); i <= maxDistance; i++ {
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
		return e.getMoves(tile, 1, knightPredicate)
	case King:
		return e.getMoves(tile, 1, joinPredicate(linearPredicate, diagonalPredicate))
	case Rook:
		return e.getMoves(tile, 7, linearPredicate)
	case Bishop:
		return e.getMoves(tile, 7, diagonalPredicate)
	case Queen:
		return e.getMoves(tile, 7, joinPredicate(linearPredicate, diagonalPredicate))
	case Pawn:
		return e.getMoves(tile, 1, makeForwardPredicate(figure.GetColor()))
	}

	return []Tile{}
}
