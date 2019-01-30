package tmdb

func (suite *TMBDTestSuite) TestGetCertificationsMovie() {
	certifications, err := suite.GetCertificationsMovie()
	suite.Nil(err)
	suite.NotNil(certifications.Certifications)
}

func (suite *TMBDTestSuite) TestGetCertificationsMovieFail() {
	suite.APIKey = ""
	_, err := suite.GetCertificationsMovie()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCertificationsTV() {
	certifications, err := suite.GetCertificationsTV()
	suite.Nil(err)
	suite.NotNil(certifications.Certifications)
}

func (suite *TMBDTestSuite) TestGetCertificationsTVFail() {
	suite.APIKey = ""
	_, err := suite.GetCertificationsTV()
	suite.NotNil(err)
}
