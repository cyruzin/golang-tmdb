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
