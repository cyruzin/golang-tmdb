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

func (suite *TMBDTestSuite) TestGetPeopleMovieCredits() {
	jasonMomoa, err := suite.GetPeopleMovieCredits(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleMovieCreditsFail() {
	suite.APIKey = ""
	_, err := suite.GetPeopleMovieCredits(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeopleMovieCreditsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.GetPeopleMovieCredits(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleTVCredits() {
	jasonMomoa, err := suite.GetPeopleTVCredits(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleTVCreditsFail() {
	suite.APIKey = ""
	_, err := suite.GetPeopleTVCredits(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeopleTVCreditsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.GetPeopleTVCredits(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleCombinedCredits() {
	jasonMomoa, err := suite.GetPeopleCombinedCredits(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleCombinedCreditsFail() {
	suite.APIKey = ""
	_, err := suite.GetPeopleCombinedCredits(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeopleCombinedCreditsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.GetPeopleCombinedCredits(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleExternalIDs() {
	jasonMomoa, err := suite.GetPeopleExternalIDs(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleExternalIDsFail() {
	suite.APIKey = ""
	_, err := suite.GetPeopleExternalIDs(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeopleExternalIDsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.GetPeopleExternalIDs(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}
