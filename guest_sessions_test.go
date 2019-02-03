package tmdb

func (suite *TMBDTestSuite) TestGetGuestSessionRatedMovies() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.CreateGuestSession()
	suite.Nil(err)
	guest, err := suite.GetGuestSessionRatedMovies(gs.GuestSessionID, o)
	suite.Nil(err)
	suite.NotNil(guest)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedMoviesFail() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.CreateGuestSession()
	suite.Nil(err)
	suite.APIKey = ""
	_, err = suite.GetGuestSessionRatedMovies(gs.GuestSessionID, o)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVShows() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.CreateGuestSession()
	suite.Nil(err)
	guest, err := suite.GetGuestSessionRatedTVShows(gs.GuestSessionID, o)
	suite.Nil(err)
	suite.NotNil(guest)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVShowsFail() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.CreateGuestSession()
	suite.Nil(err)
	suite.APIKey = ""
	_, err = suite.GetGuestSessionRatedTVShows(gs.GuestSessionID, o)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVEpisodes() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.CreateGuestSession()
	suite.Nil(err)
	guest, err := suite.GetGuestSessionRatedTVEpisodes(gs.GuestSessionID, o)
	suite.Nil(err)
	suite.NotNil(guest)
}

func (suite *TMBDTestSuite) TestGetGuestSessionRatedTVEpisodesFail() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["sort_by"] = "created_at.asc"
	gs, err := suite.CreateGuestSession()
	suite.Nil(err)
	suite.APIKey = ""
	_, err = suite.GetGuestSessionRatedTVEpisodes(gs.GuestSessionID, o)
	suite.NotNil(err)
}
