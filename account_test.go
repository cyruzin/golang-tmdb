package tmdb

func (suite *TMBDTestSuite) TestGetAccountDetails() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
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
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	response, err := suite.client.GetCreatedLists(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountCreatedListsWithOptions() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
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
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	response, err := suite.client.GetFavoriteMovies(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountFavoriteMoviesWithOptions() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
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
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	response, err := suite.client.GetFavoriteTVShows(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountFavoriteWithOptionsTVShows() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
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
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	markAsFavorite := AccountFavorite{
		MediaType: "movie",
		MediaID:   500,
		Favorite:  true,
	}
	response, err := suite.client.MarkAsFavorite(0, markAsFavorite)
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
	_, err := suite.client.MarkAsFavorite(0, markAsFavorite)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAccountRatedMovies() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	response, err := suite.client.GetRatedMovies(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountRatedMoviesWithOptions() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	options := make(map[string]string)
	options["language"] = "pt-BR"
	options["page"] = "1"
	response, err := suite.client.GetRatedMovies(0, options)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountRatedMoviesFail() {
	suite.client.sessionID = ""
	_, err := suite.client.GetRatedMovies(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAccountRatedTVShows() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	response, err := suite.client.GetRatedTVShows(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountRatedTVShowsWithOptions() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	options := make(map[string]string)
	options["language"] = "pt-BR"
	options["page"] = "1"
	response, err := suite.client.GetRatedTVShows(0, options)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountRatedTVShowsFail() {
	suite.client.sessionID = ""
	_, err := suite.client.GetRatedTVShows(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAccountRatedTVEpisodes() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	response, err := suite.client.GetRatedTVEpisodes(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountRatedTVEpisodesWithOptions() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	options := make(map[string]string)
	options["language"] = "pt-BR"
	options["page"] = "1"
	response, err := suite.client.GetRatedTVEpisodes(0, options)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountRatedTVEpisodesFail() {
	suite.client.sessionID = ""
	_, err := suite.client.GetRatedTVEpisodes(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAccountMovieWatchlist() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	response, err := suite.client.GetMovieWatchlist(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountMovieWatchlistWithOptions() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	options := make(map[string]string)
	options["language"] = "pt-BR"
	options["page"] = "1"
	response, err := suite.client.GetMovieWatchlist(0, options)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountMovieWatchlistFail() {
	suite.client.sessionID = ""
	_, err := suite.client.GetMovieWatchlist(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAccountTVShowsWatchlist() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	response, err := suite.client.GetTVShowsWatchlist(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountTVShowsWatchlistWithOptions() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	options := make(map[string]string)
	options["language"] = "pt-BR"
	options["page"] = "1"
	response, err := suite.client.GetTVShowsWatchlist(0, options)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountTVShowsWatchlistFail() {
	suite.client.sessionID = ""
	_, err := suite.client.GetTVShowsWatchlist(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAddToWatchlist() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	addToWatchlist := AccountWatchlist{
		MediaType: "tv",
		MediaID:   82856,
		Watchlist: true,
	}
	response, err := suite.client.AddToWatchlist(0, addToWatchlist)
	suite.Nil(err)
	suite.NotNil(response.StatusMessage)
}

func (suite *TMBDTestSuite) TestAddToWatchlistFail() {
	suite.client.sessionID = ""
	addToWatchlist := AccountWatchlist{
		MediaType: "tv",
		MediaID:   82856,
		Watchlist: true,
	}
	_, err := suite.client.AddToWatchlist(0, addToWatchlist)
	suite.NotNil(err)
}
