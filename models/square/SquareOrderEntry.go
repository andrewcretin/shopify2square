package square

type SquareOrderEntry struct {
	OrderID    string `json:"order_id"`
	LocationID string `json:"location_id"`
	Version    int    `json:"version,omitempty"`
}
