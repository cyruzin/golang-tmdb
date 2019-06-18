package tmdb

const starWarsID = 10

func (suite *TMBDTestSuite) TestGetCollectionDetails() {
	collection, err := suite.client.GetCollectionDetails(starWarsID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionDetailsFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetCollectionDetails(starWarsID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.client.GetCollectionDetails(starWarsID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionImages() {
	collection, err := suite.client.GetCollectionImages(starWarsID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionImagesFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetCollectionImages(starWarsID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.client.GetCollectionImages(starWarsID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionTranslations() {
	collection, err := suite.client.GetCollectionTranslations(starWarsID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionTranslationsFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetCollectionTranslations(starWarsID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionTranslationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.client.GetCollectionTranslations(starWarsID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}
