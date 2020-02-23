package tmdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImageURL(t *testing.T) {
	url := "https://image.tmdb.org/t/p/w500/bOGkgRGdhrBYJSLpXaxhXVstddV.jpg"
	generatedURL := GetImageURL("/bOGkgRGdhrBYJSLpXaxhXVstddV.jpg", "w500")
	assert.Equal(t, generatedURL, url)
}

func TestGetVideoURL(t *testing.T) {
	url := "https://www.youtube.com/watch?v=6ZfuNTqbHE8"
	generatedURL := GetVideoURL("6ZfuNTqbHE8")
	assert.Equal(t, generatedURL, url)
}
