package tmdb

import "fmt"

// GetGuestSessionRatedMovies get the rated movies for a guest session.
//
// https://developer.themoviedb.org/reference/guest-session-rated-movies
func (c *Client) GetGuestSessionRatedMovies(
	id string,
	urlOptions map[string]string,
) (*GuestSessionRatedMovies, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%s/rated/movies?api_key=%s%s",
		baseURL,
		guestSessionURL,
		id,
		c.apiKey,
		options,
	)
	guestSessionRatedMovies := GuestSessionRatedMovies{}
	if err := c.get(tmdbURL, &guestSessionRatedMovies); err != nil {
		return nil, err
	}
	return &guestSessionRatedMovies, nil
}

// GetGuestSessionRatedTVShows get the rated TV shows for a guest session.
//
// https://developer.themoviedb.org/reference/guest-session-rated-tv
func (c *Client) GetGuestSessionRatedTVShows(
	id string,
	urlOptions map[string]string,
) (*GuestSessionRatedTVShows, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%s/rated/tv?api_key=%s%s",
		baseURL,
		guestSessionURL,
		id,
		c.apiKey,
		options,
	)
	guestSessionRatedTVShows := GuestSessionRatedTVShows{}
	if err := c.get(tmdbURL, &guestSessionRatedTVShows); err != nil {
		return nil, err
	}
	return &guestSessionRatedTVShows, nil
}

// GetGuestSessionRatedTVEpisodes get the rated TV episodes for a guest session.
//
// https://developer.themoviedb.org/reference/guest-session-rated-tv-episodes
func (c *Client) GetGuestSessionRatedTVEpisodes(
	id string,
	urlOptions map[string]string,
) (*GuestSessionRatedTVEpisodes, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%s/rated/tv/episodes?api_key=%s%s",
		baseURL,
		guestSessionURL,
		id,
		c.apiKey,
		options,
	)
	guestSessionRatedTVEpisodes := GuestSessionRatedTVEpisodes{}
	if err := c.get(tmdbURL, &guestSessionRatedTVEpisodes); err != nil {
		return nil, err
	}
	return &guestSessionRatedTVEpisodes, nil
}
