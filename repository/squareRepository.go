package repository

type SquareRepository struct {
}

func NewSquareRepository() (*SquareRepository, error) {

	repo := SquareRepository{}
	return &repo, nil

}
