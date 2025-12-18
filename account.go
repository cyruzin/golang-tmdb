package tmdb

import (
	"fmt"
	"net/http"
)

// GetAccountDetails get your account details.
//
// https://developer.themoviedb.org/reference/account-details
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

// MarkAsFavorite this method allows you to mark a movie
// or TV show as a favorite item.
//
// https://developer.themoviedb.org/reference/account-add-favorite
func (c *Client) MarkAsFavorite(
	id int,
	body AccountFavorite,
) (*Response, error) {
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
		body,
		http.MethodPost,
		&markAsFavorite,
	); err != nil {
		return nil, err
	}
	return &markAsFavorite, nil
}

// AddToWatchlist add a movie or TV show to your watchlist.
//
// https://developer.themoviedb.org/reference/account-add-to-watchlist
func (c *Client) AddToWatchlist(
	id int,
	body AccountWatchlist,
) (*Response, error) {
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
		body,
		http.MethodPost,
		&addToWatchlist,
	); err != nil {
		return nil, err
	}
	return &addToWatchlist, nil
}

// GetFavoriteMovies get the list of your favorite movies.
//
// https://developer.themoviedb.org/reference/account-get-favorites
func (c *Client) GetFavoriteMovies(
	id int,
	urlOptions map[string]string,
) (*AccountFavoriteMovies, error) {
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

// GetFavoriteTVShows get the list of your favorite TV shows.
//
// https://developer.themoviedb.org/reference/account-favorite-tv
func (c *Client) GetFavoriteTVShows(
	id int,
	urlOptions map[string]string,
) (*AccountFavoriteTVShows, error) {
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

// GetCreatedLists get all of the lists created by an account.
// Will invlude private lists if you are the owner.
//
// https://developer.themoviedb.org/reference/account-lists
func (c *Client) GetCreatedLists(
	id int,
	urlOptions map[string]string,
) (*AccountCreatedLists, error) {
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

// GetRatedMovies get a list of all the movies you have rated.
//
// https://developer.themoviedb.org/reference/account-rated-movies
func (c *Client) GetRatedMovies(
	id int,
	urlOptions map[string]string,
) (*AccountRatedMovies, error) {
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

// GetRatedTVShows get a list of all the TV shows you have rated.
//
// https://developer.themoviedb.org/reference/account-rated-tv
func (c *Client) GetRatedTVShows(
	id int,
	urlOptions map[string]string,
) (*AccountRatedTVShows, error) {
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

// GetRatedTVEpisodes get a list of all the TV episodes you have rated.
//
// https://developer.themoviedb.org/reference/account-rated-tv-episodes
func (c *Client) GetRatedTVEpisodes(
	id int,
	urlOptions map[string]string,
) (*AccountRatedTVEpisodes, error) {
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

// GetMovieWatchlist get a list of all the movies you have added to your watchlist.
//
// https://developer.themoviedb.org/reference/account-watchlist-movies
func (c *Client) GetMovieWatchlist(
	id int,
	urlOptions map[string]string,
) (*AccountMovieWatchlist, error) {
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

// GetTVShowsWatchlist get a list of all the TV shows you have added to your watchlist.
//
// https://developer.themoviedb.org/reference/account-watchlist-tv
func (c *Client) GetTVShowsWatchlist(
	id int,
	urlOptions map[string]string,
) (*AccountTVShowsWatchlist, error) {
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
