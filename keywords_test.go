package tmdb

const wormWholeID = 3417

func (suite *TMBDTestSuite) TestGetKeywordDetails() {
	keyword, err := suite.client.GetKeywordDetails(wormWholeID)
	suite.NoError(err)
	suite.NotNil(keyword)
}

func (suite *TMBDTestSuite) TestGetKeywordDetailsFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetKeywordDetails(wormWholeID)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetKeywordMovies() {
	keyword, err := suite.client.GetKeywordMovies(wormWholeID, nil)
	suite.NoError(err)
	suite.NotNil(keyword)
}

func (suite *TMBDTestSuite) TestGetKeywordMoviesFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetKeywordMovies(wormWholeID, nil)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetKeywordMoviesWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	keyword, err := suite.client.GetKeywordMovies(wormWholeID, options)
	suite.NoError(err)
	suite.NotNil(keyword)
}
