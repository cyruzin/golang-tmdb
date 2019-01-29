package tmdb

import "fmt"

// PeopleDetails type is a struct for details JSON response.
type PeopleDetails struct {
	Birthday           string   `json:"birthday"`
	KnownForDepartment string   `json:"known_for_department"`
	Deathday           string   `json:"deathday"`
	ID                 int64    `json:"id"`
	Name               string   `json:"name"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Gender             int      `json:"gender"`
	Biography          string   `json:"biography"`
	Popularity         float32  `json:"popularity"`
	PlaceOfBirth       string   `json:"place_of_birth"`
	ProfilePath        string   `json:"profile_path"`
	Adult              bool     `json:"adult"`
	IMDbID             string   `json:"imdb_id"`
	Homepage           string   `json:"homepage"`
	*PeopleChangesShort
	*PeopleMovieCreditsShort
	*PeopleTVCreditsShort
	*PeopleCombinedCreditsShort
	*PeopleExternalIDsShort
}

// PeopleChanges type is a struct for changes JSON response.
type PeopleChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID            string `json:"id"`
			Action        string `json:"action"`
			Time          string `json:"time"`
			OriginalValue struct {
				Profile struct {
					FilePath string `json:"file_path"`
				} `json:"profile"`
			} `json:"original_value"`
		} `json:"items"`
	} `json:"changes"`
}

// PeopleChangesShort type is a short struct for changes JSON response.
type PeopleChangesShort struct {
	Change *PeopleChanges `json:"changes,omitempty"`
}

// PeopleMovieCredits type is a struct for movie credits JSON response.
type PeopleMovieCredits struct {
	Cast []struct {
		Character        string  `json:"character"`
		CreditID         string  `json:"credit_id"`
		PosterPath       string  `json:"poster_path"`
		ID               int64   `json:"id"`
		Video            bool    `json:"video"`
		VoteCount        int64   `json:"vote_count"`
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIds         []int   `json:"genre_ids"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Popularity       float32 `json:"popularity"`
		Title            string  `json:"title"`
		VoteAverage      float32 `json:"vote_average"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
	} `json:"cast"`
	Crew []struct {
		ID               int64   `json:"id"`
		Department       string  `json:"department"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Job              string  `json:"job"`
		Overview         string  `json:"overview"`
		VoteCount        int64   `json:"vote_count"`
		Video            bool    `json:"video"`
		ReleaseDate      string  `json:"release_date"`
		VoteAverage      float32 `json:"vote_average"`
		Title            string  `json:"title"`
		Popularity       float32 `json:"popularity"`
		GenreIDs         []int   `json:"genre_ids"`
		BackdropPath     string  `json:"backdrop_path"`
		Adult            bool    `json:"adult"`
		PosterPath       string  `json:"poster_path"`
		CreditID         string  `json:"credit_id"`
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// PeopleMovieCreditsShort type is a short struct for movie credits JSON response.
type PeopleMovieCreditsShort struct {
	MovieCredits *PeopleMovieCredits `json:"movie_credits,omitempty"`
}

// PeopleTVCredits type is a struct for tv credits JSON response.
type PeopleTVCredits struct {
	Cast []struct {
		CreditID         string   `json:"credit_id"`
		OriginalName     string   `json:"original_name"`
		ID               int64    `json:"id"`
		GenreIds         []int    `json:"genre_ids"`
		Character        string   `json:"character"`
		Name             string   `json:"name"`
		PosterPath       string   `json:"poster_path"`
		VoteCount        int64    `json:"vote_count"`
		VoteAverage      float32  `json:"vote_average"`
		Popularity       float32  `json:"popularity"`
		EpisodeCount     int      `json:"episode_count"`
		OriginalLanguage string   `json:"original_language"`
		FirstAirDate     string   `json:"first_air_date"`
		BackdropPath     string   `json:"backdrop_path"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
	} `json:"cast"`
	Crew []struct {
		ID               int64    `json:"id"`
		Department       string   `json:"department"`
		OriginalLanguage string   `json:"original_language"`
		EpisodeCount     int      `json:"episode_count"`
		Job              string   `json:"job"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
		OriginalName     string   `json:"original_name"`
		GenreIDs         []int    `json:"genre_ids"`
		Name             string   `json:"name"`
		FirstAirDate     string   `json:"first_air_date"`
		BackdropPath     string   `json:"backdrop_path"`
		Popularity       float32  `json:"popularity"`
		VoteCount        int64    `json:"vote_count"`
		VoteAverage      float32  `json:"vote_average"`
		PosterPath       string   `json:"poster_path"`
		CreditID         string   `json:"credit_id"`
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// PeopleTVCreditsShort type is a short struct for tv credits JSON response.
type PeopleTVCreditsShort struct {
	TVCredits *PeopleTVCredits `json:"tv_credits,omitempty"`
}

// PeopleCombinedCredits type is a struct for combined credits JSON response.
type PeopleCombinedCredits struct {
	Cast []struct {
		ID               int64    `json:"id"`
		Character        string   `json:"character"`
		OriginalTitle    string   `json:"original_title"`
		Overview         string   `json:"overview"`
		VoteCount        int64    `json:"vote_count"`
		Video            bool     `json:"video"`
		MediaType        string   `json:"media_type"`
		ReleaseDate      string   `json:"release_date"`
		VoteAverage      float32  `json:"vote_average"`
		Title            string   `json:"title"`
		Popularity       float32  `json:"popularity"`
		OriginalLanguage string   `json:"original_language"`
		GenreIDs         []int    `json:"genre_ids"`
		BackdropPath     string   `json:"backdrop_path"`
		Adult            bool     `json:"adult"`
		PosterPath       string   `json:"poster_path"`
		CreditID         string   `json:"credit_id"`
		EpisodeCount     int      `json:"episode_count"`
		OriginCountry    []string `json:"origin_country"`
		OriginalName     string   `json:"original_name"`
		Name             string   `json:"name"`
		FirstAirDate     string   `json:"first_air_date"`
	} `json:"cast"`
	Crew []struct {
		ID               int64   `json:"id"`
		Department       string  `json:"department"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Job              string  `json:"job"`
		Overview         string  `json:"overview"`
		VoteCount        int64   `json:"vote_count"`
		Video            bool    `json:"video"`
		MediaType        string  `json:"media_type"`
		PosterPath       string  `json:"poster_path"`
		BackdropPath     string  `json:"backdrop_path"`
		Title            string  `json:"title"`
		Popularity       float32 `json:"popularity"`
		GenreIDs         []int   `json:"genre_ids"`
		VoteAverage      float32 `json:"vote_average"`
		Adult            bool    `json:"adult"`
		ReleaseDate      string  `json:"release_date"`
		CreditID         string  `json:"credit_id"`
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// PeopleCombinedCreditsShort type is a short struct for combined credits JSON response.
type PeopleCombinedCreditsShort struct {
	CombinedCredits *PeopleCombinedCredits `json:"combined_credits,omitempty"`
}

// PeopleExternalIDs type is a struct for external ids JSON response.
type PeopleExternalIDs struct {
	ID          int64  `json:"id,omitempty"`
	TwitterID   string `json:"twitter_id"`
	FacebookID  string `json:"facebook_id"`
	TvrageID    int64  `json:"tvrage_id"`
	InstagramID string `json:"instagram_id"`
	FreebaseMid string `json:"freebase_mid"`
	ImdbID      string `json:"imdb_id"`
	FreebaseID  string `json:"freebase_id"`
}

// PeopleExternalIDsShort type is a short struct for external ids JSON response.
type PeopleExternalIDsShort struct {
	ExternalIDs *PeopleExternalIDs `json:"external_ids,omitempty"`
}

// GetPeopleDetails get the primary person details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/people/get-person-details
func (c *Client) GetPeopleDetails(id int, o map[string]string) (*PeopleDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleDetails{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleChanges get the changes for a person.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by
// using the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/people/get-person-changes
func (c *Client) GetPeopleChanges(id int, o map[string]string) (*PeopleChanges, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/changes?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleChanges{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleMovieCredits get the movie credits for a person.
//
// https://developers.themoviedb.org/3/people/get-person-movie-credits
func (c *Client) GetPeopleMovieCredits(id int, o map[string]string) (*PeopleMovieCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/movie_credits?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleMovieCredits{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleTVCredits get the TV show credits for a person.
//
// https://developers.themoviedb.org/3/people/get-person-tv-credits
func (c *Client) GetPeopleTVCredits(id int, o map[string]string) (*PeopleTVCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/tv_credits?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleTVCredits{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleCombinedCredits get the movie and TV credits together in a single response.
//
// https://developers.themoviedb.org/3/people/get-person-combined-credits
func (c *Client) GetPeopleCombinedCredits(id int, o map[string]string) (*PeopleCombinedCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/combined_credits?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleCombinedCredits{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleExternalIDs get the external ids for a person.
// We currently support the following external sources.
//
// External Sources: IMDb ID, Facebook, Freebase MID, Freebase ID,
// Instagram, TVRage ID, Twitter.
//
// https://developers.themoviedb.org/3/people/get-person-external-ids
func (c *Client) GetPeopleExternalIDs(id int, o map[string]string) (*PeopleExternalIDs, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/external_ids?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleExternalIDs{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
