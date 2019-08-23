package controller

import (
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

	return nil, nil

}
