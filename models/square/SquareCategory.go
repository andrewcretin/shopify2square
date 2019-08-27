package square

type SquareCategory struct {
	Name string `json:"name"`
}

type SquareCategoryUpdate struct {
	Object            SquareCategory `json:"object"`
	UpdatedProperties []string       `json:"updatedProperties"`
}
