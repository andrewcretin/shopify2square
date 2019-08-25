package square

import "time"

type SquareLocation struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Address      SquareAddress `json:"address"`
	Timezone     string        `json:"timezone"`
	Capabilities []string      `json:"capabilities"`
	Status       string        `json:"status"`
	CreatedAt    time.Time     `json:"created_at"`
	MerchantID   string        `json:"merchant_id"`
	Country      string        `json:"country"`
	LanguageCode string        `json:"language_code"`
	Currency     string        `json:"currency"`
	PhoneNumber  string        `json:"phone_number"`
	BusinessName string        `json:"business_name"`
}
