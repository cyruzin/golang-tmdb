package tmdb

// AccountCreatedListsResults Result Types
type AccountCreatedListsResults struct {
	Results []struct {
		Description   string `json:"description"`
		FavoriteCount int64  `json:"favorite_count"`
		ID            int64  `json:"id"`
		ItemCount     int64  `json:"item_count"`
		Iso639_1      string `json:"iso_639_1"`
		ListType      string `json:"list_type"`
		Name          string `json:"name"`
		PosterPath    string `json:"poster_path"`
	} `json:"results"`
}

// MovieResult represents the details of a movie as returned by the TMDB API.
// It includes information such as the movie's ID, titles, language, overview,
// release date, poster and backdrop paths, popularity metrics, genre IDs, and flags
// indicating whether the movie is for adults or contains video content.
type MovieResult struct {
	ID               int64   `json:"id"`
	Title            string  `json:"title"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	PosterPath       string  `json:"poster_path"`
	BackdropPath     string  `json:"backdrop_path"`
	Popularity       float32 `json:"popularity"`
	GenreIDs         []int64 `json:"genre_ids"`
	Adult            bool    `json:"adult"`
	Video            bool    `json:"video"`
	VoteMetrics
}

// AccountFavoriteMoviesResults Result Types
type AccountFavoriteMoviesResults struct {
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int   `json:"genre_ids"`
		ID               int64   `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		PosterPath       string  `json:"poster_path"`
		Popularity       float64 `json:"popularity"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteMetrics
	} `json:"results"`
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
	Adult            bool     `json:"adult"`
	VoteMetrics
}

// AccountFavoriteTVShowsResults Result Types
type AccountFavoriteTVShowsResults struct {
	Results []struct {
		BackdropPath     string   `json:"backdrop_path"`
		FirstAirDate     string   `json:"first_air_date"`
		GenreIDs         []int64  `json:"genre_ids"`
		ID               int64    `json:"id"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
		PosterPath       string   `json:"poster_path"`
		Popularity       float64  `json:"popularity"`
		Name             string   `json:"name"`
		VoteMetrics
	} `json:"results"`
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
	Results []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"results"`
}

// SearchMoviesResults Result Types
type SearchMoviesResults struct {
	Results []MovieResult `json:"results"`
}

// SearchMultiResults Result Types
type SearchMultiResults struct {
	Results []struct {
		PosterPath       string   `json:"poster_path,omitempty"`
		Popularity       float32  `json:"popularity"`
		ID               int64    `json:"id"`
		Overview         string   `json:"overview,omitempty"`
		BackdropPath     string   `json:"backdrop_path,omitempty"`
		MediaType        string   `json:"media_type"`
		FirstAirDate     string   `json:"first_air_date,omitempty"`
		OriginCountry    []string `json:"origin_country,omitempty"`
		GenreIDs         []int64  `json:"genre_ids,omitempty"`
		OriginalLanguage string   `json:"original_language,omitempty"`
		Name             string   `json:"name,omitempty"`
		OriginalName     string   `json:"original_name,omitempty"`
		Adult            bool     `json:"adult,omitempty"`
		ReleaseDate      string   `json:"release_date,omitempty"`
		OriginalTitle    string   `json:"original_title,omitempty"`
		Title            string   `json:"title,omitempty"`
		Video            bool     `json:"video,omitempty"`
		ProfilePath      string   `json:"profile_path,omitempty"`
		KnownFor         []struct {
			PosterPath       string  `json:"poster_path"`
			Adult            bool    `json:"adult"`
			Overview         string  `json:"overview"`
			ReleaseDate      string  `json:"release_date"`
			OriginalTitle    string  `json:"original_title"`
			GenreIDs         []int64 `json:"genre_ids"`
			ID               int64   `json:"id"`
			MediaType        string  `json:"media_type"`
			OriginalLanguage string  `json:"original_language"`
			Title            string  `json:"title"`
			BackdropPath     string  `json:"backdrop_path"`
			Popularity       float32 `json:"popularity"`
			Video            bool    `json:"video"`
			VoteMetrics
		} `json:"known_for,omitempty"`
		VoteMetrics
	} `json:"results"`
}

// SearchPeopleResults Result Types
type SearchPeopleResults struct {
	Results []struct {
		Adult    bool  `json:"adult"`
		Gender   int   `json:"gender,omitempty"`
		ID       int64 `json:"id"`
		KnownFor []struct {
			Adult            bool     `json:"adult,omitempty"` // Movie
			BackdropPath     string   `json:"backdrop_path"`
			FirstAirDate     string   `json:"first_air_date,omitempty"` // TV
			GenreIDs         []int64  `json:"genre_ids"`
			ID               int64    `json:"id"`
			MediaType        string   `json:"media_type"`
			Name             string   `json:"name,omitempty"` // TV
			OriginalLanguage string   `json:"original_language"`
			OriginalName     string   `json:"original_name,omitempty"`  // TV
			OriginalTitle    string   `json:"original_title,omitempty"` // Movie
			OriginCountry    []string `json:"origin_country,omitempty"` // TV
			Overview         string   `json:"overview"`
			Popularity       float32  `json:"popularity"`
			PosterPath       string   `json:"poster_path"`
			ReleaseDate      string   `json:"release_date,omitempty"` // Movie
			Title            string   `json:"title,omitempty"`        // Movie
			Video            bool     `json:"video,omitempty"`        // Movie
			VoteMetrics
		} `json:"known_for"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"results"`
}

// SearchTVShowsResults Result Types
type SearchTVShowsResults struct {
	Results []TVShowResult `json:"results"`
}

// TrendingResults Result Types
type TrendingResults struct {
	Results []struct {
		Adult              bool     `json:"adult,omitempty"`
		Gender             int      `json:"gender,omitempty"`
		BackdropPath       string   `json:"backdrop_path,omitempty"`
		GenreIDs           []int64  `json:"genre_ids,omitempty"`
		ID                 int64    `json:"id"`
		OriginalLanguage   string   `json:"original_language"`
		OriginalTitle      string   `json:"original_title,omitempty"`
		Overview           string   `json:"overview,omitempty"`
		PosterPath         string   `json:"poster_path,omitempty"`
		ReleaseDate        string   `json:"release_date,omitempty"`
		Title              string   `json:"title,omitempty"`
		Video              bool     `json:"video,omitempty"`
		Popularity         float32  `json:"popularity,omitempty"`
		FirstAirDate       string   `json:"first_air_date,omitempty"`
		Name               string   `json:"name,omitempty"`
		OriginCountry      []string `json:"origin_country,omitempty"`
		OriginalName       string   `json:"original_name,omitempty"`
		KnownForDepartment string   `json:"known_for_department,omitempty"`
		ProfilePath        string   `json:"profile_path,omitempty"`
		MediaType          string   `json:"media_type,omitempty"`
		KnownFor           []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			GenreIds         []int   `json:"genre_ids"`
			ID               int     `json:"id"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			PosterPath       string  `json:"poster_path"`
			ReleaseDate      string  `json:"release_date"`
			Title            string  `json:"title"`
			Video            bool    `json:"video"`
			Popularity       float64 `json:"popularity"`
			MediaType        string  `json:"media_type"`
			VoteMetrics
		} `json:"known_for,omitempty"`
		VoteMetrics
	} `json:"results"`
}

