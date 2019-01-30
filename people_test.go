package tmdb

const (
	tomCruiseID  = 500
	jasonMomoaID = 117642
)

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

func (suite *TMBDTestSuite) TestGetPeopleImages() {
	jasonMomoa, err := suite.GetPeopleImages(jasonMomoaID)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleImagesFail() {
	suite.APIKey = ""
	_, err := suite.GetPeopleImages(0)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeopleTaggedImages() {
	tomCruise, err := suite.GetPeopleTaggedImages(tomCruiseID, nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleTaggedImagesFail() {
	suite.APIKey = ""
	_, err := suite.GetPeopleTaggedImages(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeopleTaggedImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.GetPeopleTaggedImages(tomCruiseID, options)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleTranslations() {
	tomCruise, err := suite.GetPeopleTranslations(tomCruiseID, nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleTranslationsFail() {
	suite.APIKey = ""
	_, err := suite.GetPeopleTranslations(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeopleTranslationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.GetPeopleTranslations(tomCruiseID, options)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleLatest() {
	tomCruise, err := suite.GetPeopleLatest(nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPeopleLatestFail() {
	suite.APIKey = ""
	_, err := suite.GetPeopleLatest(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeopleLatestWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.GetPeopleLatest(options)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPeoplePopular() {
	tomCruise, err := suite.GetPeoplePopular(nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.Page)
}

func (suite *TMBDTestSuite) TestGetPeoplePopularFail() {
	suite.APIKey = ""
	_, err := suite.GetPeoplePopular(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPeoplePopularWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.GetPeoplePopular(options)
	suite.Nil(err)
	suite.NotNil(tomCruise.Page)
}
