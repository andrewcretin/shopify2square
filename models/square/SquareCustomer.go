package square

import "time"

type SquareCustomer struct {
	ID           string        `json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	GivenName    string        `json:"given_name"`
	FamilyName   string        `json:"family_name"`
	EmailAddress string        `json:"email_address"`
	Address      SquareAddress `json:"address"`
	PhoneNumber  string        `json:"phone_number"`
	ReferenceID  string        `json:"reference_id"`
	Note         string        `json:"note"`
}

type SquareCustomerUpdate struct {
	Item              SquareCustomer `json:"item"`
	UpdatedProperties []string       `json:"updatedProperties"`
}
