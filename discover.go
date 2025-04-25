package tmdb

import "fmt"

// DiscoverMovie type is a struct for movie JSON response.
type DiscoverMovie struct {
	PaginatedResultsMeta
	*DiscoverMovieResults
}

// GetDiscoverMovie discover movies by different types of data like
// average rating, number of votes, genres and certifications. You can
// get a valid list of certifications from the  method.
//
// Discover also supports a nice list of sort options.
// See below for all of the available options.
//
// Please note, when using certification \ certification.lte you must
// also specify certification_country. These two parameters work together
// in order to filter the results. You can only filter results with the
// countries we have added to our certifications list.
//
// If you specify the region parameter, the regional release date will be
// used instead of the primary release date. The date returned will be the
// first date based on your query (ie. if a with_release_type is specified).
// It's important to note the order of the release types that are used.
// Specifying "2|3" would return the limited theatrical release date as
// opposed to "3|2" which would return the theatrical date.
//
// Also note that a number of filters support being comma (,) or pipe (|)
// separated. Comma's are treated like an AND and query while pipe's
// are an OR.
//
// https://developers.themoviedb.org/3/discover/movie-discover
func (c *Client) GetDiscoverMovie(
	urlOptions map[string]string,
) (*DiscoverMovie, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smovie?api_key=%s%s",
		baseURL,
		discoverURL,
		c.apiKey,
		options,
	)
	discoverMovie := DiscoverMovie{}
	if err := c.get(tmdbURL, &discoverMovie); err != nil {
		return nil, err
	}
	return &discoverMovie, nil
}

// DiscoverTV type is a struct for tv JSON response.
type DiscoverTV struct {
	PaginatedResultsMeta
	*DiscoverTVResults
}

// GetDiscoverTV Discover TV shows by different types of data like average
// rating, number of votes, genres, the network they aired on and air dates.
//
// Discover also supports a nice list of sort options. See below for all of
// the available options.
//
// Also note that a number of filters support being comma (,) or pipe (|)
// separated. Comma's are treated like an AND and query while pipe's are an OR.
//
// https://developers.themoviedb.org/3/discover/tv-discover
func (c *Client) GetDiscoverTV(
	urlOptions map[string]string,
) (*DiscoverTV, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stv?api_key=%s%s",
		baseURL,
		discoverURL,
		c.apiKey,
		options,
	)
	discoverTV := DiscoverTV{}
	if err := c.get(tmdbURL, &discoverTV); err != nil {
		return nil, err
	}
	return &discoverTV, nil
}
