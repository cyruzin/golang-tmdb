package tmdb

func (suite *TMBDTestSuite) TestGetCertificationMovie() {
	certification, err := suite.client.GetCertificationMovie()
	suite.Nil(err)
	suite.NotNil(certification.Certifications)
}

func (suite *TMBDTestSuite) TestGetCertificationMovieFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetCertificationMovie()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCertificationTV() {
	certification, err := suite.client.GetCertificationTV()
	suite.Nil(err)
	suite.NotNil(certification.Certifications)
}

func (suite *TMBDTestSuite) TestGetCertificationTVFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetCertificationTV()
	suite.NotNil(err)
}
