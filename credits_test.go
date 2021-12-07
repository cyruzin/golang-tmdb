package tmdb

const (
	tyrionLannisterID = "5256c8b219c2956ff6047cd8"
	oliverQueenID     = "5256ce3819c2956ff608293d"
)

func (suite *TMBDTestSuite) TestGetCreditDetails() {
	credit, err := suite.client.GetCreditDetails(oliverQueenID)
	suite.Nil(err)
	suite.NotNil(credit)
}

func (suite *TMBDTestSuite) TestGetCreditDetailsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetCreditDetails(tyrionLannisterID)
	suite.NotNil(err)
}
