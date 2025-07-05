package tmdb

import (
	"fmt"

	json "github.com/goccy/go-json"
)

func (c *CreditsDetails) UnmarshalJSON(data []byte) error {
	type alias CreditsDetails
	var raw struct {
		alias
		Media json.RawMessage `json:"media"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*c = CreditsDetails(raw.alias)
	var mediaType struct {
		MediaType MediaType `json:"media_type"`
	}
	if err := json.Unmarshal(raw.Media, &mediaType); err != nil {
		return err
	}

	switch mediaType.MediaType {
	case MediaTypeMovie:
		var movie CreditMovieMedia
		if err := json.Unmarshal(raw.Media, &movie); err != nil {
			return err
		}
		c.Media = movie
	case MediaTypeTV:
		var tv CreditTVShowMedia
		if err := json.Unmarshal(raw.Media, &tv); err != nil {
			return err
		}
		c.Media = tv
	default:
		return fmt.Errorf("unknown media_type: %s", mediaType.MediaType)
	}
	return nil
}

func (p *PersonResults) UnmarshalJSON(data []byte) error {
	type alias PersonResults
	var raw struct {
		alias
		KnownFor []json.RawMessage `json:"known_for"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*p = PersonResults(raw.alias)

	err := unmarshalMediaList(raw.KnownFor, func(m Media) {
		p.KnownFor = append(p.KnownFor, m)
	})
	return err
}

func (t *Trending) UnmarshalJSON(data []byte) error {
	type alias Trending
	var raw struct {
		alias
		Results []json.RawMessage `json:"results"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*t = Trending(raw.alias)

	err := unmarshalMediaList(raw.Results, func(m Media) {
		t.Results = append(t.Results, m)
	})
	return err
}

func unmarshalMediaList(rawItems []json.RawMessage, appendFunc func(Media)) error {
	for _, item := range rawItems {
		var mediaType struct {
			MediaType MediaType `json:"media_type"`
		}
		if err := json.Unmarshal(item, &mediaType); err != nil {
			return err
		}

		switch mediaType.MediaType {
		case MediaTypeMovie:
			var movie MovieMedia
			if err := json.Unmarshal(item, &movie); err != nil {
				return err
			}
			appendFunc(movie)
		case MediaTypeTV:
			var tv TVShowMedia
			if err := json.Unmarshal(item, &tv); err != nil {
				return err
			}
			appendFunc(tv)
		default:
			return fmt.Errorf("unknown media_type: %s", mediaType.MediaType)
		}
	}
	return nil
}
