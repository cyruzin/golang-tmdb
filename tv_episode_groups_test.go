package tmdb

const epGroupID = "5acf93e60e0a26346d0000ce"

func (suite *TMBDTestSuite) TestGetTVEpisodeGroupsDetails() {
	tv, err := suite.client.GetTVEpisodeGroupsDetails(epGroupID, nil)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeGroupsDetailsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTVEpisodeGroupsDetails("", nil)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeGroupsDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.client.GetTVEpisodeGroupsDetails(epGroupID, options)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}
