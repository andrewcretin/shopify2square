package controller

import (
	"encoding/json"
	"github.com/andrewcretin/shopify2square/models"
	"github.com/andrewcretin/shopify2square/repository"
	goshopify "github.com/andrewcretin/shopify2square/src/github.com/bold-commerce/go-shopify"
	"sync"
	"time"
)

type Controller struct {
	ShopifyRepo *repository.ShopifyRepository
	SquareRepo  *repository.SquareRepository
}

func NewController() (*Controller, error) {

	shopifyRepo, err := repository.NewShopifyRepository()
	if err != nil {
		return nil, err
	}
	squareRepo, err := repository.NewSquareRepository()
	if err != nil {
		return nil, err
	}

	c := Controller{
		ShopifyRepo: shopifyRepo,
		SquareRepo:  squareRepo,
	}
	return &c, nil

}

func (c *Controller) SyncShopifyToSquare() (*models.SyncResponse, error) {

	var err error
	syncMap := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(2)

	// get all shopify sync data
	go func() {
		defer wg.Done()
		shd, e := c.GetShopifyData()
		if err != nil {
			err = e
		} else {
			syncMap.Store("shopifyData", shd)
		}
	}()

	// get all square sync data
	go func() {
		defer wg.Done()
		sqd, e := c.GetSquareData()
		if err != nil {
			err = e
		} else {
			syncMap.Store("squareData", sqd)
		}
	}()

	wg.Wait()
	if err != nil {
		return nil, err
	}

	var shopifyData models.ShopifySyncData
	var squareData models.SquareSyncData
	shopifyDataInterface, ok := syncMap.Load("shopifyData")
	if ok {
		bytes, _ := json.Marshal(shopifyDataInterface)
		_ = json.Unmarshal(bytes, &shopifyData)
	}
	squareDataInterface, ok := syncMap.Load("squareData")
	if ok {
		bytes, _ := json.Marshal(squareDataInterface)
		_ = json.Unmarshal(bytes, &squareData)
	}

	// check for new objects

	// check existing objects for updates

	// write/update objects

	return nil, nil

}

func (c *Controller) GetShopifyData() (*models.ShopifySyncData, error) {

	var err error
	syncMap := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(3)

	// get products
	go func() {
		defer wg.Done()
		products, e := c.ShopifyRepo.GetProducts()
		if e != nil {
			err = e
		} else {
			syncMap.Store("products", products)
		}
	}()

	// get customers
	go func() {
		defer wg.Done()
		customers, e := c.ShopifyRepo.GetCustomers()
		if e != nil {
			err = e
		} else {
			// get customer addresses
			e = c.ShopifyRepo.GetCustomerAddresses(customers)
			if e != nil {
				err = e
			}
			syncMap.Store("customers", customers)
		}
	}()

	// get orders
	go func() {
		defer wg.Done()
		opts := goshopify.OrderListOptions{
			Status:       "any",
			CreatedAtMax: time.Now(),
		}
		orders, e := c.ShopifyRepo.GetAllOrders([]goshopify.Order{}, opts)
		if e != nil {
			err = e
		} else {
			syncMap.Store("orders", orders)
		}
	}()

	wg.Wait()
	if err != nil {
		return nil, err
	}

	data := models.ShopifySyncData{}
	productsInterface, ok := syncMap.Load("products")
	if ok {
		var products []goshopify.Product
		bytes, _ := json.Marshal(productsInterface)
		_ = json.Unmarshal(bytes, &products)
		data.Products = products
	}
	customerInterface, ok := syncMap.Load("customers")
	if ok {
		var customers []goshopify.Customer
		bytes, _ := json.Marshal(customerInterface)
		_ = json.Unmarshal(bytes, &customers)
		data.Customers = customers
	}
	ordersInterface, ok := syncMap.Load("orders")
	if ok {
		var orders []goshopify.Order
		bytes, _ := json.Marshal(ordersInterface)
		_ = json.Unmarshal(bytes, &orders)
		data.Orders = orders
	}

	return &data, nil

}

func (c *Controller) GetSquareData() (*models.SquareSyncData, error) {

	var err error
	data := models.SquareSyncData{}

	// create wait group for each action

	wg := sync.WaitGroup{}
	wg.Add(3)

	//get products
	go func(d *models.SquareSyncData) {
		defer wg.Done()
		products, categories, e := c.SquareRepo.GetProductsAndCategories()
		if e != nil {
			err = e
		} else {
			d.Products = products
			d.Categories = categories
		}
	}(&data)

	// get customers
	go func(d *models.SquareSyncData) {
		defer wg.Done()
		customers, e := c.SquareRepo.GetAllCustomers(nil, nil)
		if e != nil {
			err = e
		} else {
			d.Customers = customers
		}
	}(&data)

	//get orders
	go func(d *models.SquareSyncData) {
		defer wg.Done()
		orders, e := c.SquareRepo.GetAllOrders()
		if e != nil {
			err = e
		} else {
			d.Orders = orders
		}
	}(&data)

	wg.Wait()
	if err != nil {
		return nil, err
	}

	return &data, nil

}
