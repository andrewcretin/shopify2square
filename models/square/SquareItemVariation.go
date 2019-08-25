package square

type SquareItemVariation struct {
	ItemID            string                   `json:"item_id"`
	Name              string                   `json:"name"`
	Sku               string                   `json:"sku"`
	Ordinal           int                      `json:"ordinal"`
	PricingType       string                   `json:"pricing_type"`
	PriceMoney        SquareMoney              `json:"price_money"`
	LocationOverrides []SquareLocationOverride `json:"location_overrides"`
}
