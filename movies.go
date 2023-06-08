package tmdb

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
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
		Iso3166_1 string `json:"iso_3166_1"`
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
	*MovieWatchProvidersAppend
}

// MovieAlternativeTitlesAppend type is a struct for alternative
// titles in append to response.
type MovieAlternativeTitlesAppend struct {
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles,omitempty"`
}

// MovieChangesAppend type is a struct for changes in append to response.
type MovieChangesAppend struct {
	Changes *MovieChanges `json:"changes,omitempty"`
}

// MovieCreditsAppend type is a struct for credits in append to response.
type MovieCreditsAppend struct {
	Credits struct {
		*MovieCredits
	} `json:"credits,omitempty"`
}

// MovieExternalIDsAppend type is a struct for external ids in append to response.
type MovieExternalIDsAppend struct {
	*MovieExternalIDs `json:"external_ids,omitempty"`
}

// MovieImagesAppend type is a struct for images in append to response.
type MovieImagesAppend struct {
	Images *MovieImages `json:"images,omitempty"`
}

// MovieReleaseDatesAppend type is a struct for release dates in append to response.
type MovieReleaseDatesAppend struct {
	ReleaseDates *MovieReleaseDates `json:"release_dates,omitempty"`
}

// MovieVideosAppend type is a struct for videos in append to response.
type MovieVideosAppend struct {
	Videos struct {
		*MovieVideos
	} `json:"videos,omitempty"`
}

// MovieTranslationsAppend type is a struct for translations in append to response.
type MovieTranslationsAppend struct {
	Translations *MovieTranslations `json:"translations,omitempty"`
}

// MovieRecommendationsAppend type is a struct for
// recommendations in append to response.
type MovieRecommendationsAppend struct {
	Recommendations *MovieRecommendations `json:"recommendations,omitempty"`
}

// MovieSimilarAppend type is a struct for similar movies in append to response.
type MovieSimilarAppend struct {
	Similar *MovieSimilar `json:"similar,omitempty"`
}

// MovieReviewsAppend type is a struct for reviews in append to response.
type MovieReviewsAppend struct {
	Reviews struct {
		*MovieReviews
	} `json:"reviews,omitempty"`
}

// MovieListsAppend type is a struct for lists in append to response.
type MovieListsAppend struct {
	Lists *MovieLists `json:"lists,omitempty"`
}

// MovieKeywordsAppend type is a struct for keywords in append to response.
type MovieKeywordsAppend struct {
	Keywords struct {
		*MovieKeywords
	} `json:"keywords,omitempty"`
}

// MovieWatchProvidersAppend type is a struct for
// watch/providers in append to response.
type MovieWatchProvidersAppend struct {
	WatchProviders *MovieWatchProviders `json:"watch/providers,omitempty"`
}

// GetMovieDetails get the primary information about a movie.
//
// https://developers.themoviedb.org/3/movies
func (c *Client) GetMovieDetails(id int64, urlOptions map[string]string) (*MovieDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieDetails := MovieDetails{}
	if err := c.get(tmdbURL, &movieDetails); err != nil {
		return nil, err
	}
	return &movieDetails, nil
}

// MovieAccountStates type is a struct for account states JSON response.
type MovieAccountStates struct {
	ID        int64               `json:"id"`
	Favorite  bool                `json:"favorite"`
	Rated     jsoniter.RawMessage `json:"rated"`
	Watchlist bool                `json:"watchlist"`
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
func (c *Client) GetMovieAccountStates(id int64, urlOptions map[string]string) (*MovieAccountStates, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/account_states?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieAccountStates := MovieAccountStates{}
	if err := c.get(tmdbURL, &movieAccountStates); err != nil {
		return nil, err
	}
	return &movieAccountStates, nil
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

// GetMovieAlternativeTitles get all of the alternative titles for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-alternative-titles
func (c *Client) GetMovieAlternativeTitles(id int64, urlOptions map[string]string) (*MovieAlternativeTitles, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/alternative_titles?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieAlternativeTitles := MovieAlternativeTitles{}
	if err := c.get(tmdbURL, &movieAlternativeTitles); err != nil {
		return nil, err
	}
	return &movieAlternativeTitles, nil
}

// MovieChanges type is a struct for changes JSON response.
type MovieChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID            jsoniter.RawMessage `json:"id"`
			Action        jsoniter.RawMessage `json:"action"`
			Time          jsoniter.RawMessage `json:"time"`
			Iso639_1      jsoniter.RawMessage `json:"iso_639_1"`
			Value         jsoniter.RawMessage `json:"value"`
			OriginalValue jsoniter.RawMessage `json:"original_value"`
		} `json:"items"`
	} `json:"changes"`
}

