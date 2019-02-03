package tmdb

const (
	lucasFilmID = 1
	disneyID    = 2
)

func (suite *TMBDTestSuite) TestGetCompanyDetails() {
	company, err := suite.GetCompanyDetails(lucasFilmID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyDetailsFail() {
	suite.APIKey = ""
	_, err := suite.GetCompanyDetails(lucasFilmID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCompanyAlternativeNames() {
	company, err := suite.GetCompanyAlternativeNames(lucasFilmID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyAlternativeNamesFail() {
	suite.APIKey = ""
	_, err := suite.GetCompanyAlternativeNames(lucasFilmID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCompanyImages() {
	company, err := suite.GetCompanyImages(disneyID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompanyImagesFail() {
	suite.APIKey = ""
	_, err := suite.GetCompanyImages(disneyID)
	suite.NotNil(err)
}
