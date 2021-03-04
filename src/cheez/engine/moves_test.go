package engine_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MovesSuite struct {
	suite.Suite
}

func TestMoves(t *testing.T) {
	suite.Run(t, new(MovesSuite))
}

func (s *MovesSuite) TestPawnForward() {

}
