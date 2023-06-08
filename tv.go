package tmdb

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

// TVDetails type is a struct for details JSON response.
type TVDetails struct {
	BackdropPath string `json:"backdrop_path"`
	CreatedBy    []struct {
		ID          int64  `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"created_by"`
	EpisodeRunTime []int  `json:"episode_run_time"`
	FirstAirDate   string `json:"first_air_date"`
	Genres         []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage         string   `json:"homepage"`
	ID               int64    `json:"id"`
	InProduction     bool     `json:"in_production"`
	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int64   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int64   `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float32 `json:"vote_average"`
		VoteCount      int64   `json:"vote_count"`
	} `json:"last_episode_to_air"`
	Name             string `json:"name"`
	NextEpisodeToAir struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int64   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int64   `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float32 `json:"vote_average"`
		VoteCount      int64   `json:"vote_count"`
	} `json:"next_episode_to_air"`
	Networks []struct {
		Name          string `json:"name"`
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes    int      `json:"number_of_episodes"`
	NumberOfSeasons     int      `json:"number_of_seasons"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalName        string   `json:"original_name"`
	Overview            string   `json:"overview"`
	Popularity          float32  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		Name          string `json:"name"`
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Name      string `json:"name"`
	} `json:"production_countries"`
	Seasons []struct {
		AirDate      string `json:"air_date"`
		EpisodeCount int    `json:"episode_count"`
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		Overview     string `json:"overview"`
		PosterPath   string `json:"poster_path"`
		SeasonNumber int    `json:"season_number"`
	} `json:"seasons"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Type        string  `json:"type"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
	*TVAggregateCreditsAppend
	*TVAlternativeTitlesAppend
	*TVChangesAppend
	*TVContentRatingsAppend
	*TVCreditsAppend
	*TVEpisodeGroupsAppend
	*TVExternalIDsAppend
	*TVImagesAppend
	*TVKeywordsAppend
	*TVRecommendationsAppend
	*TVReviewsAppend
	*TVScreenedTheatricallyAppend
	*TVSimilarAppend
	*TVTranslationsAppend
	*TVVideosAppend
	*TVWatchProvidersAppend
}

// TVAggregateCreditsAppend type is a struct
// for aggregate credits in append to response.
type TVAggregateCreditsAppend struct {
	AggregateCredits *TVAggregateCredits `json:"aggregate_credits,omitempty"`
}

// TVAlternativeTitlesAppend type is a struct
// for alternative titles in append to response.
type TVAlternativeTitlesAppend struct {
	AlternativeTitles *TVAlternativeTitles `json:"alternative_titles,omitempty"`
}

// TVChangesAppend type is a struct for changes in append to response.
type TVChangesAppend struct {
	Changes *TVChanges `json:"changes,omitempty"`
}

// TVContentRatingsAppend type is a struct for
// content ratings in append to response.
type TVContentRatingsAppend struct {
	ContentRatings *TVContentRatings `json:"content_ratings,omitempty"`
}

// TVCreditsAppend type is a struct for credits in append to response.
type TVCreditsAppend struct {
	Credits struct {
		*TVCredits
	} `json:"credits,omitempty"`
}

// TVEpisodeGroupsAppend type is a struct for
// episode groups in append to response.
type TVEpisodeGroupsAppend struct {
	EpisodeGroups *TVEpisodeGroups `json:"episode_groups,omitempty"`
}

// TVExternalIDsAppend type is a short for
// external ids in append to response.
type TVExternalIDsAppend struct {
	*TVExternalIDs `json:"external_ids,omitempty"`
}

// TVImagesAppend type is a struct for images in append to response.
type TVImagesAppend struct {
	Images *TVImages `json:"images,omitempty"`
}

// TVKeywordsAppend type is a struct for keywords in append to response.
type TVKeywordsAppend struct {
	Keywords struct {
		*TVKeywords
	} `json:"keywords,omitempty"`
}

// TVRecommendationsAppend type is a struct
// for recommendations in append to response.
type TVRecommendationsAppend struct {
	Recommendations *TVRecommendations `json:"recommendations,omitempty"`
}

// TVReviewsAppend type is a struct for reviews in append to response.
type TVReviewsAppend struct {
	Reviews struct {
		*TVReviews
	} `json:"reviews,omitempty"`
}

// TVScreenedTheatricallyAppend type is a struct
// for screened theatrically in append to response.
type TVScreenedTheatricallyAppend struct {
	ScreenedTheatrically *TVScreenedTheatrically `json:"screened_theatrically,omitempty"`
}

// TVSimilarAppend type is a struct for
// similar tv shows in append to response.
type TVSimilarAppend struct {
	Similar *TVSimilar `json:"similar,omitempty"`
}

// TVTranslationsAppend type is a struct
// for translations in append to response.
type TVTranslationsAppend struct {
	Translations *TVTranslations `json:"translations,omitempty"`
}

// TVVideosAppend type is a struct for videos in append to response.
type TVVideosAppend struct {
	Videos struct {
		*TVVideos
	} `json:"videos,omitempty"`
}

// TVWatchProvidersAppend type is a struct for
// watch/providers in append to response.
type TVWatchProvidersAppend struct {
	WatchProviders *TVWatchProviders `json:"watch/providers,omitempty"`
}

// GetTVDetails get the primary TV show details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/tv/get-tv-details
func (c *Client) GetTVDetails(id int64, urlOptions map[string]string) (*TVDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvDetails := TVDetails{}
	if err := c.get(tmdbURL, &tvDetails); err != nil {
		return nil, err
	}
	return &tvDetails, nil
}

// TVAccountStates type is a struct for account states JSON response.
type TVAccountStates struct {
	ID        int64               `json:"id"`
	Favorite  bool                `json:"favorite"`
	Rated     jsoniter.RawMessage `json:"rated"`
	Watchlist bool                `json:"watchlist"`
}

// GetTVAccountStates grab the following account states for a session:
//
// TV show rating.
//
// If it belongs to your watchlist.
//
// If it belongs to your favourite list.
//
// https://developers.themoviedb.org/3/tv/get-tv-account-states
func (c *Client) GetTVAccountStates(id int64, urlOptions map[string]string) (*TVAccountStates, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/account_states?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvAccountStates := TVAccountStates{}
	if err := c.get(tmdbURL, &tvAccountStates); err != nil {
		return nil, err
	}
	return &tvAccountStates, nil
}

// TVAggregateCredits type is a struct for aggregate credits JSON response.
type TVAggregateCredits struct {
	ID   int64 `json:"id,omitempty"`
	Cast []struct {
		ID                 int64   `json:"id"`
		Adult              bool    `json:"adult"`
		Gender             int     `json:"gender"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		Order              int     `json:"order"`
		OriginalName       string  `json:"original_name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		Roles              []struct {
			Character    string `json:"character"`
			CreditID     string `json:"credit_id"`
			EpisodeCount int    `json:"episode_count"`
		} `json:"roles"`
		TotalEpisodeCount int `json:"total_episode_count"`
	} `json:"cast"`
	Crew []struct {
		ID         int64  `json:"id"`
		Adult      bool   `json:"adult"`
		Department string `json:"department"`
		Gender     int    `json:"gender"`
		Jobs       []struct {
			CreditID     string `json:"credit_id"`
			EpisodeCount int    `json:"episode_count"`
			Job          string `json:"job"`
		} `json:"jobs"`
		TotalEpisodeCount  int     `json:"total_episode_count"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"crew"`
}

// GetTVAggregateCredits get the aggregate credits (cast and crew) that have been added to a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-aggregate-credits
func (c *Client) GetTVAggregateCredits(id int64, urlOptions map[string]string) (*TVAggregateCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/aggregate_credits?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvAggregateCredits := TVAggregateCredits{}
	if err := c.get(tmdbURL, &tvAggregateCredits); err != nil {
		return nil, err
	}
	return &tvAggregateCredits, nil
}

// TVAlternativeTitles type is a struct for alternative titles JSON response.
type TVAlternativeTitles struct {
	ID int `json:"id,omitempty"`
	*TVAlternativeTitlesResults
}

// GetTVAlternativeTitles get all of the alternative titles for a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-alternative-titles
func (c *Client) GetTVAlternativeTitles(id int64, urlOptions map[string]string) (*TVAlternativeTitles, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/alternative_titles?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvAlternativeTitles := TVAlternativeTitles{}
	if err := c.get(tmdbURL, &tvAlternativeTitles); err != nil {
		return nil, err
	}
	return &tvAlternativeTitles, nil
}

// TVChanges type is a struct for changes JSON response.
type TVChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID            string              `json:"id"`
			Action        string              `json:"action"`
			Time          string              `json:"time"`
			Iso639_1      string              `json:"iso_639_1"`
			Iso3166_1     string              `json:"iso_3166_1"`
			Value         jsoniter.RawMessage `json:"value"`
			OriginalValue jsoniter.RawMessage `json:"original_value"`
		} `json:"items"`
	} `json:"changes"`
}

// GetTVChanges get the changes for a TV show.
//
// By default only the last 24 hours are returned.
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// TV show changes are different than movie changes in that there
// are some edits on seasons and episodes that will create a change
// entry at the show level. These can be found under the season
// and episode keys. These keys will contain a series_id and episode_id.
// You can use the and methods to look these up individually.
//
// https://developers.themoviedb.org/3/tv/get-tv-changes
func (c *Client) GetTVChanges(id int64, urlOptions map[string]string) (*TVChanges, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/changes?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tVChanges := TVChanges{}
	if err := c.get(tmdbURL, &tVChanges); err != nil {
		return nil, err
	}
	return &tVChanges, nil
}

// TVContentRatings type is a struct for content ratings JSON response.
type TVContentRatings struct {
	*TVContentRatingsResults
	ID int64 `json:"id,omitempty"`
}

// GetTVContentRatings get the list of content ratings (certifications) that have been added to a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-content-ratings
func (c *Client) GetTVContentRatings(id int64, urlOptions map[string]string) (*TVContentRatings, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/content_ratings?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvContentRatings := TVContentRatings{}
	if err := c.get(tmdbURL, &tvContentRatings); err != nil {
		return nil, err
	}
	return &tvContentRatings, nil
}

// TVCredits type is a struct for credits JSON response.
type TVCredits struct {
	ID   int64 `json:"id,omitempty"`
	Cast []struct {
		Character          string  `json:"character"`
		CreditID           string  `json:"credit_id"`
		Gender             int     `json:"gender"`
		ID                 int64   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		Order              int     `json:"order"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"cast"`
	Crew []struct {
		CreditID           string  `json:"credit_id"`
		Department         string  `json:"department"`
		Gender             int     `json:"gender"`
		ID                 int64   `json:"id"`
		Job                string  `json:"job"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"crew"`
}

// GetTVCredits get the credits (cast and crew) that have been added to a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-credits
func (c *Client) GetTVCredits(id int64, urlOptions map[string]string) (*TVCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/credits?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvCredits := TVCredits{}
	if err := c.get(tmdbURL, &tvCredits); err != nil {
		return nil, err
	}
	return &tvCredits, nil
}

// TVEpisodeGroups type is a struct for episode groups JSON response.
type TVEpisodeGroups struct {
	*TVEpisodeGroupsResults
	ID int64 `json:"id,omitempty"`
}

// GetTVEpisodeGroups get all of the episode groups that have been created for a TV show.
//
// With a group ID you can call the get TV episode group details method.
//
// https://developers.themoviedb.org/3/tv/get-tv-episode-groups
func (c *Client) GetTVEpisodeGroups(id int64, urlOptions map[string]string) (*TVEpisodeGroups, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/episode_groups?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tVEpisodeGroups := TVEpisodeGroups{}
	if err := c.get(tmdbURL, &tVEpisodeGroups); err != nil {
		return nil, err
	}
	return &tVEpisodeGroups, nil
}

// TVExternalIDs type is a struct for external ids JSON response.
type TVExternalIDs struct {
	IMDbID      string `json:"imdb_id"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
	ID          int64  `json:"id,omitempty"`
}

// GetTVExternalIDs get the external ids for a TV show.
//
// We currently support the following external sources.
//
// Media Databases: IMDb ID, TVDB ID, Freebase MID*, Freebase ID* TVRage ID*.
//
// Social IDs: Facebook, Instagram and Twitter.
//
// *Defunct or no longer available as a service.
//
// https://developers.themoviedb.org/3/tv/get-tv-external-ids
func (c *Client) GetTVExternalIDs(id int64, urlOptions map[string]string) (*TVExternalIDs, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/external_ids?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvExternalIDs := TVExternalIDs{}
	if err := c.get(tmdbURL, &tvExternalIDs); err != nil {
		return nil, err
	}
	return &tvExternalIDs, nil
}

// TVImages type is a struct for images JSON response.
type TVImages struct {
	ID        int64 `json:"id,omitempty"`
	Backdrops []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		Iso639_1    string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int64   `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"backdrops"`
	Posters []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		Iso639_1    string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int64   `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"posters"`
}

// GetTVImages get the images that belong to a TV show.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter. This should be a comma
// separated value like so: include_image_language=en,null.
//
// https://developers.themoviedb.org/3/tv/get-tv-images
func (c *Client) GetTVImages(id int64, urlOptions map[string]string) (*TVImages, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvImages := TVImages{}
	if err := c.get(tmdbURL, &tvImages); err != nil {
		return nil, err
	}
	return &tvImages, nil
}

// TVKeywords type is a struct for keywords JSON response.
type TVKeywords struct {
	ID int64 `json:"id,omitempty"`
	*TVKeywordsResults
}

// GetTVKeywords get the keywords that have been added to a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-keywords
func (c *Client) GetTVKeywords(id int64) (*TVKeywords, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/keywords?api_key=%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
	)
	tvKeywords := TVKeywords{}
	if err := c.get(tmdbURL, &tvKeywords); err != nil {
		return nil, err
	}
	return &tvKeywords, nil
}

