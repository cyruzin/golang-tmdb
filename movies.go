package tmdb

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
)

// MovieDetails type is a struct for movie details JSON response.
type MovieDetails struct {
	Adult               bool                `json:"adult"`
	BackdropPath        string              `json:"backdrop_path"`
	BelongsToCollection BelongsToCollection `json:"belongs_to_collection"`
	Budget              int64               `json:"budget"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	ID                  int64               `json:"id"`
	IMDbID              string              `json:"imdb_id"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalTitle       string              `json:"original_title"`
	Overview            string              `json:"overview"`
	Popularity          float32             `json:"popularity"`
	PosterPath          string              `json:"poster_path"`
	OriginCountry       []string            `json:"origin_country"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	ReleaseDate         string              `json:"release_date"`
	Revenue             int64               `json:"revenue"`
	Runtime             int                 `json:"runtime"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Title               string              `json:"title"`
	Video               bool                `json:"video"`
	VoteMetrics
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
	Videos *VideoResults `json:"videos"`
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
	WatchProviders *WatchProviderResults `json:"watch/providers"`
}

// GetMovieDetails get the primary information about a movie.
//
// https://developers.themoviedb.org/3/movies
func (c *Client) GetMovieDetails(id int, urlOptions map[string]string) (*MovieDetails, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id))
	var movieDetails MovieDetails
	if err := c.get(url, urlOptions, &movieDetails); err != nil {
		return nil, err
	}
	return &movieDetails, nil
}

// MovieAccountStates type is a struct for account states JSON response.
type MovieAccountStates struct {
	ID        int64           `json:"id"`
	Favorite  bool            `json:"favorite"`
	Rated     json.RawMessage `json:"rated"`
	Watchlist bool            `json:"watchlist"`
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
func (c *Client) GetMovieAccountStates(id int, urlOptions map[string]string) (*MovieAccountStates, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "account_states")
	var movieAccountStates MovieAccountStates
	if err := c.get(url, urlOptions, &movieAccountStates); err != nil {
		return nil, err
	}
	return &movieAccountStates, nil
}

// MovieAlternativeTitles type is a struct for alternative titles JSON response.
type MovieAlternativeTitles struct {
	ID     int                `json:"id"`
	Titles []AlternativeTitle `json:"titles"`
}

// GetMovieAlternativeTitles get all of the alternative titles for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-alternative-titles
func (c *Client) GetMovieAlternativeTitles(id int, urlOptions map[string]string) (*MovieAlternativeTitles, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "alternative_titles")
	var movieAlternativeTitles MovieAlternativeTitles
	if err := c.get(url, urlOptions, &movieAlternativeTitles); err != nil {
		return nil, err
	}
	return &movieAlternativeTitles, nil
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

// GetMovieChanges get the changes for a movie.
//
// By default only the last 24 hours are returned.
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/movies/get-movie-changes
func (c *Client) GetMovieChanges(id int, urlOptions map[string]string) (*MovieChanges, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "changes")
	var movieChanges MovieChanges
	if err := c.get(url, urlOptions, &movieChanges); err != nil {
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
func (c *Client) GetMovieCredits(id int, urlOptions map[string]string) (*MovieCredits, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "credits")
	var movieCredits MovieCredits
	if err := c.get(url, urlOptions, &movieCredits); err != nil {
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
	WikiDataID  string `json:"wikidata_id,omitempty"`
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
func (c *Client) GetMovieExternalIDs(id int, urlOptions map[string]string) (*MovieExternalIDs, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "external_ids")
	var movieExternalIDs MovieExternalIDs
	if err := c.get(url, urlOptions, &movieExternalIDs); err != nil {
		return nil, err
	}
	return &movieExternalIDs, nil
}

// MovieImage type is a struct for a single image.
type MovieImage struct {
	ImageBase
	Iso639_1 string `json:"iso_639_1"`
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
func (c *Client) GetMovieImages(id int, urlOptions map[string]string) (*MovieImages, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "images")
	var movieImages MovieImages
	if err := c.get(url, urlOptions, &movieImages); err != nil {
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
func (c *Client) GetMovieKeywords(id int) (*MovieKeywords, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "keywords")
	var movieKeywords MovieKeywords
	if err := c.get(url, nil, &movieKeywords); err != nil {
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
func (c *Client) GetMovieReleaseDates(id int) (*MovieReleaseDates, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "release_dates")
	var movieReleaseDates MovieReleaseDates
	if err := c.get(url, nil, &movieReleaseDates); err != nil {
		return nil, err
	}
	return &movieReleaseDates, nil
}

// GetMovieVideos get the videos that have been added to a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-videos
func (c *Client) GetMovieVideos(id int, urlOptions map[string]string) (*VideoResults, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "videos")
	var movieVideos VideoResults
	if err := c.get(url, urlOptions, &movieVideos); err != nil {
		return nil, err
	}
	return &movieVideos, nil
}

// GetMovieWatchProviders get a list of the availabilities per country by provider for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-watch-providers
func (c *Client) GetMovieWatchProviders(id int, urlOptions map[string]string) (*WatchProviderResults, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "watch", "providers")
	var movieWatchProviders WatchProviderResults
	if err := c.get(url, urlOptions, &movieWatchProviders); err != nil {
		return nil, err
	}
	return &movieWatchProviders, nil
}

// MovieTranslations type is a struct for translations JSON response.
type MovieTranslations struct {
	ID           int64         `json:"id,omitempty"`
	Translations []Translation `json:"translations"`
}

// GetMovieTranslations get a list of translations that have been created for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-translations
func (c *Client) GetMovieTranslations(id int, urlOptions map[string]string) (*MovieTranslations, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "translations")
	var movieTranslations MovieTranslations
	if err := c.get(url, urlOptions, &movieTranslations); err != nil {
		return nil, err
	}
	return &movieTranslations, nil
}

