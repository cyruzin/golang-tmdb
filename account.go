package tmdb

import (
	"fmt"
	"net/http"
)

// AccountDetails type is a struct for details JSON response.
type AccountDetails struct {
	Avatar struct {
		Gravatar struct {
			Hash string `json:"hash"`
		} `json:"gravatar"`
		TMDB struct {
			AvatarPath string `json:"avatar_path"`
		} `json:"tmdb"`
	} `json:"avatar"`
	ID           int64  `json:"id"`
	Iso639_1     string `json:"iso_639_1"`
	Iso3166_1    string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
}

// GetAccountDetails get your account details.
//
// https://developers.themoviedb.org/3/account/get-account-details
func (c *Client) GetAccountDetails() (*AccountDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s/account?api_key=%s&session_id=%s",
		baseURL,
		c.apiKey,
		c.sessionID,
	)
	details := AccountDetails{}
	if err := c.get(tmdbURL, &details); err != nil {
		return nil, err
	}
	return &details, nil
}

// AccountCreatedLists type is a struct for created lists JSON response.
type AccountCreatedLists struct {
	Page int64 `json:"page"`
	*AccountCreatedListsResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetCreatedLists get all of the lists created by an account.
// Will invlude private lists if you are the owner.
//
// https://developers.themoviedb.org/3/account/get-created-lists
func (c *Client) GetCreatedLists(id int64, urlOptions map[string]string) (*AccountCreatedLists, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/lists?api_key=%s&session_id=%s%s",
		baseURL,
		accountURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	createdLists := AccountCreatedLists{}
	if err := c.get(tmdbURL, &createdLists); err != nil {
		return nil, err
	}
	return &createdLists, nil
}

// AccountFavoriteMovies type is a struct for favorite movies JSON response.
type AccountFavoriteMovies struct {
	Page int64 `json:"page"`
	*AccountFavoriteMoviesResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetFavoriteMovies get the list of your favorite movies.
//
// https://developers.themoviedb.org/3/account/get-favorite-movies
func (c *Client) GetFavoriteMovies(id int64, urlOptions map[string]string) (*AccountFavoriteMovies, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/favorite/movies?api_key=%s&session_id=%s%s",
		baseURL,
		accountURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	favoriteMovies := AccountFavoriteMovies{}
	if err := c.get(tmdbURL, &favoriteMovies); err != nil {
		return nil, err
	}
	return &favoriteMovies, nil
}

// AccountFavoriteTVShows type is a struct for favorite tv shows JSON response.
type AccountFavoriteTVShows struct {
	Page int64 `json:"page"`
	*AccountFavoriteTVShowsResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetFavoriteTVShows get the list of your favorite TV shows.
//
// https://developers.themoviedb.org/3/account/get-favorite-tv-shows
func (c *Client) GetFavoriteTVShows(id int64, urlOptions map[string]string) (*AccountFavoriteTVShows, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/favorite/tv?api_key=%s&session_id=%s%s",
		baseURL,
		accountURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	favoriteTVShows := AccountFavoriteTVShows{}
	if err := c.get(tmdbURL, &favoriteTVShows); err != nil {
		return nil, err
	}
	return &favoriteTVShows, nil
}

// AccountFavorite type is a struct for movies or TV shows
// favorite JSON request.
type AccountFavorite struct {
	MediaType string `json:"media_type"`
	MediaID   int64  `json:"media_id"`
	Favorite  bool   `json:"favorite"`
}

// MarkAsFavorite this method allows you to mark a movie
// or TV show as a favorite item.
//
// https://developers.themoviedb.org/3/account/mark-as-favorite
func (c *Client) MarkAsFavorite(id int64, title *AccountFavorite) (*Response, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/favorite?api_key=%s&session_id=%s",
		baseURL,
		accountURL,
		id,
		c.apiKey,
		c.sessionID,
	)
	markAsFavorite := Response{}
	if err := c.request(
		tmdbURL,
		title,
		http.MethodPost,
		&markAsFavorite,
	); err != nil {
		return nil, err
	}
	return &markAsFavorite, nil
}

// AccountRatedMovies type is a struct for rated movies JSON response.
type AccountRatedMovies struct {
	*AccountFavoriteMovies
}

// GetRatedMovies get a list of all the movies you have rated.
//
// https://developers.themoviedb.org/3/account/get-rated-movies
func (c *Client) GetRatedMovies(id int64, urlOptions map[string]string) (*AccountRatedMovies, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rated/movies?api_key=%s&session_id=%s%s",
		baseURL,
		accountURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	ratedMovies := AccountRatedMovies{}
	if err := c.get(tmdbURL, &ratedMovies); err != nil {
		return nil, err
	}
	return &ratedMovies, nil
}

// AccountRatedTVShows type is a struct for rated TV shows JSON response.
type AccountRatedTVShows struct {
	*AccountFavoriteTVShows
}

// GetRatedTVShows get a list of all the TV shows you have rated.
//
// https://developers.themoviedb.org/3/account/get-rated-tv-shows
func (c *Client) GetRatedTVShows(id int64, urlOptions map[string]string) (*AccountRatedTVShows, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rated/tv?api_key=%s&session_id=%s%s",
		baseURL,
		accountURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	ratedTVShows := AccountRatedTVShows{}
	if err := c.get(tmdbURL, &ratedTVShows); err != nil {
		return nil, err
	}
	return &ratedTVShows, nil
}

// AccountRatedTVEpisodes type is a struct for rated TV episodes JSON response.
type AccountRatedTVEpisodes struct {
	Page int64 `json:"page"`
	*AccountRatedTVEpisodesResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetRatedTVEpisodes get a list of all the TV episodes you have rated.
//
// https://developers.themoviedb.org/3/account/get-rated-tv-episodes
func (c *Client) GetRatedTVEpisodes(id int64, urlOptions map[string]string) (*AccountRatedTVEpisodes, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rated/tv/episodes?api_key=%s&session_id=%s%s",
		baseURL,
		accountURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	ratedTVEpisodes := AccountRatedTVEpisodes{}
	if err := c.get(tmdbURL, &ratedTVEpisodes); err != nil {
		return nil, err
	}
	return &ratedTVEpisodes, nil
}

// AccountMovieWatchlist type is a struct for movie watchlist JSON response.
type AccountMovieWatchlist struct {
	*AccountFavoriteMovies
}

// GetMovieWatchlist get a list of all the movies you have added to your watchlist.
//
// https://developers.themoviedb.org/3/account/get-movie-watchlist
func (c *Client) GetMovieWatchlist(id int64, urlOptions map[string]string) (*AccountMovieWatchlist, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/watchlist/movies?api_key=%s&session_id=%s%s",
		baseURL,
		accountURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	movieWatchlist := AccountMovieWatchlist{}
	if err := c.get(tmdbURL, &movieWatchlist); err != nil {
		return nil, err
	}
	return &movieWatchlist, nil
}

// AccountTVShowsWatchlist type is a struct for tv shows watchlist JSON response.
type AccountTVShowsWatchlist struct {
	*AccountFavoriteTVShows
}

// GetTVShowsWatchlist get a list of all the TV shows you have added to your watchlist.
//
// https://developers.themoviedb.org/3/account/get-tv-show-watchlist
func (c *Client) GetTVShowsWatchlist(id int64, urlOptions map[string]string) (*AccountTVShowsWatchlist, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/watchlist/tv?api_key=%s&session_id=%s%s",
		baseURL,
		accountURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	tvShowsWatchlist := AccountTVShowsWatchlist{}
	if err := c.get(tmdbURL, &tvShowsWatchlist); err != nil {
		return nil, err
	}
	return &tvShowsWatchlist, nil
}

// AccountWatchlist type is a struct for movies or TV shows
// watchlist JSON request.
type AccountWatchlist struct {
	MediaType string `json:"media_type"`
	MediaID   int64  `json:"media_id"`
	Watchlist bool   `json:"watchlist"`
}

// AddToWatchlist add a movie or TV show to your watchlist.
//
// https://developers.themoviedb.org/3/account/add-to-watchlist
func (c *Client) AddToWatchlist(id int64, title *AccountWatchlist) (*Response, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/watchlist?api_key=%s&session_id=%s",
		baseURL,
		accountURL,
		id,
		c.apiKey,
		c.sessionID,
	)
	addToWatchlist := Response{}
	if err := c.request(
		tmdbURL,
		title,
		http.MethodPost,
		&addToWatchlist,
	); err != nil {
		return nil, err
	}
	return &addToWatchlist, nil
}
