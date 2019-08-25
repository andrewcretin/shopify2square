package models

import "github.com/andrewcretin/shopify2square/models/square"

type SearchResp struct {
	OrderEntries []square.SquareOrderEntry `json:"order_entries,omitempty"`
	Orders       []square.SquareOrder      `json:"orders,omitempty"`
	Cursor       *string                   `json:"cursor,omitempty"`
}
