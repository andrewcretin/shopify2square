package controller

import (
	"fmt"
	"github.com/andrewcretin/shopify2square/models"
	"github.com/andrewcretin/shopify2square/repository"
)

type Controller struct {
	ShopifyRepo *repository.ShopifyRepository
	SquareRepo  *repository.SquareRepository
}

func NewController() (*Controller, error) {

	shopifyRepo, err := repository.NewShopifyRepository()
	if err != nil {
		return nil, err
	}
	squareRepo, err := repository.NewSquareRepository()
	if err != nil {
		return nil, err
	}

	c := Controller{
		ShopifyRepo: shopifyRepo,
		SquareRepo:  squareRepo,
	}
	return &c, nil

}

func (c *Controller) SyncShopifyToSquare() (*models.SyncResponse, error) {

	shopifyData, err := c.GetShopifyData()
	if err != nil {
		return nil, err
	}
	fmt.Print(shopifyData)

	return nil, nil

}

func (c *Controller) GetShopifyData() (*models.ShopifySyncData, error) {

	var data models.ShopifySyncData

	// get products
	products, err := c.ShopifyRepo.GetProducts()
	if err != nil {
		return nil, err
	} else {
		data.Products = products
	}

	// get customers
	customers, err := c.ShopifyRepo.GetCustomers()
	if err != nil {
		return nil, err
	} else {
		// get customer addresses
		err = c.ShopifyRepo.GetCustomerAddresses(customers)
		if err != nil {
			return nil, err
		}
		data.Customers = customers
	}

	return &data, nil

}
