package engine

import (
	"time"
)

// Engine is the Chess game engine of Cheez
type Engine struct {
	// board is the representation of current state of game
	board Board

	// MoveCount tells us how many moves there are
	MoveCount uint
	// MoveHistory is algebraic representation of move history
	MoveHistory string
	// UpNext tells us who's up next (black or white)
	UpNext Piece
	// Timers are saying how much time each player still has
	// [0] = Light
	// [1] = Dark
	Timers [2]time.Duration
}

// NewEngine returns a new instance of Engine
func NewEngine(duration time.Duration) *Engine {
	return &Engine{
		board:       NewGameFENString.Parse(),
		MoveCount:   0,
		MoveHistory: "",
		UpNext:      Light,
		Timers:      [2]time.Duration{duration, duration},
	}
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

// GetPiece returns a piece from the X, Y tile
func (e *Engine) GetPiece(x, y uint8) Piece {
	return e.board[x][y]
}

// GetTile returns a piece from X, Y tile
// but the input is Tile
func (e *Engine) GetTile(t Tile) Piece {
	return e.GetPiece(uint8(t.X), uint8(t.Y))
}

// GetValidMoves returns all valid moves that can be made on a specific tile,
// based on what Piece is on that tile
func (e *Engine) GetValidMoves(tile Tile) []Tile {
	figure := e.GetTile(tile)

	// TODO:
	// * handle en-passant
	// * handle pawn takes
	// * handle pinning
	// * handle check (only unpin, king moves allowed)

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

// MovePiece performs a move of a Piece on Board
// It confirms that the move is valid, changes which side is up next,
// and increases any counters and times that it needs to
func (e *Engine) MovePiece(from, to Tile) bool {
	if !e.isValidMove(from, to) {
		return false
	}

	// we increase the counter after every black move
	if e.UpNext == Dark {
		e.MoveCount++
		e.UpNext = Light
	} else {
		e.UpNext = Dark
	}

	e.board[to.X][to.Y] = e.board[from.X][from.Y]
	e.board[from.X][from.Y] = 0

	return true
}
