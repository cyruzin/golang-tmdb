package tmdb

const (
	lucasFilmID = 1
	disneyID    = 2
)

func (suite *TMBDTestSuite) TestGetCompaniesDetails() {
	company, err := suite.GetCompaniesDetails(lucasFilmID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompaniesDetailsFail() {
	suite.APIKey = ""
	_, err := suite.GetCompaniesDetails(lucasFilmID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCompaniesAlternativeNames() {
	company, err := suite.GetCompaniesAlternativeNames(lucasFilmID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompaniesAlternativeNamesFail() {
	suite.APIKey = ""
	_, err := suite.GetCompaniesAlternativeNames(lucasFilmID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetCompaniesImages() {
	company, err := suite.GetCompaniesImages(disneyID)
	suite.Nil(err)
	suite.NotNil(company.ID)
}

func (suite *TMBDTestSuite) TestGetCompaniesImagesFail() {
	suite.APIKey = ""
	_, err := suite.GetCompaniesImages(disneyID)
	suite.NotNil(err)
}
