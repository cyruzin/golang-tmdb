package tmdb

import json "github.com/goccy/go-json"

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsMovieMediaType() {
	jsonData := `{
  "media": {
    "title": "Title",
    "media_type": "movie"
  }
}`
	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Nil(err)
	media := cd.Media
	suite.NotNil(media)
	suite.Equal(MediaTypeMovie, media.GetMediaType())
}

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsMovieMediaTypeFail() {
	jsonData := `{
  "media": {
    "title": ["Name"],
    "media_type": "movie"
  }
}`
	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsTVMediaType() {
	jsonData := `{
  "media": {
    "name": "Name",
    "media_type": "tv"
  }
}`
	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Nil(err)
	media := cd.Media
	suite.NotNil(media)
	suite.Equal(MediaTypeTV, media.GetMediaType())
}

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsTVMediaTypeFail() {
	jsonData := `{
  "media": {
    "name": ["Name"],
    "media_type": "tv"
  }
}`
	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsUnknownMediaType() {
	jsonData := `{
  "media": {
    "media_type": "unknown"
  }
}`

	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Error(err)
	suite.Contains(err.Error(), "unknown media_type")
}

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsInvalidMediaType() {
	jsonData := `{
  "media": {
    media_type: []
  }
}`

	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsInvalidMedia() {
	jsonData := `{invalid_data}`

	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonInvalidResult() {
	jsonData := `{invalid_data}`

	var t PersonMedia
	err := json.Unmarshal([]byte(jsonData), &t)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPaginatedMediaResultsInvalidResult() {
	jsonData := `{invalid_data}`

	var pmr PaginatedMediaResults
	err := json.Unmarshal([]byte(jsonData), &pmr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonTaggedImageInvalidResult() {
	jsonData := `{invalid_data}`

	var t PersonTaggedImage
	err := json.Unmarshal([]byte(jsonData), &t)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonMediaInvalidMediaType() {
	jsonData := `{
  "known_for": [
    {
      media_type: []
    }
  ]
}`

	var pr PersonMedia
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonMediaUnknownMediaType() {
	jsonData := `{
  "known_for": [
    {
      "media_type": "unknown"
    }
  ]
}`

	var pr PersonMedia
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
	suite.Contains(err.Error(), "unknown media_type")
}

func (suite *TMBDTestSuite) TestUnmarshalPersonMediaMovieMediaTypeFail() {
	jsonData := `{
  "known_for": [
    {
      "title": [
        "Title"
      ],
      "media_type": "movie"
    }
  ]
}`

	var pr PersonMedia
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonMediaTVMediaTypeFail() {
	jsonData := `{
  "known_for": [
    {
      "name": [
        "Name"
      ],
      "media_type": "tv"
    }
  ]
}`

	var pr PersonMedia
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonMediaPersonMediaTypeFail() {
	jsonData := `{
  "known_for": [
    {
      "name": [
        "Name"
      ],
      "media_type": "person"
    }
  ]
}`

	var pr PersonMedia
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonMediaMediaType() {
	jsonData := `{
  "known_for": [
    {
      "name": "Name",
      "media_type": "tv"
    },
    {
      "title": "Title",
      "media_type": "movie"
    },
	{
	  "name": "Name",
	  "media_type": "person"
	}
  ]
}`

	var pr PersonMedia
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Nil(err)
	knownFor := pr.KnownFor
	suite.NotNil(knownFor)
	suite.Equal(MediaTypeTV, knownFor[0].GetMediaType())
	suite.Equal(MediaTypeMovie, knownFor[1].GetMediaType())
	suite.Equal(MediaTypePerson, knownFor[2].GetMediaType())
}
