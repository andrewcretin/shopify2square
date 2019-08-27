package square

import (
	"fmt"
	goshopify "github.com/andrewcretin/shopify2square/src/github.com/bold-commerce/go-shopify"
	"time"
)

type SquareCustomer struct {
	Id           string        `json:"id"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	GivenName    string        `json:"given_name,omitempty"`
	FamilyName   string        `json:"family_name,omitempty"`
	EmailAddress string        `json:"email_address,omitempty"`
	Address      SquareAddress `json:"address,omitempty"`
	PhoneNumber  string        `json:"phone_number,omitempty"`
	Note         string        `json:"note,omitempty"`
}

type SquareCustomerUpdate struct {
	Object            SquareCustomer `json:"object"`
	UpdatedProperties []string       `json:"updatedProperties"`
}

func (c *SquareCustomer) InitFromShopifyCustomer(cc goshopify.Customer) {

	c.Id = fmt.Sprintf("%d", cc.ID)
	if cc.CreatedAt != nil {
		c.CreatedAt = *cc.CreatedAt
	}
	if cc.UpdatedAt != nil {
		c.UpdatedAt = *cc.UpdatedAt
	}
	c.GivenName = cc.FirstName
	c.FamilyName = cc.LastName
	c.EmailAddress = cc.Email
	if len(cc.Addresses) > 0 && cc.Addresses[0] != nil {
		c.Address = SquareAddress{}
		c.Address.InitFromShopifyAddress(*cc.Addresses[0])
	} else {
		c.Address = DefaultSquareAddress()
	}
	c.PhoneNumber = cc.Phone
	c.Note = cc.Note

}

func (c *SquareCustomer) UpdatedProperties(cc SquareCustomer) []string {
	var updatedProperties []string
	if c.Id != cc.Id {
		updatedProperties = append(updatedProperties, "Id")
	}
	if !c.CreatedAt.Equal(cc.CreatedAt) {
		updatedProperties = append(updatedProperties, "CreatedAt")
	}
	if !c.UpdatedAt.Equal(cc.UpdatedAt) {
		updatedProperties = append(updatedProperties, "UpdatedAt")
	}
	if c.GivenName != cc.GivenName {
		updatedProperties = append(updatedProperties, "GivenName")
	}
	if c.FamilyName != cc.FamilyName {
		updatedProperties = append(updatedProperties, "FamilyName")
	}
	if c.EmailAddress != cc.EmailAddress {
		updatedProperties = append(updatedProperties, "EmailAddress")
	}
	if c.PhoneNumber != cc.PhoneNumber {
		updatedProperties = append(updatedProperties, "PhoneNumber")
	}
	if c.Note != cc.Note {
		updatedProperties = append(updatedProperties, "Note")
	}
	if len(c.Address.UpdatedProperties(cc.Address)) > 0 {
		updatedProperties = append(updatedProperties, "Address")
	}
	return updatedProperties
}
