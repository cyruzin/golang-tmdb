package tmdb

const reviewID = "5488c29bc3a3686f4a00004a"

func (suite *TMBDTestSuite) TestGetReviewDetails() {
	review, err := suite.client.GetReviewDetails(reviewID)
	suite.NoError(err)
	suite.NotNil(review)
}

func (suite *TMBDTestSuite) TestGetReviewDetailsFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetReviewDetails(reviewID)
	suite.Error(err)
}
