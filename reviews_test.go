package tmdb

const reviewID = "5488c29bc3a3686f4a00004a"

func (suite *TMBDTestSuite) TestGetReviewDetails() {
	review, err := suite.GetReviewDetails(reviewID)
	suite.Nil(err)
	suite.NotNil(review)
}

func (suite *TMBDTestSuite) TestGetReviewDetailsFail() {
	suite.APIKey = ""
	_, err := suite.GetReviewDetails(reviewID)
	suite.NotNil(err)
}
