package tmdb

func (suite *TMBDTestSuite) TestGetGenreMovieList() {
	genres, err := suite.client.GetGenreMovieList(nil)
	suite.Nil(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenreMovieListFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetGenreMovieList(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGenreMovieListWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	genres, err := suite.client.GetGenreMovieList(options)
	suite.Nil(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenreTVList() {
	genres, err := suite.client.GetGenreTVList(nil)
	suite.Nil(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenreTVListFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetGenreTVList(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGenreTVListWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	genres, err := suite.client.GetGenreTVList(options)
	suite.Nil(err)
	suite.NotNil(genres)
}
