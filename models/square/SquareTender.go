package square

import "time"

type SquareTender struct {
	ID                 string            `json:"id"`
	LocationID         string            `json:"location_id"`
	TransactionID      string            `json:"transaction_id"`
	CreatedAt          time.Time         `json:"created_at"`
	AmountMoney        SquareMoney       `json:"amount_money"`
	ProcessingFeeMoney SquareMoney       `json:"processing_fee_money"`
	CustomerID         string            `json:"customer_id"`
	Type               string            `json:"type"`
	CardDetails        SquareCardDetails `json:"card_details"`
	TipMoney           SquareMoney       `json:"tip_money"`
}
