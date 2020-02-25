package tmdb

const imageURL = "https://image.tmdb.org/t/p/"

// All supported sizes including logos, posters, backdrops,
// profiles and still sizes.
var imageSize = map[string]string{
	"w45":      imageURL + "w45",
	"w92":      imageURL + "w92",
	"w154":     imageURL + "w154",
	"w185":     imageURL + "w185",
	"w300":     imageURL + "w300",
	"w342":     imageURL + "w342",
	"w500":     imageURL + "w500",
	"w780":     imageURL + "w780",
	"w1280":    imageURL + "w1280",
	"h632":     imageURL + "632",
	"original": imageURL + "original",
}

// GetImageURL accepts two parameters, the key and the size
// and returns the complete URL of the image.
//
// Available sizes:
//
// w45      - logo/profile
// w92      - logo/poster/still
// w154     - logo/poster
// w185     - logo/poster/profile/still
// w300     - backdrop/logo/still
// w342     - poster
// w500     - logo/poster
// w780     - backdrop/poster
// w1280    - backdrop
// h632     - profile
// original - backdrop/logo/poster/profile/still
//
// https://developers.themoviedb.org/3/configuration/get-api-configuration
func GetImageURL(key string, size string) string {
	return imageSize[size] + key
}

const videoURL = "https://www.youtube.com/watch?v="

// GetVideoURL accepts one parameter, the key and returns
// the complete URL of the video.
func GetVideoURL(key string) string {
	return videoURL + key
}