// TVRecommendations type is a struct for recommendations JSON response.
type TVRecommendations struct {
	Page int64 `json:"page"`
	*TVRecommendationsResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetTVRecommendations get the list of TV show recommendations for this item.
//
// https://developers.themoviedb.org/3/tv/get-tv-recommendations
func (c *Client) GetTVRecommendations(id int64, urlOptions map[string]string) (*TVRecommendations, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/recommendations?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvRecommendations := TVRecommendations{}
	if err := c.get(tmdbURL, &tvRecommendations); err != nil {
		return nil, err
	}
	return &tvRecommendations, nil
}

// TVReviews type is a struct for reviews JSON response.
type TVReviews struct {
	ID   int64 `json:"id,omitempty"`
	Page int64 `json:"page"`
	*TVReviewsResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetTVReviews get the reviews for a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-reviews
func (c *Client) GetTVReviews(id int64, urlOptions map[string]string) (*TVReviews, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/reviews?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvReviews := TVReviews{}
	if err := c.get(tmdbURL, &tvReviews); err != nil {
		return nil, err
	}
	return &tvReviews, nil
}

// TVScreenedTheatrically type is a struct for screened theatrically JSON response.
type TVScreenedTheatrically struct {
	ID int64 `json:"id,omitempty"`
	*TVScreenedTheatricallyResults
}

// GetTVScreenedTheatrically get a list of seasons or episodes that
// have been screened in a film festival or theatre.
//
// https://developers.themoviedb.org/3/tv/get-screened-theatrically
func (c *Client) GetTVScreenedTheatrically(id int64) (*TVScreenedTheatrically, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/screened_theatrically?api_key=%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
	)
	tvScreenedTheatrically := TVScreenedTheatrically{}
	if err := c.get(tmdbURL, &tvScreenedTheatrically); err != nil {
		return nil, err
	}
	return &tvScreenedTheatrically, nil
}

// TVSimilar type is a struct for similar tv shows JSON response.
type TVSimilar struct {
	*TVRecommendations
}

// GetTVSimilar a list of similar TV shows.
// These items are assembled by looking at keywords and genres.
//
// https://developers.themoviedb.org/3/tv/get-similar-tv-shows
func (c *Client) GetTVSimilar(id int64, urlOptions map[string]string) (*TVSimilar, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/similar?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tVSimilar := TVSimilar{}
	if err := c.get(tmdbURL, &tVSimilar); err != nil {
		return nil, err
	}
	return &tVSimilar, nil
}

// TVWatchProviders type is a struct for watch/providers JSON response.
type TVWatchProviders struct {
	ID int64 `json:"id,omitempty"`
	*TVWatchProvidersResults
}

// GetTVWatchProviders get a list of the availabilities per country by provider for a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-watch-providers
func (c *Client) GetTVWatchProviders(id int64, urlOptions map[string]string) (*TVWatchProviders, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/watch/providers?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvWatchProviders := TVWatchProviders{}
	if err := c.get(tmdbURL, &tvWatchProviders); err != nil {
		return nil, err
	}
	return &tvWatchProviders, nil
}

// TVTranslations type is a struct for translations JSON response.
type TVTranslations struct {
	ID           int64 `json:"id,omitempty"`
	Translations []struct {
		Iso3166_1   string `json:"iso_3166_1"`
		Iso639_1    string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Name     string `json:"name"`
			Overview string `json:"overview"`
			Tagline  string `json:"tagline"`
			Homepage string `json:"homepage"`
		} `json:"data"`
	} `json:"translations"`
}

