package tmdb

const (
	bumblebeeID   = 424783
	jackReacherID = 75780
	aquamanID     = 297802
)

func (suite *TMBDTestSuite) TestGetMovieDetails() {
	bumblebee, err := suite.client.GetMovieDetails(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal("Bumblebee", bumblebee.Title)
}

func (suite *TMBDTestSuite) TestGetMovieDetailsFail() {
	_, err := suite.client.GetMovieDetails(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jackreacher, err := suite.client.GetMovieDetails(jackReacherID, options)
	suite.Nil(err)
	suite.Equal("Jack Reacher: O Último Tiro", jackreacher.Title)
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitles() {
	bumblebee, err := suite.client.GetMovieAlternativeTitles(bumblebeeID, nil)
	suite.Nil(err)
	for _, v := range bumblebee.Titles {
		if v.Iso3166_1 == "US" {
			suite.NotNil(v.Title)
		}
	}
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitlesFail() {
	_, err := suite.client.GetMovieAlternativeTitles(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitlesWithOptions() {
	options := make(map[string]string)
	options["country"] = "RU"
	bumblebee, err := suite.client.GetMovieAlternativeTitles(bumblebeeID, options)
	suite.Nil(err)
	suite.NotNil(bumblebee.ID)
}

func (suite *TMBDTestSuite) TestGetMovieChanges() {
	bumblebee, err := suite.client.GetMovieChanges(bumblebeeID, nil)
	suite.Nil(err)
	for _, v := range bumblebee.Changes {
		for _, v := range v.Items {
			suite.NotNil(v.ID)
		}
	}
}

// The API isn't handling zero values for this end-point.
// TODO: Fix this test later.
func (suite *TMBDTestSuite) TestGetMovieChangesFail() {
	_, err := suite.client.GetMovieChanges(0, nil)
	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestGetMovieChangesWithOptions() {
	options := make(map[string]string)
	options["start_date"] = "2019-01-01"
	options["end_date"] = "2019-01-12"
	options["page"] = "1"
	bumblebee, err := suite.client.GetMovieChanges(bumblebeeID, options)
	suite.Nil(err)
	for _, v := range bumblebee.Changes {
		for _, v := range v.Items {
			suite.NotNil(v.ID)
		}
	}
}

func (suite *TMBDTestSuite) TestGetMovieCredits() {
	bumblebee, err := suite.client.GetMovieCredits(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal(int64(bumblebeeID), bumblebee.ID)
}

func (suite *TMBDTestSuite) TestGetMovieCreditsFail() {
	_, err := suite.client.GetMovieCredits(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieExternalIDs() {
	bumblebee, err := suite.client.GetMovieExternalIDs(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal("bumblebeemovie", bumblebee.FacebookID)
}

func (suite *TMBDTestSuite) TestGetMovieExternalIDsFail() {
	_, err := suite.client.GetMovieExternalIDs(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieImages() {
	bumblebee, err := suite.client.GetMovieImages(bumblebeeID, nil)
	suite.Nil(err)
	suite.NotNil(bumblebee.Backdrops[0].FilePath)
}

func (suite *TMBDTestSuite) TestGetMovieImagesFail() {
	_, err := suite.client.GetMovieImages(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	bumblebee, err := suite.client.GetMovieImages(bumblebeeID, options)
	suite.Nil(err)
	suite.Equal(int64(bumblebeeID), bumblebee.ID)
}

func (suite *TMBDTestSuite) TestGetMovieKeywords() {
	bumblebee, err := suite.client.GetMovieKeywords(bumblebeeID)
	suite.Nil(err)
	suite.Equal("technology", bumblebee.Keywords[0].Name)
}

func (suite *TMBDTestSuite) TestGetMovieKeywordsFail() {
	_, err := suite.client.GetMovieKeywords(0)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieReleaseDates() {
	bumblebee, err := suite.client.GetMovieReleaseDates(bumblebeeID)
	suite.Nil(err)
	suite.NotEmpty(bumblebee.Results[0].Iso3166_1)
}

func (suite *TMBDTestSuite) TestGetMovieReleaseDatesFail() {
	_, err := suite.client.GetMovieReleaseDates(0)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieAccountStatesFail() {
	_, err := suite.client.GetMovieAccountStates(0, nil)
	suite.Equal("Authentication failed: You do not have permissions to access the service.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieAccountStatesWithOptions() {
	suite.client.apiKey = ""
	options := make(map[string]string)
	options["session_id"] = "koQubnkaZ"
	_, err := suite.client.GetMovieAccountStates(jackReacherID, options)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieVideos() {
	bumblebee, err := suite.client.GetMovieVideos(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal("5b1638bc0e0a262df10146ed", bumblebee.Results[0].ID)
}

func (suite *TMBDTestSuite) TestGetMovieVideosFail() {
	_, err := suite.client.GetMovieVideos(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieVideosWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	bumblebee, err := suite.client.GetMovieVideos(bumblebeeID, options)
	suite.Nil(err)
	suite.Equal("5b1638bc0e0a262df10146ed", bumblebee.Results[0].ID)
}

func (suite *TMBDTestSuite) TestGetMovieTranslations() {
	bumblebee, err := suite.client.GetMovieTranslations(bumblebeeID, nil)
	suite.Nil(err)
	suite.NotEmpty(bumblebee.Translations[0].Iso3166_1)
}

func (suite *TMBDTestSuite) TestGetMovieTranslationsFail() {
	_, err := suite.client.GetMovieTranslations(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieTranslationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	bumblebee, err := suite.client.GetMovieTranslations(bumblebeeID, options)
	suite.Nil(err)
	suite.Equal(int64(bumblebeeID), bumblebee.ID)
}

func (suite *TMBDTestSuite) TestGetMovieRecommendations() {
	bumblebee, err := suite.client.GetMovieRecommendations(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), bumblebee.Page)
}

func (suite *TMBDTestSuite) TestGetMovieRecommendationsFail() {
	_, err := suite.client.GetMovieRecommendations(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieRecommendationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jackreacher, err := suite.client.GetMovieRecommendations(jackReacherID, options)
	suite.Nil(err)
	suite.Equal(int64(1), jackreacher.Page)
}

func (suite *TMBDTestSuite) TestGetMovieSimilar() {
	bumblebee, err := suite.client.GetMovieSimilar(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), bumblebee.Page)
}

func (suite *TMBDTestSuite) TestGetMovieSimilarFail() {
	_, err := suite.client.GetMovieSimilar(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieSimilarWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jackreacher, err := suite.client.GetMovieSimilar(jackReacherID, options)
	suite.Nil(err)
	suite.Equal(int64(1), jackreacher.Page)
}

func (suite *TMBDTestSuite) TestGetMovieReviews() {
	aquaman, err := suite.client.GetMovieReviews(aquamanID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), aquaman.Page)
}

func (suite *TMBDTestSuite) TestGetMovieReviewsFail() {
	_, err := suite.client.GetMovieReviews(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieReviewsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	aquaman, err := suite.client.GetMovieReviews(aquamanID, options)
	suite.Nil(err)
	suite.Equal(int64(1), aquaman.Page)
}

func (suite *TMBDTestSuite) TestGetMovieLists() {
	aquaman, err := suite.client.GetMovieLists(aquamanID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), aquaman.Page)
}

func (suite *TMBDTestSuite) TestGetMovieListsFail() {
	_, err := suite.client.GetMovieLists(0, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetMovieListsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	aquaman, err := suite.client.GetMovieLists(aquamanID, options)
	suite.Nil(err)
	suite.Equal(int64(1), aquaman.Page)
}

func (suite *TMBDTestSuite) TestGetMovieLatest() {
	movie, err := suite.client.GetMovieLatest(nil)
	suite.Nil(err)
	suite.NotNil(movie.VoteCount)
}

func (suite *TMBDTestSuite) TestGetMovieLatestFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetMovieLatest(nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieLatestWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movie, err := suite.client.GetMovieLatest(options)
	suite.Nil(err)
	suite.NotNil(movie.VoteCount)
}

func (suite *TMBDTestSuite) TestGetMovieNowPlaying() {
	movies, err := suite.client.GetMovieNowPlaying(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieNowPlayingFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetMovieNowPlaying(nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieNowPlayingWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.client.GetMovieNowPlaying(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMoviePopular() {
	movies, err := suite.client.GetMoviePopular(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMoviePopularFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetMoviePopular(nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMoviePopularWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.client.GetMoviePopular(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieTopRated() {
	movies, err := suite.client.GetMovieTopRated(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieTopRatedFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetMovieTopRated(nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieTopRatedWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.client.GetMovieTopRated(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieUpcoming() {
	movies, err := suite.client.GetMovieUpcoming(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieUpcomingFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetMovieUpcoming(nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieUpcomingWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.client.GetMovieUpcoming(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestPostMovieRating() {
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.PostMovieRating(aquamanID, 10, nil)
	suite.Nil(err)
	suite.NotNil(response.StatusCode)
}

func (suite *TMBDTestSuite) TestPostMovieRatingFail() {
	suite.client.sessionID = ""
	_, err := suite.client.PostMovieRating(aquamanID, 10, nil)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestDeleteMovieRating() {
	suite.client.SetSessionID(sessionID)
	response, err := suite.client.DeleteMovieRating(aquamanID, nil)
	suite.Nil(err)
	suite.Equal(13, response.StatusCode)
}

func (suite *TMBDTestSuite) TestDeleteMovieRatingFail() {
	suite.client.sessionID = ""
	_, err := suite.client.DeleteMovieRating(aquamanID, nil)
	suite.NotNil(err)
}
