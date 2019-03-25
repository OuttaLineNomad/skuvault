package inventory

// SetItemQuantities is a automatically generated struct from json provided by sku vault's api docs.
type SetItemQuantities struct {
	Items []struct {
		LocationCode string `json:"LocationCode"`
		Quantity     int64  `json:"Quantity"`
		Sku          string `json:"Sku"`
		WarehouseID  int64  `json:"WarehouseId"`
	} `json:"Items"`
}

// SetItemQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type SetItemQuantitiesResponse struct {
	Errors []struct {
		Code  string `json:"Code"`
		Error string `json:"Error"`
		Sku   string `json:"SKU"`
	} `json:"Errors"`
	Status string `json:"Status"`
}
