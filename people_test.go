package tmdb

const (
	tomCruiseID  = 500
	jasonMomoaID = 117642
)

func (suite *TMBDTestSuite) TestGetPersonDetails() {
	jasonMomoa, err := suite.GetPersonDetails(jasonMomoaID, nil)
	suite.Nil(err)
	suite.Equal("Jason Momoa", jasonMomoa.Name)
}

func (suite *TMBDTestSuite) TestGetPersonDetailsFail() {
	_, err := suite.GetPersonDetails(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.GetPersonDetails(jasonMomoaID, options)
	suite.Nil(err)
	suite.Equal("Jason Momoa", jasonMomoa.Name)
}

func (suite *TMBDTestSuite) TestGetPersonChanges() {
	jasonMomoa, err := suite.GetPersonChanges(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.Changes)
}

func (suite *TMBDTestSuite) TestGetPersonChangesFail() {
	suite.APIKey = ""
	_, err := suite.GetPersonChanges(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonChangesWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	jasonMomoa, err := suite.GetPersonChanges(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.Changes)
}

func (suite *TMBDTestSuite) TestGetPersonMovieCredits() {
	jasonMomoa, err := suite.GetPersonMovieCredits(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonMovieCreditsFail() {
	suite.APIKey = ""
	_, err := suite.GetPersonMovieCredits(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonMovieCreditsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.GetPersonMovieCredits(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonTVCredits() {
	jasonMomoa, err := suite.GetPersonTVCredits(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonTVCreditsFail() {
	suite.APIKey = ""
	_, err := suite.GetPersonTVCredits(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonTVCreditsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.GetPersonTVCredits(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonCombinedCredits() {
	jasonMomoa, err := suite.GetPersonCombinedCredits(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonCombinedCreditsFail() {
	suite.APIKey = ""
	_, err := suite.GetPersonCombinedCredits(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonCombinedCreditsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.GetPersonCombinedCredits(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonExternalIDs() {
	jasonMomoa, err := suite.GetPersonExternalIDs(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonExternalIDsFail() {
	suite.APIKey = ""
	_, err := suite.GetPersonExternalIDs(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonExternalIDsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.GetPersonExternalIDs(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonImages() {
	jasonMomoa, err := suite.GetPersonImages(jasonMomoaID)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonImagesFail() {
	suite.APIKey = ""
	_, err := suite.GetPersonImages(0)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonTaggedImages() {
	tomCruise, err := suite.GetPersonTaggedImages(tomCruiseID, nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonTaggedImagesFail() {
	suite.APIKey = ""
	_, err := suite.GetPersonTaggedImages(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonTaggedImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.GetPersonTaggedImages(tomCruiseID, options)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonTranslations() {
	tomCruise, err := suite.GetPersonTranslations(tomCruiseID, nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonTranslationsFail() {
	suite.APIKey = ""
	_, err := suite.GetPersonTranslations(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonTranslationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.GetPersonTranslations(tomCruiseID, options)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonLatest() {
	tomCruise, err := suite.GetPersonLatest(nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonLatestFail() {
	suite.APIKey = ""
	_, err := suite.GetPersonLatest(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonLatestWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.GetPersonLatest(options)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonPopular() {
	tomCruise, err := suite.GetPersonPopular(nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.Page)
}

func (suite *TMBDTestSuite) TestGetPersonPopularFail() {
	suite.APIKey = ""
	_, err := suite.GetPersonPopular(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonPopularWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.GetPersonPopular(options)
	suite.Nil(err)
	suite.NotNil(tomCruise.Page)
}
