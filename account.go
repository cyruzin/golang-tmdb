package tmdb

import (
	"fmt"
)

// AccountDetails type is a struct for details JSON response.
type AccountDetails struct {
	Avatar struct {
		Gravatar struct {
			Hash string `json:"hash"`
		} `json:"gravatar"`
	} `json:"avatar"`
	ID           int64  `json:"id"`
	Iso639_1     string `json:"iso_639_1"`
	Iso3166_1    string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
}

// AccountCreatedLists type is a struct for created lists JSON response.
type AccountCreatedLists struct {
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

// AccountFavoriteMovies type is a struct for favorite movies JSON response.
type AccountFavoriteMovies struct {
	Page    int64 `json:"page"`
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
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int64   `json:"vote_count"`
	} `json:"results"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// AccountFavoriteTVShows type is a struct for favorite tv shows JSON response.
type AccountFavoriteTVShows struct {
	Page    int64 `json:"page"`
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
		VoteAverage      float64  `json:"vote_average"`
		VoteCount        int64    `json:"vote_count"`
	} `json:"results"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// AccountFavorite type is a struct for a favorite movie or TV shows JSON request.
type AccountFavorite struct {
	MediaType string `json:"media_type"`
	MediaID   int64  `json:"media_id"`
	Favorite  bool   `json:"favorite"`
}

// AccountRatedMovies type is a struct for rated movies JSON response.
type AccountRatedMovies struct {
	*AccountFavoriteMovies
}

// GetAccountDetails get your account details.
//
// https://developers.themoviedb.org/3/account/get-account-details
func (c *Client) GetAccountDetails() (*AccountDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s/account?api_key=%s&session_id=%s",
		baseURL, c.apiKey, c.sessionID,
	)
	details := AccountDetails{}
	err := c.get(tmdbURL, &details)
	if err != nil {
		return nil, err
	}
	return &details, nil
}

// GetCreatedLists get all of the lists created by an account.
// Will invlude private lists if you are the owner.
//
// https://developers.themoviedb.org/3/account/get-created-lists
func (c *Client) GetCreatedLists(id int, o map[string]string) (*AccountCreatedLists, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/lists?api_key=%s&session_id=%s%s",
		baseURL, accountURL, id, c.apiKey, c.sessionID, options,
	)
	createdLists := AccountCreatedLists{}
	err := c.get(tmdbURL, &createdLists)
	if err != nil {
		return nil, err
	}
	return &createdLists, nil
}

// GetFavoriteMovies get the list of your favorite movies.
//
// https://developers.themoviedb.org/3/account/get-favorite-movies
func (c *Client) GetFavoriteMovies(id int, o map[string]string) (*AccountFavoriteMovies, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/favorite/movies?api_key=%s&session_id=%s%s",
		baseURL, accountURL, id, c.apiKey, c.sessionID, options,
	)
	favoriteMovies := AccountFavoriteMovies{}
	err := c.get(tmdbURL, &favoriteMovies)
	if err != nil {
		return nil, err
	}
	return &favoriteMovies, nil
}

// GetFavoriteTVShows get the list of your favorite TV shows.
//
// https://developers.themoviedb.org/3/account/get-favorite-tv-shows
func (c *Client) GetFavoriteTVShows(id int, o map[string]string) (*AccountFavoriteTVShows, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/favorite/tv?api_key=%s&session_id=%s%s",
		baseURL, accountURL, id, c.apiKey, c.sessionID, options,
	)
	favoriteTVShows := AccountFavoriteTVShows{}
	err := c.get(tmdbURL, &favoriteTVShows)
	if err != nil {
		return nil, err
	}
	return &favoriteTVShows, nil
}

// AccountMarkAsFavorite this method allows you to mark a movie
// or TV show as a favorite item.
//
// https://developers.themoviedb.org/3/account/mark-as-favorite
func (c *Client) AccountMarkAsFavorite(id int, title *AccountFavorite) (*Response, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/favorite?api_key=%s&session_id=%s",
		baseURL, accountURL, id, c.apiKey, c.sessionID,
	)
	markAsFavorite := Response{}
	err := c.request(tmdbURL, title, "POST", &markAsFavorite)
	if err != nil {
		return nil, err
	}
	return &markAsFavorite, nil
}

// GetRatedMovies get a list of all the movies you have rated.
//
// https://developers.themoviedb.org/3/account/get-rated-movies
func (c *Client) GetRatedMovies(id int, o map[string]string) (*AccountRatedMovies, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rated/movies?api_key=%s&session_id=%s%s",
		baseURL, accountURL, id, c.apiKey, c.sessionID, options,
	)
	ratedMovies := AccountRatedMovies{}
	err := c.get(tmdbURL, &ratedMovies)
	if err != nil {
		return nil, err
	}
	return &ratedMovies, nil
}
