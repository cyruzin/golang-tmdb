package tmdb

func (suite *TMBDTestSuite) TestCreateGuestSession() {
	rt, err := suite.client.CreateGuestSession()
	suite.Nil(err)
	suite.Equal(true, rt.Success)
}

func (suite *TMBDTestSuite) TestCreateGuestSessionFail() {
	suite.client.apiKey = ""
	_, err := suite.client.CreateGuestSession()
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestCreateRequestToken() {
	rt, err := suite.client.CreateRequestToken()
	suite.Nil(err)
	suite.Equal(true, rt.Success)
}

func (suite *TMBDTestSuite) TestCreateRequestTokenFail() {
	suite.client.apiKey = ""
	_, err := suite.client.CreateRequestToken()
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

// TODO: How to validate a token that needs a browser?
// func (suite *TMBDTestSuite) TestCreateSession() {
// }

// func (suite *TMBDTestSuite) TestCreateSessionDenied() {
// 	session, err := suite.CreateSession("kpaishQpkpfVmbi")
// 	suite.Nil(session)
// 	suite.Equal("Session denied.", err.Error())
// }

// func (suite *TMBDTestSuite) TestCreateSessionFail() {
// 	suite.client.apiKey = ""
// 	_, err := suite.CreateRequestToken()
// 	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
// }
