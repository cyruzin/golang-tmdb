package tmdb

func (suite *TMBDTestSuite) TestGetCertificationMovie() {
	certification, err := suite.GetCertificationMovie()
	suite.Nil(err)
	suite.NotNil(certification.Certifications)
}

func (suite *TMBDTestSuite) TestGetCertificationMovieFail() {
	suite.APIKey = ""
	_, err := suite.GetCertificationMovie()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCertificationTV() {
	certification, err := suite.GetCertificationTV()
	suite.Nil(err)
	suite.NotNil(certification.Certifications)
}

func (suite *TMBDTestSuite) TestGetCertificationTVFail() {
	suite.APIKey = ""
	_, err := suite.GetCertificationTV()
	suite.NotNil(err)
}
