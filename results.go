package tmdb

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

type MovieVideosResults struct {
	Results []struct {
		ID        string `json:"id"`
		Iso639_1  string `json:"iso_639_1"`
		Iso3166_1 string `json:"iso_3166_1"`
		Key       string `json:"key"`
		Name      string `json:"name"`
		Site      string `json:"site"`
		Size      int    `json:"size"`
		Type      string `json:"type"`
	} `json:"results"`
}

type MovieWatchProvidersResults struct {
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
		VoteCount        int64   `json:"vote_count"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
	} `json:"results"`
}

type MovieReviewsResults struct {
	Results []struct {
		ID      string `json:"id"`
		Author  string `json:"author"`
		Content string `json:"content"`
		URL     string `json:"url"`
	} `json:"results"`
}

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

type MovieNowPlayingResults struct {
	Results []struct {
		PosterPath  string `json:"poster_path"`
		Adult       bool   `json:"adult"`
		Overview    string `json:"overview"`
		ReleaseDate string `json:"release_date"`
		Genres      []struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"genres"`
		ID               int64   `json:"id"`
		OriginalTitle    string  `json:"original_title"`
		OriginalLanguage string  `json:"original_language"`
		Title            string  `json:"title"`
		BackdropPath     string  `json:"backdrop_path"`
		Popularity       float32 `json:"popularity"`
		VoteCount        int64   `json:"vote_count"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
	} `json:"results"`
}

type MoviePopularResults struct {
	Results []struct {
		PosterPath  string `json:"poster_path"`
		Adult       bool   `json:"adult"`
		Overview    string `json:"overview"`
		ReleaseDate string `json:"release_date"`
		Genres      []struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"genres"`
		ID               int64   `json:"id"`
		OriginalTitle    string  `json:"original_title"`
		OriginalLanguage string  `json:"original_language"`
		Title            string  `json:"title"`
		BackdropPath     string  `json:"backdrop_path"`
		Popularity       float32 `json:"popularity"`
		VoteCount        int64   `json:"vote_count"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
	} `json:"results"`
}
