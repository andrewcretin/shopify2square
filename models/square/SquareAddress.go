package square

type SquareAddress struct {
	AddressLine1                 string `json:"address_line_1"`
	AddressLine2                 string `json:"address_line_2"`
	Locality                     string `json:"locality"`
	AdministrativeDistrictLevel1 string `json:"administrative_district_level_1"`
	PostalCode                   string `json:"postal_code"`
	Country                      string `json:"country"`
}
