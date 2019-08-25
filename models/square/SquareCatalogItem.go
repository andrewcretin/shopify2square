package square

import "time"

type SquareCatalogItem struct {
	Type                  SquareCatalogItemType `json:"type"`
	ID                    string                `json:"id"`
	UpdatedAt             time.Time             `json:"updated_at"`
	Version               int64                 `json:"version"`
	IsDeleted             bool                  `json:"is_deleted"`
	PresentAtAllLocations bool                  `json:"present_at_all_locations"`
	CategoryData          interface{}           `json:"category_data"`
	ItemData              interface{}           `json:"item_data"`
	ItemVariationData     interface{}           `json:"item_variation_data"`
}

type SquareCatalogItemType string

//noinspection ALL
const (
	ITEM           SquareCatalogItemType = "ITEM"
	ITEM_VARIATION SquareCatalogItemType = "ITEM_VARIATION"
	CATEGORY       SquareCatalogItemType = "CATEGORY"
)
