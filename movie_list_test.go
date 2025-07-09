package tmdb

func (suite *TMBDTestSuite) TestGetMovieNowPlaying() {
	movies, err := suite.client.GetMovieNowPlaying(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieNowPlayingFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetMovieNowPlaying(nil)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieNowPlayingWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.client.GetMovieNowPlaying(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMoviePopular() {
	movies, err := suite.client.GetMoviePopular(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMoviePopularFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetMoviePopular(nil)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMoviePopularWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.client.GetMoviePopular(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieTopRated() {
	movies, err := suite.client.GetMovieTopRated(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieTopRatedFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetMovieTopRated(nil)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieTopRatedWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.client.GetMovieTopRated(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieUpcoming() {
	movies, err := suite.client.GetMovieUpcoming(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieUpcomingFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetMovieUpcoming(nil)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieUpcomingWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.client.GetMovieUpcoming(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}
