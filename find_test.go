package tmdb

const (
	theWalkingDeadID = "tt1520211"
	interstellarID   = "tt0816692"
	alPacinoID       = "nm0000199"
)

var options = map[string]string{
	"external_source": "imdb_id",
	"language":        "en-US",
}

func (suite *TMBDTestSuite) TestGetFindMovieByID() {
	results, err := suite.client.GetFindByID(interstellarID, options)
	suite.Nil(err)
	suite.NotNil(results)
}

func (suite *TMBDTestSuite) TestGetFindTVByID() {
	results, err := suite.client.GetFindByID(theWalkingDeadID, options)
	suite.Nil(err)
	suite.NotNil(results)
}

func (suite *TMBDTestSuite) TestGetFindPersonByID() {
	results, err := suite.client.GetFindByID(alPacinoID, options)
	suite.Nil(err)
	suite.NotNil(results)
}

func (suite *TMBDTestSuite) TestGetFindByIDFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetFindByID(theWalkingDeadID, options)
	suite.NotNil(err)
}
