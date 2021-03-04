package engine_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FENSuite struct {
	suite.Suite
}

func TestFENSuite(t *testing.T) {
	suite.Run(t, new(FENSuite))
}

func (s *FENSuite) TestInitialString() {

}
