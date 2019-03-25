package purchaseorders

// GetIncomingItems is a automatically generated struct from json provided by sku vault's api docs.
type GetIncomingItems struct {
	PageNumber int64 `json:"PageNumber"`
}

// GetIncomingItemsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetIncomingItemsResponse struct {
	Errors    []interface{} `json:"Errors"`
	SoldItems []struct {
		Date       string `json:"Date"`
		Quantity   int64  `json:"Quantity"`
		Sku        string `json:"SKU"`
		TotalPrice int64  `json:"TotalPrice"`
	} `json:"SoldItems"`
	Status string `json:"Status"`
}
