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
	response, err := suite.client.MarkAsFavorite(0, &markAsFavorite)
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
	_, err := suite.client.MarkAsFavorite(0, &markAsFavorite)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAccountRatedMovies() {
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.GetRatedMovies(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountRatedMoviesWithOptions() {
	suite.client.SetSessionID(sessionID)
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
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.GetRatedTVShows(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountRatedTVShowsWithOptions() {
	suite.client.SetSessionID(sessionID)
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
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.GetRatedTVEpisodes(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountRatedTVEpisodesWithOptions() {
	suite.client.SetSessionID(sessionID)
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
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.GetMovieWatchlist(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountMovieWatchlistWithOptions() {
	suite.client.SetSessionID(sessionID)
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
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.GetTVShowsWatchlist(0, nil)
	suite.Nil(err)
	suite.NotNil(response.Page)
}

func (suite *TMBDTestSuite) TestAccountTVShowsWatchlistWithOptions() {
	suite.client.SetSessionID(sessionID)
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
	suite.client.SetSessionID(sessionID)
	addToWatchlist := AccountWatchlist{
		MediaType: "tv",
		MediaID:   82856,
		Watchlist: true,
	}
	response, err := suite.client.AddToWatchlist(0, &addToWatchlist)
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
	_, err := suite.client.AddToWatchlist(0, &addToWatchlist)
	suite.NotNil(err)
}
