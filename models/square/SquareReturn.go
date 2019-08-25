package square

type SquareOrderMoneyAmount struct {
	TotalMoney         SquareMoney `json:"total_money"`
	TaxMoney           SquareMoney `json:"tax_money"`
	DiscountMoney      SquareMoney `json:"discount_money"`
	TipMoney           SquareMoney `json:"tip_money"`
	ServiceChargeMoney SquareMoney `json:"service_charge_money"`
}
