package inventory

// AddItemBulk is a automatically generated struct from json provided by sku vault's api docs.
type AddItemBulk struct {
	Items []struct {
		Code         string `json:"Code"`
		LocationCode string `json:"LocationCode"`
		Note         string `json:"Note"`
		Quantity     int  `json:"Quantity"`
		Reason       string `json:"Reason"`
		Sku          string `json:"Sku"`
		WarehouseID  int  `json:"WarehouseId"`
	} `json:"Items"`
}

// AddItemBulkResponse is a automatically generated struct from json provided by sku vault's api docs.
type AddItemBulkResponse struct {
	Errors  []interface{} `json:"Errors"`
	Results []string      `json:"Results"`
	Status  string        `json:"Status"`
}