// GetMovieChanges get the changes for a movie.
//
// By default only the last 24 hours are returned.
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/movies/get-movie-changes
func (c *Client) GetMovieChanges(id int64, urlOptions map[string]string) (*MovieChanges, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/changes?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieChanges := MovieChanges{}
	if err := c.get(tmdbURL, &movieChanges); err != nil {
		return nil, err
	}
	return &movieChanges, nil
}

// MovieCredits type is a struct for credits JSON response.
type MovieCredits struct {
	ID   int64 `json:"id,omitempty"`
	Cast []struct {
		Adult              bool    `json:"adult"`
		CastID             int64   `json:"cast_id"`
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
		Adult              bool    `json:"adult"`
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

// GetMovieCredits get the cast and crew for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-credits
func (c *Client) GetMovieCredits(id int64, urlOptions map[string]string) (*MovieCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf("%s%s%d/credits?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieCredits := MovieCredits{}
	if err := c.get(tmdbURL, &movieCredits); err != nil {
		return nil, err
	}
	return &movieCredits, nil
}

// MovieExternalIDs type is a struct for external ids JSON response.
type MovieExternalIDs struct {
	IMDbID      string `json:"imdb_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
	ID          int64  `json:"id,omitempty"`
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
func (c *Client) GetMovieExternalIDs(id int64, urlOptions map[string]string) (*MovieExternalIDs, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/external_ids?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieExternalIDs := MovieExternalIDs{}
	if err := c.get(tmdbURL, &movieExternalIDs); err != nil {
		return nil, err
	}
	return &movieExternalIDs, nil
}

// MovieImage type is a struct for a single image.
type MovieImage struct {
	AspectRatio float32 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	Iso639_1    string  `json:"iso_639_1"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
	Width       int     `json:"width"`
}

// MovieImages type is a struct for images JSON response.
type MovieImages struct {
	ID        int64        `json:"id,omitempty"`
	Backdrops []MovieImage `json:"backdrops"`
	Logos     []MovieImage `json:"logos"`
	Posters   []MovieImage `json:"posters"`
}

// GetMovieImages get the images that belong to a movie.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter. This should be a comma
// separated value like so: include_image_language=en,null.
//
// https://developers.themoviedb.org/3/movies/get-movie-images
func (c *Client) GetMovieImages(id int64, urlOptions map[string]string) (*MovieImages, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieImages := MovieImages{}
	if err := c.get(tmdbURL, &movieImages); err != nil {
		return nil, err
	}
	return &movieImages, nil
}

// MovieKeywords type is a struct for keywords JSON response.
type MovieKeywords struct {
	ID       int64 `json:"id,omitempty"`
	Keywords []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"keywords"`
}

// GetMovieKeywords get the keywords that have been added to a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-keywords
func (c *Client) GetMovieKeywords(id int64) (*MovieKeywords, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/keywords?api_key=%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
	)
	movieKeywords := MovieKeywords{}
	if err := c.get(tmdbURL, &movieKeywords); err != nil {
		return nil, err
	}
	return &movieKeywords, nil
}

// MovieReleaseDates type is a struct for release dates JSON response.
type MovieReleaseDates struct {
	ID int64 `json:"id,omitempty"`
	*MovieReleaseDatesResults
}

// GetMovieReleaseDates get the release date along with the certification for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-release-dates
func (c *Client) GetMovieReleaseDates(id int64) (*MovieReleaseDates, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/release_dates?api_key=%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
	)
	movieReleaseDates := MovieReleaseDates{}
	if err := c.get(tmdbURL, &movieReleaseDates); err != nil {
		return nil, err
	}
	return &movieReleaseDates, nil
}

// MovieVideos type is a struct for videos JSON response.
type MovieVideos struct {
	ID int64 `json:"id,omitempty"`
	*MovieVideosResults
}

// GetMovieVideos get the videos that have been added to a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-videos
func (c *Client) GetMovieVideos(id int64, urlOptions map[string]string) (*MovieVideos, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/videos?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieVideos := MovieVideos{}
	if err := c.get(tmdbURL, &movieVideos); err != nil {
		return nil, err
	}
	return &movieVideos, nil
}

// MovieWatchProviders type is a struct for watch/providers JSON response.
type MovieWatchProviders struct {
	ID int64 `json:"id,omitempty"`
	*MovieWatchProvidersResults
}

// GetMovieWatchProviders get a list of the availabilities per country by provider for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-watch-providers
func (c *Client) GetMovieWatchProviders(id int64, urlOptions map[string]string) (*MovieWatchProviders, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/watch/providers?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieWatchProviders := MovieWatchProviders{}
	if err := c.get(tmdbURL, &movieWatchProviders); err != nil {
		return nil, err
	}
	return &movieWatchProviders, nil
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
			Runtime  int    `json:"runtime"`
			Tagline  string `json:"tagline"`
			Homepage string `json:"homepage"`
		} `json:"data"`
	} `json:"translations"`
}

// GetMovieTranslations get a list of translations that have been created for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-translations
func (c *Client) GetMovieTranslations(id int64, urlOptions map[string]string) (*MovieTranslations, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/translations?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieTranslations := MovieTranslations{}
	if err := c.get(tmdbURL, &movieTranslations); err != nil {
		return nil, err
	}
	return &movieTranslations, nil
}

// MovieRecommendations type is a struct for recommendations JSON response.
type MovieRecommendations struct {
	Page int64 `json:"page"`
	*MovieRecommendationsResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetMovieRecommendations get a list of recommended movies for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-recommendations
func (c *Client) GetMovieRecommendations(id int64, urlOptions map[string]string) (*MovieRecommendations, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/recommendations?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieRecommendations := MovieRecommendations{}
	if err := c.get(tmdbURL, &movieRecommendations); err != nil {
		return nil, err
	}
	return &movieRecommendations, nil
}

// MovieSimilar type is a struct for similar movies JSON response.
type MovieSimilar struct {
	*MovieRecommendations
}

// GetMovieSimilar get a list of similar movies.
//
// This is not the same as the "Recommendation" system you see on the website.
// These items are assembled by looking at keywords and genres.
//
// https://developers.themoviedb.org/3/movies/get-similar-movies
func (c *Client) GetMovieSimilar(id int64, urlOptions map[string]string) (*MovieSimilar, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/similar?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieSimilar := MovieSimilar{}
	if err := c.get(tmdbURL, &movieSimilar); err != nil {
		return nil, err
	}
	return &movieSimilar, nil
}

// MovieReviews type is a struct for reviews JSON response.
type MovieReviews struct {
	ID   int64 `json:"id,omitempty"`
	Page int64 `json:"page"`
	*MovieReviewsResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetMovieReviews get the user reviews for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-reviews
func (c *Client) GetMovieReviews(id int64, urlOptions map[string]string) (*MovieReviews, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/reviews?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieReviews := MovieReviews{}
	if err := c.get(tmdbURL, &movieReviews); err != nil {
		return nil, err
	}
	return &movieReviews, nil
}

// MovieLists type is a struct for lists JSON response.
type MovieLists struct {
	ID   int64 `json:"id"`
	Page int64 `json:"page"`
	*MovieListsResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetMovieLists get a list of lists that this movie belongs to.
//
// https://developers.themoviedb.org/3/movies/get-movie-lists
func (c *Client) GetMovieLists(id int64, urlOptions map[string]string) (*MovieLists, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/lists?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieLists := MovieLists{}
	if err := c.get(tmdbURL, &movieLists); err != nil {
		return nil, err
	}
	return &movieLists, nil
}

// MovieLatest type is a struct for latest JSON response.
type MovieLatest struct {
	*MovieDetails
}

// GetMovieLatest get the most newly created movie.
//
// This is a live response and will continuously change.
//
// https://developers.themoviedb.org/3/movies/get-latest-movie
func (c *Client) GetMovieLatest(
	urlOptions map[string]string,
) (*MovieLatest, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%slatest?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	movieLastest := MovieLatest{}
	if err := c.get(tmdbURL, &movieLastest); err != nil {
		return nil, err
	}
	return &movieLastest, nil
}

// MovieNowPlaying type is a struct for now playing JSON response.
type MovieNowPlaying struct {
	Page int64 `json:"page"`
	*MovieNowPlayingResults
	Dates struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	} `json:"dates"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
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
func (c *Client) GetMovieNowPlaying(
	urlOptions map[string]string,
) (*MovieNowPlaying, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%snow_playing?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	movieNowPlaying := MovieNowPlaying{}
	if err := c.get(tmdbURL, &movieNowPlaying); err != nil {
		return nil, err
	}
	return &movieNowPlaying, nil
}

// MoviePopular type is a struct for popular JSON response.
type MoviePopular struct {
	Page int64 `json:"page"`
	*MoviePopularResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetMoviePopular get a list of the current popular movies on TMDb.
//
// This list updates daily.
//
// https://developers.themoviedb.org/3/movies/get-popular-movies
func (c *Client) GetMoviePopular(
	urlOptions map[string]string,
) (*MoviePopular, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%spopular?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	moviePopular := MoviePopular{}
	if err := c.get(tmdbURL, &moviePopular); err != nil {
		return nil, err
	}
	return &moviePopular, nil
}

// MovieTopRated type is a struct for top rated JSON response.
type MovieTopRated struct {
	*MoviePopular
}

// GetMovieTopRated get the top rated movies on TMDb.
//
// https://developers.themoviedb.org/3/movies/get-top-rated-movies
func (c *Client) GetMovieTopRated(
	urlOptions map[string]string,
) (*MovieTopRated, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stop_rated?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	movieTopRated := MovieTopRated{}
	if err := c.get(tmdbURL, &movieTopRated); err != nil {
		return nil, err
	}
	return &movieTopRated, nil
}

// MovieUpcoming type is a struct for upcoming JSON response.
type MovieUpcoming struct {
	*MovieNowPlaying
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
func (c *Client) GetMovieUpcoming(
	urlOptions map[string]string,
) (*MovieUpcoming, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%supcoming?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	movieUpcoming := MovieUpcoming{}
	if err := c.get(tmdbURL, &movieUpcoming); err != nil {
		return nil, err
	}
	return &movieUpcoming, nil
}

// PostMovieRating rate a movie.
//
// A valid session or guest session ID is required.
//
// You can read more about how this works:
// https://developers.themoviedb.org/3/authentication/how-do-i-generate-a-session-id
//
// https://developers.themoviedb.org/3/movies/rate-movie
func (c *Client) PostMovieRating(id int64, rating float32, urlOptions map[string]string) (*Response, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rating?api_key=%s&session_id=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	body := struct {
		Value float32 `json:"value"`
	}{Value: rating}
	response := Response{}
	if err := c.request(
		tmdbURL,
		body,
		http.MethodPost,
		&response,
	); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteMovieRating remove your rating for a movie.
//
// A valid session or guest session ID is required.
//
// You can read more about how this works:
// https://developers.themoviedb.org/3/authentication/how-do-i-generate-a-session-id
//
// https://developers.themoviedb.org/3/movies/delete-movie-rating
func (c *Client) DeleteMovieRating(id int64, urlOptions map[string]string) (*Response, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rating?api_key=%s&session_id=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	response := Response{}
	if err := c.request(
		tmdbURL,
		[]byte{},
		http.MethodDelete,
		&response,
	); err != nil {
		return nil, err
	}
	return &response, nil
}
