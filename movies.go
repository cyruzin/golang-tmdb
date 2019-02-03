package tmdb

import (
	"encoding/json"
	"fmt"
)

// MovieDetails type is a struct for movie details JSON response.
type MovieDetails struct {
	Adult               bool   `json:"adult"`
	BackdropPath        string `json:"backdrop_path"`
	BelongsToCollection struct {
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
		BackdropPath string `json:"backdrop_path"`
	} `json:"belongs_to_collection"`
	Budget int64 `json:"budget"`
	Genres []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            string  `json:"homepage"`
	ID                  int64   `json:"id"`
	IMDbID              string  `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float32 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ProductionCompanies []struct {
		Name          string `json:"name"`
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_31661_1"`
		Name      string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         int64  `json:"revenue"`
	Runtime         int    `json:"runtime"`
	SpokenLanguages []struct {
		Iso639_1 string `json:"iso_639_1"`
		Name     string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
	*MovieAlternativeTitlesAppend
	*MovieChangesAppend
	*MovieCreditsAppend
	*MovieExternalIDsAppend
	*MovieImagesAppend
	*MovieKeywordsAppend
	*MovieReleaseDatesAppend
	*MovieVideosAppend
	*MovieTranslationsAppend
	*MovieRecommendationsAppend
	*MovieSimilarAppend
	*MovieReviewsAppend
	*MovieListsAppend
}

// MovieAccountStates type is a struct for account states JSON response.
type MovieAccountStates struct {
	ID        int64           `json:"id"`
	Favorite  bool            `json:"favorite"`
	Rated     json.RawMessage `json:"rated"`
	Watchlist bool            `json:"watchlist"`
}

// MovieAlternativeTitles type is a struct for alternative titles JSON response.
type MovieAlternativeTitles struct {
	ID     int `json:"id,omitempty"`
	Titles []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Title     string `json:"title"`
		Type      string `json:"type"`
	} `json:"titles"`
}

// MovieAlternativeTitlesAppend type is a struct for alternative
// titles in append to response.
type MovieAlternativeTitlesAppend struct {
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
}

// MovieChanges type is a struct for changes JSON response.
type MovieChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID            json.RawMessage `json:"id"`
			Action        json.RawMessage `json:"action"`
			Time          json.RawMessage `json:"time"`
			Iso639_1      json.RawMessage `json:"iso_639_1"`
			Value         json.RawMessage `json:"value"`
			OriginalValue json.RawMessage `json:"original_value"`
		} `json:"items"`
	} `json:"changes"`
}

// MovieChangesAppend type is a struct for changes in append to response.
type MovieChangesAppend struct {
	Changes *MovieChanges `json:"changes,omitempty"`
}

