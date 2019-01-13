package tmdb

const (
	bumblebeeID   = 424783
	jackReacherID = 75780
)

func (suite *TMBDTestSuite) TestGetMovieDetails() {
	bumblebee, err := suite.GetMovieDetails(bumblebeeID, nil)

	suite.Nil(err)
	suite.Equal("Bumblebee", bumblebee.Title)
}

func (suite *TMBDTestSuite) TestGetMovieDetailsWithOptions() {
	options := make(map[string]string)

	options["language"] = "pt-BR"

	jackreacher, err := suite.GetMovieDetails(jackReacherID, options)

	suite.Nil(err)
	suite.Equal("Jack Reacher: O Último Tiro", jackreacher.Title)
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitles() {
	bumblebee, err := suite.GetMovieAlternativeTitles(bumblebeeID, nil)

	suite.Nil(err)

	for _, v := range bumblebee.Titles {
		if v.Iso3166_1 == "US" {
			suite.Equal("Transformers 6", v.Title)
		}
	}
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitlesWithOptions() {
	options := make(map[string]string)

	options["country"] = "RU"

	bumblebee, err := suite.GetMovieAlternativeTitles(bumblebeeID, options)

	suite.Nil(err)

	suite.Equal("RU", bumblebee.Titles[0].Iso3166_1)
}
