package tmdb

func (suite *TMBDTestSuite) TestGetGenreMovieList() {
	genres, err := suite.client.GetGenreMovieList(nil)
	suite.NoError(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenreMovieListFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetGenreMovieList(nil)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetGenreMovieListWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	genres, err := suite.client.GetGenreMovieList(options)
	suite.NoError(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenreTVList() {
	genres, err := suite.client.GetGenreTVList(nil)
	suite.NoError(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenreTVListFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetGenreTVList(nil)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetGenreTVListWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	genres, err := suite.client.GetGenreTVList(options)
	suite.NoError(err)
	suite.NotNil(genres)
}
