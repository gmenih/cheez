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

	if t.X >= 1 {
		forwardLeft := t.Add(-1, direction)

		if e.GetTile(forwardLeft) != 0 && e.GetTile(forwardLeft).GetColor() != piece.GetColor() {
			moves = append(moves, forwardLeft)
		}
	}

	if t.X <= 6 {
		forwardRight := t.Add(1, direction)

		if e.GetTile(forwardRight) != 0 && e.GetTile(forwardRight).GetColor() != piece.GetColor() {
			moves = append(moves, forwardRight)
		}
	}

	return moves
}

func (e *Engine) isValidMove(from, to Tile) bool {
	moves := e.GetValidMoves(from)

	for _, t := range moves {
		if t.Equals(to) {
			return true
		}
	}

	return false
}

// GetValidMoves returns all valid moves that can be made on a specific tile,
// based on what Piece is on that tile
func (e *Engine) GetValidMoves(tile Tile) []Tile {
	figure := e.GetTile(tile)

	// TODO:
	// * handle en-passant
	// * handle pinning
	// * handle check (only unpin, king moves allowed)
	// * handle castling

	if figure.GetColor() == e.UpNext {
		switch figure.GetPlain() {
		case Pawn:
			return e.getPawnMoves(tile)
		case Knight:
			return e.getMoves(tile, 1, knightMover)
		case King:
			return e.getMoves(tile, 1, joinPredicate(linearMover, diagonalMover))
		case Rook:
			return e.getMoves(tile, 7, linearMover)
		case Bishop:
			return e.getMoves(tile, 7, diagonalMover)
		case Queen:
			return e.getMoves(tile, 7, joinPredicate(linearMover, diagonalMover))
		}
	}

	return []Tile{}
}
