package repository

import (
	"fmt"
	"github.com/andrewcretin/shopify2square/envConfig"
	goshopify "github.com/andrewcretin/shopify2square/src/github.com/bold-commerce/go-shopify"
	"github.com/joncalhoun/drip"
	"sort"
	"time"
)

//noinspection ALL
const (
	DEFAULT_ORDER_LIMIT = 250
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
	b := drip.Bucket{
		Capacity:     len(customers),
		DripInterval: 1 * time.Second,
		PerDrip:      2,
	}

	for i := range customers {
		e := b.Consume(1)
		if e != nil {
			fmt.Println("Sleep 1s.")
			time.Sleep(1 * time.Second)
		}
		ads, e := r.Client.CustomerAddress.List(customers[i].ID, nil)
		if e != nil {
			err = e
		} else {
			var addresses []*goshopify.CustomerAddress
			for i := range ads {
				addresses = append(addresses, &ads[i])
			}
			customers[i].Addresses = addresses
		}
	}
	defer func() {
		_ = b.Stop()
	}()
	_ = b.Start()
	if err != nil {
		return err
	}
	return nil

}

func (r *ShopifyRepository) GetAllOrders(orders []goshopify.Order, opts goshopify.OrderListOptions) ([]goshopify.Order, error) {

	opts.Limit = DEFAULT_ORDER_LIMIT
	tempOrders, err := r.Client.Order.List(opts)
	if err != nil {
		return nil, err
	} else {
		orders = append(orders, tempOrders...)
	}
	fmt.Printf("\n retrieved %d orders", len(tempOrders))
	if len(tempOrders) > 0 {
		sort.Slice(tempOrders, func(i, j int) bool {
			return tempOrders[i].CreatedAt.After(*tempOrders[j].CreatedAt)
		})
		latestCreatedAt := *tempOrders[len(tempOrders)-1].CreatedAt
		opts.CreatedAtMax = latestCreatedAt.Add(-1 * time.Second)
		return r.GetAllOrders(orders, opts)
	} else {
		return orders, nil
	}

}
