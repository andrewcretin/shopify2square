package repository

import (
	"fmt"
	"github.com/andrewcretin/shopify2square/envConfig"
	goshopify "github.com/andrewcretin/shopify2square/src/github.com/bold-commerce/go-shopify"
	"sync"
)

type ShopifyRepository struct {
	Client goshopify.Client
}

func NewShopifyRepository() (*ShopifyRepository, error) {

	// Create an app somewhere.
	app := goshopify.App{
		ApiKey:   envConfig.CurrentEnvironment().ShopifyKey,
		Password: envConfig.CurrentEnvironment().ShopifyPassword,
	}
	client := goshopify.NewClient(app, "good-spirit-kombucha", "")

	repo := ShopifyRepository{
		Client: *client,
	}
	return &repo, nil

}

func (r *ShopifyRepository) GetProducts() ([]goshopify.Product, error) {

	products, err := r.Client.Product.List(nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n retrieved %d products", len(products))
	return products, nil

}

func (r *ShopifyRepository) GetCustomers() ([]goshopify.Customer, error) {

	customers, err := r.Client.Customer.List(nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n retrieved %d customers", len(customers))
	return customers, nil

}

func (r *ShopifyRepository) GetCustomerAddresses(customers []goshopify.Customer) error {

	var err error
	wg := sync.WaitGroup{}
	wg.Add(len(customers))

	concurrentOperations := 2
	sem := make(chan bool, concurrentOperations)

	for i := range customers {
		sem <- true
		go func(c goshopify.Customer) {
			defer func() {
				wg.Done()
				<-sem
			}()
			ads, e := r.Client.CustomerAddress.List(c.ID, nil)
			if e != nil {
				err = e
			} else {
				var addresses []*goshopify.CustomerAddress
				for i := range ads {
					addresses = append(addresses, &ads[i])
				}
				customers[i].Addresses = addresses
			}
		}(customers[i])
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}

	wg.Wait()
	if err != nil {
		return err
	}
	return nil

}
