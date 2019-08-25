package square

type SquareItem struct {
	Name                  string                `json:"name"`
	Description           string                `json:"description"`
	LabelColor            string                `json:"label_color"`
	Visibility            string                `json:"visibility"`
	CategoryID            string                `json:"category_id"`
	VariationCatalogItems []SquareCatalogItem   `json:"variations"`
	Variations            []SquareItemVariation `json:"parsed_variations"`
	ProductType           string                `json:"product_type"`
	SkipModifierScreen    bool                  `json:"skip_modifier_screen"`
}
