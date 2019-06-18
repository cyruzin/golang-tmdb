package tmdb

const (
	lucasFilmID = 1
	disneyID    = 2
)

func (suite *TMBDTestSuite) TestGetCompanyDetails() {
	company, err := suite.client.GetCompanyDetails(lucasFilmID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyDetailsFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetCompanyDetails(lucasFilmID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCompanyAlternativeNames() {
	company, err := suite.client.GetCompanyAlternativeNames(lucasFilmID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyAlternativeNamesFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetCompanyAlternativeNames(lucasFilmID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCompanyImages() {
	company, err := suite.client.GetCompanyImages(disneyID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyImagesFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetCompanyImages(disneyID)
	suite.NotNil(err)
}
