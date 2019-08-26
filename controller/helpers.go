package controller

import (
	"github.com/andrewcretin/shopify2square/models"
	"github.com/andrewcretin/shopify2square/models/square"
)

//noinspection ALL
func ParseCustomerModifications(shd models.ShopifySyncData, sqd models.SquareSyncData) ([]square.SquareCustomer, []square.SquareCustomerUpdate) {
	return nil, nil
}

//noinspection ALL
func ParseProductModifications(shd models.ShopifySyncData, sqd models.SquareSyncData) []square.SquareItem {
	return nil
}

//noinspection ALL
func ParseCategoryModifications(shd models.ShopifySyncData, sqd models.SquareSyncData) []square.SquareCategory {
	return nil
}

//noinspection ALL
func ParseOrderModifications(shd models.ShopifySyncData, sqd models.SquareSyncData) []square.SquareOrder {
	return nil
}
