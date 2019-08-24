package square

import "time"

type SquareItemVariationResponse struct {
	Type                  string              `json:"type"`
	ID                    string              `json:"id"`
	UpdatedAt             time.Time           `json:"updated_at"`
	Version               int64               `json:"version"`
	IsDeleted             bool                `json:"is_deleted"`
	PresentAtAllLocations bool                `json:"present_at_all_locations"`
	ItemVariationData     SquareItemVariation `json:"item_variation_data"`
}
