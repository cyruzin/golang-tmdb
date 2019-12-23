package tmdb

func (suite *TMBDTestSuite) TestGetAccountDetails() {
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.GetAccountDetails()
	suite.Nil(err)
	suite.NotNil(response.ID)
}

func (suite *TMBDTestSuite) TestGetAccountDetailsFail() {
	suite.client.sessionID = ""
	_, err := suite.client.GetAccountDetails()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAccountCreatedLists() {
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.GetCreatedLists(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountCreatedListsWithOptions() {
	suite.client.SetSessionID(sessionID)
	options := make(map[string]string)
	options["language"] = "pt-BR"
	options["page"] = "1"
	response, err := suite.client.GetCreatedLists(0, options)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountCreatedListsFail() {
	suite.client.sessionID = ""
	_, err := suite.client.GetCreatedLists(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAccountFavoriteMovies() {
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.GetFavoriteMovies(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountFavoriteMoviesWithOptions() {
	suite.client.SetSessionID(sessionID)
	options := make(map[string]string)
	options["language"] = "pt-BR"
	options["page"] = "1"
	response, err := suite.client.GetFavoriteMovies(0, options)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountFavoriteMovieFail() {
	suite.client.sessionID = ""
	_, err := suite.client.GetFavoriteMovies(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAccountFavoriteTVShows() {
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.GetFavoriteTVShows(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountFavoriteWithOptionsTVShows() {
	suite.client.SetSessionID(sessionID)
	options := make(map[string]string)
	options["language"] = "pt-BR"
	options["page"] = "1"
	response, err := suite.client.GetFavoriteTVShows(0, options)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountFavoriteFailTVShows() {
	suite.client.sessionID = ""
	_, err := suite.client.GetFavoriteTVShows(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestMarkAsFavorite() {
	suite.client.SetSessionID(sessionID)
	markAsFavorite := AccountFavorite{
		MediaType: "movie",
		MediaID:   500,
		Favorite:  true,
	}
	response, err := suite.client.AccountMarkAsFavorite(0, &markAsFavorite)
	suite.Nil(err)
	suite.NotNil(response.StatusMessage)
}

func (suite *TMBDTestSuite) TestMarkAsFavoriteFail() {
	suite.client.sessionID = ""
	markAsFavorite := AccountFavorite{
		MediaType: "movie",
		MediaID:   500,
		Favorite:  true,
	}
	_, err := suite.client.AccountMarkAsFavorite(0, &markAsFavorite)
	suite.NotNil(err)
}
