package controller

import (
	"github.com/andrewcretin/shopify2square/models/square"
	goshopify "github.com/andrewcretin/shopify2square/src/github.com/bold-commerce/go-shopify"
)

func ArrayContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//noinspection ALL
func ParseCustomerModifications(shopifyCustomers []goshopify.Customer, squareCustomers []square.SquareCustomer) ([]square.SquareCustomer, []square.SquareCustomerUpdate) {

	var newCustomers []square.SquareCustomer
	var updatedCustomers []square.SquareCustomerUpdate

	// map all shopify customers to square customers
	var incomingShopifyCustomers []square.SquareCustomer
	for i := range shopifyCustomers {
		tempCustomer := square.SquareCustomer{}
		tempCustomer.InitFromShopifyCustomer(shopifyCustomers[i])
		incomingShopifyCustomers = append(incomingShopifyCustomers, tempCustomer)
	}

	// put square customers into a map
	squareCustomersMap := make(map[string]square.SquareCustomer, len(squareCustomers))
	for i := range squareCustomers {
		squareCustomersMap[squareCustomers[i].ReferenceId] = squareCustomers[i]
	}

	for i := range incomingShopifyCustomers {
		existingCustomer, ok := squareCustomersMap[incomingShopifyCustomers[i].ReferenceId]
		if ok {
			// the incoming shopify customer's id matches an existing square customer, check for update
			updatedProperties := existingCustomer.UpdatedProperties(incomingShopifyCustomers[i])
			if len(updatedProperties) > 0 {
				// updates required
				update := square.SquareCustomerUpdate{
					Object:            existingCustomer,
					UpdatedProperties: updatedProperties,
				}
				updatedCustomers = append(updatedCustomers, update)
			}
		} else {
			newCustomers = append(newCustomers, incomingShopifyCustomers[i])
		}
	}

	return newCustomers, updatedCustomers
}

//noinspection ALL
func ParseProductModifications(shopifyProducts []goshopify.Product, squareItems []square.SquareItem) ([]square.SquareItem, []square.SquareItemUpdate) {
	return nil, nil
}

//noinspection ALL
func ParseCategoryModifications(shopifyProductTypes []string, squareCategories []square.SquareCategory) ([]square.SquareCategory, []square.SquareCategoryUpdate) {
	return nil, nil
}

//noinspection ALL
func ParseOrderModifications(shopifyOrders []goshopify.Order, squareOrders []square.SquareOrder) ([]square.SquareOrder, []square.SquareOrderUpdate) {
	return nil, nil
}
