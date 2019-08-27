package httpClient

import (
	"encoding/json"
	"fmt"
	"github.com/andrewcretin/shopify2square/envConfig"
	"github.com/andrewcretin/shopify2square/httpClient/models"
	"github.com/andrewcretin/shopify2square/models/square"
	"github.com/parnurzeal/gorequest"
	"github.com/satori/go.uuid"
)

func ParseSquareResponse(resp gorequest.Response, body string, responseKeys []string, errs []error, results ...interface{}) error {

	if len(responseKeys) > 0 && len(results) != len(responseKeys) {
		return fmt.Errorf("must provide one result interface for every given response key")
	}

	if len(errs) > 0 {
		return errs[0]
	}
	if resp.StatusCode == 404 {
		return fmt.Errorf("could not find any results")
	}

	baseObjectMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(body), &baseObjectMap)
	if err != nil {
		return err
	}

	var errors []models.SquareError
	errorsInterface, ok := baseObjectMap["errors"]
	errorInterfaceBytes, err := json.Marshal(errorsInterface)
	if ok {
		err := json.Unmarshal(errorInterfaceBytes, &errors)
		if err != nil {
			return err
		}
	}

	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("\n encountered unexpected error: %+v \n", errors[i])
		}
		return fmt.Errorf(errors[0].Detail)
	}

	if len(responseKeys) > 0 {
		for i := range responseKeys {
			objectInterface, ok := baseObjectMap[responseKeys[i]]
			objectInterfaceBytes, err := json.Marshal(objectInterface)
			if err != nil {
				return err
			} else if ok {
				err := json.Unmarshal(objectInterfaceBytes, &results[i])
				if err != nil {
					return err
				}
			} else {
				fmt.Printf("\n unable to map response from key: %s \n", responseKeys[i])
			}
		}
	} else {
		err := json.Unmarshal([]byte(body), &results)
		if err != nil {
			return err
		}
	}

	return nil

}

func GetSquareCustomers(cursor *string) ([]square.SquareCustomer, *string, error) {

	url := "https://connect.squareup.com/v2/customers"
	if cursor != nil {
		url = fmt.Sprintf("%s?cursor=%s", url, *cursor)
	}
	authToken := fmt.Sprintf("Bearer %s", envConfig.CurrentEnvironment().SquareAccessToken)

	resp, body, errs := gorequest.New().Get(url).
		Set("Accept", `application/json`).
		Set("Authorization", authToken).
		End()

	var customers []square.SquareCustomer
	var respCursor *string
	responseInterfaces := []interface{}{&customers, &respCursor}
	keys := []string{"customers", "cursor"}
	err := ParseSquareResponse(resp, body, keys, errs, responseInterfaces...)
	if err != nil {
		return nil, nil, err
	}
	return customers, respCursor, nil

}

func GetSquareCatalog(cursor *string) ([]square.SquareCatalogItem, *string, error) {

	url := "https://connect.squareup.com/v2/catalog/list"
	if cursor != nil {
		url = fmt.Sprintf("%s?cursor=%s", url, *cursor)
	}
	authToken := fmt.Sprintf("Bearer %s", envConfig.CurrentEnvironment().SquareAccessToken)

	resp, body, errs := gorequest.New().Get(url).
		Set("Accept", `application/json`).
		Set("Authorization", authToken).
		End()

	var catalogItems []square.SquareCatalogItem
	var respCursor *string
	responseInterfaces := []interface{}{&catalogItems, &respCursor}
	keys := []string{"objects", "cursor"}
	err := ParseSquareResponse(resp, body, keys, errs, responseInterfaces...)
	if err != nil {
		return nil, nil, err
	}
	return catalogItems, respCursor, nil

}

func GetLocations() ([]square.SquareLocation, error) {

	url := "https://connect.squareup.com/v2/locations"
	authToken := fmt.Sprintf("Bearer %s", envConfig.CurrentEnvironment().SquareAccessToken)

	resp, body, errs := gorequest.New().Get(url).
		Set("Accept", `application/json`).
		Set("Authorization", authToken).
		End()

	var locations []square.SquareLocation
	responseInterfaces := []interface{}{&locations}
	keys := []string{"locations"}
	err := ParseSquareResponse(resp, body, keys, errs, responseInterfaces...)
	if err != nil {
		return nil, err
	}
	return locations, nil

}

func BatchGetOrders(locationId string, orderIds []string) ([]square.SquareOrder, error) {

	url := fmt.Sprintf("https://connect.squareup.com/v2/locations/%s/orders/batch-retrieve", locationId)
	authToken := fmt.Sprintf("Bearer %s", envConfig.CurrentEnvironment().SquareAccessToken)

	bodyData := struct {
		OrderIds []string `json:"order_ids"`
	}{
		OrderIds: orderIds,
	}

	resp, body, errs := gorequest.New().Post(url).
		Set("Accept", `application/json`).
		Set("Authorization", authToken).
		Send(bodyData).
		End()

	var orders []square.SquareOrder
	responseInterfaces := []interface{}{&orders}
	keys := []string{"orders"}
	err := ParseSquareResponse(resp, body, keys, errs, responseInterfaces...)
	if err != nil {
		return nil, err
	}
	return orders, nil

}

func SearchOrders(req models.SearchOrderReq) (*models.SearchResp, error) {

	url := "https://connect.squareup.com/v2/orders/search"
	authToken := fmt.Sprintf("Bearer %s", envConfig.CurrentEnvironment().SquareAccessToken)

	resp, body, errs := gorequest.New().Post(url).
		Set("Accept", `application/json`).
		Set("Authorization", authToken).
		Send(req).
		End()

	searchResponse := models.SearchResp{}
	var err error
	var respCursor *string
	if req.ReturnEntries {
		// returns order entry
		var orderEntries []square.SquareOrderEntry
		responseInterfaces := []interface{}{&orderEntries, &respCursor}
		keys := []string{"order_entries", "cursor"}
		err = ParseSquareResponse(resp, body, keys, errs, responseInterfaces...)
		if err != nil {
			return nil, err
		}
		searchResponse.OrderEntries = orderEntries
	} else {
		// returns complete order
		var orders []square.SquareOrder
		responseInterfaces := []interface{}{&orders, &respCursor}
		keys := []string{"orders", "cursor"}
		err = ParseSquareResponse(resp, body, keys, errs, responseInterfaces...)
		if err != nil {
			return nil, err
		}
		searchResponse.Orders = orders
	}
	searchResponse.Cursor = respCursor

	return &searchResponse, nil

}

func WriteCustomer(c square.SquareCustomer) error {

	url := "https://connect.squareup.com/v2/customers"
	authToken := fmt.Sprintf("Bearer %s", envConfig.CurrentEnvironment().SquareAccessToken)

	respBody := struct {
		square.SquareCustomer
		IdempotencyKey string `json:"idempotency_key"`
	}{
		c,
		uuid.NewV4().String(),
	}

	resp, body, errs := gorequest.New().Post(url).
		Set("Accept", `application/json`).
		Set("Authorization", authToken).
		Send(respBody).
		End()

	var newCustomer square.SquareCustomer
	responseInterfaces := []interface{}{&newCustomer}
	keys := []string{"customer"}
	err := ParseSquareResponse(resp, body, keys, errs, responseInterfaces...)
	if err != nil {
		return err
	}
	fmt.Print("\n successfully wrote new customer \n")
	return nil

}
