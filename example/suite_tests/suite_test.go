package suite_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	SuiteVariable int
}

func (suite *ExampleTestSuite) SetupTest() {
	suite.SuiteVariable = 1
}

func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 1, suite.SuiteVariable)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
