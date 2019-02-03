package tmdb

const starWarsID = 10

func (suite *TMBDTestSuite) TestGetCollectionDetails() {
	collection, err := suite.GetCollectionDetails(starWarsID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionDetailsFail() {
	suite.APIKey = ""
	_, err := suite.GetCollectionDetails(starWarsID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.GetCollectionDetails(starWarsID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionImages() {
	collection, err := suite.GetCollectionImages(starWarsID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionImagesFail() {
	suite.APIKey = ""
	_, err := suite.GetCollectionImages(starWarsID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.GetCollectionImages(starWarsID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionTranslations() {
	collection, err := suite.GetCollectionTranslations(starWarsID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionTranslationsFail() {
	suite.APIKey = ""
	_, err := suite.GetCollectionTranslations(starWarsID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionTranslationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.GetCollectionTranslations(starWarsID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}
