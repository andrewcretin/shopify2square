package models

type SquareError struct {
	Category string `json:"category"`
	Code     string `json:"code"`
	Detail   string `json:"detail"`
	Field    string `json:"field"`
}
