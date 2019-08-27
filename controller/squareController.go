package controller

import (
	"github.com/andrewcretin/shopify2square/models"
	"github.com/andrewcretin/shopify2square/models/square"
)

func (c *Controller) SyncSquareCustomers(new []square.SquareCustomer, updated []square.SquareCustomerUpdate, resp *models.SyncResponse) error {

	// update response object
	resp.NewCustomers = int64(len(new))
	resp.UpdatedCustomers = int64(len(updated))

	// write new customers
	err := c.SquareRepo.WriteCustomers(new)
	if err != nil {
		return err
	}

	// update existing customers

	return nil

}

func (c *Controller) SyncSquareItems(new []square.SquareItem, updated []square.SquareItemUpdate, resp *models.SyncResponse) error {

	// update response object
	resp.NewItems = int64(len(new))
	resp.UpdatedItems = int64(len(updated))

	return nil

}

func (c *Controller) SyncSquareCategories(new []square.SquareCategory, updated []square.SquareCategoryUpdate, resp *models.SyncResponse) error {

	// update response object
	resp.NewCategories = int64(len(new))
	resp.UpdatedCategories = int64(len(updated))

	return nil

}

func (c *Controller) SyncSquareOrders(new []square.SquareOrder, updated []square.SquareOrderUpdate, resp *models.SyncResponse) error {

	// update response object
	resp.NewOrders = int64(len(new))
	resp.UpdatedOrders = int64(len(updated))

	return nil

}
