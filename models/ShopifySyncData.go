package models

import goshopify "github.com/andrewcretin/shopify2square/src/github.com/bold-commerce/go-shopify"

type ShopifySyncData struct {
	Products  []goshopify.Product  `json:"products"`
	Customers []goshopify.Customer `json:"products"`
}
