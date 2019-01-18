package tmdb

const (
	vikingsID = 44217
)

func (suite *TMBDTestSuite) TestGetTVDetails() {
	vikings, err := suite.GetTVDetails(vikingsID, nil)
	suite.Nil(err)
	suite.Equal("Vikings", vikings.Name)
}

func (suite *TMBDTestSuite) TestGetTVDetailsFail() {
	_, err := suite.GetTVDetails(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	vikings, err := suite.GetTVDetails(vikingsID, options)
	suite.Nil(err)
	suite.Equal("Vikings", vikings.Name)
}