// MovieReleaseDatesResults Result Types
type MovieReleaseDatesResults struct {
	Results []struct {
		Iso3166_1    string `json:"iso_3166_1"`
		ReleaseDates []struct {
			Certification string `json:"certification"`
			Iso639_1      string `json:"iso_639_1"`
			ReleaseDate   string `json:"release_date"`
			Type          int    `json:"type"`
			Note          string `json:"note"`
		} `json:"release_dates"`
	} `json:"results"`
}

// MovieVideosResults Result Types
type MovieVideosResults struct {
	Results []struct {
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
	} `json:"results"`
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
	Results []struct {
		PosterPath       string  `json:"poster_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		OriginalTitle    string  `json:"original_title"`
		OriginalLanguage string  `json:"original_language"`
		Title            string  `json:"title"`
		BackdropPath     string  `json:"backdrop_path"`
		Popularity       float32 `json:"popularity"`
		Video            bool    `json:"video"`
		VoteMetrics
	} `json:"results"`
}

// MovieReviewsResults Result Types
type MovieReviewsResults struct {
	Results []struct {
		ID      string `json:"id"`
		Author  string `json:"author"`
		Content string `json:"content"`
		URL     string `json:"url"`
	} `json:"results"`
}

// MovieListsResults Result Types
type MovieListsResults struct {
	Results []struct {
		Description   string `json:"description"`
		FavoriteCount int64  `json:"favorite_count"`
		ID            int64  `json:"id"`
		ItemCount     int64  `json:"item_count"`
		Iso639_1      string `json:"iso_639_1"`
		ListType      string `json:"list_type"`
		Name          string `json:"name"`
		PosterPath    string `json:"poster_path"`
	} `json:"results"`
}

// MovieNowPlayingResults Result Types
type MovieNowPlayingResults struct {
	Results []struct {
		PosterPath       string  `json:"poster_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		Genres           []Genre `json:"genres"`
		ID               int64   `json:"id"`
		OriginalTitle    string  `json:"original_title"`
		OriginalLanguage string  `json:"original_language"`
		Title            string  `json:"title"`
		BackdropPath     string  `json:"backdrop_path"`
		Popularity       float32 `json:"popularity"`
		Video            bool    `json:"video"`
		VoteMetrics
	} `json:"results"`
}

