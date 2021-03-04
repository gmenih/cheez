package engine_test

import (
	"gmenih341/cheez/src/cheez/engine"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TileSuite struct {
	suite.Suite
}

func TestTiles(t *testing.T) {
	suite.Run(t, new(TileSuite))
}

func (s *TileSuite) TestT() {
	t := engine.T(1, 1)

	assert.Equal(s.T(), t.X, uint8(1))
	assert.Equal(s.T(), t.Y, uint8(1))
}

func (s *TileSuite) TestAdd() {
	t := engine.T(5, 5)

	t1 := t.Add(1, 1)
	t2 := t.Add(8, 0)
	t3 := t.Add(-4, -4)

	assert.Equal(s.T(), t1.X, uint8(6))
	assert.Equal(s.T(), t1.Y, uint8(6))

	assert.Equal(s.T(), t2.X, uint8(13))
	assert.Equal(s.T(), t2.Y, uint8(5))

	assert.Equal(s.T(), t3.X, uint8(1))
	assert.Equal(s.T(), t3.Y, uint8(1))
}

func (s *TileSuite) TestEquals() {
	t := engine.T(5, 5)

	assert.Equal(s.T(), t.Equals(engine.T(5, 5)), true)
	assert.Equal(s.T(), t.Equals(engine.T(6, 5)), false)

}
