package controller

import (
	"fmt"
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
	var shopifyData *models.ShopifySyncData
	var squareData *models.SquareSyncData
	wg := sync.WaitGroup{}
	wg.Add(2)

	// TODO : return objects in channels

	// get all shopify sync data
	go func(shd *models.ShopifySyncData) {
		defer wg.Done()
		tempShd, e := c.GetShopifyData()
		if err != nil {
			err = e
		} else {
			shd = tempShd
		}
	}(shopifyData)

	// get all square sync data
	go func(sqd *models.SquareSyncData) {
		defer wg.Done()
		tempSqd, e := c.GetSquareData()
		if err != nil {
			err = e
		} else {
			sqd = tempSqd
		}
	}(squareData)

	wg.Wait()
	if err != nil {
		return nil, err
	}

	fmt.Print(shopifyData)
	fmt.Print(squareData)

	// handle comparison of shopifyData to squareData

	// update existing data

	// write new data

	return nil, nil

}

func (c *Controller) GetShopifyData() (*models.ShopifySyncData, error) {

	var err error
	data := models.ShopifySyncData{}

	// create wait group for each action

	wg := sync.WaitGroup{}
	wg.Add(3)

	// get products
	go func(d *models.ShopifySyncData) {
		defer wg.Done()
		products, e := c.ShopifyRepo.GetProducts()
		if e != nil {
			err = e
		} else {
			d.Products = products
		}
	}(&data)

	// get customers
	go func(d *models.ShopifySyncData) {
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
			d.Customers = customers
		}
	}(&data)

	// get orders
	go func(d *models.ShopifySyncData) {
		defer wg.Done()
		opts := goshopify.OrderListOptions{
			Status:       "any",
			CreatedAtMax: time.Now(),
		}
		orders, e := c.ShopifyRepo.GetAllOrders([]goshopify.Order{}, opts)
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
			d.Catgories = categories
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
