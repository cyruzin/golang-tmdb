package tmdb

import json "github.com/goccy/go-json"

func (suite *TMBDTestSuite) TestUnmarshalJSONCrewTVUnmarshalError() {
	input := []byte(`{
		"id": 1,
		"crew": [
			{
				"media_type": "tv",
				"episode_count": "invalid"
			}
		]
	}`)

	var credits PersonMediaCredits
	err := json.Unmarshal(input, &credits)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalJSONCastTVUnmarshalError() {
	input := []byte(`{
		"id": 1,
		"cast": [
			{
				"media_type": "tv",
				"episode_count": "invalid"
			}
		]
	}`)

	var credits PersonMediaCredits
	err := json.Unmarshal(input, &credits)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalJSONCrewMovieUnmarshalError() {
	input := []byte(`{
		"id": 1,
		"crew": [
			{
				"media_type": "movie",
				"id": "invalid"
			}
		]
	}`)

	var credits PersonMediaCredits
	err := json.Unmarshal(input, &credits)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalJSONCastMovieUnmarshalError() {
	input := []byte(`{
		"id": 1,
		"cast": [
			{
				"media_type": "movie",
				"id": "invalid"
			}
		]
	}`)

	var credits PersonMediaCredits
	err := json.Unmarshal(input, &credits)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalJSONPersonMediaCreditsError() {
	jsonData := []byte(`{
  "cast": "",
  "crew": "",
  "ID": 12234
}`)
	var cd PersonMediaCredits
	err := json.Unmarshal(jsonData, &cd)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestUnmarshalJSONPersonMediaCreditsCastError() {
	jsonData := []byte(`{
  "cast": ["test"],
  "crew": [],
  "ID": 12234
}`)
	var cd PersonMediaCredits
	err := json.Unmarshal(jsonData, &cd)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalJSONPersonMediaCreditsCrewError() {
	jsonData := []byte(`{
  "cast": [],
  "crew": ["test"],
  "ID": 12234
}`)
	var cd PersonMediaCredits
	err := json.Unmarshal(jsonData, &cd)
	suite.Error(err)
}
func (suite *TMBDTestSuite) TestUnmarshalJSONPersonMediaCredits() {
	input := []byte(`{
		"id": 1234,
		"cast": [
			{
				"media_type": "movie",
				"id": 1234,
				"title": "Test",
				"credit_id": "abc"
			},
			{
				"media_type": "tv",
				"id": 1234,
				"name": "Test",
				"credit_id": "abc"
			}
		],
		"crew": [
			{
				"media_type": "movie",
				"id": 1234,
				"job": "Test",
				"credit_id": "abc"
			},
			{
				"media_type": "tv",
				"id": 1234,
				"job": "Test",
				"credit_id": "abc"
			}
		]
	}`)

	var credits PersonMediaCredits
	err := json.Unmarshal(input, &credits)
	suite.NoError(err)
	suite.Equal(MediaTypeMovie, credits.Cast[0].GetMediaCreditType())
	suite.Equal(MediaTypeTV, credits.Cast[1].GetMediaCreditType())
	suite.Equal(MediaTypeMovie, credits.Crew[0].GetMediaCreditType())
	suite.Equal(MediaTypeTV, credits.Crew[1].GetMediaCreditType())
}

func (suite *TMBDTestSuite) TestUnmarshalJSONPersonMediaCreditsCastUnknownMediaType() {
	input := []byte(`{
		"id": 1234,
		"cast": [
			{
				"media_type": "abc",
				"id": 1234
			}
		]
	}`)

	var credits PersonMediaCredits
	err := json.Unmarshal(input, &credits)

	suite.Error(err)
	suite.Contains(err.Error(), "unknown media_type")
}
func (suite *TMBDTestSuite) TestUnmarshalJSONPersonMediaCreditsCrewUnknownMediaType() {
	input := []byte(`{
		"id": 1234,
		"crew": [
			{
				"media_type": "abc",
				"id": 1234
			}
		]
	}`)

	var credits PersonMediaCredits
	err := json.Unmarshal(input, &credits)

	suite.Error(err)
	suite.Contains(err.Error(), "unknown media_type")
}

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

func (suite *TMBDTestSuite) TestUnmarshalPersonMediaInvalidResult() {
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

func (suite *TMBDTestSuite) TestUnmarshalPersonResultInvalidResult() {
	jsonData := `{invalid_data}`

	var t PersonResult
	err := json.Unmarshal([]byte(jsonData), &t)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPaginatedPersonResultsInvalidResult() {
	jsonData := `{invalid_data}`

	var pmr PaginatedPersonResults
	err := json.Unmarshal([]byte(jsonData), &pmr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultInvalidMediaType() {
	jsonData := `{
  "known_for": [
    {
      media_type: []
    }
  ]
}`

	var pr PersonResult
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultUnknownMediaType() {
	jsonData := `{
  "known_for": [
    {
      "media_type": "unknown"
    }
  ]
}`

	var pr PersonResult
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
	suite.Contains(err.Error(), "unknown media_type")
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultMovieMediaTypeFail() {
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

	var pr PersonResult
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultTVMediaTypeFail() {
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

	var pr PersonResult
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultPersonMediaTypeFail() {
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

	var pr PersonResult
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultMediaType() {
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

	var pr PersonResult
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Nil(err)
	knownFor := pr.KnownFor
	suite.NotNil(knownFor)
	suite.Equal(MediaTypeTV, knownFor[0].GetMediaType())
	suite.Equal(MediaTypeMovie, knownFor[1].GetMediaType())
	suite.Equal(MediaTypePerson, knownFor[2].GetMediaType())
}
