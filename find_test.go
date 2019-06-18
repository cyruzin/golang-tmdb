package tmdb

const theWalkingDeadID = "tt1520211"

func (suite *TMBDTestSuite) TestGetFindByID() {
	options := make(map[string]string)
	options["external_source"] = "imdb_id"
	options["language"] = "en-US"
	genres, err := suite.client.GetFindByID(theWalkingDeadID, options)
	suite.Nil(err)
	suite.NotNil(genres)
}

func (suite *TMBDTestSuite) TestGetFindByIDFail() {
	suite.client.APIKey = ""
	options := make(map[string]string)
	options["external_source"] = "imdb_id"
	options["language"] = "en-US"
	_, err := suite.client.GetFindByID(theWalkingDeadID, options)
	suite.NotNil(err)
}
