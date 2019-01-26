package tmdb

func (suite *TMBDTestSuite) TestGetTVEpisodeDetails() {
	got, err := suite.GetTVEpisodeDetails(gotID, 1, 1, nil)
	suite.Nil(err)
	suite.Equal("2011-04-17", got.AirDate)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeDetailsFail() {
	_, err := suite.GetTVEpisodeDetails(0, 1, 1, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	got, err := suite.GetTVEpisodeDetails(gotID, 1, 1, options)
	suite.Nil(err)
	suite.Equal("2011-04-17", got.AirDate)
}
