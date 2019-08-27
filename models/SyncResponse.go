package models

type SyncResponse struct {
	NewCategories     int64 `json:"newCustomers"`
	UpdatedCategories int64 `json:"updatedCustomers"`
	NewCustomers      int64 `json:"newCustomers"`
	UpdatedCustomers  int64 `json:"updatedCustomers"`
	NewItems          int64 `json:"newItems"`
	UpdatedItems      int64 `json:"updatedItems"`
	NewOrders         int64 `json:"newOrders"`
	UpdatedOrders     int64 `json:"updatedOrders"`
}
