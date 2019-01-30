package tmdb

const (
	defaultQuery = "a"
	warner       = "Warner Bros"
)

func (suite *TMBDTestSuite) TestGetSearchCompanies() {
	search, err := suite.GetSearchCompanies(warner, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchCompaniesFail() {
	suite.APIKey = ""
	_, err := suite.GetSearchCompanies(warner, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchCompaniesWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.GetSearchCompanies(warner, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchCollections() {
	search, err := suite.GetSearchCollections(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchCollectionsFail() {
	suite.APIKey = ""
	_, err := suite.GetSearchCollections(defaultQuery, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchCollectionsWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.GetSearchCollections(defaultQuery, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchKeywords() {
	search, err := suite.GetSearchKeywords(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchKeywordsFail() {
	suite.APIKey = ""
	_, err := suite.GetSearchKeywords(defaultQuery, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchKeywordsWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.GetSearchKeywords(defaultQuery, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchMovies() {
	search, err := suite.GetSearchMovies(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchMoviesFail() {
	suite.APIKey = ""
	_, err := suite.GetSearchMovies(defaultQuery, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchMoviesWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.GetSearchMovies(defaultQuery, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}
