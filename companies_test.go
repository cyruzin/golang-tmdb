package tmdb

const (
	lucasFilmID = 1
	randomID    = 3
)

func (suite *TMBDTestSuite) TestGetCompanyDetails() {
	company, err := suite.client.GetCompanyDetails(lucasFilmID)
	suite.NoError(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyDetailsFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetCompanyDetails(lucasFilmID)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetCompanyAlternativeNames() {
	company, err := suite.client.GetCompanyAlternativeNames(lucasFilmID)
	suite.NoError(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyAlternativeNamesFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetCompanyAlternativeNames(lucasFilmID)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetCompanyImages() {
	company, err := suite.client.GetCompanyImages(randomID)
	suite.NoError(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyImagesFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetCompanyImages(randomID)
	suite.Error(err)
}
