package repository

type ShopifyRepository struct {
}

func NewShopifyRepository() (*ShopifyRepository, error) {

	repo := ShopifyRepository{}
	return &repo, nil

}
