package tmdb

const (
	americanMovieClassics = 174
	netflixID             = 213
)

func (suite *TMBDTestSuite) TestGetNetworkDetails() {
	network, err := suite.client.GetNetworkDetails(netflixID)
	suite.NoError(err)
	suite.NotNil(network)
}

func (suite *TMBDTestSuite) TestGetNetworkDetailsFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetNetworkDetails(netflixID)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetNetworkAlternativeNames() {
	network, err := suite.client.GetNetworkAlternativeNames(americanMovieClassics)
	suite.NoError(err)
	suite.NotNil(network)
}

func (suite *TMBDTestSuite) TestGetNetworkAlternativeNamesFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetNetworkAlternativeNames(americanMovieClassics)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestGetNetworkImages() {
	network, err := suite.client.GetNetworkImages(netflixID)
	suite.NoError(err)
	suite.NotNil(network)
}

func (suite *TMBDTestSuite) TestGetNetworkImagesFail() {
	suite.client.apiKey = "badkey"
	_, err := suite.client.GetNetworkImages(netflixID)
	suite.Error(err)
}
