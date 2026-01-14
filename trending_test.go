package tmdb

func (suite *TMBDTestSuite) TestGetTrending() {
	trending, err := suite.client.GetTrending("movie", "day", nil)
	suite.NoError(err)
	suite.NotNil(trending)
}

func (suite *TMBDTestSuite) TestGetTrendingWithOptions() {
	options := map[string]string{
		"language": "en-US",
		"region":   "US",
	}
	trending, err := suite.client.GetTrending("tv", "week", options)
	suite.NoError(err)
	suite.NotNil(trending)
}

func (suite *TMBDTestSuite) TestGetTrendingFail() {
	_, err := suite.client.GetTrending("movie", "bad", nil)
	suite.Error(err)
}
