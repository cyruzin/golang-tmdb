package tmdb

const wormWholeID = 3417

func (suite *TMBDTestSuite) TestGetKeywordDetails() {
	keyword, err := suite.GetKeywordDetails(wormWholeID)
	suite.Nil(err)
	suite.NotNil(keyword)
}

func (suite *TMBDTestSuite) TestGetKeywordDetailsFail() {
	suite.APIKey = ""
	_, err := suite.GetKeywordDetails(wormWholeID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetKeywordMovies() {
	keyword, err := suite.GetKeywordMovies(wormWholeID, nil)
	suite.Nil(err)
	suite.NotNil(keyword)
}

func (suite *TMBDTestSuite) TestGetKeywordMoviesFail() {
	suite.APIKey = ""
	_, err := suite.GetKeywordMovies(wormWholeID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetKeywordMoviesWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	keyword, err := suite.GetKeywordMovies(wormWholeID, options)
	suite.Nil(err)
	suite.NotNil(keyword)
}
