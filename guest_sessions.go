package tmdb

// GuestSessionRatedMovies type is a struct for rated movies JSON response.
type GuestSessionRatedMovies struct {
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		PosterPath       string  `json:"poster_path"`
		Popularity       float32 `json:"popularity"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		Rating           float32 `json:"rating"`
		VoteMetrics
	} `json:"results"`
	PaginatedResultsMeta
}

// GetGuestSessionRatedMovies get the rated movies for a guest session.
//
// https://developers.themoviedb.org/3/guest-sessions/get-guest-session-rated-movies
func (c *Client) GetGuestSessionRatedMovies(id string, urlOptions map[string]string) (*GuestSessionRatedMovies, error) {
	url := c.baseURL.JoinPath(guestSessionPath, id, "rated", "movies")
	var guestSessionRatedMovies GuestSessionRatedMovies
	if err := c.get(url, urlOptions, &guestSessionRatedMovies); err != nil {
		return nil, err
	}
	return &guestSessionRatedMovies, nil
}

// GuestSessionRatedTVShows type is a struct for rated tv shows JSON response.
type GuestSessionRatedTVShows struct {
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
		Popularity       float32  `json:"popularity"`
		Name             string   `json:"name"`
		Rating           float32  `json:"rating"`
		VoteMetrics
	} `json:"results"`
	PaginatedResultsMeta
}

// GetGuestSessionRatedTVShows get the rated TV shows for a guest session.
//
// https://developers.themoviedb.org/3/guest-sessions/get-guest-session-rated-tv-shows
func (c *Client) GetGuestSessionRatedTVShows(id string, urlOptions map[string]string) (*GuestSessionRatedTVShows, error) {
	url := c.baseURL.JoinPath(guestSessionPath, id, "rated", "tv")
	var guestSessionRatedTVShows GuestSessionRatedTVShows
	if err := c.get(url, urlOptions, &guestSessionRatedTVShows); err != nil {
		return nil, err
	}
	return &guestSessionRatedTVShows, nil
}

// GuestSessionRatedTVEpisodes type is a struct for rated tv episodes JSON response.
type GuestSessionRatedTVEpisodes struct {
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
	PaginatedResultsMeta
}

// GetGuestSessionRatedTVEpisodes get the rated TV episodes for a guest session.
//
// https://developers.themoviedb.org/3/guest-sessions/get-gest-session-rated-tv-episodes
func (c *Client) GetGuestSessionRatedTVEpisodes(id string, urlOptions map[string]string) (*GuestSessionRatedTVEpisodes, error) {
	url := c.baseURL.JoinPath(guestSessionPath, id, "rated", "tv", "episodes")
	var guestSessionRatedTVEpisodes GuestSessionRatedTVEpisodes
	if err := c.get(url, urlOptions, &guestSessionRatedTVEpisodes); err != nil {
		return nil, err
	}
	return &guestSessionRatedTVEpisodes, nil
}
