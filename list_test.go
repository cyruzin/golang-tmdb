package tmdb

const listID = "50941077760ee35e1500000c"

func (suite *TMBDTestSuite) TestGetListDetails() {
	list, err := suite.client.GetListDetails(listID, nil)
	suite.Nil(err)
	suite.NotNil(list.ID)
}

func (suite *TMBDTestSuite) TestGetListDetailsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetListDetails(listID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetListDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	list, err := suite.client.GetListDetails(listID, options)
	suite.Nil(err)
	suite.NotNil(list.ID)
}

func (suite *TMBDTestSuite) TestGetListItemStatus() {
	options := make(map[string]string)
	options["movie_id"] = "10658"
	list, err := suite.client.GetListItemStatus(listID, options)
	suite.Nil(err)
	suite.NotNil(list.ID)
}

func (suite *TMBDTestSuite) TestGetListItemStatusFail() {
	suite.client.apiKey = ""
	options := make(map[string]string)
	options["movie_id"] = "10658"
	_, err := suite.client.GetListItemStatus(listID, options)
	suite.NotNil(err)
}
