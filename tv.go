package tmdb

import (
	"encoding/json"
	"fmt"
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
	Type        string  `json:"type"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
	*TVAlternativeTitlesShort
	*TVChangesShort
	*TVContentRatingsShort
	*TVCreditsShort
	*TVEpisodeGroupsShort
	*TVExternalIDsShort
	*TVImagesShort
	*TVKeywordsShort
	*TVRecommendationsShort
	*TVReviewsShort
	*TVScreenedTheatricallyShort
	*TVSimilarShort
	*TVTranslations
}

// TVAccountStates type is a struct for account states JSON response.
type TVAccountStates struct {
	ID        int64           `json:"id"`
	Favorite  bool            `json:"favorite"`
	Rated     json.RawMessage `json:"rated"`
	Watchlist bool            `json:"watchlist"`
}

// TVAlternativeTitles type is a struct for alternative titles JSON response.
type TVAlternativeTitles struct {
	ID      int `json:"id,omitempty"`
	Results []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Title     string `json:"title"`
		Type      string `json:"type"`
	} `json:"results"`
}

// TVAlternativeTitlesShort type is a short struct for alternative titles JSON response.
type TVAlternativeTitlesShort struct {
	AlternativeTitles *TVAlternativeTitles `json:"alternative_titles,omitempty"`
}

// TVChanges type is a struct for changes JSON response.
type TVChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID     string `json:"id"`
			Action string `json:"action"`
			Time   string `json:"time"`
			Value  struct {
				SeasonID     int64 `json:"season_id"`
				SeasonNumber int   `json:"season_number"`
			} `json:"value"`
		} `json:"items"`
	} `json:"changes"`
}

// TVChangesShort type is a short struct for changes JSON response.
type TVChangesShort struct {
	Changes *TVChanges `json:"changes,omitempty"`
}

// TVContentRatings type is a struct for content ratings JSON response.
type TVContentRatings struct {
	Results []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Rating    string `json:"rating"`
	} `json:"results"`
	ID int64 `json:"id,omitempty"`
}

// TVContentRatingsShort type is a short struct for content ratings JSON response.
type TVContentRatingsShort struct {
	ContentRatings *TVContentRatings `json:"content_ratings,omitempty"`
}

// TVCredits type is a struct for credits JSON response.
type TVCredits struct {
	ID   int64 `json:"id,omitempty"`
	Cast []struct {
		Character   string `json:"character"`
		CreditID    string `json:"credit_id"`
		Gender      int    `json:"gender"`
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Order       int    `json:"order"`
		ProfilePath string `json:"profile_path"`
	} `json:"cast"`
	Crew []struct {
		CreditID    string `json:"credit_id"`
		Department  string `json:"department"`
		Gender      int    `json:"gender"`
		ID          int64  `json:"id"`
		Job         string `json:"job"`
		Name        string `json:"name"`
		ProfilePath string `json:"profile_path"`
	} `json:"crew"`
}

// TVCreditsShort type is a short struct for credits JSON response.
type TVCreditsShort struct {
	Credits struct {
		*TVCredits
	} `json:"credits,omitempty"`
}

// TVEpisodeGroups type is a struct for episode groups JSON response.
type TVEpisodeGroups struct {
	Results []struct {
		Description  string `json:"description"`
		EpisodeCount int    `json:"episode_count"`
		GroupCount   int    `json:"group_count"`
		ID           string `json:"id"`
		Name         string `json:"name"`
		Network      struct {
			ID            int64  `json:"id"`
			LogoPath      string `json:"logo_path"`
			Name          string `json:"name"`
			OriginCountry string `json:"origin_country"`
		} `json:"network"`
		Type int `json:"type"`
	} `json:"results"`
	ID int64 `json:"id,omitempty"`
}

// TVEpisodeGroupsShort type is a short struct for episode groups JSON response.
type TVEpisodeGroupsShort struct {
	EpisodeGroups *TVEpisodeGroups `json:"episode_groups,omitempty"`
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

// TVExternalIDsShort type is a short struct for external ids JSON response.
type TVExternalIDsShort struct {
	*TVExternalIDs `json:"external_ids,omitempty"`
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

// TVImagesShort type is a short struct for images JSON response.
type TVImagesShort struct {
	Images *TVImages `json:"images,omitempty"`
}

// TVKeywords type is a struct for keywords JSON response.
type TVKeywords struct {
	ID      int64 `json:"id,omitempty"`
	Results []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"results"`
}

// TVKeywordsShort type is a short struct for keywords JSON response.
type TVKeywordsShort struct {
	Keywords struct {
		*TVKeywords
	} `json:"keywords,omitempty"`
}

