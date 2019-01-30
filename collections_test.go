package tmdb

const starWarsID = 10

func (suite *TMBDTestSuite) TestGetCollectionsDetails() {
	collection, err := suite.GetCollectionsDetails(starWarsID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionsDetailsFail() {
	suite.APIKey = ""
	_, err := suite.GetCollectionsDetails(starWarsID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionsDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.GetCollectionsDetails(starWarsID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionsImages() {
	collection, err := suite.GetCollectionsImages(starWarsID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionsImagesFail() {
	suite.APIKey = ""
	_, err := suite.GetCollectionsImages(starWarsID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionsImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.GetCollectionsImages(starWarsID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionsTranslations() {
	collection, err := suite.GetCollectionsTranslations(starWarsID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionsTranslationsFail() {
	suite.APIKey = ""
	_, err := suite.GetCollectionsTranslations(starWarsID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionsTranslationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.GetCollectionsTranslations(starWarsID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}
