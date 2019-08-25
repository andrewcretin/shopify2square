package square

type SquareLineItem struct {
	UID            string      `json:"uid"`
	Name           string      `json:"name"`
	Quantity       string      `json:"quantity"`
	BasePriceMoney SquareMoney `json:"base_price_money"`
	TotalMoney     SquareMoney `json:"total_money"`
}
