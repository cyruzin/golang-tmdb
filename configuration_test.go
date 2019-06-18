package tmdb

func (suite *TMBDTestSuite) TestGetConfigurationAPI() {
	configuration, err := suite.client.GetConfigurationAPI()
	suite.Nil(err)
	suite.NotNil(configuration.Images.BackdropSizes)
}

func (suite *TMBDTestSuite) TestGetConfigurationAPIFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetConfigurationAPI()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationCountries() {
	configuration, err := suite.client.GetConfigurationCountries()
	suite.Nil(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationCountriesFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetConfigurationCountries()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationJobs() {
	configuration, err := suite.client.GetConfigurationJobs()
	suite.Nil(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationJobsFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetConfigurationJobs()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationLanguages() {
	configuration, err := suite.client.GetConfigurationLanguages()
	suite.Nil(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationLanguagesFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetConfigurationLanguages()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationPrimaryTranslations() {
	configuration, err := suite.client.GetConfigurationPrimaryTranslations()
	suite.Nil(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationPrimaryTranslationsFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetConfigurationPrimaryTranslations()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationTimezones() {
	configuration, err := suite.client.GetConfigurationTimezones()
	suite.Nil(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationTimezonesFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetConfigurationTimezones()
	suite.NotNil(err)
}
