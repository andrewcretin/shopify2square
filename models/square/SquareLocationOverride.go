package square

type SquareLocationOverride struct {
	LocationID              string `json:"location_id"`
	TrackInventory          bool   `json:"track_inventory"`
	InventoryAlertType      string `json:"inventory_alert_type"`
	InventoryAlertThreshold int    `json:"inventory_alert_threshold"`
}
