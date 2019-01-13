package tmdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TMBDTestSuite struct {
	suite.Suite
	Client
}

func (suite *TMBDTestSuite) SetupTest() {
	suite.APIKey = os.Getenv("APIKey")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TMBDTestSuite))
}

func (suite *TMBDTestSuite) TestGetFail() {
	err := suite.get("http://www.testfakewebsite.org", nil)

	suite.NotNil(err)

	err = suite.get("https://api.themoviedb.org/3/movieeee/75780?language=en-US", nil)

	suite.NotNil(err)
}
