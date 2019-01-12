package tmdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmdb = Client{APIKey: os.Getenv("APIKey")}

func TestGetMovieDetails(t *testing.T) {
	bumblebee, err := tmdb.GetMovieDetails(424783, nil)

	assert.Nil(t, err)
	assert.Equal(t, "Bumblebee", bumblebee.Title)
}

func TestGetMovieDetailsWithOptions(t *testing.T) {
	options := make(map[string]string)

	options["language"] = "pt-BR"

	jackreacher, err := tmdb.GetMovieDetails(75780, options)

	assert.Nil(t, err)
	assert.Equal(t, "Jack Reacher: O Último Tiro", jackreacher.Title)
}
