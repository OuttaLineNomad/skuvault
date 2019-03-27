package sales

import "time"

// GetSoldItems is a automatically generated struct from json provided by sku vault's api docs.
type GetSoldItems struct {
	BreakDownKits bool      `json:"BreakDownKits"`
	EndDateUtc    time.Time `json:"EndDateUtc"`
	StartDateUtc  time.Time `json:"StartDateUtc"`
}

// GetSoldItemsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetSoldItemsResponse struct {
	SoldItems []struct {
		Date       time.Time `json:"Date"`
		Quantity   int       `json:"Quantity"`
		Sku        string    `json:"SKU"`
		TotalPrice float64   `json:"TotalPrice"`
	} `json:"SoldItems"`
}
