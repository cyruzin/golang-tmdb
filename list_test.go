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

func (suite *TMBDTestSuite) TestCreateList() {
	suite.client.SetSessionID(sessionID)
	createList := ListCreate{
		Name:        "Test list!",
		Description: "Description is not enough...",
		Language:    "en",
	}
	response, err := suite.client.CreateList(&createList)
	suite.Nil(err)
	suite.NotNil(response.StatusMessage)
}

func (suite *TMBDTestSuite) TestCreateListFail() {
	suite.client.sessionID = ""
	createList := ListCreate{
		Name:        "Wrapper test list",
		Description: "This is a description",
		Language:    "en",
	}
	_, err := suite.client.CreateList(&createList)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestClearList() {
	suite.client.SetSessionID(sessionID)
	_, err := suite.client.ClearList(128331, true)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestClearListFail() {
	suite.client.sessionID = ""
	_, err := suite.client.ClearList(128331, true)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestDeleteList() {
	suite.client.SetSessionID(sessionID)
	_, err := suite.client.DeleteList(128331)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestDeleteListFail() {
	suite.client.sessionID = ""
	createList := ListCreate{
		Name:        "Wrapper test list",
		Description: "This is a description",
		Language:    "en",
	}
	_, err := suite.client.CreateList(&createList)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestAddMovie() {
	suite.client.SetSessionID(sessionID)
	mediaID := ListMedia{MediaID: 500}
	_, err := suite.client.AddMovie(128331, &mediaID)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestAddMovieFail() {
	suite.client.sessionID = ""
	mediaID := ListMedia{MediaID: 500}
	_, err := suite.client.AddMovie(128331, &mediaID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestRemoveMovie() {
	suite.client.SetSessionID(sessionID)
	mediaID := ListMedia{MediaID: 500}
	_, err := suite.client.RemoveMovie(128331, &mediaID)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestRemoveMovieFail() {
	suite.client.sessionID = ""
	mediaID := ListMedia{MediaID: 500}
	_, err := suite.client.RemoveMovie(128331, &mediaID)
	suite.NotNil(err)
}
