package tmdb

func (suite *TMBDTestSuite) TestGetConfigurationAPI() {
	configuration, err := suite.client.GetConfigurationAPI()
	suite.NoError(err)
	suite.NotNil(configuration.Images.BackdropSizes)
}

func (suite *TMBDTestSuite) TestGetConfigurationAPIFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetConfigurationAPI()
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationCountries() {
	configuration, err := suite.client.GetConfigurationCountries()
	suite.NoError(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationCountriesFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetConfigurationCountries()
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationJobs() {
	configuration, err := suite.client.GetConfigurationJobs()
	suite.NoError(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationJobsFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetConfigurationJobs()
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationLanguages() {
	configuration, err := suite.client.GetConfigurationLanguages()
	suite.NoError(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationLanguagesFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetConfigurationLanguages()
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationPrimaryTranslations() {
	configuration, err := suite.client.GetConfigurationPrimaryTranslations()
	suite.NoError(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationPrimaryTranslationsFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetConfigurationPrimaryTranslations()
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationTimezones() {
	configuration, err := suite.client.GetConfigurationTimezones()
	suite.NoError(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationTimezonesFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetConfigurationTimezones()
	suite.Error(err)
}
