package engine_test

import (
	"gmenih341/cheez/src/cheez/engine"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PieceSuite struct {
	suite.Suite
}

func TestPieceSuite(t *testing.T) {
	suite.Run(t, new(PieceSuite))
}

func (suite *PieceSuite) TestGetPlain() {
	p := engine.Piece(engine.Pawn | engine.Dark)
	assert.Equal(suite.T(), engine.Pawn, p.GetPlain())

	p = engine.Piece(engine.Bishop | engine.Dark)
	assert.Equal(suite.T(), engine.Bishop, p.GetPlain())

	p = engine.Piece(engine.Knight | engine.Dark)
	assert.Equal(suite.T(), engine.Knight, p.GetPlain())

	p = engine.Piece(engine.Rook | engine.Dark)
	assert.Equal(suite.T(), engine.Rook, p.GetPlain())

	p = engine.Piece(engine.Queen | engine.Dark)
	assert.Equal(suite.T(), engine.Queen, p.GetPlain())

	p = engine.Piece(engine.King | engine.Dark)
	assert.Equal(suite.T(), engine.King, p.GetPlain())

	p = engine.Piece(engine.Pawn | engine.Dark)
	assert.Equal(suite.T(), engine.Pawn, p.GetPlain())

	p = engine.Piece(engine.Pawn | engine.Light)
	assert.Equal(suite.T(), engine.Pawn, p.GetPlain())

	p = engine.Piece(engine.Bishop | engine.Light)
	assert.Equal(suite.T(), engine.Bishop, p.GetPlain())

	p = engine.Piece(engine.Knight | engine.Light)
	assert.Equal(suite.T(), engine.Knight, p.GetPlain())

	p = engine.Piece(engine.Rook | engine.Light)
	assert.Equal(suite.T(), engine.Rook, p.GetPlain())

	p = engine.Piece(engine.Queen | engine.Light)
	assert.Equal(suite.T(), engine.Queen, p.GetPlain())

	p = engine.Piece(engine.King | engine.Light)
	assert.Equal(suite.T(), engine.King, p.GetPlain())

	p = engine.Piece(engine.Pawn | engine.Light)
	assert.Equal(suite.T(), engine.Pawn, p.GetPlain())
}

func (suite *PieceSuite) TestGetColor() {
	p := engine.Piece(engine.Pawn | engine.Dark)
	assert.Equal(suite.T(), engine.Dark, p.GetColor())

	p = engine.Piece(engine.Bishop | engine.Dark)
	assert.Equal(suite.T(), engine.Dark, p.GetColor())

	p = engine.Piece(engine.Knight | engine.Dark)
	assert.Equal(suite.T(), engine.Dark, p.GetColor())

	p = engine.Piece(engine.Rook | engine.Dark)
	assert.Equal(suite.T(), engine.Dark, p.GetColor())

	p = engine.Piece(engine.Queen | engine.Dark)
	assert.Equal(suite.T(), engine.Dark, p.GetColor())

	p = engine.Piece(engine.King | engine.Dark)
	assert.Equal(suite.T(), engine.Dark, p.GetColor())

	p = engine.Piece(engine.Pawn | engine.Dark)
	assert.Equal(suite.T(), engine.Dark, p.GetColor())

	p = engine.Piece(engine.Pawn | engine.Light)
	assert.Equal(suite.T(), engine.Light, p.GetColor())

	p = engine.Piece(engine.Bishop | engine.Light)
	assert.Equal(suite.T(), engine.Light, p.GetColor())

	p = engine.Piece(engine.Knight | engine.Light)
	assert.Equal(suite.T(), engine.Light, p.GetColor())

	p = engine.Piece(engine.Rook | engine.Light)
	assert.Equal(suite.T(), engine.Light, p.GetColor())

	p = engine.Piece(engine.Queen | engine.Light)
	assert.Equal(suite.T(), engine.Light, p.GetColor())

	p = engine.Piece(engine.King | engine.Light)
	assert.Equal(suite.T(), engine.Light, p.GetColor())

	p = engine.Piece(engine.Pawn | engine.Light)
	assert.Equal(suite.T(), engine.Light, p.GetColor())
}

func (suite *PieceSuite) TestSameColor() {
	p1 := engine.Piece(engine.King | engine.Dark)
	p2 := engine.Piece(engine.Queen | engine.Dark)

	assert.Equal(suite.T(), true, p1.SameColor(p2))

	p2 = engine.Piece(engine.Queen | engine.Light)
	assert.Equal(suite.T(), false, p1.SameColor(p2))
}
