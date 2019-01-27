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

func (suite *TMBDTestSuite) TestGetTVEpisodeChanges() {
	got, err := suite.GetTVEpisodeChanges(gotEpisodeID, nil)
	suite.Nil(err)
	suite.NotNil(got.Changes)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeChangesFail() {
	options := make(map[string]string)
	options["start_date"] = "2019-01-14"
	options["end_date"] = "2019-01-25"
	options["page"] = "1"
	_, err := suite.GetTVEpisodeChanges(0, options)
	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeChangesWithOptions() {
	options := make(map[string]string)
	options["start_date"] = "2019-01-14"
	options["end_date"] = "2019-01-25"
	options["page"] = "1"
	got, err := suite.GetTVEpisodeChanges(gotEpisodeID, options)
	suite.Nil(err)
	suite.Equal("overview", got.Changes[0].Key)
}
