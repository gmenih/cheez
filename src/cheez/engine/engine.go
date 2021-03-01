package engine

import "time"

// Tile represents X and Y on the board
type Tile struct {
	X uint8
	Y uint8
}

func T(x, y uint8) Tile {
	return Tile{x, y}
}

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
	return false
}

func (e *Engine) GetPiece(x, y uint8) Piece {
	return e.board[x][y]
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

	return false
}
