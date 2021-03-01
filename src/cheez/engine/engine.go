package engine

import (
	"fmt"
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
	if from.Equals(to) {
		return false
	}

	pieceFrom := e.GetTile(from)
	pieceTo := e.GetTile(to)

	fmt.Printf("%v", pieceFrom.GetColor() == Light)

	if pieceFrom.GetColor() != e.UpNext {
		return false
	}

	if pieceTo != 0 && pieceFrom.SameColor(pieceTo) {
		return false
	}

	return true
}

// GetPiece returns a piece from the X, Y tile
func (e *Engine) GetPiece(x, y uint8) Piece {
	return e.board[x][y]
}

// GetTile returns a piece from X, Y tile
// but the input is Tile
func (e *Engine) GetTile(tile Tile) Piece {
	return e.GetPiece(tile.X, tile.Y)
}

// GetValidMoves returns all valid moves that can be made on a specific tile,
// based on what Piece is on that tile
func (e *Engine) GetValidMoves(v Tile) []Tile {
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
