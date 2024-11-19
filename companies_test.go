package tmdb

const (
	lucasFilmID = 1
	randomID    = 3
)

func (suite *TMBDTestSuite) TestGetCompanyDetails() {
	company, err := suite.client.GetCompanyDetails(lucasFilmID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyDetailsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetCompanyDetails(lucasFilmID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCompanyAlternativeNames() {
	company, err := suite.client.GetCompanyAlternativeNames(lucasFilmID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyAlternativeNamesFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetCompanyAlternativeNames(lucasFilmID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCompanyImages() {
	company, err := suite.client.GetCompanyImages(randomID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyImagesFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetCompanyImages(randomID)
	suite.NotNil(err)
}
