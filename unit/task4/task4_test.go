package task4

import (
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestSuites(t *testing.T) {
	suite.Run(t, new(OrderServiceSuite))
}

type OrderServiceSuite struct {
	suite.Suite
}

func (s *OrderServiceSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
}

func (s *OrderServiceSuite) TearDownTest() {
	s.T().Log("TearDownTest")
}