// MovieRecommendations type is a struct for recommendations JSON response.
type MovieRecommendations struct {
	*MovieRecommendationsResults
	PaginatedResultsMeta
}

// GetMovieRecommendations get a list of recommended movies for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-recommendations
func (c *Client) GetMovieRecommendations(id int, urlOptions map[string]string) (*MovieRecommendations, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "recommendations")
	var movieRecommendations MovieRecommendations
	if err := c.get(url, urlOptions, &movieRecommendations); err != nil {
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
func (c *Client) GetMovieSimilar(id int, urlOptions map[string]string) (*MovieSimilar, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "similar")
	var movieSimilar MovieSimilar
	if err := c.get(url, urlOptions, &movieSimilar); err != nil {
		return nil, err
	}
	return &movieSimilar, nil
}

// MovieReviews type is a struct for reviews JSON response.
type MovieReviews struct {
	ID int64 `json:"id,omitempty"`
	*MovieReviewsResults
	PaginatedResultsMeta
}

// GetMovieReviews get the user reviews for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-reviews
func (c *Client) GetMovieReviews(id int, urlOptions map[string]string) (*MovieReviews, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "reviews")
	var movieReviews MovieReviews
	if err := c.get(url, urlOptions, &movieReviews); err != nil {
		return nil, err
	}
	return &movieReviews, nil
}

// MovieLists type is a struct for lists JSON response.
type MovieLists struct {
	ID int64 `json:"id"`
	*MovieListsResults
	PaginatedResultsMeta
}

// GetMovieLists get a list of lists that this movie belongs to.
//
// https://developers.themoviedb.org/3/movies/get-movie-lists
func (c *Client) GetMovieLists(id int, urlOptions map[string]string) (*MovieLists, error) {
	url := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "lists")
	var movieLists MovieLists
	if err := c.get(url, urlOptions, &movieLists); err != nil {
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
func (c *Client) GetMovieLatest(urlOptions map[string]string) (*MovieLatest, error) {
	url := c.baseURL.JoinPath(movieURL, "latest")
	var movieLatest MovieLatest
	if err := c.get(url, urlOptions, &movieLatest); err != nil {
		return nil, err
	}
	return &movieLatest, nil
}

// MovieNowPlaying type is a struct for now playing JSON response.
type MovieNowPlaying struct {
	*MovieNowPlayingResults
	Dates struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	} `json:"dates"`
	PaginatedResultsMeta
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
func (c *Client) GetMovieNowPlaying(urlOptions map[string]string) (*MovieNowPlaying, error) {
	url := c.baseURL.JoinPath(movieURL, "now_playing")
	var movieNowPlaying MovieNowPlaying
	if err := c.get(url, urlOptions, &movieNowPlaying); err != nil {
		return nil, err
	}
	return &movieNowPlaying, nil
}

// MoviePopular type is a struct for popular JSON response.
type MoviePopular struct {
	*MoviePopularResults
	PaginatedResultsMeta
}

// GetMoviePopular get a list of the current popular movies on TMDb.
//
// This list updates daily.
//
// https://developers.themoviedb.org/3/movies/get-popular-movies
func (c *Client) GetMoviePopular(urlOptions map[string]string) (*MoviePopular, error) {
	url := c.baseURL.JoinPath(movieURL, "popular")
	var moviePopular MoviePopular
	if err := c.get(url, urlOptions, &moviePopular); err != nil {
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
func (c *Client) GetMovieTopRated(urlOptions map[string]string) (*MovieTopRated, error) {
	url := c.baseURL.JoinPath(movieURL, "top_rated")
	var movieTopRated MovieTopRated
	if err := c.get(url, urlOptions, &movieTopRated); err != nil {
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
func (c *Client) GetMovieUpcoming(urlOptions map[string]string) (*MovieUpcoming, error) {
	url := c.baseURL.JoinPath(movieURL, "upcoming")
	var movieUpcoming MovieUpcoming
	if err := c.get(url, urlOptions, &movieUpcoming); err != nil {
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
func (c *Client) PostMovieRating(id int, rating float32, urlOptions map[string]string) (*Response, error) {
	u := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "rating")
	q := u.Query()
	q.Add("session_id", c.sessionID)
	for key, value := range urlOptions {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()
	body := struct {
		Value float32 `json:"value"`
	}{Value: rating}
	response := Response{}
	if err := c.request(
		u.String(),
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
func (c *Client) DeleteMovieRating(id int, urlOptions map[string]string) (*Response, error) {
	u := c.baseURL.JoinPath(movieURL, fmt.Sprintf("%d", id), "rating")
	q := u.Query()
	q.Add("session_id", c.sessionID)
	for key, value := range urlOptions {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()
	response := Response{}
	if err := c.request(
		u.String(),
		[]byte{},
		http.MethodDelete,
		&response,
	); err != nil {
		return nil, err
	}
	return &response, nil
}
