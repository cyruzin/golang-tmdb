package tmdb

func (suite *TMBDTestSuite) TestGetConfigurationAPI() {
	configuration, err := suite.GetConfigurationAPI()
	suite.Nil(err)
	suite.NotNil(configuration.Images.BackdropSizes)
}

func (suite *TMBDTestSuite) TestGetConfigurationAPIFail() {
	suite.APIKey = ""
	_, err := suite.GetConfigurationAPI()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationCountries() {
	configuration, err := suite.GetConfigurationCountries()
	suite.Nil(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationCountriesFail() {
	suite.APIKey = ""
	_, err := suite.GetConfigurationCountries()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationJobs() {
	configuration, err := suite.GetConfigurationJobs()
	suite.Nil(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationJobsFail() {
	suite.APIKey = ""
	_, err := suite.GetConfigurationJobs()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationLanguages() {
	configuration, err := suite.GetConfigurationLanguages()
	suite.Nil(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationLanguagesFail() {
	suite.APIKey = ""
	_, err := suite.GetConfigurationLanguages()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationPrimaryTranslations() {
	configuration, err := suite.GetConfigurationPrimaryTranslations()
	suite.Nil(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationPrimaryTranslationsFail() {
	suite.APIKey = ""
	_, err := suite.GetConfigurationPrimaryTranslations()
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetConfigurationTimezones() {
	configuration, err := suite.GetConfigurationTimezones()
	suite.Nil(err)
	suite.NotNil(configuration)
}

func (suite *TMBDTestSuite) TestGetConfigurationTimezonesFail() {
	suite.APIKey = ""
	_, err := suite.GetConfigurationTimezones()
	suite.NotNil(err)
}
