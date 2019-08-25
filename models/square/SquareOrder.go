package square

import "time"

type SquareOrder struct {
	ID                      string                 `json:"id"`
	LocationID              string                 `json:"location_id"`
	LineItems               []SquareLineItem       `json:"line_items"`
	CreatedAt               time.Time              `json:"created_at"`
	UpdatedAt               time.Time              `json:"updated_at"`
	State                   string                 `json:"state"`
	TotalTaxMoney           SquareMoney            `json:"total_tax_money"`
	TotalDiscountMoney      SquareMoney            `json:"total_discount_money"`
	TotalTipMoney           SquareMoney            `json:"total_tip_money"`
	TotalMoney              SquareMoney            `json:"total_money"`
	ClosedAt                time.Time              `json:"closed_at"`
	Tenders                 []SquareTender         `json:"tenders"`
	TotalServiceChargeMoney SquareMoney            `json:"total_service_charge_money"`
	ReturnAmounts           SquareOrderMoneyAmount `json:"return_amounts"`
	NetAmounts              SquareOrderMoneyAmount `json:"net_amounts"`
	CustomerID              string                 `json:"customer_id"`
}