// GetTVTranslations get a list fo translations that have been created for a TV Show.
//
// https://developers.themoviedb.org/3/tv/get-tv-translations
func (c *Client) GetTVTranslations(id int64, urlOptions map[string]string) (*TVTranslations, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/translations?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvTranslations := TVTranslations{}
	if err := c.get(tmdbURL, &tvTranslations); err != nil {
		return nil, err
	}
	return &tvTranslations, nil
}

// TVVideos type is a struct for videos JSON response.
type TVVideos struct {
	ID int64 `json:"id,omitempty"`
	*TVVideosResults
}

// GetTVVideos get the videos that have been added to a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-videos
func (c *Client) GetTVVideos(id int64, urlOptions map[string]string) (*TVVideos, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/videos?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvVideos := TVVideos{}
	if err := c.get(tmdbURL, &tvVideos); err != nil {
		return nil, err
	}
	return &tvVideos, nil
}

// TVLatest type is a struct for latest JSON response.
type TVLatest struct {
	*TVDetails
}

// GetTVLatest get the most newly created TV show.
//
// This is a live response and will continuously change.
//
// https://developers.themoviedb.org/3/tv/get-latest-tv
func (c *Client) GetTVLatest(
	urlOptions map[string]string,
) (*TVLatest, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%slatest?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	tvLatest := TVLatest{}
	if err := c.get(tmdbURL, &tvLatest); err != nil {
		return nil, err
	}
	return &tvLatest, nil
}