// MovieCredits type is a struct for credits JSON response.
type MovieCredits struct {
	ID   int64 `json:"id,omitempty"`
	Cast []struct {
		CastID      int64  `json:"cast_id"`
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

// MovieCreditsAppend type is a struct for credits in append to response.
type MovieCreditsAppend struct {
	Credits struct {
		*MovieCredits
	} `json:"credits,omitempty"`
}

// MovieExternalIDs type is a struct for external ids JSON response.
type MovieExternalIDs struct {
	IMDbID      string `json:"imdb_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
	ID          int64  `json:"id,omitempty"`
}

// MovieExternalIDsAppend type is a struct for external ids in append to response.
type MovieExternalIDsAppend struct {
	*MovieExternalIDs `json:"external_ids,omitempty"`
}

// MovieImages type is a struct for images JSON response.
type MovieImages struct {
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

// MovieImagesAppend type is a struct for images in append to response.
type MovieImagesAppend struct {
	Images *MovieImages `json:"images,omitempty"`
}

// MovieKeywords type is a struct for keywords JSON response.
type MovieKeywords struct {
	ID       int64 `json:"id,omitempty"`
	Keywords []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"keywords"`
}

// MovieKeywordsAppend type is a struct for keywords in append to response.
type MovieKeywordsAppend struct {
	Keywords struct {
		*MovieKeywords
	} `json:"keywords,omitempty"`
}

// MovieReleaseDates type is a struct for release dates JSON response.
type MovieReleaseDates struct {
	ID      int64 `json:"id,omitempty"`
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

// MovieReleaseDatesAppend type is a struct for release dates in append to response.
type MovieReleaseDatesAppend struct {
	ReleaseDates *MovieReleaseDates `json:"release_dates,omitempty"`
}

// MovieVideos type is a struct for videos JSON response.
type MovieVideos struct {
	ID      int64 `json:"id,omitempty"`
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

// MovieVideosAppend type is a struct for videos in append to response.
type MovieVideosAppend struct {
	Videos struct {
		*MovieVideos
	} `json:"videos,omitempty"`
}

// MovieTranslations type is a struct for translations JSON response.
type MovieTranslations struct {
	ID           int64 `json:"id,omitempty"`
	Translations []struct {
		Iso639_1    string `json:"iso_639_1"`
		Iso3166_1   string `json:"iso_3166_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Title    string `json:"title"`
			Overview string `json:"overview"`
			Homepage string `json:"homepage"`
		} `json:"data"`
	} `json:"translations"`
}

// MovieTranslationsAppend type is a struct for translations in append to response.
type MovieTranslationsAppend struct {
	Translations *MovieTranslations `json:"translations,omitempty"`
}

// MovieRecommendations type is a struct for recommendations JSON response.
type MovieRecommendations struct {
	Page    int64 `json:"page"`
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
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// MovieRecommendationsAppend type is a struct for
// recommendations in append to response.
type MovieRecommendationsAppend struct {
	Recommendations *MovieRecommendations `json:"recommendations,omitempty"`
}

// MovieSimilar type is a struct for similar movies JSON response.
type MovieSimilar struct {
	*MovieRecommendations
}

// MovieSimilarAppend type is a struct for similar movies in append to response.
type MovieSimilarAppend struct {
	Similar *MovieSimilar `json:"similar,omitempty"`
}

// MovieReviews type is a struct for reviews JSON response.
type MovieReviews struct {
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

// MovieReviewsAppend type is a struct for reviews in append to response.
type MovieReviewsAppend struct {
	Reviews struct {
		*MovieReviews
	} `json:"reviews,omitempty"`
}

// MovieLists type is a struct for lists JSON response.
type MovieLists struct {
	ID      int64 `json:"id"`
	Page    int64 `json:"page"`
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
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// MovieListsAppend type is a struct for lists in append to response.
type MovieListsAppend struct {
	Lists *MovieLists `json:"lists,omitempty"`
}

// MovieLatest type is a struct for latest JSON response.
type MovieLatest struct {
	*MovieDetails
}

// MovieNowPlaying type is a struct for now playing JSON response.
type MovieNowPlaying struct {
	Page    int64 `json:"page"`
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
	Dates struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	} `json:"dates"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// MoviePopular type is a struct for popular JSON response.
type MoviePopular struct {
	Page    int64 `json:"page"`
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
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// MovieTopRated type is a struct for top rated JSON response.
type MovieTopRated struct {
	*MoviePopular
}

// MovieUpcoming type is a struct for upcoming JSON response.
type MovieUpcoming struct {
	*MovieNowPlaying
}

// GetMovieDetails get the primary information about a movie.
//
// https://developers.themoviedb.org/3/movies
//
func (c *Client) GetMovieDetails(id int, o map[string]string) (*MovieDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieDetails{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieAccountStates grab the following account states for a session:
//
// Movie rating.
//
// If it belongs to your watchlist.
//
// If it belongs to your favourite list.
//
// https://developers.themoviedb.org/3/movies/get-movie-account-states
//
func (c *Client) GetMovieAccountStates(id int, o map[string]string) (*MovieAccountStates, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/account_states?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieAccountStates{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieAlternativeTitles get all of the alternative titles for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-alternative-titles
func (c *Client) GetMovieAlternativeTitles(id int, o map[string]string) (*MovieAlternativeTitles, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/alternative_titles?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieAlternativeTitles{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieChanges get the changes for a movie.
//
// By default only the last 24 hours are returned.
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/movies/get-movie-changes
func (c *Client) GetMovieChanges(id int, o map[string]string) (*MovieChanges, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/changes?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieChanges{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieCredits get the cast and crew for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-credits
func (c *Client) GetMovieCredits(id int, o map[string]string) (*MovieCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/credits?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieCredits{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieExternalIDs get the external ids for a movie.
//
// We currently support the following external sources.
//
// Media Databases: IMDb ID.
//
// Social IDs: Facebook, Instagram and Twitter.
//
// https://developers.themoviedb.org/3/movies/get-movie-external-ids
func (c *Client) GetMovieExternalIDs(id int, o map[string]string) (*MovieExternalIDs, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/external_ids?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieExternalIDs{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieImages get the images that belong to a movie.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter. This should be a comma
// separated value like so: include_image_language=en,null.
//
// https://developers.themoviedb.org/3/movies/get-movie-images
func (c *Client) GetMovieImages(id int, o map[string]string) (*MovieImages, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/images?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieImages{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieKeywords get the keywords that have been added to a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-keywords
func (c *Client) GetMovieKeywords(id int) (*MovieKeywords, error) {
	tmdbURL := fmt.Sprintf("%s%s%d/keywords?api_key=%s", baseURL, movieURL, id, c.APIKey)
	m := MovieKeywords{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieReleaseDates get the release date along with the certification for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-release-dates
func (c *Client) GetMovieReleaseDates(id int) (*MovieReleaseDates, error) {
	tmdbURL := fmt.Sprintf("%s%s%d/release_dates?api_key=%s", baseURL, movieURL, id, c.APIKey)
	m := MovieReleaseDates{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieVideos get the videos that have been added to a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-videos
func (c *Client) GetMovieVideos(id int, o map[string]string) (*MovieVideos, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/videos?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieVideos{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieTranslations get a list of translations that have been created for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-translations
func (c *Client) GetMovieTranslations(id int, o map[string]string) (*MovieTranslations, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/translations?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieTranslations{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieRecommendations get a list of recommended movies for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-recommendations
func (c *Client) GetMovieRecommendations(id int, o map[string]string) (*MovieRecommendations, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/recommendations?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieRecommendations{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieSimilar get a list of similar movies.
//
// This is not the same as the "Recommendation" system you see on the website.
// These items are assembled by looking at keywords and genres.
//
// https://developers.themoviedb.org/3/movies/get-similar-movies
func (c *Client) GetMovieSimilar(id int, o map[string]string) (*MovieSimilar, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/similar?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieSimilar{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieReviews get the user reviews for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-reviews
func (c *Client) GetMovieReviews(id int, o map[string]string) (*MovieReviews, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/reviews?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieReviews{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieLists get a list of lists that this movie belongs to.
//
// https://developers.themoviedb.org/3/movies/get-movie-lists
func (c *Client) GetMovieLists(id int, o map[string]string) (*MovieLists, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/lists?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)
	m := MovieLists{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieLatest get the most newly created movie.
//
// This is a live response and will continuously change.
//
// https://developers.themoviedb.org/3/movies/get-latest-movie
func (c *Client) GetMovieLatest(o map[string]string) (*MovieLatest, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%slatest?api_key=%s%s", baseURL, movieURL, c.APIKey, options)
	m := MovieLatest{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieNowPlaying get a list of movies in theatres.
//
// This is a release type query that looks for all movies that
// have a release type of 2 or 3 within the specified date range.
//
// You can optionally specify a region prameter which will narrow
// the search to only look for theatrical release dates within the
// specified country.
//
// https://developers.themoviedb.org/3/movies/get-now-playing
func (c *Client) GetMovieNowPlaying(o map[string]string) (*MovieNowPlaying, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%snow_playing?api_key=%s%s", baseURL, movieURL, c.APIKey, options)
	m := MovieNowPlaying{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMoviePopular get a list of the current popular movies on TMDb.
//
// This list updates daily.
//
// https://developers.themoviedb.org/3/movies/get-popular-movies
func (c *Client) GetMoviePopular(o map[string]string) (*MoviePopular, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%spopular?api_key=%s%s", baseURL, movieURL, c.APIKey, options)
	m := MoviePopular{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieTopRated get the top rated movies on TMDb.
//
// https://developers.themoviedb.org/3/movies/get-top-rated-movies
func (c *Client) GetMovieTopRated(o map[string]string) (*MovieTopRated, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%stop_rated?api_key=%s%s", baseURL, movieURL, c.APIKey, options)
	m := MovieTopRated{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMovieUpcoming get a list of upcoming movies in theatres.
//
// This is a release type query that looks for all movies that
// have a release type of 2 or 3 within the specified date range.
//
// You can optionally specify a region prameter which will narrow
// the search to only look for theatrical release dates within
// the specified country.
//
// https://developers.themoviedb.org/3/movies/get-upcoming
func (c *Client) GetMovieUpcoming(o map[string]string) (*MovieUpcoming, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%supcoming?api_key=%s%s", baseURL, movieURL, c.APIKey, options)
	m := MovieUpcoming{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// TODO: Rate Movie (POST Request) and Delete Rating (DELETE Request)
