package tmdb

func (suite *TMBDTestSuite) TestGetCertificationMovie() {
	certification, err := suite.client.GetCertificationMovie()
	suite.NoError(err)
	suite.NotNil(certification.Certifications)
}

func (suite *TMBDTestSuite) TestGetCertificationMovieFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetCertificationMovie()
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetCertificationTV() {
	certification, err := suite.client.GetCertificationTV()
	suite.NoError(err)
	suite.NotNil(certification.Certifications)
}

func (suite *TMBDTestSuite) TestGetCertificationTVFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetCertificationTV()
	suite.Error(err)
}