// TVAiringToday type is a struct for airing today JSON response.
type TVAiringToday struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	*TVAiringTodayResults
}

// GetTVAiringToday get a list of TV shows that are airing today.
// This query is purely day based as we do not currently support
// airing times.
//
// You can specify a  to offset the day calculation.
// Without a specified timezone, this query defaults
// to EST (Eastern Time UTC-05:00).
//
// https://developers.themoviedb.org/3/tv/get-tv-airing-today
func (c *Client) GetTVAiringToday(
	urlOptions map[string]string,
) (*TVAiringToday, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sairing_today?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	tvAiringToday := TVAiringToday{}
	if err := c.get(tmdbURL, &tvAiringToday); err != nil {
		return nil, err
	}
	return &tvAiringToday, nil
}

// TVOnTheAir type is a struct for on the air JSON response.
type TVOnTheAir struct {
	*TVAiringToday
}

// GetTVOnTheAir get a list of shows that are currently on the air.
//
// This query looks for any TV show that has an episode with an
// air date in the next 7 days.
//
// https://developers.themoviedb.org/3/tv/get-tv-on-the-air
func (c *Client) GetTVOnTheAir(
	urlOptions map[string]string,
) (*TVOnTheAir, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%son_the_air?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	tvOnTheAir := TVOnTheAir{}
	if err := c.get(tmdbURL, &tvOnTheAir); err != nil {
		return nil, err
	}
	return &tvOnTheAir, nil
}

