package tmdb

import json "github.com/goccy/go-json"

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
	CompanyInfo
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
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	VoteAverage  float32 `json:"vote_average"`
	AirDate      string  `json:"air_date"`
	SeasonNumber int     `json:"season_number"`
	EpisodeCount int     `json:"episode_count"`
}

// Network represents a television network with its identifying information,
// including name, unique ID, logo path, and country of origin.
type Network struct {
	CompanyInfo
}

// CreatedBy represents a person who created a particular media item, such as a TV show or movie.
// It includes identifying information such as ID, credit ID, name, gender, and profile image path.
type CreatedBy struct {
	PersonBase
	CreditID string `json:"credit_id"`
}

type VoteMetrics struct {
	VoteCount   int64   `json:"vote_count"`
	VoteAverage float32 `json:"vote_average"`
}

type ImageBase struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	Width       int     `json:"width"`
	VoteMetrics
}

type WatchProvider struct {
	LogoPath        string `json:"logo_path"`
	ProviderID      int    `json:"provider_id"`
	ProviderName    string `json:"provider_name"`
	DisplayPriority int    `json:"display_priority"`
}

type AlternativeTitle struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Title     string `json:"title"`
	Type      string `json:"type"`
}

type PersonBase struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	OriginalName string `json:"original_name"`
	Gender       int    `json:"gender"`
	ProfilePath  string `json:"profile_path"`
}

type Person struct {
	PersonBase
	Adult              bool    `json:"adult"`
	Popularity         float32 `json:"popularity"`
	KnownForDepartment string  `json:"known_for_department"`
}

type PersonMedia struct {
	Person
	MediaType MediaType `json:"media_type"`
}

type PersonResults struct {
	PersonMedia
	KnownFor []Media `json:"known_for"`
}

type Media interface {
	GetMediaType() MediaType
}

func (movie MovieMedia) GetMediaType() MediaType { return movie.MediaType }
func (tv TVShowMedia) GetMediaType() MediaType   { return tv.MediaType }

type CreditMedia interface {
	GetMediaType() MediaType
}

func (movie CreditMovieMedia) GetMediaType() MediaType { return movie.MediaType }
func (tv CreditTVShowMedia) GetMediaType() MediaType   { return tv.MediaType }

type CreditTVShowMedia struct {
	TVShowMedia
	Character string           `json:"character"`
	Episodes  []TVEpisodeMedia `json:"episodes"`
	Seasons   []TVSeasonMedia  `json:"seasons"`
}

type CreditMovieMedia struct {
	MovieMedia
	Character string `json:"character"`
}

type VideoBase struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	ID               int64   `json:"id"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	OriginalLanguage string  `json:"original_language"`
	Popularity       float32 `json:"popularity"`
	VoteMetrics
}

type MovieMeta struct {
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	ReleaseDate   string `json:"release_date"`
	Video         bool   `json:"video"`
}

type Movie struct {
	VideoBase
	MovieMeta
	GenreIDs []int64 `json:"genre_ids"`
}

type MovieMedia struct {
	Movie
	MediaType MediaType `json:"media_type"`
}

type MediaType string

const (
	MediaTypeTV    MediaType = "tv"
	MediaTypeMovie MediaType = "movie"
)

type TVShowMeta struct {
	Name          string   `json:"name"`
	OriginalName  string   `json:"original_name"`
	FirstAirDate  string   `json:"first_air_date"`
	OriginCountry []string `json:"origin_country"`
}

type TVShow struct {
	VideoBase
	TVShowMeta
	GenreIDs []int64 `json:"genre_ids"`
}

type TVShowMedia struct {
	TVShow
	MediaType MediaType `json:"media_type"`
}

type VideoDetails struct {
	VideoBase
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
}

type TVEpisode struct {
	ScreenedTheatrically
	Name           string `json:"name"`
	Overview       string `json:"overview"`
	AirDate        string `json:"air_date"`
	EpisodeType    string `json:"episode_type"`
	ProductionCode string `json:"production_code"`
	Runtime        int    `json:"runtime"`
	ShowID         int64  `json:"show_id"`
	StillPath      string `json:"still_path"`
	VoteMetrics
}

type TVEpisodeMedia struct {
	TVEpisode
	MediaType MediaType `json:"media_type"`
}

type TVSeason struct {
	AirDate      string  `json:"air_date"`
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	VoteAverage  float32 `json:"vote_average"`
}

type TVSeasonMedia struct {
	TVSeason
	EpisodeCount int       `json:"episode_count"`
	MediaType    MediaType `json:"media_type"`
	ShowID       int64     `json:"show_id"`
}

type PersonCredit struct {
	Person
	CreditID string `json:"credit_id"`
}

type Crew struct {
	PersonCredit
	Department string `json:"department"`
	Job        string `json:"job"`
}

type CastBase struct {
	PersonCredit
	Character string `json:"character"`
	Order     int    `json:"order"`
}

type Cast struct {
	CastBase
	CastID int `json:"cast_id"`
}

type PeopleMovieCreditCast struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDs         []int64 `json:"genre_ids"`
	ID               int64   `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteMetrics
	Character string `json:"character"`
	CreditID  string `json:"credit_id"`
	Order     int    `json:"order"`
}

