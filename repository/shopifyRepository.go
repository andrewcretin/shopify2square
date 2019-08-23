package repository

import (
	"fmt"
	"github.com/andrewcretin/shopify2square/envConfig"
	goshopify "github.com/andrewcretin/shopify2square/src/github.com/bold-commerce/go-shopify"
)

type ShopifyRepository struct {
	Client goshopify.Client
}

func NewShopifyRepository() (*ShopifyRepository, error) {

	// Create an app somewhere.
	app := goshopify.App{
		ApiKey:   envConfig.CurrentEnvironment().ShopifyKey,
		Password: envConfig.CurrentEnvironment().ShopifyPassword,
	}
	client := goshopify.NewClient(app, "good-spirit-kombucha", "")

	repo := ShopifyRepository{
		Client: *client,
	}
	return &repo, nil

}

func (r *ShopifyRepository) GetShopifyProducts() error {

	products, err := r.Client.Product.List(nil)
	if err != nil {
		return err
	}
	fmt.Print(products)
	return nil

}
