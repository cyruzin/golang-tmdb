package tmdb

// AccountCreatedListsResults Result Types
type AccountCreatedListsResults struct {
	Results []List `json:"results"`
}

// MovieResult represents the details of a movie as returned by the TMDB API.
// It includes information such as the movie's ID, titles, language, overview,
// release date, poster and backdrop paths, popularity metrics, genre IDs, and flags
// indicating whether the movie is for adults or contains video content.
type MovieResult struct {
	MovieMedia
}

// AccountFavoriteMoviesResults Result Types
type AccountFavoriteMoviesResults struct {
	Results []Movie `json:"results"`
}

// TVShowResult represents the details of a TV show as returned by the TMDB API.
// It includes information such as the show's ID, name, original name and language,
// overview, first air date, poster and backdrop paths, popularity score, vote count
// and average, associated genre IDs, and the countries of origin.
type TVShowResult struct {
	ID               int64    `json:"id"`
	Name             string   `json:"name"`
	OriginalName     string   `json:"original_name"`
	OriginalLanguage string   `json:"original_language"`
	Overview         string   `json:"overview"`
	FirstAirDate     string   `json:"first_air_date"`
	PosterPath       string   `json:"poster_path"`
	BackdropPath     string   `json:"backdrop_path"`
	Popularity       float32  `json:"popularity"`
	GenreIDs         []int64  `json:"genre_ids"`
	OriginCountry    []string `json:"origin_country"`
	VoteMetrics
}

// AccountFavoriteTVShowsResults Result Types
type AccountFavoriteTVShowsResults struct {
	Results []TVShow `json:"results"`
}

// AccountRatedTVEpisodesResults Result Types
type AccountRatedTVEpisodesResults struct {
	Results []struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int64   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int64   `json:"show_id"`
		StillPath      string  `json:"still_path"`
		Rating         float32 `json:"rating"`
		VoteMetrics
	} `json:"results"`
}

// ChangesMovieResults Result Types
type ChangesMovieResults struct {
	Results []struct {
		ID    int64 `json:"id"`
		Adult bool  `json:"adult"`
	} `json:"results"`
}

// CompanyAlternativeNamesResult Result Types
type CompanyAlternativeNamesResult struct {
	Results []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"results"`
}

// DiscoverMovieResults Result Types
type DiscoverMovieResults struct {
	Results []MovieResult `json:"results"`
}

// DiscoverTVResults Result Types
type DiscoverTVResults struct {
	Results []TVShowResult `json:"results"`
}

// SearchCompaniesResults Result Types
type SearchCompaniesResults struct {
	Results []struct {
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"results"`
}

// SearchCollectionsResults Result Types
type SearchCollectionsResults struct {
	Results []struct {
		Adult            bool   `json:"adult"`
		BackdropPath     string `json:"backdrop_path"`
		ID               int64  `json:"id"`
		Name             string `json:"name"`
		OriginalLanguage string `json:"original_language"`
		OriginalName     string `json:"original_name"`
		Overview         string `json:"overview"`
		PosterPath       string `json:"poster_path"`
	} `json:"results"`
}

// SearchKeywordsResults Result Types
type SearchKeywordsResults struct {
	Results []KeywordDetails `json:"results"`
}

// SearchMoviesResults Result Types
type SearchMoviesResults struct {
	Results []MovieResult `json:"results"`
}

