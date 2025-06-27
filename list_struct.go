package tmdb

// PaginatedListDetails type is a struct for details JSON response.
type PaginatedListDetails struct {
	List
	CreatedBy string       `json:"created_by"`
	Items     []MovieMedia `json:"items"`
	PaginatedResultsMeta
}

// ListItemStatus type is a struct for item status JSON response.
type ListItemStatus struct {
	ID          string `json:"id"`
	ItemPresent bool   `json:"item_present"`
}

// ListResponse type is a struct for list creation JSON response.
type IDListResponse struct {
	ListResponse
	ListID int64 `json:"list_id"`
}

type ListResponse struct {
	Response
	Success bool `json:"success"`
}

// ListCreate type is a struct for list creation JSON request.
type ListCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
}

// ListMedia type is a struct for list media JSON request.
type ListMedia struct {
	MediaID int64 `json:"media_id"`
}
