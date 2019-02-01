package tmdb

func (suite *TMBDTestSuite) TestGetTrending() {
	trending, err := suite.GetTrending("movie", "day")
	suite.Nil(err)
	suite.NotNil(trending)
}

func (suite *TMBDTestSuite) TestGetTrendingFail() {
	suite.APIKey = ""
	_, err := suite.GetTrending("movie", "day")
	suite.NotNil(err)
}
