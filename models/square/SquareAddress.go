package square

import goshopify "github.com/andrewcretin/shopify2square/src/github.com/bold-commerce/go-shopify"

type SquareAddress struct {
	AddressLine1                 string `json:"address_line_1,omitempty"`
	AddressLine2                 string `json:"address_line_2,omitempty"`
	Locality                     string `json:"locality,omitempty"`
	PostalCode                   string `json:"postal_code,omitempty"`
	AdministrativeDistrictLevel1 string `json:"administrative_district_level_1,omitempty"`
	Country                      string `json:"country,omitempty"`
	FirstName                    string `json:"firstName,omitempty"`
	LastName                     string `json:"lastName,omitempty"`
	Organization                 string `json:"organization,omitempty"`
}

func DefaultSquareAddress() SquareAddress {
	return SquareAddress{
		Country: DEFAULT_COUNTRY_CODE,
	}
}

func (a *SquareAddress) InitFromShopifyAddress(sh goshopify.CustomerAddress) {

	a.AddressLine1 = sh.Address1
	a.AddressLine2 = sh.Address2
	a.Locality = sh.City
	a.PostalCode = sh.Zip
	a.AdministrativeDistrictLevel1 = sh.Province
	a.Country = sh.CountryCode
	if a.Country == "" {
		// set default country code, as it is required to write square api
		a.Country = DEFAULT_COUNTRY_CODE
	}
	a.FirstName = sh.FirstName
	a.LastName = sh.LastName
	a.Organization = sh.Company

}

func (a *SquareAddress) UpdatedProperties(aa SquareAddress) []string {
	var updatedProperties []string
	if a.AddressLine1 != aa.AddressLine1 {
		updatedProperties = append(updatedProperties, "AddressLine1")
	}
	if a.AddressLine2 != aa.AddressLine2 {
		updatedProperties = append(updatedProperties, "AddressLine2")
	}
	if a.Locality != aa.Locality {
		updatedProperties = append(updatedProperties, "Locality")
	}
	if a.PostalCode != aa.PostalCode {
		updatedProperties = append(updatedProperties, "PostalCode")
	}
	if a.AdministrativeDistrictLevel1 != aa.AdministrativeDistrictLevel1 {
		updatedProperties = append(updatedProperties, "AdministrativeDistrictLevel1")
	}
	if a.Country != aa.Country {
		updatedProperties = append(updatedProperties, "Country")
	}
	if a.FirstName != aa.FirstName {
		updatedProperties = append(updatedProperties, "FirstName")
	}
	if a.LastName != aa.LastName {
		updatedProperties = append(updatedProperties, "LastName")
	}
	if a.Organization != aa.Organization {
		updatedProperties = append(updatedProperties, "Organization")
	}
	return updatedProperties
}
