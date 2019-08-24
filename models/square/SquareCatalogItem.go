package square

import "time"

type SquareCatalogItem struct {
	Type                  string      `json:"type"`
	ID                    string      `json:"id"`
	UpdatedAt             time.Time   `json:"updated_at"`
	Version               int64       `json:"version"`
	IsDeleted             bool        `json:"is_deleted"`
	PresentAtAllLocations bool        `json:"present_at_all_locations"`
	CategoryData          interface{} `json:"category_data"`
}
