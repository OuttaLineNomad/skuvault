package sales

// GetSoldItems is a automatically generated struct from json provided by sku vault's api docs.
type GetSoldItems struct {
	BreakDownKits bool   `json:"BreakDownKits"`
	EndDateUtc    string `json:"EndDateUtc"`
	StartDateUtc  string `json:"StartDateUtc"`
}

// GetSoldItemsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetSoldItemsResponse struct {
	SoldItems []struct {
		Date       string `json:"Date"`
		Quantity   int64  `json:"Quantity"`
		Sku        string `json:"SKU"`
		TotalPrice int64  `json:"TotalPrice"`
	} `json:"SoldItems"`
}
