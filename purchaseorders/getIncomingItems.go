package purchaseorders

import "time"

// GetIncomingItems is a automatically generated struct from json provided by sku vault's api docs.
type GetIncomingItems struct {
	PageNumber int `json:"PageNumber"`
}

// GetIncomingItemsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetIncomingItemsResponse struct {
	Errors    []interface{} `json:"Errors"`
	SoldItems []struct {
		Date       time.Time `json:"Date"`
		Quantity   int       `json:"Quantity"`
		Sku        string    `json:"SKU"`
		TotalPrice float32   `json:"TotalPrice"`
	} `json:"SoldItems"`
	Status string `json:"Status"`
}
