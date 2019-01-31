package tmdb

const tyrionLannisterID = "5256c8b219c2956ff6047cd8"

func (suite *TMBDTestSuite) TestGetCreditDetails() {
	credit, err := suite.GetCreditDetails(tyrionLannisterID)
	suite.Nil(err)
	suite.NotNil(credit)
}

func (suite *TMBDTestSuite) TestGetCreditDetailsFail() {
	suite.APIKey = ""
	_, err := suite.GetCreditDetails(tyrionLannisterID)
	suite.NotNil(err)
}
