package repository

import (
	"encoding/json"
	"fmt"
	"github.com/andrewcretin/shopify2square/httpClient"
	"github.com/andrewcretin/shopify2square/httpClient/models"
	"github.com/andrewcretin/shopify2square/models/square"
	"github.com/joncalhoun/drip"
	"github.com/parnurzeal/gorequest"
	"sync"
	"time"
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
			fmt.Printf("\n unable to map response from key: %s \n", *responseKey)
		}
	} else {
		err := json.Unmarshal([]byte(body), &results)
		if err != nil {
			return err
		}
	}

	return nil

}

func (r *SquareRepository) GetAllCustomers(customers []square.SquareCustomer, cursor *string) ([]square.SquareCustomer, error) {
	tempCustomers, tempCursor, err := httpClient.GetSquareCustomers(cursor)
	if err != nil {
		return nil, err
	} else {
		customers = append(customers, tempCustomers...)
		if tempCursor != nil {
			return r.GetAllCustomers(customers, tempCursor)
		} else {
			return customers, nil
		}
	}
}

func (r *SquareRepository) GetProductsAndCategories() ([]square.SquareItem, []square.SquareCategory, error) {
	productCatalogItems, err := r.GetEntireProductCatalog(nil, nil)
	if err != nil {
		return nil, nil, err
	} else {
		// parse product catalog items
		var items []square.SquareItem
		var categories []square.SquareCategory
		itemTypes := []square.SquareCatalogItemType{square.ITEM, square.CATEGORY}
		interfaces := []interface{}{&items, &categories}
		err = r.ParseCatalogItems(productCatalogItems, itemTypes, interfaces...)
		if err != nil {
			return nil, nil, err
		}
		for i := range items {
			var variations []square.SquareItemVariation
			err = r.ParseCatalogItems(items[i].VariationCatalogItems, []square.SquareCatalogItemType{square.ITEM_VARIATION}, &variations)
			if err != nil {
				return nil, nil, err
			}
			items[i].Variations = variations
		}
		return items, categories, nil
	}
}

func (r *SquareRepository) GetEntireProductCatalog(catalogItems []square.SquareCatalogItem, cursor *string) ([]square.SquareCatalogItem, error) {
	tempCatalogItems, tempCursor, err := httpClient.GetSquareCatalog(cursor)
	if err != nil {
		return nil, err
	} else {
		catalogItems = append(catalogItems, tempCatalogItems...)
		if tempCursor != nil {
			return r.GetEntireProductCatalog(catalogItems, tempCursor)
		} else {
			return catalogItems, nil
		}
	}
}

func (r *SquareRepository) ParseCatalogItems(catalogItems []square.SquareCatalogItem, itemTypes []square.SquareCatalogItemType, results ...interface{}) error {

	if len(itemTypes) != len(results) {
		return fmt.Errorf("must provide one result interface for every given item type")
	}

	for i := range itemTypes {
		var filteredCatalogItemInterfaces []interface{}
		for j := range catalogItems {
			if catalogItems[j].Type == itemTypes[i] {
				switch itemTypes[i] {
				case square.CATEGORY:
					filteredCatalogItemInterfaces = append(filteredCatalogItemInterfaces, catalogItems[j].CategoryData)
				case square.ITEM:
					filteredCatalogItemInterfaces = append(filteredCatalogItemInterfaces, catalogItems[j].ItemData)
				case square.ITEM_VARIATION:
					filteredCatalogItemInterfaces = append(filteredCatalogItemInterfaces, catalogItems[j].ItemVariationData)
				}
			}
		}
		bytes, err := json.Marshal(filteredCatalogItemInterfaces)
		if err != nil {
			return fmt.Errorf("error marshalling item interfaces to bytes. err: %+v", err)
		}
		stringData := string(bytes)
		fmt.Print(stringData)
		err = json.Unmarshal(bytes, results[i])
		if err != nil {
			return fmt.Errorf("error unmarshalling from item interfaces bytes to results interface. err: %+v", err)
		}
	}

	return nil

}

func (r *SquareRepository) BatchGetOrders(locationId string, orderIds []string) ([]square.SquareOrder, error) {

	if len(orderIds) > 100 {
		return nil, fmt.Errorf("only 100 orders can be retrieved at a time")
	}

	orders, err := httpClient.BatchGetOrders(locationId, orderIds)
	if err != nil {
		return nil, err
	}
	return orders, nil

}

func (r *SquareRepository) SearchOrders(req models.SearchOrderReq, resp *models.SearchResp) (*models.SearchResp, error) {

	tempResp, err := httpClient.SearchOrders(req)
	if err != nil {
		return nil, err
	} else {
		// append search results
		if resp != nil {
			resp.Orders = append(resp.Orders, tempResp.Orders...)
			resp.OrderEntries = append(resp.OrderEntries, tempResp.OrderEntries...)
			resp.Cursor = tempResp.Cursor
		} else {
			resp = tempResp
		}

		// recurse if necesssary
		if resp.Cursor != nil {
			req.Cursor = resp.Cursor
			resp.Cursor = nil
			return r.SearchOrders(req, resp)
		} else {
			return resp, nil
		}
	}

}

func (r *SquareRepository) GetAllOrders() ([]square.SquareOrder, error) {

	allLocations, err := httpClient.GetLocations()
	if err != nil {
		return nil, err
	}
	var locationIds []string
	for i := range allLocations {
		locationIds = append(locationIds, allLocations[i].ID)
	}

	req := models.SearchOrderReq{
		ReturnEntries: false,
		Limit:         500,
		LocationIds:   locationIds,
		Cursor:        nil,
	}
	searchResp, err := r.SearchOrders(req, nil)
	if err != nil {
		return nil, err
	}
	return searchResp.Orders, nil

}

func (r *SquareRepository) WriteCustomers(customers []square.SquareCustomer) error {

	var err error
	wg := sync.WaitGroup{}
	wg.Add(len(customers))

	b := drip.Bucket{
		Capacity:     len(customers),
		DripInterval: 1 * time.Second,
		PerDrip:      2,
	}

	for i := range customers {
		go func(c square.SquareCustomer) {
			defer func() {
				err := b.Consume(1)
				if err != nil {
					fmt.Println("Sleep 500ms.")
					time.Sleep(500 * time.Millisecond)
				}
				wg.Done()
			}()
			e := httpClient.WriteCustomer(c)
			if e != nil {
				err = e
			}
		}(customers[i])
	}

	defer func() {
		_ = b.Stop()
	}()
	_ = b.Start()
	wg.Wait()
	if err != nil {
		return err
	}
	return nil

}
