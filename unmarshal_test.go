package tmdb

import json "github.com/goccy/go-json"

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsMovieMediaType() {
	jsonData := `{
  "credit_type": "cast",
  "department": "Acting",
  "job": "Actor",
  "media": {
    "adult": false,
    "backdrop_path": "/bfh9Z3Ghz4FOJAfLOAhmc3ccnHU.jpg",
    "id": 671,
    "title": "Harry Potter and the Philosopher's Stone",
    "original_title": "Harry Potter and the Philosopher's Stone",
    "overview": "Harry Potter has lived under the stairs at his aunt and uncle's house his whole life. But on his 11th birthday, he learns he's a powerful wizard—with a place waiting for him at the Hogwarts School of Witchcraft and Wizardry. As he learns to harness his newfound powers with the help of the school's kindly headmaster, Harry uncovers the truth about his parents' deaths—and about the villain who's to blame.",
    "poster_path": "/wuMc08IPKEatf9rnMNXvIDxqP4W.jpg",
    "media_type": "movie",
    "original_language": "en",
    "genre_ids": [
      12,
      14
    ],
    "popularity": 35.7234,
    "release_date": "2001-11-16",
    "video": false,
    "vote_average": 7.901,
    "vote_count": 28318,
    "character": "Harry Potter"
  },
  "media_type": "movie",
  "id": "52fe4267c3a36847f801be91",
  "person": {
    "adult": false,
    "id": 10980,
    "name": "Daniel Radcliffe",
    "original_name": "Daniel Radcliffe",
    "media_type": "person",
    "popularity": 5.7284,
    "gender": 2,
    "known_for_department": "Acting",
    "profile_path": "/iPg0J9UzAlPj1fLEJNllpW9IhGe.jpg"
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
  "credit_type": "cast",
  "department": "Acting",
  "job": "Actor",
  "media": {
    "adult": false,
    "backdrop_path": "/bfh9Z3Ghz4FOJAfLOAhmc3ccnHU.jpg",
    "id": 671,
    "title": ["Harry Potter and the Philosopher's Stone"],
    "original_title": "Harry Potter and the Philosopher's Stone",
    "overview": "Harry Potter has lived under the stairs at his aunt and uncle's house his whole life. But on his 11th birthday, he learns he's a powerful wizard—with a place waiting for him at the Hogwarts School of Witchcraft and Wizardry. As he learns to harness his newfound powers with the help of the school's kindly headmaster, Harry uncovers the truth about his parents' deaths—and about the villain who's to blame.",
    "poster_path": "/wuMc08IPKEatf9rnMNXvIDxqP4W.jpg",
    "media_type": "movie",
    "original_language": "en",
    "genre_ids": [
      12,
      14
    ],
    "popularity": 35.7234,
    "release_date": "2001-11-16",
    "video": false,
    "vote_average": 7.901,
    "vote_count": 28318,
    "character": "Harry Potter"
  },
  "media_type": "movie",
  "id": "52fe4267c3a36847f801be91",
  "person": {
    "adult": false,
    "id": 10980,
    "name": "Daniel Radcliffe",
    "original_name": "Daniel Radcliffe",
    "media_type": "person",
    "popularity": 5.7284,
    "gender": 2,
    "known_for_department": "Acting",
    "profile_path": "/iPg0J9UzAlPj1fLEJNllpW9IhGe.jpg"
  }
}`
	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsTVMediaType() {
	jsonData := `{
  "credit_type": "cast",
  "department": "Acting",
  "job": "Actor",
  "media": {
    "adult": false,
    "backdrop_path": "/vNnLAKmoczRlNarxyGrrw0KSOeX.jpg",
    "id": 1412,
    "name": "Arrow",
    "original_name": "Arrow",
    "overview": "Spoiled billionaire playboy Oliver Queen is missing and presumed dead when his yacht is lost at sea. He returns five years later a changed man, determined to clean up the city as a hooded vigilante armed with a bow.",
    "poster_path": "/u8ZHFj1jC384JEkTt3vNg1DfWEb.jpg",
    "media_type": "tv",
    "original_language": "en",
    "genre_ids": [
      80,
      18,
      10759
    ],
    "popularity": 89.2053,
    "first_air_date": "2012-10-10",
    "vote_average": 6.826,
    "vote_count": 6086,
    "origin_country": [
      "US"
    ],
    "character": "Oliver Queen / Green Arrow",
    "episodes": [],
    "seasons": []
  },
  "media_type": "tv",
  "id": "5256ce3819c2956ff608293d",
  "person": {
    "adult": false,
    "id": 76110,
    "name": "Stephen Amell",
    "original_name": "Stephen Amell",
    "media_type": "person",
    "popularity": 1.8679,
    "gender": 2,
    "known_for_department": "Acting",
    "profile_path": "/iyYZeOsJkE7D8Wsamv6NML9WlL8.jpg"
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
  "credit_type": "cast",
  "department": "Acting",
  "job": "Actor",
  "media": {
    "adult": false,
    "backdrop_path": "/vNnLAKmoczRlNarxyGrrw0KSOeX.jpg",
    "id": 1412,
    "name": ["Arrow"],
    "original_name": "Arrow",
    "overview": "Spoiled billionaire playboy Oliver Queen is missing and presumed dead when his yacht is lost at sea. He returns five years later a changed man, determined to clean up the city as a hooded vigilante armed with a bow.",
    "poster_path": "/u8ZHFj1jC384JEkTt3vNg1DfWEb.jpg",
    "media_type": "tv",
    "original_language": "en",
    "genre_ids": [
      80,
      18,
      10759
    ],
    "popularity": 89.2053,
    "first_air_date": "2012-10-10",
    "vote_average": 6.826,
    "vote_count": 6086,
    "origin_country": [
      "US"
    ],
    "character": "Oliver Queen / Green Arrow",
    "episodes": [],
    "seasons": []
  },
  "media_type": "tv",
  "id": "5256ce3819c2956ff608293d",
  "person": {
    "adult": false,
    "id": 76110,
    "name": "Stephen Amell",
    "original_name": "Stephen Amell",
    "media_type": "person",
    "popularity": 1.8679,
    "gender": 2,
    "known_for_department": "Acting",
    "profile_path": "/iyYZeOsJkE7D8Wsamv6NML9WlL8.jpg"
  }
}`
	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsUnknownMediaType() {
	jsonData := `{
  "credit_type": "cast",
  "department": "Acting",
  "job": "Actor",
  "media": {
    "adult": false,
    "backdrop_path": "/bfh9Z3Ghz4FOJAfLOAhmc3ccnHU.jpg",
    "id": 671,
    "title": "Harry Potter and the Philosopher's Stone",
    "original_title": "Harry Potter and the Philosopher's Stone",
    "overview": "Harry Potter has lived under the stairs at his aunt and uncle's house his whole life. But on his 11th birthday, he learns he's a powerful wizard—with a place waiting for him at the Hogwarts School of Witchcraft and Wizardry. As he learns to harness his newfound powers with the help of the school's kindly headmaster, Harry uncovers the truth about his parents' deaths—and about the villain who's to blame.",
    "poster_path": "/wuMc08IPKEatf9rnMNXvIDxqP4W.jpg",
    "media_type": "unknown",
    "original_language": "en",
    "genre_ids": [
      12,
      14
    ],
    "popularity": 35.7234,
    "release_date": "2001-11-16",
    "video": false,
    "vote_average": 7.901,
    "vote_count": 28318,
    "character": "Harry Potter"
  },
  "media_type": "movie",
  "id": "52fe4267c3a36847f801be91",
  "person": {
    "adult": false,
    "id": 10980,
    "name": "Daniel Radcliffe",
    "original_name": "Daniel Radcliffe",
    "media_type": "person",
    "popularity": 5.7284,
    "gender": 2,
    "known_for_department": "Acting",
    "profile_path": "/iPg0J9UzAlPj1fLEJNllpW9IhGe.jpg"
  }
}`

	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Error(err)
	suite.Contains(err.Error(), "unknown media_type")
}

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsInvalidMediaType() {
	jsonData := `{
  "credit_type": "cast",
  "department": "Acting",
  "job": "Actor",
  "media": {
    "adult": false,
    "backdrop_path": "/bfh9Z3Ghz4FOJAfLOAhmc3ccnHU.jpg",
    "id": 671,
    "title": "Harry Potter and the Philosopher's Stone",
    "original_title": "Harry Potter and the Philosopher's Stone",
    "overview": "Harry Potter has lived under the stairs at his aunt and uncle's house his whole life. But on his 11th birthday, he learns he's a powerful wizard—with a place waiting for him at the Hogwarts School of Witchcraft and Wizardry. As he learns to harness his newfound powers with the help of the school's kindly headmaster, Harry uncovers the truth about his parents' deaths—and about the villain who's to blame.",
    "poster_path": "/wuMc08IPKEatf9rnMNXvIDxqP4W.jpg",
    media_type: [],
    "original_language": "en",
    "genre_ids": [
      12,
      14
    ],
    "popularity": 35.7234,
    "release_date": "2001-11-16",
    "video": false,
    "vote_average": 7.901,
    "vote_count": 28318,
    "character": "Harry Potter"
  },
  "media_type": "movie",
  "id": "52fe4267c3a36847f801be91",
  "person": {
    "adult": false,
    "id": 10980,
    "name": "Daniel Radcliffe",
    "original_name": "Daniel Radcliffe",
    "media_type": "person",
    "popularity": 5.7284,
    "gender": 2,
    "known_for_department": "Acting",
    "profile_path": "/iPg0J9UzAlPj1fLEJNllpW9IhGe.jpg"
  }
}`

	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalCreditsDetailsInvalidMedia() {
	jsonData := `{
  "credit_type": "cast",
  "department": "Acting",
  "job": "Actor",
  media: {
    "adult": false,
    "backdrop_path": "/vNnLAKmoczRlNarxyGrrw0KSOeX.jpg",
    "id": 1412,
    "name": "Arrow",
    "original_name": "Arrow",
    "overview": "Spoiled billionaire playboy Oliver Queen is missing and presumed dead when his yacht is lost at sea. He returns five years later a changed man, determined to clean up the city as a hooded vigilante armed with a bow.",
    "poster_path": "/u8ZHFj1jC384JEkTt3vNg1DfWEb.jpg",
    "media_type": "tv",
    "original_language": "en",
    "genre_ids": [
      80,
      18,
      10759
    ],
    "popularity": 89.2053,
    "first_air_date": "2012-10-10",
    "vote_average": 6.826,
    "vote_count": 6086,
    "origin_country": [
      "US"
    ],
    "character": "Oliver Queen / Green Arrow",
    "episodes": [],
    "seasons": []
  },
  "media_type": "movie",
  "id": "52fe4267c3a36847f801be91",
  "person": {
    "adult": false,
    "id": 10980,
    "name": "Daniel Radcliffe",
    "original_name": "Daniel Radcliffe",
    "media_type": "person",
    "popularity": 5.7284,
    "gender": 2,
    "known_for_department": "Acting",
    "profile_path": "/iPg0J9UzAlPj1fLEJNllpW9IhGe.jpg"
  }
}`

	var cd CreditsDetails
	err := json.Unmarshal([]byte(jsonData), &cd)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultsMediaType() {
	jsonData := `{
      "adult": false,
      "id": 23458,
      "name": "Charlie Cox",
      "original_name": "Charlie Cox",
      "media_type": "person",
      "popularity": 2.3871,
      "gender": 2,
      "known_for_department": "Acting",
      "profile_path": "/jBHDZ8MA4I7krNQx4IfqdfPfleD.jpg",
      "known_for": [
        {
          "adult": false,
          "backdrop_path": "/qsnXwGS7KBbX4JLqHvICngtR8qg.jpg",
          "id": 61889,
          "name": "Marvel's Daredevil",
          "original_name": "Marvel's Daredevil",
          "overview": "Lawyer-by-day Matt Murdock uses his heightened senses from being blinded as a young boy to fight crime at night on the streets of Hell’s Kitchen as Daredevil.",
          "poster_path": "/QWbPaDxiB6LW2LjASknzYBvjMj.jpg",
          "media_type": "tv",
          "original_language": "en",
          "genre_ids": [
            80,
            18,
            10759
          ],
          "popularity": 89.5761,
          "first_air_date": "2015-04-10",
          "vote_average": 8.2,
          "vote_count": 4781,
          "origin_country": [
            "US"
          ]
        },
		{
          "adult": false,
          "backdrop_path": "/x2LTysjFSpeK0d40zHw6Lqd4oAH.jpg",
          "id": 2270,
          "title": "Stardust",
          "original_title": "Stardust",
          "overview": "In a countryside town bordering on a magical land, a young man makes a promise to his beloved that he'll retrieve a fallen star by venturing into the magical realm. His journey takes him into a world beyond his wildest dreams and reveals his true identity.",
          "poster_path": "/7zbFmxy3DqKYL2M8Hop6uylp2Uy.jpg",
          "media_type": "movie",
          "original_language": "en",
          "genre_ids": [
            12,
            14,
            10749,
            10751
          ],
          "popularity": 5.5908,
          "release_date": "2007-08-10",
          "video": false,
          "vote_average": 7.284,
          "vote_count": 4176
        }
      ]
    }`

	var pr PersonResults
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Nil(err)
	knownFor := pr.KnownFor
	suite.NotNil(knownFor)
	suite.Equal(MediaTypeTV, knownFor[0].GetMediaType())
	suite.Equal(MediaTypeMovie, knownFor[1].GetMediaType())
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultsTVMediaTypeFail() {
	jsonData := `{
      "adult": false,
      "id": 23458,
      "name": "Charlie Cox",
      "original_name": "Charlie Cox",
      "media_type": "person",
      "popularity": 2.3871,
      "gender": 2,
      "known_for_department": "Acting",
      "profile_path": "/jBHDZ8MA4I7krNQx4IfqdfPfleD.jpg",
      "known_for": [
        {
          "adult": false,
          "backdrop_path": "/qsnXwGS7KBbX4JLqHvICngtR8qg.jpg",
          "id": 61889,
          "name": "Marvel's Daredevil",
          "original_name": ["Marvel's Daredevil"],
          "overview": "Lawyer-by-day Matt Murdock uses his heightened senses from being blinded as a young boy to fight crime at night on the streets of Hell’s Kitchen as Daredevil.",
          "poster_path": "/QWbPaDxiB6LW2LjASknzYBvjMj.jpg",
          "media_type": "tv",
          "original_language": "en",
          "genre_ids": [
            80,
            18,
            10759
          ],
          "popularity": 89.5761,
          "first_air_date": "2015-04-10",
          "vote_average": 8.2,
          "vote_count": 4781,
          "origin_country": [
            "US"
          ]
        }
      ]
    }`

	var pr PersonResults
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultsMovieMediaTypeFail() {
	jsonData := `{
      "adult": false,
      "id": 23458,
      "name": "Charlie Cox",
      "original_name": "Charlie Cox",
      "media_type": "person",
      "popularity": 2.3871,
      "gender": 2,
      "known_for_department": "Acting",
      "profile_path": "/jBHDZ8MA4I7krNQx4IfqdfPfleD.jpg",
      "known_for": [
		{
          "adult": false,
          "backdrop_path": "/x2LTysjFSpeK0d40zHw6Lqd4oAH.jpg",
          "id": 2270,
          "title": ["Stardust"],
          "original_title": "Stardust",
          "overview": "In a countryside town bordering on a magical land, a young man makes a promise to his beloved that he'll retrieve a fallen star by venturing into the magical realm. His journey takes him into a world beyond his wildest dreams and reveals his true identity.",
          "poster_path": "/7zbFmxy3DqKYL2M8Hop6uylp2Uy.jpg",
          "media_type": "movie",
          "original_language": "en",
          "genre_ids": [
            12,
            14,
            10749,
            10751
          ],
          "popularity": 5.5908,
          "release_date": "2007-08-10",
          "video": false,
          "vote_average": 7.284,
          "vote_count": 4176
        }
      ]
    }`

	var pr PersonResults
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultsUnknownMediaType() {
	jsonData := `{
      "adult": false,
      "id": 23458,
      "name": "Charlie Cox",
      "original_name": "Charlie Cox",
      "media_type": "person",
      "popularity": 2.3871,
      "gender": 2,
      "known_for_department": "Acting",
      "profile_path": "/jBHDZ8MA4I7krNQx4IfqdfPfleD.jpg",
      "known_for": [
        {
          "media_type": "unknown"
        }
      ]
    }`

	var pr PersonResults
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
	suite.Contains(err.Error(), "unknown media_type")
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultsInvalidMediaType() {
	jsonData := `{
      "adult": false,
      "id": 23458,
      "name": "Charlie Cox",
      "original_name": "Charlie Cox",
      "media_type": "person",
      "popularity": 2.3871,
      "gender": 2,
      "known_for_department": "Acting",
      "profile_path": "/jBHDZ8MA4I7krNQx4IfqdfPfleD.jpg",
      "known_for": [
        {
          "id": 1234,
          media_type: [],
        }
      ]
    }`

	var pr PersonResults
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalPersonResultsInvalidKnownFor() {
	jsonData := `{
      "adult": false,
      "id": 23458,
      "name": "Charlie Cox",
      "original_name": "Charlie Cox",
      "media_type": "person",
      "popularity": 2.3871,
      "gender": 2,
      "known_for_department": "Acting",
      "profile_path": "/jBHDZ8MA4I7krNQx4IfqdfPfleD.jpg",
      known_for: [
        {
          "adult": false,
          "backdrop_path": "/qsnXwGS7KBbX4JLqHvICngtR8qg.jpg",
          "id": 61889,
          "name": "Marvel's Daredevil",
          "original_name": "Marvel's Daredevil",
          "overview": "Lawyer-by-day Matt Murdock uses his heightened senses from being blinded as a young boy to fight crime at night on the streets of Hell’s Kitchen as Daredevil.",
          "poster_path": "/QWbPaDxiB6LW2LjASknzYBvjMj.jpg",
          "media_type": "tv",
          "original_language": "en",
          "genre_ids": [
            80,
            18,
            10759
          ],
          "popularity": 89.5761,
          "first_air_date": "2015-04-10",
          "vote_average": 8.2,
          "vote_count": 4781,
          "origin_country": [
            "US"
          ]
        },
      ]
    }`

	var pr PersonResults
	err := json.Unmarshal([]byte(jsonData), &pr)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestUnmarshalTrendingInvalidResult() {
	jsonData := `{
  "page":1,
  "total_results":1,
	"total_pages":1,
      result: [
        {
          "adult": false,
          "backdrop_path": "/qsnXwGS7KBbX4JLqHvICngtR8qg.jpg",
          "id": 61889,
          "name": "Marvel's Daredevil",
          "original_name": "Marvel's Daredevil",
          "overview": "Lawyer-by-day Matt Murdock uses his heightened senses from being blinded as a young boy to fight crime at night on the streets of Hell’s Kitchen as Daredevil.",
          "poster_path": "/QWbPaDxiB6LW2LjASknzYBvjMj.jpg",
          "media_type": "tv",
          "original_language": "en",
          "genre_ids": [
            80,
            18,
            10759
          ],
          "popularity": 89.5761,
          "first_air_date": "2015-04-10",
          "vote_average": 8.2,
          "vote_count": 4781,
          "origin_country": [
            "US"
          ]
        },
      ]
    }`

	var t Trending
	err := json.Unmarshal([]byte(jsonData), &t)
	suite.Error(err)
}
