package tmdb

// BelongsToCollection represents information about a collection to which a movie belongs,
// including its unique identifier, name, poster image path, and backdrop image path.
type BelongsToCollection struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

// ProductionCompany represents a company involved in the production of a movie or TV show.
// It includes the company's name, unique identifier, logo path, and country of origin.
type ProductionCompany struct {
	Name          string `json:"name"`
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

// Genre represents a movie or TV show genre with its unique identifier and name.
type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ProductionCountry represents a country involved in the production of a movie or TV show,
// containing its ISO 3166-1 code and the country's name.
type ProductionCountry struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Name      string `json:"name"`
}

// SpokenLanguage represents a language spoken in a media item, including its ISO 639-1 code and name.
type SpokenLanguage struct {
	Iso639_1 string `json:"iso_639_1"`
	Name     string `json:"name"`
}

// TranslationData represents the translated information for a media item,
// including its title, name, overview, homepage, runtime, and tagline.
// Fields are mapped to their respective JSON keys and may be omitted if empty.
type TranslationData struct {
	Title    string `json:"title,omitempty"`
	Name     string `json:"name,omitempty"`
	Overview string `json:"overview,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Runtime  *int   `json:"runtime,omitempty"`
	Tagline  string `json:"tagline,omitempty"`
}

// Translation represents a translation entry with language and country codes,
// localized and English names, and associated translation data.
type Translation struct {
	Iso3166_1   string          `json:"iso_3166_1"`
	Iso639_1    string          `json:"iso_639_1"`
	Name        string          `json:"name"`
	EnglishName string          `json:"english_name"`
	Data        TranslationData `json:"data"`
}

// PaginatedResultsMeta represents metadata for paginated API responses,
// including the current page, total number of results, and total number of pages.
type PaginatedResultsMeta struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
}

// Season represents a TV show season with details such as air date, episode count, name, overview, poster path, season number, vote average, and associated show ID.
type Season struct {
	AirDate      string  `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	VoteAverage  float32 `json:"vote_average,omitempty"`
	ShowID       int64   `json:"show_id,omitempty"`
}

// LastEpisodeToAir represents the details of the most recently aired episode of a TV show.
// It includes information such as air date, episode and season numbers, production code,
// episode overview, voting statistics, and related media paths.
type LastEpisodeToAir struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int64   `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float32 `json:"vote_average"`
	VoteCount      int64   `json:"vote_count"`
}

// NextEpisodeToAir represents the details of the next episode scheduled to air for a TV show.
// It includes information such as air date, episode and season numbers, show and episode IDs,
// episode name and overview, production code, still image path, and voting statistics.
type NextEpisodeToAir struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int64   `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float32 `json:"vote_average"`
	VoteCount      int64   `json:"vote_count"`
}

// Network represents a television network with its identifying information,
// including name, unique ID, logo path, and country of origin.
type Network struct {
	Name          string `json:"name"`
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

// CreatedBy represents a person who created a particular media item, such as a TV show or movie.
// It includes identifying information such as ID, credit ID, name, gender, and profile image path.
type CreatedBy struct {
	ID          int64  `json:"id"`
	CreditID    string `json:"credit_id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profile_path"`
}