// SearchMultiResults Result Types
type SearchMultiResults struct {
	Results []struct {
		PosterPath       string    `json:"poster_path,omitempty"`
		Popularity       float32   `json:"popularity"`
		ID               int64     `json:"id"`
		Overview         string    `json:"overview,omitempty"`
		BackdropPath     string    `json:"backdrop_path,omitempty"`
		MediaType        MediaType `json:"media_type"`
		FirstAirDate     string    `json:"first_air_date,omitempty"`
		OriginCountry    []string  `json:"origin_country,omitempty"`
		GenreIDs         []int64   `json:"genre_ids,omitempty"`
		OriginalLanguage string    `json:"original_language,omitempty"`
		Name             string    `json:"name,omitempty"`
		OriginalName     string    `json:"original_name,omitempty"`
		Adult            bool      `json:"adult,omitempty"`
		ReleaseDate      string    `json:"release_date,omitempty"`
		OriginalTitle    string    `json:"original_title,omitempty"`
		Title            string    `json:"title,omitempty"`
		Video            bool      `json:"video,omitempty"`
		ProfilePath      string    `json:"profile_path,omitempty"`
		KnownFor         []struct {
			PosterPath       string    `json:"poster_path"`
			Adult            bool      `json:"adult"`
			Overview         string    `json:"overview"`
			ReleaseDate      string    `json:"release_date"`
			OriginalTitle    string    `json:"original_title"`
			GenreIDs         []int64   `json:"genre_ids"`
			ID               int64     `json:"id"`
			MediaType        MediaType `json:"media_type"`
			OriginalLanguage string    `json:"original_language"`
			Title            string    `json:"title"`
			BackdropPath     string    `json:"backdrop_path"`
			Popularity       float32   `json:"popularity"`
			Video            bool      `json:"video"`
			VoteMetrics
		} `json:"known_for,omitempty"`
		VoteMetrics
	} `json:"results"`
}

// SearchPeopleResults Result Types
type SearchPeopleResults struct {
	Results []PersonResults `json:"results"`
}

// SearchTVShowsResults Result Types
type SearchTVShowsResults struct {
	Results []TVShowResult `json:"results"`
}

// TrendingResults Result Types
type TrendingResults struct {
	Results []Media `json:"results"`
}

// MovieReleaseDatesResults Result Types
type MovieReleaseDatesResults struct {
	Results []MovieResultDate `json:"results"`
}

type WatchProviderResult struct {
	Link     string           `json:"link"`
	Flatrate *[]WatchProvider `json:"flatrate"`
	Rent     *[]WatchProvider `json:"rent"`
	Buy      *[]WatchProvider `json:"buy"`
}

type WatchProviderResults struct {
	ID      int64                          `json:"id"`
	Results map[string]WatchProviderResult `json:"results"`
}

// MovieRecommendationsResults Result Types
type MovieRecommendationsResults struct {
	Results []Movie `json:"results"`
}

// MovieReviewsResults Result Types
type MovieReviewsResults struct {
	Results []Review `json:"results"`
}

// MovieListsResults Result Types
type MovieListsResults struct {
	Results []MovieList `json:"results"`
}

// MovieNowPlayingResults Result Types
type MovieNowPlayingResults struct {
	Results []Movie `json:"results"`
}

// MoviePopularResults Result Types
type MoviePopularResults struct {
	Results []Movie `json:"results"`
}

// TVContentRatingsResults Result Types
type TVContentRatingsResults struct {
	Results []ContentRatings `json:"results"`
}

// TVEpisodeGroupsResults Result Types
type TVEpisodeGroupsResults struct {
	Results []TVEpisodeGroup `json:"results"`
}

// TVKeywordsResults Result Types
type TVKeywordsResults struct {
	Results []KeywordDetails `json:"results"`
}

type TVRecommendationsMediaResults struct {
	Results []TVShowMedia `json:"results"`
}

// TVReviewsResults Result Types
type TVReviewsResults struct {
	Results []Review `json:"results"`
}

// TVScreenedTheatricallyResults Result Types
type TVScreenedTheatricallyResults struct {
	Results []ScreenedTheatrically `json:"results"`
}

// TVAiringTodayResults Result Types
type TVAiringTodayResults struct {
	Results []TVShow `json:"results"`
}

type VideoResult struct {
	ID          string `json:"id"`
	Iso639_1    string `json:"iso_639_1"`
	Iso3166_1   string `json:"iso_3166_1"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Official    bool   `json:"official"`
	PublishedAt string `json:"published_at"`
	Site        string `json:"site"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
}

type VideoResults struct {
	ID      int64         `json:"id"`
	Results []VideoResult `json:"results"`
}
