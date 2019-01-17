package tmdb

func (suite *TMBDTestSuite) TestCreateGuestSession() {
	rt, err := suite.CreateGuestSession()
	suite.Nil(err)
	suite.Equal(true, rt.Success)
}

func (suite *TMBDTestSuite) TestCreateGuestSessionFail() {
	suite.Client.APIKey = ""
	_, err := suite.CreateGuestSession()
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestCreateRequestToken() {
	rt, err := suite.CreateRequestToken()
	suite.Nil(err)
	suite.Equal(true, rt.Success)
}

func (suite *TMBDTestSuite) TestCreateRequestTokenFail() {
	suite.Client.APIKey = ""
	_, err := suite.CreateRequestToken()
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

// TODO: How to validate a token that needs a browser?
// func (suite *TMBDTestSuite) TestCreateSession() {
// }

func (suite *TMBDTestSuite) TestCreateSessionDenied() {
	session, err := suite.CreateSession("kpaishQpkpfVmbi")
	suite.Nil(session)
	suite.Equal("Session denied.", err.Error())
}

func (suite *TMBDTestSuite) TestCreateSessionFail() {
	suite.Client.APIKey = ""
	_, err := suite.CreateRequestToken()
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}
