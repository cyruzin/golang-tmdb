package tmdb

func (suite *TMBDTestSuite) TestGetAvailableWatchProviderRegions() {
	watchProviderRegions, err := suite.client.GetAvailableWatchProviderRegions(nil)
	suite.NoError(err)
	suite.NotNil(watchProviderRegions)
}

func (suite *TMBDTestSuite) TestGetAvailableWatchProviderRegionsFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetAvailableWatchProviderRegions(nil)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetAvailableWatchProviderRegionsWithOptions() {
	options := make(map[string]string)
	options["language"] = "de-DE"
	watchProviderRegions, err := suite.client.GetAvailableWatchProviderRegions(options)
	suite.NoError(err)
	suite.NotNil(watchProviderRegions)
}

func (suite *TMBDTestSuite) TestGetWatchProvidersMovie() {
	watchProviders, err := suite.client.GetWatchProvidersMovie(nil)
	suite.NoError(err)
	suite.NotNil(watchProviders)
}

func (suite *TMBDTestSuite) TestGetWatchProvidersMovieFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetWatchProvidersMovie(nil)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetWatchProvidersTv() {
	watchProviders, err := suite.client.GetWatchProvidersTv(nil)
	suite.NoError(err)
	suite.NotNil(watchProviders)
}

func (suite *TMBDTestSuite) TestGetWatchProvidersTvFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetWatchProvidersTv(nil)
	suite.Error(err)
}