type PeopleMovieCreditCrew struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDs         []int64 `json:"genre_ids"`
	ID               int64   `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	CreditID         string  `json:"credit_id"`
	Department       string  `json:"department"`
	Job              string  `json:"job"`
	VoteMetrics
}

type PeopleTVCreditCrew struct {
	ID               int64    `json:"id"`
	Department       string   `json:"department"`
	OriginalLanguage string   `json:"original_language"`
	EpisodeCount     int      `json:"episode_count"`
	Job              string   `json:"job"`
	Overview         string   `json:"overview"`
	OriginCountry    []string `json:"origin_country"`
	OriginalName     string   `json:"original_name"`
	GenreIDs         []int64  `json:"genre_ids"`
	Name             string   `json:"name"`
	FirstAirDate     string   `json:"first_air_date"`
	BackdropPath     string   `json:"backdrop_path"`
	Popularity       float32  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	CreditID         string   `json:"credit_id"`
	VoteMetrics
}
type PeopleTVCreditCast struct {
	CreditID         string   `json:"credit_id"`
	OriginalName     string   `json:"original_name"`
	ID               int64    `json:"id"`
	GenreIDs         []int64  `json:"genre_ids"`
	Character        string   `json:"character"`
	Name             string   `json:"name"`
	PosterPath       string   `json:"poster_path"`
	Popularity       float32  `json:"popularity"`
	EpisodeCount     int      `json:"episode_count"`
	OriginalLanguage string   `json:"original_language"`
	FirstAirDate     string   `json:"first_air_date"`
	BackdropPath     string   `json:"backdrop_path"`
	Overview         string   `json:"overview"`
	OriginCountry    []string `json:"origin_country"`
	VoteMetrics
}

type TVEpisodeDetailsBase struct {
	TVEpisode
	Crew       []Crew     `json:"crew"`
	GuestStars []CastBase `json:"guest_stars"`
}

type CompanyInfo struct {
	Name          string `json:"name"`
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

type Job struct {
	CreditID     string `json:"credit_id"`
	Job          string `json:"job"`
	EpisodeCount int    `json:"episode_count"`
}

type Role struct {
	CreditID     string `json:"credit_id"`
	Character    string `json:"character"`
	EpisodeCount int    `json:"episode_count"`
}

type CrewCredit struct {
	Person
	Department        string `json:"department"`
	Jobs              []Job  `json:"jobs"`
	TotalEpisodeCount int    `json:"total_episode_count"`
}

type CastCredit struct {
	Person
	Order             int    `json:"order"`
	Roles             []Role `json:"roles"`
	TotalEpisodeCount int    `json:"total_episode_count"`
}

type Change struct {
	Key   string       `json:"key"`
	Items []ChangeItem `json:"items"`
}

type ChangeItem struct {
	ID            string           `json:"id"`
	Action        string           `json:"action"`
	Time          string           `json:"time"`
	Iso639_1      string           `json:"iso_639_1"`
	Iso3166_1     string           `json:"iso_3166_1"`
	Value         *json.RawMessage `json:"value"`
	OriginalValue *json.RawMessage `json:"original_value"`
}

type ContentRatings struct {
	Descriptors []any  `json:"descriptors"`
	Iso3166_1   string `json:"iso_3166_1"`
	Rating      string `json:"rating"`
}

type TVEpisodeGroup struct {
	Description  string  `json:"description"`
	EpisodeCount int     `json:"episode_count"`
	GroupCount   int     `json:"group_count"`
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Network      Network `json:"network"`
	Type         int     `json:"type"`
}

type TVEpisodeOrder struct {
	TVEpisode
	Order int `json:"order"`
}

type Group struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Order    int              `json:"order"`
	Episodes []TVEpisodeOrder `json:"episodes"`
	Locked   bool             `json:"locked"`
}

type Review struct {
	Author        string `json:"author"`
	AuthorDetails Author `json:"author_details"`
	Content       string `json:"content"`
	CreatedAt     string `json:"created_at"`
	ID            string `json:"id"`
	UpdatedAt     string `json:"updated_at"`
	URL           string `json:"url"`
}

type Author struct {
	Name       string  `json:"name"`
	Username   string  `json:"username"`
	AvatarPath string  `json:"avatar_path"`
	Rating     float32 `json:"rating"`
}

type ScreenedTheatrically struct {
	ID            int64 `json:"id"`
	EpisodeNumber int   `json:"episode_number"`
	SeasonNumber  int   `json:"season_number"`
}

type List struct {
	Description   string `json:"description"`
	FavoriteCount int64  `json:"favorite_count"`
	ID            int64  `json:"id"`
	ItemCount     int64  `json:"item_count"`
	Iso639_1      string `json:"iso_639_1"`
	ListType      string `json:"list_type"`
	Name          string `json:"name"`
	PosterPath    string `json:"poster_path"`
}

type MovieList struct {
	List
	Iso3166_1 string `json:"iso_3166_1"`
}

type Country struct {
	Iso3166_1   string `json:"iso_3166_1"`
	EnglishName string `json:"english_name"`
	NativeName  string `json:"native_name"`
}

type ReleaseDate struct {
	Certification string `json:"certification"`
	Iso639_1      string `json:"iso_639_1"`
	Note          string `json:"note"`
	ReleaseDate   string `json:"release_date"`
	Type          int    `json:"type"`
}

type MovieResultDate struct {
	Iso3166_1    string        `json:"iso_3166_1"`
	ReleaseDates []ReleaseDate `json:"release_dates"`
}
