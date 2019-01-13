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
	suite.Contains(err.Error(), "no such host")
	err = suite.get("https://api.themoviedb.org/3/movie/7578000?language=en-US", nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
	err = suite.get("", nil)
	suite.Equal("url field is empty", err.Error())
}