// MoviePopularResults Result Types
type MoviePopularResults struct {
	Results []struct {
		PosterPath       string  `json:"poster_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		Genres           []Genre `json:"genres"`
		ID               int64   `json:"id"`
		OriginalTitle    string  `json:"original_title"`
		OriginalLanguage string  `json:"original_language"`
		Title            string  `json:"title"`
		BackdropPath     string  `json:"backdrop_path"`
		Popularity       float32 `json:"popularity"`
		Video            bool    `json:"video"`
		VoteMetrics
	} `json:"results"`
}

// TVContentRatingsResults Result Types
type TVContentRatingsResults struct {
	Results []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Rating    string `json:"rating"`
	} `json:"results"`
}

// TVEpisodeGroupsResults Result Types
type TVEpisodeGroupsResults struct {
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
}

// TVKeywordsResults Result Types
type TVKeywordsResults struct {
	Results []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"results"`
}

// TVRecommendationsResults Result Types
type TVRecommendationsResults struct {
	Results []struct {
		PosterPath       string   `json:"poster_path"`
		Popularity       float32  `json:"popularity"`
		ID               int64    `json:"id"`
		BackdropPath     string   `json:"backdrop_path"`
		Overview         string   `json:"overview"`
		FirstAirDate     string   `json:"first_air_date"`
		OriginCountry    []string `json:"origin_country"`
		GenreIDs         []int64  `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		Name             string   `json:"name"`
		OriginalName     string   `json:"original_name"`
		VoteMetrics
	} `json:"results"`
}

// TVReviewsResults Result Types
type TVReviewsResults struct {
	Results []struct {
		ID      string `json:"id"`
		Author  string `json:"author"`
		Content string `json:"content"`
		URL     string `json:"url"`
	} `json:"results"`
}

// TVScreenedTheatricallyResults Result Types
type TVScreenedTheatricallyResults struct {
	Results []struct {
		ID            int64 `json:"id"`
		EpisodeNumber int   `json:"episode_number"`
		SeasonNumber  int   `json:"season_number"`
	} `json:"results"`
}

// TVWatchProvidersResults Result Types
type TVWatchProvidersResults struct {
	Results map[string]struct {
		Link     string `json:"link"`
		Flatrate []struct {
			DisplayPriority int64  `json:"display_priority"`
			LogoPath        string `json:"logo_path"`
			ProviderID      int64  `json:"provider_id"`
			ProviderName    string `json:"provider_name"`
		} `json:"flatrate,omitempty"`
		Rent []struct {
			DisplayPriority int64  `json:"display_priority"`
			LogoPath        string `json:"logo_path"`
			ProviderID      int64  `json:"provider_id"`
			ProviderName    string `json:"provider_name"`
		} `json:"rent,omitempty"`
		Buy []struct {
			DisplayPriority int64  `json:"display_priority"`
			LogoPath        string `json:"logo_path"`
			ProviderID      int64  `json:"provider_id"`
			ProviderName    string `json:"provider_name"`
		} `json:"buy,omitempty"`
	} `json:"results"`
}

// TVAiringTodayResults Result Types
type TVAiringTodayResults struct {
	Results []struct {
		OriginalName     string   `json:"original_name"`
		GenreIDs         []int64  `json:"genre_ids"`
		Name             string   `json:"name"`
		Popularity       float32  `json:"popularity"`
		OriginCountry    []string `json:"origin_country"`
		FirstAirDate     string   `json:"first_air_date"`
		BackdropPath     string   `json:"backdrop_path"`
		OriginalLanguage string   `json:"original_language"`
		ID               int64    `json:"id"`
		Overview         string   `json:"overview"`
		PosterPath       string   `json:"poster_path"`
		VoteMetrics
	} `json:"results"`
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
