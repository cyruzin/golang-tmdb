package tmdb

const reviewID = "5488c29bc3a3686f4a00004a"

func (suite *TMBDTestSuite) TestGetReviewDetails() {
	review, err := suite.client.GetReviewDetails(reviewID)
	suite.Nil(err)
	suite.NotNil(review)
}

func (suite *TMBDTestSuite) TestGetReviewDetailsFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetReviewDetails(reviewID)
	suite.NotNil(err)
}
