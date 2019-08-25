package square

type SquareCardDetails struct {
	Status      string     `json:"status"`
	Card        SquareCard `json:"card"`
	EntryMethod string     `json:"entry_method"`
}

type SquareCard struct {
	CardBrand string `json:"card_brand"`
	Last4     string `json:"last_4"`
}
