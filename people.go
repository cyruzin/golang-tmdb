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
