package tmdb

const collectionID = 10
const collectionIDImage = 10

func (suite *TMBDTestSuite) TestGetCollectionDetails() {
	collection, err := suite.client.GetCollectionDetails(collectionID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionDetailsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetCollectionDetails(collectionID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.client.GetCollectionDetails(collectionID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionImages() {
	collection, err := suite.client.GetCollectionImages(collectionIDImage, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionImagesFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetCollectionImages(collectionID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	options["language"] = "en"
	collection, err := suite.client.GetCollectionImages(collectionIDImage, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionTranslations() {
	collection, err := suite.client.GetCollectionTranslations(collectionID, nil)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}

func (suite *TMBDTestSuite) TestGetCollectionTranslationsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetCollectionTranslations(collectionID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCollectionTranslationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	collection, err := suite.client.GetCollectionTranslations(collectionID, options)
	suite.Nil(err)
	suite.NotNil(collection.ID)
}
