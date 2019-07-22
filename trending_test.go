package tmdb

func (suite *TMBDTestSuite) TestGetTrending() {
	trending, err := suite.client.GetTrending("movie", "day")
	suite.Nil(err)
	suite.NotNil(trending)
}

func (suite *TMBDTestSuite) TestGetTrendingFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTrending("movie", "day")
	suite.NotNil(err)
}
