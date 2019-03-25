package inventory

// RemoveItemBulk is a automatically generated struct from json provided by sku vault's api docs.
type RemoveItemBulk struct {
	Items []struct {
		Code         string `json:"Code"`
		LocationCode string `json:"LocationCode"`
		Quantity     int64  `json:"Quantity"`
		Reason       string `json:"Reason"`
		Sku          string `json:"Sku"`
		WarehouseID  int64  `json:"WarehouseId"`
	} `json:"Items"`
}

// RemoveItemBulkResponse is a automatically generated struct from json provided by sku vault's api docs.
type RemoveItemBulkResponse struct {
	Errors  []string `json:"Errors"`
	Results []string `json:"Results"`
	Status  string   `json:"Status"`
}
