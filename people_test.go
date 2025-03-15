package tmdb

const (
	tomCruiseID  = 500
	jasonMomoaID = 117642
)

func (suite *TMBDTestSuite) TestGetPersonDetails() {
	jasonMomoa, err := suite.client.GetPersonDetails(jasonMomoaID, nil)
	suite.Nil(err)
	suite.Equal("Jason Momoa", jasonMomoa.Name)
}

func (suite *TMBDTestSuite) TestGetPersonDetailsFail() {
	_, err := suite.client.GetPersonDetails(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.client.GetPersonDetails(jasonMomoaID, options)
	suite.Nil(err)
	suite.Equal("Jason Momoa", jasonMomoa.Name)
}

func (suite *TMBDTestSuite) TestGetPersonDetailsWithAppend() {
	options := make(map[string]string)
	options["append_to_response"] = "translations"
	jasonMomoa, err := suite.client.GetPersonDetails(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.PersonTranslationsAppend)
	suite.NotNil(jasonMomoa.Translations)
	suite.NotEmpty(jasonMomoa.Translations.Translations)
	suite.Equal("Jason Momoa", jasonMomoa.Translations.Translations[0].Data.Name)
}

func (suite *TMBDTestSuite) TestGetPersonChanges() {
	jasonMomoa, err := suite.client.GetPersonChanges(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.Changes)
}

func (suite *TMBDTestSuite) TestGetPersonChangesFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetPersonChanges(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonChangesWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	jasonMomoa, err := suite.client.GetPersonChanges(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.Changes)
}

func (suite *TMBDTestSuite) TestGetPersonMovieCredits() {
	jasonMomoa, err := suite.client.GetPersonMovieCredits(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonMovieCreditsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetPersonMovieCredits(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonMovieCreditsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.client.GetPersonMovieCredits(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonTVCredits() {
	jasonMomoa, err := suite.client.GetPersonTVCredits(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonTVCreditsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetPersonTVCredits(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonTVCreditsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.client.GetPersonTVCredits(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonCombinedCredits() {
	jasonMomoa, err := suite.client.GetPersonCombinedCredits(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonCombinedCreditsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetPersonCombinedCredits(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonCombinedCreditsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.client.GetPersonCombinedCredits(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonExternalIDs() {
	jasonMomoa, err := suite.client.GetPersonExternalIDs(jasonMomoaID, nil)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonExternalIDsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetPersonExternalIDs(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonExternalIDsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jasonMomoa, err := suite.client.GetPersonExternalIDs(jasonMomoaID, options)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonImages() {
	jasonMomoa, err := suite.client.GetPersonImages(jasonMomoaID)
	suite.Nil(err)
	suite.NotNil(jasonMomoa.ID)
}

func (suite *TMBDTestSuite) TestGetPersonImagesFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetPersonImages(0)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonTaggedImages() {
	tomCruise, err := suite.client.GetPersonTaggedImages(tomCruiseID, nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonTaggedImagesFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetPersonTaggedImages(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonTaggedImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.client.GetPersonTaggedImages(tomCruiseID, options)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonTranslations() {
	tomCruise, err := suite.client.GetPersonTranslations(tomCruiseID, nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonTranslationsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetPersonTranslations(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonTranslationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.client.GetPersonTranslations(tomCruiseID, options)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
	suite.NotEmpty(tomCruise.Translations)
	suite.NotEmpty(tomCruise.Translations[0].Data.Name)
}

func (suite *TMBDTestSuite) TestGetPersonLatest() {
	tomCruise, err := suite.client.GetPersonLatest(nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonLatestFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetPersonLatest(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonLatestWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.client.GetPersonLatest(options)
	suite.Nil(err)
	suite.NotNil(tomCruise.ID)
}

func (suite *TMBDTestSuite) TestGetPersonPopular() {
	tomCruise, err := suite.client.GetPersonPopular(nil)
	suite.Nil(err)
	suite.NotNil(tomCruise.Page)
}

func (suite *TMBDTestSuite) TestGetPersonPopularFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetPersonPopular(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetPersonPopularWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tomCruise, err := suite.client.GetPersonPopular(options)
	suite.Nil(err)
	suite.NotNil(tomCruise.Page)
}
