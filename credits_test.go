package tmdb

const tyrionLannisterID = "5256c8b219c2956ff6047cd8"
const oliverQueenID = "5256ce3819c2956ff608293d"

// Commented until they fix this issue
//
//--- FAIL: TestSuite/TestGetCreditDetails (0.02s)
//credits_test.go:8:
//		Error Trace:	credits_test.go:8
//		Error:      	Expected nil, but got: &errors.errorString{s:"could not decode the data: json: cannot unmarshal number 18.0 into Go struct field .person.known_for.genre_ids of type int64"}
//		Test:       	TestSuite/TestGetCreditDetails
//credits_test.go:9:
//		Error Trace:	credits_test.go:9
//		Error:      	Expected value not to be nil.
//		Test:       	TestSuite/TestGetCreditDetails
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
