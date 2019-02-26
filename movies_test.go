package tmdb

const (
	bumblebeeID   = 424783
	jackReacherID = 75780
	aquamanID     = 297802
)

func (suite *TMBDTestSuite) TestGetMovieDetails() {
	bumblebee, err := suite.GetMovieDetails(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal("Bumblebee", bumblebee.Title)
}

func (suite *TMBDTestSuite) TestGetMovieDetailsFail() {
	_, err := suite.GetMovieDetails(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jackreacher, err := suite.GetMovieDetails(jackReacherID, options)
	suite.Nil(err)
	suite.Equal("Jack Reacher: O Último Tiro", jackreacher.Title)
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitles() {
	bumblebee, err := suite.GetMovieAlternativeTitles(bumblebeeID, nil)
	suite.Nil(err)
	for _, v := range bumblebee.Titles {
		if v.Iso3166_1 == "US" {
			suite.NotNil(v.Title)
		}
	}
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitlesFail() {
	_, err := suite.GetMovieAlternativeTitles(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitlesWithOptions() {
	options := make(map[string]string)
	options["country"] = "RU"
	bumblebee, err := suite.GetMovieAlternativeTitles(bumblebeeID, options)
	suite.Nil(err)
	suite.Equal("RU", bumblebee.Titles[0].Iso3166_1)
}

func (suite *TMBDTestSuite) TestGetMovieChanges() {
	bumblebee, err := suite.GetMovieChanges(bumblebeeID, nil)
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
	_, err := suite.GetMovieChanges(0, nil)
	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestGetMovieChangesWithOptions() {
	options := make(map[string]string)
	options["start_date"] = "2019-01-01"
	options["end_date"] = "2019-01-12"
	options["page"] = "1"
	bumblebee, err := suite.GetMovieChanges(bumblebeeID, options)
	suite.Nil(err)
	for _, v := range bumblebee.Changes {
		for _, v := range v.Items {
			suite.NotNil(v.ID)
		}
	}
}

func (suite *TMBDTestSuite) TestGetMovieCredits() {
	bumblebee, err := suite.GetMovieCredits(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal(int64(bumblebeeID), bumblebee.ID)
}

func (suite *TMBDTestSuite) TestGetMovieCreditsFail() {
	_, err := suite.GetMovieCredits(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieExternalIDs() {
	bumblebee, err := suite.GetMovieExternalIDs(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal("bumblebeemovie", bumblebee.FacebookID)
}

func (suite *TMBDTestSuite) TestGetMovieExternalIDsFail() {
	_, err := suite.GetMovieExternalIDs(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieImages() {
	bumblebee, err := suite.GetMovieImages(bumblebeeID, nil)
	suite.Nil(err)
	suite.NotNil(bumblebee.Backdrops[0].FilePath)
}

func (suite *TMBDTestSuite) TestGetMovieImagesFail() {
	_, err := suite.GetMovieImages(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	bumblebee, err := suite.GetMovieImages(bumblebeeID, options)
	suite.Nil(err)
	suite.Equal(int64(bumblebeeID), bumblebee.ID)
}

func (suite *TMBDTestSuite) TestGetMovieKeywords() {
	bumblebee, err := suite.GetMovieKeywords(bumblebeeID)
	suite.Nil(err)
	suite.Equal("technology", bumblebee.Keywords[0].Name)
}

func (suite *TMBDTestSuite) TestGetMovieKeywordsFail() {
	_, err := suite.GetMovieKeywords(0)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieReleaseDates() {
	bumblebee, err := suite.GetMovieReleaseDates(bumblebeeID)
	suite.Nil(err)
	suite.NotNil(bumblebee.Results[0].Iso3166_1)
}

func (suite *TMBDTestSuite) TestGetMovieReleaseDatesFail() {
	_, err := suite.GetMovieReleaseDates(0)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieAccountStatesFail() {
	_, err := suite.GetMovieAccountStates(0, nil)
	suite.Equal("Authentication failed: You do not have permissions to access the service.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieAccountStatesWithOptions() {
	suite.Client.APIKey = ""
	options := make(map[string]string)
	options["session_id"] = "koQubnkaZ"
	_, err := suite.GetMovieAccountStates(jackReacherID, options)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieVideos() {
	bumblebee, err := suite.GetMovieVideos(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal("5b1638bc0e0a262df10146ed", bumblebee.Results[0].ID)
}

func (suite *TMBDTestSuite) TestGetMovieVideosFail() {
	_, err := suite.GetMovieVideos(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieVideosWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	bumblebee, err := suite.GetMovieVideos(bumblebeeID, options)
	suite.Nil(err)
	suite.Equal("5b1638bc0e0a262df10146ed", bumblebee.Results[0].ID)
}

func (suite *TMBDTestSuite) TestGetMovieTranslations() {
	bumblebee, err := suite.GetMovieTranslations(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal("DK", bumblebee.Translations[0].Iso3166_1)
}

func (suite *TMBDTestSuite) TestGetMovieTranslationsFail() {
	_, err := suite.GetMovieTranslations(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieTranslationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	bumblebee, err := suite.GetMovieTranslations(bumblebeeID, options)
	suite.Nil(err)
	suite.Equal(int64(bumblebeeID), bumblebee.ID)
}

func (suite *TMBDTestSuite) TestGetMovieRecommendations() {
	bumblebee, err := suite.GetMovieRecommendations(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), bumblebee.Page)
}

func (suite *TMBDTestSuite) TestGetMovieRecommendationsFail() {
	_, err := suite.GetMovieRecommendations(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieRecommendationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jackreacher, err := suite.GetMovieRecommendations(jackReacherID, options)
	suite.Nil(err)
	suite.Equal(int64(1), jackreacher.Page)
}

func (suite *TMBDTestSuite) TestGetMovieSimilar() {
	bumblebee, err := suite.GetMovieSimilar(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), bumblebee.Page)
}

func (suite *TMBDTestSuite) TestGetMovieSimilarFail() {
	_, err := suite.GetMovieSimilar(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieSimilarWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jackreacher, err := suite.GetMovieSimilar(jackReacherID, options)
	suite.Nil(err)
	suite.Equal(int64(1), jackreacher.Page)
}

func (suite *TMBDTestSuite) TestGetMovieReviews() {
	aquaman, err := suite.GetMovieReviews(aquamanID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), aquaman.Page)
}

func (suite *TMBDTestSuite) TestGetMovieReviewsFail() {
	_, err := suite.GetMovieReviews(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieReviewsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	aquaman, err := suite.GetMovieReviews(aquamanID, options)
	suite.Nil(err)
	suite.Equal(int64(1), aquaman.Page)
}

func (suite *TMBDTestSuite) TestGetMovieLists() {
	aquaman, err := suite.GetMovieLists(aquamanID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), aquaman.Page)
}

func (suite *TMBDTestSuite) TestGetMovieListsFail() {
	_, err := suite.GetMovieLists(0, nil)
	suite.Equal("Internal error: Something went wrong, contact TMDb.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieListsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	aquaman, err := suite.GetMovieLists(aquamanID, options)
	suite.Nil(err)
	suite.Equal(int64(1), aquaman.Page)
}

func (suite *TMBDTestSuite) TestGetMovieLatest() {
	movie, err := suite.GetMovieLatest(nil)
	suite.Nil(err)
	suite.NotNil(movie.VoteCount)
}

func (suite *TMBDTestSuite) TestGetMovieLatestFail() {
	suite.Client.APIKey = ""
	_, err := suite.GetMovieLatest(nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieLatestWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movie, err := suite.GetMovieLatest(options)
	suite.Nil(err)
	suite.NotNil(movie.VoteCount)
}

func (suite *TMBDTestSuite) TestGetMovieNowPlaying() {
	movies, err := suite.GetMovieNowPlaying(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieNowPlayingFail() {
	suite.Client.APIKey = ""
	_, err := suite.GetMovieNowPlaying(nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieNowPlayingWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.GetMovieNowPlaying(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMoviePopular() {
	movies, err := suite.GetMoviePopular(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMoviePopularFail() {
	suite.Client.APIKey = ""
	_, err := suite.GetMoviePopular(nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMoviePopularWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.GetMoviePopular(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieTopRated() {
	movies, err := suite.GetMovieTopRated(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieTopRatedFail() {
	suite.Client.APIKey = ""
	_, err := suite.GetMovieTopRated(nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieTopRatedWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.GetMovieTopRated(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieUpcoming() {
	movies, err := suite.GetMovieUpcoming(nil)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}

func (suite *TMBDTestSuite) TestGetMovieUpcomingFail() {
	suite.Client.APIKey = ""
	_, err := suite.GetMovieUpcoming(nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieUpcomingWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	movies, err := suite.GetMovieUpcoming(options)
	suite.Nil(err)
	suite.Equal(int64(1), movies.Page)
}
