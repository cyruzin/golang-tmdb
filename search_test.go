package tmdb

import (
	"os"
	"testing"
)

const (
	defaultQuery = "a"
	warner       = "Warner Bros"
)

func (suite *TMBDTestSuite) TestGetSearchCompanies() {
	search, err := suite.client.GetSearchCompanies(warner, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchCompaniesFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetSearchCompanies(warner, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchCompaniesWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.client.GetSearchCompanies(warner, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchCollections() {
	search, err := suite.client.GetSearchCollections(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchCollectionsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetSearchCollections(defaultQuery, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchCollectionsWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.client.GetSearchCollections(defaultQuery, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchKeywords() {
	search, err := suite.client.GetSearchKeywords(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchKeywordsFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetSearchKeywords(defaultQuery, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchKeywordsWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.client.GetSearchKeywords(defaultQuery, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchMovies() {
	search, err := suite.client.GetSearchMovies(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchMoviesFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetSearchMovies(defaultQuery, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchMoviesWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.client.GetSearchMovies(defaultQuery, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchMulti() {
	search, err := suite.client.GetSearchMulti(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchMultiFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetSearchMulti(defaultQuery, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchMultiWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.client.GetSearchMulti(defaultQuery, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchPeople() {
	search, err := suite.client.GetSearchPeople(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchPeopleFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetSearchPeople(defaultQuery, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchPeopleWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.client.GetSearchPeople(defaultQuery, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchTVShow() {
	search, err := suite.client.GetSearchTVShow(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func (suite *TMBDTestSuite) TestGetSearchTVShowFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetSearchTVShow(defaultQuery, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetSearchTVShowWithOptions() {
	options := make(map[string]string)
	options["page"] = "1"
	search, err := suite.client.GetSearchTVShow(defaultQuery, options)
	suite.Nil(err)
	suite.NotNil(search.Page)
}

func BenchmarkGetSearchMulti(b *testing.B) {
	var tmdbClient Client
	tmdbClient.apiKey = os.Getenv("APIKey")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := tmdbClient.GetSearchMulti(defaultQuery, nil)
		if err != nil {
			b.Error(err)
		}
	}
}
