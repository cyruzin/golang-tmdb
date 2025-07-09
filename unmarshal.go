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

func (p *PersonMedia) UnmarshalJSON(data []byte) error {
	type alias PersonMedia
	var raw struct {
		alias
		KnownFor []json.RawMessage `json:"known_for"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*p = PersonMedia(raw.alias)
	return unmarshalMediaList(raw.KnownFor, func(m Media) {
		p.KnownFor = append(p.KnownFor, m)
	})
}

func (pti *PersonTaggedImage) UnmarshalJSON(data []byte) error {
	type alias PersonTaggedImage
	var raw struct {
		alias
		Media json.RawMessage `json:"media"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*pti = PersonTaggedImage(raw.alias)
	return unmarshalMedia(raw.Media, func(m Media) {
		pti.Media = m
	})
}

func (pmr *PaginatedMediaResults) UnmarshalJSON(data []byte) error {
	type alias PaginatedMediaResults
	var raw struct {
		alias
		Results []json.RawMessage `json:"results"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*pmr = PaginatedMediaResults(raw.alias)
	return unmarshalMediaList(raw.Results, func(m Media) {
		pmr.Results = append(pmr.Results, m)
	})
}

func unmarshalMedia(item json.RawMessage, handleMedia func(Media)) error {
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
		handleMedia(movie)
	case MediaTypeTV:
		var tv TVShowMedia
		if err := json.Unmarshal(item, &tv); err != nil {
			return err
		}
		handleMedia(tv)
	case MediaTypePerson:
		var person PersonMedia
		if err := json.Unmarshal(item, &person); err != nil {
			return err
		}
		handleMedia(person)
	default:
		return fmt.Errorf("unknown media_type: %s", mediaType.MediaType)
	}
	return nil
}

func unmarshalMediaList(rawItems []json.RawMessage, handleMedia func(Media)) error {
	for _, item := range rawItems {
		if err := unmarshalMedia(item, handleMedia); err != nil {
			return err
		}
	}
	return nil
}
