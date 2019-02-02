package tmdb

const (
	americanMovieClassics = 174
	netflixID             = 213
)

func (suite *TMBDTestSuite) TestGetNetworkDetails() {
	network, err := suite.GetNetworkDetails(netflixID)
	suite.Nil(err)
	suite.NotNil(network)
}

func (suite *TMBDTestSuite) TestGetNetworkDetailsFail() {
	suite.APIKey = ""
	_, err := suite.GetNetworkDetails(netflixID)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetNetworkAlternativeNames() {
	network, err := suite.GetNetworkAlternativeNames(americanMovieClassics)
	suite.Nil(err)
	suite.NotNil(network)
}

func (suite *TMBDTestSuite) TestGetNetworkAlternativeNamesFail() {
	suite.APIKey = ""
	_, err := suite.GetNetworkAlternativeNames(americanMovieClassics)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetNetworkImages() {
	network, err := suite.GetNetworkImages(netflixID)
	suite.Nil(err)
	suite.NotNil(network)
}

func (suite *TMBDTestSuite) TestGetNetworkImagesFail() {
	suite.APIKey = ""
	_, err := suite.GetNetworkImages(netflixID)
	suite.NotNil(err)
}
