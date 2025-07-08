package tmdb

const (
	vikingsID = 44217
	flashID   = 60735
	gotID     = 1399
)

const (
	gotSeasonID  = 3624
	gotEpisodeID = 63056
)

func (suite *TMBDTestSuite) TestGetTVDetails() {
	vikings, err := suite.client.GetTVDetails(vikingsID, nil)
	suite.Nil(err)
	suite.Equal("Vikings", vikings.Name)
}

func (suite *TMBDTestSuite) TestGetTVDetailsFail() {
	_, err := suite.client.GetTVDetails(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	vikings, err := suite.client.GetTVDetails(vikingsID, options)
	suite.Nil(err)
	suite.Equal("Vikings", vikings.Name)
}

func (suite *TMBDTestSuite) TestGetTVAccountStates() {
	options := make(map[string]string)
	options["session_id"] = sessionID
	_, err := suite.client.GetTVAccountStates(vikingsID, options)
	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestGetTVAccountStatesFail() {
	suite.client.apiKey = ""
	options := make(map[string]string)
	options["session_id"] = "koQubnkaZ"
	_, err := suite.client.GetTVAccountStates(vikingsID, options)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVAggregateCredits() {
	vikings, err := suite.client.GetTVAggregateCredits(vikingsID, nil)
	suite.Nil(err)
	suite.Equal(int64(vikingsID), vikings.ID)
}

func (suite *TMBDTestSuite) TestGetTVAggregateFail() {
	_, err := suite.client.GetTVAggregateCredits(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVAlternativeTitles() {
	flash, err := suite.client.GetTVAlternativeTitles(flashID, nil)
	suite.Nil(err)
	for _, v := range flash.Results {
		if v.Iso3166_1 == "GR" {
			suite.Equal("Ο Φλας", v.Title)
		}
	}
}

func (suite *TMBDTestSuite) TestGetTVAlternativeTitlesFail() {
	_, err := suite.client.GetTVAlternativeTitles(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVAlternativeTitlesWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	flash, err := suite.client.GetTVAlternativeTitles(flashID, options)
	suite.Nil(err)
	suite.NotNil(flash.Results)
}

func (suite *TMBDTestSuite) TestGetTVChanges() {
	changes, err := suite.client.GetTVChanges(vikingsID, nil)
	suite.Nil(err)
	suite.NotNil(changes)
}

func (suite *TMBDTestSuite) TestGetTVChangesFail() {
	_, err := suite.client.GetTVChanges(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVChangesWithOptions() {
	options := make(map[string]string)
	options["start_date"] = "2019-01-01"
	options["end_date"] = "2019-01-12"
	options["page"] = "1"
	flash, err := suite.client.GetTVChanges(flashID, options)
	suite.Nil(err)
	for _, v := range flash.Changes {
		for _, v := range v.Items {
			suite.NotNil(v.ID)
		}
	}
}

func (suite *TMBDTestSuite) TestGetTVContentRatings() {
	vikings, err := suite.client.GetTVContentRatings(vikingsID, nil)
	suite.Nil(err)
	suite.NotNil(vikings.Results[0].Iso3166_1)
}

func (suite *TMBDTestSuite) TestGetTVContentRatingsFail() {
	_, err := suite.client.GetTVContentRatings(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVContentRatingsWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	vikings, err := suite.client.GetTVContentRatings(vikingsID, options)
	suite.Nil(err)
	suite.NotNil(vikings.Results[0].Iso3166_1)
}

func (suite *TMBDTestSuite) TestGetTVCredits() {
	vikings, err := suite.client.GetTVCredits(vikingsID, nil)
	suite.Nil(err)
	suite.Equal(int64(vikingsID), vikings.ID)
}

func (suite *TMBDTestSuite) TestGetTVCreditsFail() {
	_, err := suite.client.GetTVCredits(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeGroups() {
	vikings, err := suite.client.GetTVEpisodeGroups(vikingsID, nil)
	suite.Nil(err)
	suite.Equal(int64(vikingsID), vikings.ID)
}

func (suite *TMBDTestSuite) TestGetTVEpisodeGroupsFail() {
	_, err := suite.client.GetTVEpisodeGroups(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVEpisodeGroupsWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	vikings, err := suite.client.GetTVEpisodeGroups(vikingsID, options)
	suite.Nil(err)
	suite.Equal(int64(vikingsID), vikings.ID)
}

func (suite *TMBDTestSuite) TestGetTVExternalIDs() {
	flash, err := suite.client.GetTVExternalIDs(flashID, nil)
	suite.Nil(err)
	suite.Equal("CWTheFlash", flash.FacebookID)
}

func (suite *TMBDTestSuite) TestGetTVExternalIDsFail() {
	_, err := suite.client.GetTVExternalIDs(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVImages() {
	flash, err := suite.client.GetTVImages(flashID, nil)
	suite.Nil(err)
	suite.NotNil(flash.Backdrops[0].FilePath)
}

func (suite *TMBDTestSuite) TestGetTVImagesFail() {
	_, err := suite.client.GetTVImages(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVImagesWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	flash, err := suite.client.GetTVImages(flashID, options)
	suite.Nil(err)
	suite.Equal(int64(flashID), flash.ID)
}

func (suite *TMBDTestSuite) TestGetTVKeywords() {
	flash, err := suite.client.GetTVKeywords(flashID)
	suite.Nil(err)
	suite.NotEmpty(flash.Results[0].Name)
}

func (suite *TMBDTestSuite) TestGetTVKeywordsFail() {
	_, err := suite.client.GetTVKeywords(0)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVRecommendations() {
	flash, err := suite.client.GetTVRecommendations(flashID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), flash.Page)
}

func (suite *TMBDTestSuite) TestGetTVRecommendationsFail() {
	_, err := suite.client.GetTVRecommendations(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVRecommendationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	flash, err := suite.client.GetTVRecommendations(flashID, options)
	suite.Nil(err)
	suite.Equal(int64(1), flash.Page)
}

func (suite *TMBDTestSuite) TestGetTVScreenedTheatrically() {
	flash, err := suite.client.GetTVScreenedTheatrically(flashID)
	suite.Nil(err)
	suite.Equal(int64(flashID), flash.ID)
}

func (suite *TMBDTestSuite) TestGetTVScreenedTheatricallyFail() {
	_, err := suite.client.GetTVScreenedTheatrically(0)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVReviews() {
	flash, err := suite.client.GetTVReviews(flashID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), flash.Page)
}

func (suite *TMBDTestSuite) TestGetTVReviewsFail() {
	_, err := suite.client.GetTVReviews(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVReviewsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	flash, err := suite.client.GetTVReviews(flashID, options)
	suite.Nil(err)
	suite.Equal(int64(1), flash.Page)
}

func (suite *TMBDTestSuite) TestGetTVWatchProviders() {
	gameOfThrones, err := suite.client.GetTVWatchProviders(gotID, nil)
	suite.Nil(err)
	suite.Equal("https://www.themoviedb.org/tv/1399-game-of-thrones/watch?locale=AR", gameOfThrones.Results["AR"].Link)
}

func (suite *TMBDTestSuite) TestGetTVWatchProvidersFail() {
	_, err := suite.client.GetTVWatchProviders(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVSimilar() {
	flashID, err := suite.client.GetTVSimilar(flashID, nil)
	suite.Nil(err)
	suite.Equal(int64(1), flashID.Page)
}

func (suite *TMBDTestSuite) TestGetTVSimilarFail() {
	_, err := suite.client.GetTVSimilar(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVSimilarWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	flash, err := suite.client.GetTVSimilar(flashID, options)
	suite.Nil(err)
	suite.Equal(int64(1), flash.Page)
}

func (suite *TMBDTestSuite) TestGetTVTranslations() {
	flash, err := suite.client.GetTVTranslations(flashID, nil)
	suite.Nil(err)
	suite.NotEmpty(flash.Translations[0].Iso3166_1)
}

func (suite *TMBDTestSuite) TestGetTVTranslationsFail() {
	_, err := suite.client.GetTVTranslations(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVTranslationsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	flash, err := suite.client.GetTVTranslations(flashID, options)
	suite.Nil(err)
	suite.Equal(int64(flashID), flash.ID)
}

func (suite *TMBDTestSuite) TestGetTVVideos() {
	flash, err := suite.client.GetTVVideos(flashID, nil)
	suite.Nil(err)
	suite.NotNil(flash.Results[0].ID)
}

func (suite *TMBDTestSuite) TestGetTVVideosFail() {
	_, err := suite.client.GetTVVideos(0, nil)
	suite.Equal("code: 6 | success: false | message: Invalid id: The pre-requisite id is invalid or not found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVVideosWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	flash, err := suite.client.GetTVVideos(flashID, options)
	suite.Nil(err)
	suite.NotNil(flash.Results[0].ID)
}

func (suite *TMBDTestSuite) TestGetTVLatest() {
	tv, err := suite.client.GetTVLatest(nil)
	suite.Nil(err)
	suite.NotNil(tv.VoteCount)
}

func (suite *TMBDTestSuite) TestGetTVLatestFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTVLatest(nil)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVLatestWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.client.GetTVLatest(options)
	suite.Nil(err)
	suite.NotNil(tv.VoteCount)
}

func (suite *TMBDTestSuite) TestGetTVAiringToday() {
	tv, err := suite.client.GetTVAiringToday(nil)
	suite.Nil(err)
	suite.NotNil(tv.TotalPages)
}

func (suite *TMBDTestSuite) TestGetTVAiringTodayFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTVAiringToday(nil)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVAiringTodayWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.client.GetTVAiringToday(options)
	suite.Nil(err)
	suite.NotNil(tv.TotalPages)
}

func (suite *TMBDTestSuite) TestGetTVOnTheAir() {
	tv, err := suite.client.GetTVOnTheAir(nil)
	suite.Nil(err)
	suite.NotNil(tv.TotalPages)
}

func (suite *TMBDTestSuite) TestGetTVOnTheAirFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTVOnTheAir(nil)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVOnTheAirWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.client.GetTVOnTheAir(options)
	suite.Nil(err)
	suite.NotNil(tv.TotalPages)
}

func (suite *TMBDTestSuite) TestGetTVPopular() {
	tv, err := suite.client.GetTVPopular(nil)
	suite.Nil(err)
	suite.NotNil(tv.TotalPages)
}

func (suite *TMBDTestSuite) TestGetTVPopularFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTVPopular(nil)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVPopularWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.client.GetTVPopular(options)
	suite.Nil(err)
	suite.NotNil(tv.TotalPages)
}

func (suite *TMBDTestSuite) TestGetTVTopRated() {
	tv, err := suite.client.GetTVTopRated(nil)
	suite.Nil(err)
	suite.NotNil(tv.TotalPages)
}

func (suite *TMBDTestSuite) TestGetTVTopRatedFail() {
	suite.client.apiKey = ""
	_, err := suite.client.GetTVTopRated(nil)
	suite.Equal("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
}

func (suite *TMBDTestSuite) TestGetTVTopRatedWithOptions() {
	options := make(map[string]string)
	options["language"] = "en-US"
	tv, err := suite.client.GetTVTopRated(options)
	suite.Nil(err)
	suite.NotNil(tv.TotalPages)
}

func (suite *TMBDTestSuite) TestPostTVShowRating() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	response, err := suite.client.PostTVShowRating(vikingsID, 10, nil)
	suite.Nil(err)
	suite.NotNil(response.StatusCode)
}

func (suite *TMBDTestSuite) TestPostTVShowRatingFail() {
	_, err := suite.client.PostTVShowRating(0, 10, nil)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestDeleteTVShowRating() {
	err := suite.client.SetSessionID(sessionID)
	suite.Nil(err)
	response, err := suite.client.DeleteTVShowRating(vikingsID, nil)
	suite.Nil(err)
	suite.Equal(13, response.StatusCode)
}

func (suite *TMBDTestSuite) TestDeleteTVShowRatingFail() {
	suite.client.sessionID = ""
	_, err := suite.client.DeleteTVShowRating(vikingsID, nil)
	suite.NotNil(err)
}
