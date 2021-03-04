package engine

type moverFunc func(t Tile, i int8) []Tile

func isInBounds(t Tile) bool {
	return t.X >= 0 && t.Y >= 0 && t.X <= 7 && t.Y <= 7
}

func linearMover(t Tile, i int8) []Tile {
	return []Tile{
		t.Add(0, i),
		t.Add(0, -i),
		t.Add(-i, 0),
		t.Add(i, 0),
	}
}

func diagonalMover(t Tile, i int8) []Tile {
	return []Tile{
		t.Add(i, i),
		t.Add(i, -i),
		t.Add(-i, i),
		t.Add(-i, -i),
	}
}

func knightMover(t Tile, _ int8) []Tile {
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

func joinPredicate(predicates ...moverFunc) moverFunc {
	return func(t Tile, i int8) []Tile {
		moves := []Tile{}
		for _, p := range predicates {
			moves = append(moves, p(t, i)...)
		}
		return moves
	}
}

func (e Engine) getMoves(t Tile, maxDistance int8, moverFn moverFunc) []Tile {
	moves := []Tile{}
	blockedDirections := map[int]bool{}
	sourceColor := e.GetTile(t).GetColor()

	for i := int8(1); i <= maxDistance; i++ {
		directions := moverFn(t, i)

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

				// TODO: check if pinned
				// if collides with own king in one direction
				// and with a diagonal/linear piece in another
				moves = append(moves, direction)
			}
		}
	}

	return moves
}

func (e Engine) getPawnMoves(t Tile) []Tile {
	piece := e.GetTile(t)
	moves := []Tile{}

	direction := int8(1)
	if piece.IsDark() {
		direction = -1
	}

	// check if forward possible
	forward := t.Add(0, direction)

	if e.GetTile(forward) == 0 {
		moves = append(moves, forward)

		forward2 := t.Add(0, direction*2)
		if (piece.IsLight() && t.Y == 1) || (piece.IsDark() && t.Y == 6) && e.GetTile(forward2) == 0 {
			moves = append(moves, forward2)
		}
	}

	if t.X >= 1 && t.X <= 6 {
		forwardLeft := t.Add(-1, direction)
		forwardRight := t.Add(1, direction)

		if e.GetTile(forwardLeft) != 0 && e.GetTile(forwardLeft).GetColor() != piece.GetColor() {
			moves = append(moves, forwardLeft)
		}

		if e.GetTile(forwardRight) != 0 && e.GetTile(forwardRight).GetColor() != piece.GetColor() {
			moves = append(moves, forwardRight)
		}
	}

	// if t.X < 6 && e.GetTile(fr) != 0 && e.GetTile(fr).GetColor() != piece.GetColor() {
	// 	moves = append(moves, fr)
	// }
	// if (piece.IsDark() && t.Y == 6) || (piece.IsLight() && t.Y == 2) {
	// 	moves = append(moves, t.Add(0, direction), t.Add(0, direction*2))
	// }
	return moves
}
