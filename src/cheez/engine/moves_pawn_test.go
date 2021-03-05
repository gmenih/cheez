package engine_test

import (
	"gmenih341/cheez/src/cheez/engine"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MovesSuite struct {
	suite.Suite
}

func TestPawnMoves(t *testing.T) {
	suite.Run(t, new(MovesSuite))
}

// Tests if Pawn can move forward twice if it's on it's
// second rank, and is not blocked by another piece
func (s *MovesSuite) TestPawnForwardTwice() {
	moves := s.validateMovesByFEN(
		"8/8/8/8/8/8/4P3/8 w - - 0 1",
		engine.TT("e2"),
		[]engine.Tile{
			engine.TT("e3"),
			engine.TT("e4"),
		},
	)
	assert.Equal(s.T(), 2, len(moves))
}

// Tests if pawn can only move forward once, if blocked
// by another piece but still on it's second
func (s *MovesSuite) TestPawnForwardTwiceBlocked() {
	// blocked by enemy queen
	moves := s.validateMovesByFEN(
		"8/8/8/8/4q3/8/4P3/8 w - - 0 1",
		engine.TT("e2"),
		[]engine.Tile{
			engine.TT("e3"),
		},
	)

	assert.Equal(s.T(), 1, len(moves))

	// blocked by own knight
	moves = s.validateMovesByFEN(
		"8/8/8/8/4N3/8/4P3/8 w - - 0 1",
		engine.TT("e2"),
		[]engine.Tile{
			engine.TT("e3"),
		},
	)

	assert.Equal(s.T(), 1, len(moves))
}

// Tests that pawn can't move forward if blocked by another piece
func (s *MovesSuite) TestPawnForwardBlocked() {
	// blocked by enemy knight
	moves := s.validateMovesByFEN(
		"8/8/8/8/8/4n3/4P3/8 w - - 0 1",
		engine.TT("e2"),
		[]engine.Tile{},
	)

	assert.Equal(s.T(), 0, len(moves))

	// blocked by own knight
	moves = s.validateMovesByFEN(
		"8/8/8/8/8/4N3/4P3/8 w - - 0 1",
		engine.TT("e2"),
		[]engine.Tile{},
	)

	assert.Equal(s.T(), 0, len(moves))
}

// Tests that pawn can only move forward once if it is no longer on the second rank
func (s *MovesSuite) TestPawnForwardAfterMoved() {
	moves := s.validateMovesByFEN(
		"8/8/8/8/8/4P3/8/8 w - - 0 1",
		engine.TT("e3"),
		[]engine.Tile{
			engine.TT("e4"),
		},
	)

	// can move 2 pieces forward
	assert.Equal(s.T(), 1, len(moves))
}

// Tests that pawn can take left
func (s *MovesSuite) TestPawnTakesLeft() {
	moves := s.validateMovesByFEN(
		"8/8/8/3n4/4P3/8/8/8 w - - 0 1",
		engine.TT("e4"),
		[]engine.Tile{
			engine.TT("e5"),
			engine.TT("d5"),
		},
	)

	// can move 2 pieces forward
	assert.Equal(s.T(), len(moves), 2)

	moves = s.validateMovesByFEN(
		"8/8/8/3nQ3/4P3/8/8/8 w - - 0 1",
		engine.TT("e4"),
		[]engine.Tile{
			engine.TT("d5"),
		},
	)

	// can move 2 pieces forward
	assert.Equal(s.T(), 1, len(moves))
}

// Tests that pawn can take right
func (s *MovesSuite) TestPawnTakesRight() {
	moves := s.validateMovesByFEN(
		"8/8/8/8/5n2/4P3/8/8 w - - 0 1",
		engine.TT("e3"),
		[]engine.Tile{
			engine.TT("e4"),
			engine.TT("f4"),
		},
	)

	// can move 2 pieces forward
	assert.Equal(s.T(), len(moves), 2)

	moves = s.validateMovesByFEN(
		"8/8/8/8/4Qn2/4P3/8/8 w - - 0 1",
		engine.TT("e3"),
		[]engine.Tile{
			engine.TT("f4"),
		},
	)

	// can move 2 pieces forward
	assert.Equal(s.T(), 1, len(moves))
}

// Tests that pawn can't move it is pinend
func (s *MovesSuite) TestPawnMovePinned() {
	moves := s.validateMovesByFEN(
		"8/8/8/8/8/K3P2q/8/8 w - - 0 1",
		engine.T(4, 2),
		[]engine.Tile{},
	)

	// can move 2 pieces forward
	assert.Equal(s.T(), 0, len(moves))
}

// Tests that pawn can't take if it is pinned
func (s *MovesSuite) TestPawnTakesPinned() {
	moves := s.validateMovesByFEN(
		"8/8/8/8/8/K3P2q/8/8 w - - 0 1",
		engine.TT("e3"),
		[]engine.Tile{},
	)

	// can move 2 pieces forward
	assert.Equal(s.T(), 0, len(moves))

	// can take if it clears the pin
	moves = s.validateMovesByFEN(
		"8/8/8/8/5q2/4P3/8/2K5 w - - 0 1",
		engine.TT("e3"),
		[]engine.Tile{
			engine.TT("f4"),
		},
	)

	// can move 2 pieces forward
	assert.Equal(s.T(), 1, len(moves))
}

// Tests that pawn can do en passant move
func (s *MovesSuite) TestPawnEnPassant() {
	// TODO: figure out
}

func (s *MovesSuite) validateMovesByFEN(fenString engine.FENString, tile engine.Tile, expectedMoves []engine.Tile) []engine.Tile {
	e := fenString.ParseToEngine()

	validMoves := e.GetValidMoves(tile)

	// can move 2 pieces forward
	for i, m := range expectedMoves {
		assert.True(s.T(), assert.ObjectsAreEqual(m, validMoves[i]))
	}

	return validMoves
}
