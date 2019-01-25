package tmdb

func (suite *TMBDTestSuite) TestGetTVSeasonsDetails() {
	got, err := suite.GetTVSeasonsDetails(gotID, 1, nil)
	suite.Nil(err)
	suite.Equal("Season 1", got.Name)
}

func (suite *TMBDTestSuite) TestGetTVSeasonsDetailsFail() {
	_, err := suite.GetTVSeasonsDetails(0, 1, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVSeasonsDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	got, err := suite.GetTVSeasonsDetails(gotID, 1, options)
	suite.Nil(err)
	suite.Equal("1ª Temporada", got.Name)
}

func (suite *TMBDTestSuite) TestGetTVSeasonsChange() {
	options := make(map[string]string)
	options["start_date"] = "2019-01-14"
	options["end_date"] = "2019-01-25"
	options["page"] = "1"
	got, err := suite.GetTVSeasonsChanges(gotSeasonID, options)
	suite.Nil(err)
	suite.Equal("5c423aaf925141344cb32a9d", got.Changes[0].Items[0].ID)
}
