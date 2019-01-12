package tmdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmdb = Client{APIKey: os.Getenv("APIKey")}

func TestGetFail(t *testing.T) {
	err := tmdb.get("http://www.testfakewebsite.org", nil)

	assert.NotNil(t, err)

	err = tmdb.get("https://api.themoviedb.org/3/movie/75780?language=en-US", nil)

	assert.NotNil(t, err)
}
