package models

import (
	"github.com/andrewcretin/shopify2square/models/square"
)

type SquareSyncData struct {
	//Products  []goshopify.Product  `json:"products"`
	Customers []square.SquareCustomer `json:"customers"`
	//Orders    []goshopify.Order    `json:"orders"`
}
