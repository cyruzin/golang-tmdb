package tmdb

const jasonMomoaID = 117642

func (suite *TMBDTestSuite) TestGetPeopleDetails() {
	jasonMomoa, err := suite.GetPeopleDetails(jasonMomoaID, nil)
	suite.Nil(err)
	suite.Equal("Jason Momoa", jasonMomoa.Name)
}

func (suite *TMBDTestSuite) TestGetPeopleDetailsFail() {
	_, err := suite.GetPeopleDetails(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeopleDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.GetPeopleDetails(jasonMomoaID, options)
	suite.Nil(err)
	suite.Equal("Jason Momoa", jasonMomoa.Name)
}

func (suite *TMBDTestSuite) TestGetPeopleChanges() {
	jasonMomoa, err := suite.GetPeopleChanges(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.Changes)
}

func (suite *TMBDTestSuite) TestGetPeopleChangesFail() {
	suite.APIKey = ""
	_, err := suite.GetPeopleChanges(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeopleChangesWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	jasonMomoa, err := suite.GetPeopleChanges(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.Changes)
}
