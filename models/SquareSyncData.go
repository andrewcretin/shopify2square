package models

import (
	"github.com/andrewcretin/shopify2square/models/square"
)

type SquareSyncData struct {
	Products   []square.SquareItem     `json:"products"`
	Categories []square.SquareCategory `json:"categories"`
	Customers  []square.SquareCustomer `json:"customers"`
	Orders     []square.SquareOrder    `json:"orders"`
}
