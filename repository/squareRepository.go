package repository

import (
	"encoding/json"
	"fmt"
	"github.com/andrewcretin/shopify2square/envConfig"
	"github.com/andrewcretin/shopify2square/models/square"
	"github.com/parnurzeal/gorequest"
)

type SquareRepository struct {
}

func NewSquareRepository() (*SquareRepository, error) {

	repo := SquareRepository{}
	return &repo, nil

}

func (r *SquareRepository) ParseResponse(resp gorequest.Response, body string, responseKey *string, errs []error, results interface{}) error {

	if len(errs) > 0 {
		return errs[0]
	}
	if resp.StatusCode == 404 {
		return fmt.Errorf("could not find any results")
	}

	if responseKey != nil {
		baseObjectMap := map[string]interface{}{}
		err := json.Unmarshal([]byte(body), &baseObjectMap)
		if err != nil {
			return err
		}
		objectInterface, ok := baseObjectMap[*responseKey]
		objectInterfaceBytes, err := json.Marshal(objectInterface)
		if err != nil {
			return err
		} else if ok {
			err := json.Unmarshal(objectInterfaceBytes, &results)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unable to map response from key: %s", *responseKey)
		}
	} else {
		err := json.Unmarshal([]byte(body), &results)
		if err != nil {
			return err
		}
	}

	return nil

}

func (r *SquareRepository) GetCustomers() ([]square.SquareCustomer, error) {

	url := "https://connect.squareup.com/v2/customers"
	authToken := fmt.Sprintf("Bearer %s", envConfig.CurrentEnvironment().SquareAccessToken)

	request := gorequest.New()
	resp, body, errs := request.Get(url).
		Set("Accept", `application/json`).
		Set("Authorization", authToken).
		End()

	var customers []square.SquareCustomer
	responseKey := "customers"
	err := r.ParseResponse(resp, body, &responseKey, errs, &customers)
	if err != nil {
		return nil, err
	}
	return customers, nil

}
