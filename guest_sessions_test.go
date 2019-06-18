package tmdb

func (suite *TMBDTestSuite) TestGetGuestSessionRatedMovies() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	guest, err := suite.client.GetGuestSessionRatedMovies(gs.GuestSessionID, o)
	suite.Nil(err)
	suite.NotNil(guest)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedMoviesFail() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	suite.client.APIKey = ""
	_, err = suite.client.GetGuestSessionRatedMovies(gs.GuestSessionID, o)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVShows() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	guest, err := suite.client.GetGuestSessionRatedTVShows(gs.GuestSessionID, o)
	suite.Nil(err)
	suite.NotNil(guest)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVShowsFail() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	suite.client.APIKey = ""
	_, err = suite.client.GetGuestSessionRatedTVShows(gs.GuestSessionID, o)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVEpisodes() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	guest, err := suite.client.GetGuestSessionRatedTVEpisodes(gs.GuestSessionID, o)
	suite.Nil(err)
	suite.NotNil(guest)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVEpisodesFail() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	suite.client.APIKey = ""
	_, err = suite.client.GetGuestSessionRatedTVEpisodes(gs.GuestSessionID, o)
	suite.NotNil(err)
}
