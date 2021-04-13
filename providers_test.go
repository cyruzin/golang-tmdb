package tmdb

func (suite *TMBDTestSuite) TestGetAvailableWatchProviderRegions() {
	watchProviderRegions, err := suite.client.GetAvailableWatchProviderRegions(nil)
	suite.Nil(err)
	suite.NotNil(watchProviderRegions)
}

func (suite *TMBDTestSuite) TestGetAvailableWatchProviderRegionsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetAvailableWatchProviderRegions(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetAvailableWatchProviderRegionsWithOptions() {
	options := make(map[string]string)
	options["language"] = "de-DE"
	watchProviderRegions, err := suite.client.GetAvailableWatchProviderRegions(options)
	suite.Nil(err)
	suite.NotNil(watchProviderRegions)

	suite.Equal(watchProviderRegions.Regions[0].EnglishName, "Argentina")
	suite.Equal(watchProviderRegions.Regions[0].NativeName, "Argentinien")
}

func (suite *TMBDTestSuite) TestGetWatchProvidersMovie() {
	watchProviders, err := suite.client.GetWatchProvidersMovie(nil)
	suite.Nil(err)
	suite.NotNil(watchProviders)
}

func (suite *TMBDTestSuite) TestGetWatchProvidersMovieFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetWatchProvidersMovie(nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetWatchProvidersTv() {
	watchProviders, err := suite.client.GetWatchProvidersTv(nil)
	suite.Nil(err)
	suite.NotNil(watchProviders)
}

func (suite *TMBDTestSuite) TestGetWatchProvidersTvFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetWatchProvidersTv(nil)
	suite.NotNil(err)
}
