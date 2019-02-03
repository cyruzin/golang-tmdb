package tmdb

const listID = "50941077760ee35e1500000c"

func (suite *TMBDTestSuite) TestGetListDetails() {
	list, err := suite.GetListDetails(listID, nil)
	suite.Nil(err)
	suite.NotNil(list.ID)
}

func (suite *TMBDTestSuite) TestGetListDetailsFail() {
	suite.Client.APIKey = ""
	_, err := suite.GetListDetails(listID, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetListDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	list, err := suite.GetListDetails(listID, options)
	suite.Nil(err)
	suite.NotNil(list.ID)
}

func (suite *TMBDTestSuite) TestGetListItemStatus() {
	options := make(map[string]string)
	options["movie_id"] = "10658"
	list, err := suite.GetListItemStatus(listID, options)
	suite.Nil(err)
	suite.NotNil(list.ID)
}

func (suite *TMBDTestSuite) TestGetListItemStatusFail() {
	suite.Client.APIKey = ""
	options := make(map[string]string)
	options["movie_id"] = "10658"
	_, err := suite.GetListItemStatus(listID, options)
	suite.NotNil(err)
}
