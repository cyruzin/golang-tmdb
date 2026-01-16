package tmdb

func (suite *TMBDTestSuite) TestGetChangesMovie() {
	changes, err := suite.client.GetChangesMovie(nil)
	suite.NoError(err)
	suite.NotNil(changes.Page)
}

func (suite *TMBDTestSuite) TestGetChangesMovieFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetChangesMovie(nil)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetChangesMovieWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	changes, err := suite.client.GetChangesMovie(options)
	suite.NoError(err)
	suite.NotNil(changes.Page)
}

func (suite *TMBDTestSuite) TestGetChangesTV() {
	changes, err := suite.client.GetChangesTV(nil)
	suite.NoError(err)
	suite.NotNil(changes.Page)
}

func (suite *TMBDTestSuite) TestGetChangesTVFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetChangesTV(nil)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetChangesTVWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	changes, err := suite.client.GetChangesTV(options)
	suite.NoError(err)
	suite.NotNil(changes.Page)
}

func (suite *TMBDTestSuite) TestGetChangesPerson() {
	changes, err := suite.client.GetChangesPerson(nil)
	suite.NoError(err)
	suite.NotNil(changes.Page)
}

func (suite *TMBDTestSuite) TestGetChangesPersonFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetChangesPerson(nil)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetChangesPersonWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	changes, err := suite.client.GetChangesPerson(options)
	suite.NoError(err)
	suite.NotNil(changes.Page)
}
