package tmdb

func (suite *TMBDTestSuite) TestGetGenreMovieList() {
	genres, err := suite.GetGenreMovieList(nil)
	suite.Nil(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenreMovieListFail() {
	suite.APIKey = ""
	_, err := suite.GetGenreMovieList(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGenreMovieListWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	genres, err := suite.GetGenreMovieList(options)
	suite.Nil(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenreTVList() {
	genres, err := suite.GetGenreTVList(nil)
	suite.Nil(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenreTVListFail() {
	suite.APIKey = ""
	_, err := suite.GetGenreTVList(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGenreTVListWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	genres, err := suite.GetGenreTVList(options)
	suite.Nil(err)
	suite.NotNil(genres)
}
