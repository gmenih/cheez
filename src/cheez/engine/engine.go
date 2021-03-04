package engine

import (
	"time"
)

// Engine is the Chess game engine of Cheez
type Engine struct {
	// Board is the representation of current state of game
	Board Board

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

// NewEngine returns a new game with empty board
func NewEngine(duration time.Duration) *Engine {
	return &Engine{
		Board:       Board{},
		MoveCount:   0,
		MoveHistory: "",
		UpNext:      Light,
		Timers:      [2]time.Duration{duration, duration},
	}
}

// SetBoard sets the game board
func (e *Engine) SetBoard(board Board) {
	e.Board = board
}

// GetPiece returns a piece from the X, Y tile
func (e *Engine) GetPiece(x, y uint8) Piece {
	return e.Board[x][y]
}

// GetTile returns a piece from X, Y tile
// but the input is Tile
func (e *Engine) GetTile(t Tile) Piece {
	return e.GetPiece(uint8(t.X), uint8(t.Y))
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

	e.Board[to.X][to.Y] = e.Board[from.X][from.Y]
	e.Board[from.X][from.Y] = 0

	return true
}