// TVRecommendations type is a struct for recommendations JSON response.
type TVRecommendations struct {
	Page    int64 `json:"page"`
	Results []struct {
		PosterPath       string   `json:"poster_path"`
		Popularity       float32  `json:"popularity"`
		ID               int64    `json:"id"`
		BackdropPath     string   `json:"backdrop_path"`
		VoteAverage      float32  `json:"vote_average"`
		Overview         string   `json:"overview"`
		FirstAirDate     string   `json:"first_air_date"`
		OriginCountry    []string `json:"origin_country"`
		GenreIDs         []int64  `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		VoteCount        int64    `json:"vote_count"`
		Name             string   `json:"name"`
		OriginalName     string   `json:"original_name"`
	} `json:"results"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// TVRecommendationsShort type is a short struct for recommendations JSON response.
type TVRecommendationsShort struct {
	Recommendations *TVRecommendations `json:"recommendations,omitempty"`
}

// TVReviews type is a struct for reviews JSON response.
type TVReviews struct {
	ID      int64 `json:"id,omitempty"`
	Page    int64 `json:"page"`
	Results []struct {
		ID      string `json:"id"`
		Author  string `json:"author"`
		Content string `json:"content"`
		URL     string `json:"url"`
	} `json:"results"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// TVReviewsShort type is a short struct for reviews JSON response.
type TVReviewsShort struct {
	Reviews struct {
		*TVReviews
	} `json:"reviews,omitempty"`
}

// TVScreenedTheatrically type is a struct for screened theatrically JSON response.
type TVScreenedTheatrically struct {
	ID      int64 `json:"id,omitempty"`
	Results []struct {
		ID            int64 `json:"id"`
		EpisodeNumber int   `json:"episode_number"`
		SeasonNumber  int   `json:"season_number"`
	} `json:"results"`
}

// TVScreenedTheatricallyShort type is a short struct
// for screened theatrically JSON response.
type TVScreenedTheatricallyShort struct {
	ScreenedTheatrically *TVScreenedTheatrically `json:"screened_theatrically,omitempty"`
}

// TVSimilar type is a struct for similar tv shows JSON response.
type TVSimilar struct {
	*TVRecommendations
}

// TVSimilarShort type is a short struct for similar tv shows JSON response.
type TVSimilarShort struct {
	Similar *TVSimilar `json:"similar,omitempty"`
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
			Homepage string `json:"homepage"`
		} `json:"data"`
	} `json:"translations"`
}

// TVTranslationsShort type is a short struct for translations JSON response.
type TVTranslationsShort struct {
	Translations struct {
		*TVTranslations
	} `json:"translations,omitempty"`
}

// GetTVDetails get the primary TV show details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/tv/get-tv-details
func (c *Client) GetTVDetails(id int, o map[string]string) (*TVDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVDetails{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
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
//
func (c *Client) GetTVAccountStates(id int, o map[string]string) (*TVAccountStates, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/account_states?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVAccountStates{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVAlternativeTitles get all of the alternative titles for a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-alternative-titles
func (c *Client) GetTVAlternativeTitles(id int, o map[string]string) (*TVAlternativeTitles, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/alternative_titles?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVAlternativeTitles{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
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
func (c *Client) GetTVChanges(id int, o map[string]string) (*TVChanges, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/changes?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVChanges{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVContentRatings get the list of content ratings (certifications) that have been added to a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-content-ratings
func (c *Client) GetTVContentRatings(id int, o map[string]string) (*TVContentRatings, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/content_ratings?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVContentRatings{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVCredits get the credits (cast and crew) that have been added to a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-credits
func (c *Client) GetTVCredits(id int, o map[string]string) (*TVCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/credits?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVCredits{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVEpisodeGroups get all of the episode groups that have been created for a TV show.
//
// With a group ID you can call the get TV episode group details method.
//
// https://developers.themoviedb.org/3/tv/get-tv-episode-groups
func (c *Client) GetTVEpisodeGroups(id int, o map[string]string) (*TVEpisodeGroups, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/episode_groups?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVEpisodeGroups{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
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
func (c *Client) GetTVExternalIDs(id int, o map[string]string) (*TVExternalIDs, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/external_ids?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVExternalIDs{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVImages get the images that belong to a TV show.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter. This should be a comma
// separated value like so: include_image_language=en,null.
//
// https://developers.themoviedb.org/3/tv/get-tv-images
func (c *Client) GetTVImages(id int, o map[string]string) (*TVImages, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/images?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVImages{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVKeywords get the keywords that have been added to a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-keywords
func (c *Client) GetTVKeywords(id int) (*TVKeywords, error) {
	tmdbURL := fmt.Sprintf("%s%s%d/keywords?api_key=%s", baseURL, tvURL, id, c.APIKey)
	t := TVKeywords{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVRecommendations get the list of TV show recommendations for this item.
//
// https://developers.themoviedb.org/3/tv/get-tv-recommendations
func (c *Client) GetTVRecommendations(id int, o map[string]string) (*TVRecommendations, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/recommendations?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVRecommendations{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVReviews get the reviews for a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-reviews
func (c *Client) GetTVReviews(id int, o map[string]string) (*TVReviews, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/reviews?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVReviews{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVScreenedTheatrically get a list of seasons or episodes that
// have been screened in a film festival or theatre.
//
// https://developers.themoviedb.org/3/tv/get-screened-theatrically
func (c *Client) GetTVScreenedTheatrically(id int) (*TVScreenedTheatrically, error) {
	tmdbURL := fmt.Sprintf("%s%s%d/screened_theatrically?api_key=%s", baseURL, tvURL, id, c.APIKey)
	t := TVScreenedTheatrically{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVSimilar a list of similar TV shows.
// These items are assembled by looking at keywords and genres.
//
// https://developers.themoviedb.org/3/tv/get-similar-tv-shows
func (c *Client) GetTVSimilar(id int, o map[string]string) (*TVSimilar, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/similar?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVSimilar{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVTranslations get a list of translations that have been created for a TV Show.
//
// https://developers.themoviedb.org/3/tv/get-tv-translations
func (c *Client) GetTVTranslations(id int, o map[string]string) (*TVTranslations, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/translations?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVTranslations{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
