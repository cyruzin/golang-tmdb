package tmdb

const (
	americanMovieClassics = 174
	netflixID             = 213
)

func (suite *TMBDTestSuite) TestGetNetworkDetails() {
	network, err := suite.client.GetNetworkDetails(netflixID)
	suite.Nil(err)
	suite.NotNil(network)
}

func (suite *TMBDTestSuite) TestGetNetworkDetailsFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetNetworkDetails(netflixID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetNetworkAlternativeNames() {
	network, err := suite.client.GetNetworkAlternativeNames(americanMovieClassics)
	suite.Nil(err)
	suite.NotNil(network)
}

func (suite *TMBDTestSuite) TestGetNetworkAlternativeNamesFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetNetworkAlternativeNames(americanMovieClassics)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetNetworkImages() {
	network, err := suite.client.GetNetworkImages(netflixID)
	suite.Nil(err)
	suite.NotNil(network)
}

func (suite *TMBDTestSuite) TestGetNetworkImagesFail() {
	suite.client.APIKey = ""
	_, err := suite.client.GetNetworkImages(netflixID)
	suite.NotNil(err)
}
