package tmdb

func (suite *TMBDTestSuite) TestGetGenresMovieList() {
	genres, err := suite.GetGenresMovieList(nil)
	suite.Nil(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenresMovieListFail() {
	suite.APIKey = ""
	_, err := suite.GetGenresMovieList(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGenresMovieListWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	genres, err := suite.GetGenresMovieList(options)
	suite.Nil(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenresTVList() {
	genres, err := suite.GetGenresTVList(nil)
	suite.Nil(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetGenresTVListFail() {
	suite.APIKey = ""
	_, err := suite.GetGenresTVList(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGenresTVListWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	genres, err := suite.GetGenresTVList(options)
	suite.Nil(err)
	suite.NotNil(genres)
}
