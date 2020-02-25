package tmdb

func (suite *TMBDTestSuite) TestGetTVEpisodeDetails() {
	got, err := suite.client.GetTVEpisodeDetails(gotID, 1, 1, nil)
	suite.Nil(err)
	suite.Equal("2011-04-17", got.AirDate)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeDetailsFail() {
	_, err := suite.client.GetTVEpisodeDetails(0, 1, 1, nil)
	suite.Equal("code: 34 | success: false | message: The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	got, err := suite.client.GetTVEpisodeDetails(gotID, 1, 1, options)
	suite.Nil(err)
	suite.Equal("2011-04-17", got.AirDate)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeChanges() {
	got, err := suite.client.GetTVEpisodeChanges(gotEpisodeID, nil)
	suite.Nil(err)
	suite.NotNil(got.Changes)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeChangesFail() {
	options := make(map[string]string)
	options["start_date"] = "2019-01-14"
	options["end_date"] = "2019-01-25"
	options["page"] = "1"
	_, err := suite.client.GetTVEpisodeChanges(0, options)
	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeChangesWithOptions() {
	options := make(map[string]string)
	options["start_date"] = "2019-01-14"
	options["end_date"] = "2019-01-25"
	options["page"] = "1"
	got, err := suite.client.GetTVEpisodeChanges(gotEpisodeID, options)
	suite.Nil(err)
	suite.Equal("overview", got.Changes[0].Key)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeCredits() {
	got, err := suite.client.GetTVEpisodeCredits(gotID, 1, 1)
	suite.Nil(err)
	suite.Equal(int64(gotEpisodeID), got.ID)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeCreditsFail() {
	_, err := suite.client.GetTVEpisodeCredits(0, 1, 1)
	suite.Equal("code: 34 | success: false | message: The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeExternalIDs() {
	got, err := suite.client.GetTVEpisodeExternalIDs(gotID, 1, 1)
	suite.Nil(err)
	suite.Equal(int64(gotEpisodeID), got.ID)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeExternalIDsFail() {
	_, err := suite.client.GetTVEpisodeExternalIDs(0, 1, 1)
	suite.Equal("code: 34 | success: false | message: The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeImages() {
	got, err := suite.client.GetTVEpisodeImages(gotID, 1, 1)
	suite.Nil(err)
	suite.Equal(int64(gotEpisodeID), got.ID)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeImagesFail() {
	_, err := suite.client.GetTVEpisodeImages(0, 1, 1)
	suite.Equal("code: 34 | success: false | message: The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeTranslations() {
	got, err := suite.client.GetTVEpisodeTranslations(gotID, 1, 1)
	suite.Nil(err)
	suite.Equal(int64(gotEpisodeID), got.ID)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeTranslationsFail() {
	_, err := suite.client.GetTVEpisodeTranslations(0, 1, 1)
	suite.Equal("code: 34 | success: false | message: The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeVideos() {
	got, err := suite.client.GetTVEpisodeVideos(gotID, 1, 1, nil)
	suite.Nil(err)
	suite.Equal(int64(gotEpisodeID), got.ID)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeVideosFail() {
	_, err := suite.client.GetTVEpisodeVideos(0, 1, 2, nil)
	suite.Equal("code: 34 | success: false | message: The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeVideosWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	got, err := suite.client.GetTVEpisodeVideos(gotID, 1, 2, options)
	suite.Nil(err)
	suite.Equal(int64(63057), got.ID)
}