// TVPopular type is a struct for popular JSON response.
type TVPopular struct {
	*TVAiringToday
}

// GetTVPopular get a list of the current popular TV shows on TMDb.
// This list updates daily.
//
// https://developers.themoviedb.org/3/tv/get-popular-tv-shows
func (c *Client) GetTVPopular(
	urlOptions map[string]string,
) (*TVPopular, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%spopular?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	tvPopular := TVPopular{}
	if err := c.get(tmdbURL, &tvPopular); err != nil {
		return nil, err
	}
	return &tvPopular, nil
}

// TVTopRated type is a struct for top rated JSON response.
type TVTopRated struct {
	*TVAiringToday
}

// GetTVTopRated get a list of the top rated TV shows on TMDb.
//
// https://developers.themoviedb.org/3/tv/get-top-rated-tv
func (c *Client) GetTVTopRated(
	urlOptions map[string]string,
) (*TVTopRated, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stop_rated?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	tvTopRated := TVTopRated{}
	if err := c.get(tmdbURL, &tvTopRated); err != nil {
		return nil, err
	}
	return &tvTopRated, nil
}

// PostTVShowRating rate a TV show.
//
// A valid session or guest session ID is required.
//
// You can read more about how this works:
// https://developers.themoviedb.org/3/authentication/how-do-i-generate-a-session-id
//
// https://developers.themoviedb.org/3/tv/rate-tv-show
func (c *Client) PostTVShowRating(id int64, rating float32, urlOptions map[string]string) (*Response, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rating?api_key=%s&session_id=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	body := struct {
		Value float32 `json:"value"`
	}{Value: rating}
	tvShowRating := Response{}
	if err := c.request(
		tmdbURL,
		body,
		http.MethodPost,
		&tvShowRating,
	); err != nil {
		return nil, err
	}
	return &tvShowRating, nil
}

// DeleteTVShowRating remove your rating for a TV show.
//
// A valid session or guest session ID is required.
//
// You can read more about how this works:
// https://developers.themoviedb.org/3/authentication/how-do-i-generate-a-session-id
//
// https://developers.themoviedb.org/3/tv/delete-tv-show-rating
func (c *Client) DeleteTVShowRating(id int64, urlOptions map[string]string) (*Response, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rating?api_key=%s&session_id=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	tvShowRating := Response{}
	if err := c.request(
		tmdbURL,
		[]byte{},
		http.MethodDelete,
		&tvShowRating,
	); err != nil {
		return nil, err
	}
	return &tvShowRating, nil
}
