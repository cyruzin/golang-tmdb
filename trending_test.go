package tmdb

func (suite *TMBDTestSuite) TestGetTrending() {
	trending, err := suite.client.GetTrending("movie", "day", nil)
	suite.Nil(err)
	suite.NotNil(trending)
}

func (suite *TMBDTestSuite) TestGetTrendingWithOptions() {
	options := map[string]string{
		"language": "en-US",
		"region":   "US",
	}
	trending, err := suite.client.GetTrending("tv", "week", options)
	suite.Nil(err)
	suite.NotNil(trending)
}

func (suite *TMBDTestSuite) TestGetTrendingFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTrending("movie", "day", nil)
	suite.NotNil(err)
}
