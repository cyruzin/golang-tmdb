package tmdb

func (suite *TMBDTestSuite) TestGetGuestSessionRatedMovies() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	_, err = suite.client.GetGuestSessionRatedMovies(gs.GuestSessionID, o)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedMoviesFail() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	suite.client.apiKey = ""
	_, err = suite.client.GetGuestSessionRatedMovies(gs.GuestSessionID, o)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVShows() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	_, err = suite.client.GetGuestSessionRatedTVShows(gs.GuestSessionID, o)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVShowsFail() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	suite.client.apiKey = ""
	_, err = suite.client.GetGuestSessionRatedTVShows(gs.GuestSessionID, o)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVEpisodes() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	_, err = suite.client.GetGuestSessionRatedTVEpisodes(gs.GuestSessionID, o)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVEpisodesFail() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	suite.client.apiKey = ""
	_, err = suite.client.GetGuestSessionRatedTVEpisodes(gs.GuestSessionID, o)
	suite.NotNil(err)
}
