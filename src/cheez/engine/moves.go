package engine

import (
	"math"
)

type moverFunc func(t Tile, i int8) []Tile

func isInBounds(t Tile) bool {
	return t.X >= 0 && t.Y >= 0 && t.X <= 7 && t.Y <= 7
}

func genericMover(tile Tile, mul int8, num int8, angleDiff float64) []Tile {
	moves := []Tile{}
	delta := math.Pi/(float64(num)/2) + angleDiff

	for i := float64(0); i < float64(num); i++ {
		x := int8(math.Round(math.Cos(delta*i))) * mul
		y := int8(math.Round(math.Sin(delta*i))) * mul

		moves = append(moves, tile.Add(x, y))
	}

	return moves
}

func linearMover(tile Tile, mul int8) []Tile {
	return genericMover(tile, mul, 4, 0)
}

func diagonalMover(tile Tile, mul int8) []Tile {
	return genericMover(tile, mul, 4, math.Pi/4)
}

func omniDirectionalMover(tile Tile, mul int8) []Tile {
	return genericMover(tile, mul, 8, 0)
}

func knightMover(tile Tile, _ int8) []Tile {
	return []Tile{
		tile.Add(-1, -2),
		tile.Add(-1, 2),
		tile.Add(-2, -1),
		tile.Add(-2, 1),
		tile.Add(1, -2),
		tile.Add(1, 2),
		tile.Add(2, -1),
		tile.Add(2, 1),
	}
}

func intersectMoves(movesA []Tile, movesB []Tile) []Tile {
	intersection := []Tile{}

	for _, a := range movesA {
		for _, b := range movesB {
			if a.Equals(b) {
				intersection = append(intersection, b)
			}
		}
	}

	return intersection
}

func (e Engine) filterPinnedMoves(tile Tile, moves []Tile) []Tile {
	sourcePiece := e.GetTile(tile)
	kingDirection := int8(-1)
	blockedDirections := map[int]bool{}
	directionEnemies := map[int]Piece{}
	moveMap := map[int][]Tile{}

	for i := int8(1); i < 7; i++ {
		directions := omniDirectionalMover(tile, i)

		for dIndex, direction := range directions {
			if blockedDirections[dIndex] == true {
				continue
			}

			if isInBounds(direction) {
				targetPiece := e.GetTile(direction)
				moveMap[dIndex] = append(moveMap[dIndex], direction)

				if targetPiece != 0 {
					blockedDirections[dIndex] = true

					if targetPiece.GetColor() != sourcePiece.GetColor() {
						directionEnemies[dIndex] = targetPiece.GetPlain()
						continue
					} else if targetPiece.GetPlain() == King {
						kingDirection = int8(dIndex)
					}
				}
			}
		}
	}

	if kingDirection != -1 {
		oppositeDirection := ((kingDirection + 8) - 4) % 8
		if v, ok := directionEnemies[int(oppositeDirection)]; ok && v == Queen || v == Bishop || v == Rook {
			return intersectMoves(moves, append(moveMap[int(oppositeDirection)], moveMap[int(kingDirection)]...))
		}
	}

	return moves
}

func (e Engine) getMoves(tile Tile, maxDistance int8, moverFn moverFunc) []Tile {
	moves := []Tile{}
	blockedDirections := map[int]bool{}
	sourceColor := e.GetTile(tile).GetColor()

	for i := int8(1); i <= maxDistance; i++ {
		directions := moverFn(tile, i)

		for dIndex, direction := range directions {
			if blockedDirections[dIndex] == true {
				continue
			}

			if isInBounds(direction) {
				targetColor := e.GetTile(direction).GetColor()
				if targetColor != 0 {
					blockedDirections[dIndex] = true

					if targetColor == sourceColor {
						continue
					}
				}

				moves = append(moves, direction)
			} else {
				blockedDirections[dIndex] = true
			}
		}
	}

	return moves
}

func (e Engine) getPawnMoves(tile Tile) []Tile {
	// TODO: handle en passant
	piece := e.GetTile(tile)
	moves := []Tile{}

	direction := int8(1)
	if piece.IsDark() {
		direction = -1
	}

	// check if forward possible
	forward := tile.Add(0, direction)

	if e.GetTile(forward) == 0 {
		moves = append(moves, forward)

		forward2 := tile.Add(0, direction*2)
		if ((piece.IsLight() && tile.Y == 1) || (piece.IsDark() && tile.Y == 6)) && e.GetTile(forward2) == 0 {
			moves = append(moves, forward2)
		}
	}

	if tile.X >= 1 {
		forwardLeft := tile.Add(-1, direction)

		if e.GetTile(forwardLeft) != 0 && e.GetTile(forwardLeft).GetColor() != piece.GetColor() {
			moves = append(moves, forwardLeft)
		}
	}

	if tile.X <= 6 {
		forwardRight := tile.Add(1, direction)

		if e.GetTile(forwardRight) != 0 && e.GetTile(forwardRight).GetColor() != piece.GetColor() {
			moves = append(moves, forwardRight)
		}
	}

	return moves
}

func (e Engine) isValidMove(from, to Tile) bool {
	moves := e.GetValidMoves(from)

	for _, tile := range moves {
		if tile.Equals(to) {
			return true
		}
	}

	return false
}

func (e Engine) getMovesByPiece(tile Tile) []Tile {
	piece := e.GetTile(tile)

	if piece.GetColor() == e.UpNext {
		switch piece.GetPlain() {
		case Pawn:
			return e.getPawnMoves(tile)
		case Knight:
			return e.getMoves(tile, 1, knightMover)
		case King:
			// TODO: handle castling
			return e.getMoves(tile, 1, omniDirectionalMover)
		case Rook:
			return e.getMoves(tile, 7, linearMover)
		case Bishop:
			return e.getMoves(tile, 7, diagonalMover)
		case Queen:
			return e.getMoves(tile, 7, omniDirectionalMover)
		}
	}

	return []Tile{}
}

// GetValidMoves returns all valid moves that can be made on a specific tile,
// based on what Piece is on that tile
func (e *Engine) GetValidMoves(tile Tile) []Tile {
	moves := e.getMovesByPiece(tile)
	moves = e.filterPinnedMoves(tile, moves)
	// TODO:
	// * handle check (only unpin, king moves allowed)

	return moves
}
