package tmdb

func (suite *TMBDTestSuite) TestGetChangesMovie() {
	changes, err := suite.GetChangesMovie(nil)
	suite.Nil(err)
	suite.NotNil(changes.Page)
}

func (suite *TMBDTestSuite) TestGetChangesMovieFail() {
	suite.APIKey = ""
	_, err := suite.GetChangesMovie(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetChangesMovieWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	changes, err := suite.GetChangesMovie(options)
	suite.Nil(err)
	suite.NotNil(changes.Page)
}

func (suite *TMBDTestSuite) TestGetChangesTV() {
	changes, err := suite.GetChangesTV(nil)
	suite.Nil(err)
	suite.NotNil(changes.Page)
}

func (suite *TMBDTestSuite) TestGetChangesTVFail() {
	suite.APIKey = ""
	_, err := suite.GetChangesTV(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetChangesTVWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	changes, err := suite.GetChangesTV(options)
	suite.Nil(err)
	suite.NotNil(changes.Page)
}

func (suite *TMBDTestSuite) TestGetChangesPerson() {
	changes, err := suite.GetChangesPerson(nil)
	suite.Nil(err)
	suite.NotNil(changes.Page)
}

func (suite *TMBDTestSuite) TestGetChangesPersonFail() {
	suite.APIKey = ""
	_, err := suite.GetChangesPerson(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetChangesPersonWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	changes, err := suite.GetChangesPerson(options)
	suite.Nil(err)
	suite.NotNil(changes.Page)
}
