package tmdb

const epGroupID = "5acf93e60e0a26346d0000ce"

func (suite *TMBDTestSuite) TestGetTVEpisodeGroupsDetails() {
	tv, err := suite.GetTVEpisodeGroupsDetails(epGroupID, nil)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeGroupsDetailsFail() {
	suite.Client.APIKey = ""
	_, err := suite.GetTVEpisodeGroupsDetails("", nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeGroupsDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.GetTVEpisodeGroupsDetails(epGroupID, options)
	suite.Nil(err)
	suite.NotNil(tv.ID)
}
