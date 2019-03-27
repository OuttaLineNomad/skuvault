package inventory

// PickItemBulk is a automatically generated struct from json provided by sku vault's api docs.
type PickItemBulk struct {
	Items []struct {
		Code          string `json:"Code"`
		IsExpressPick bool   `json:"IsExpressPick"`
		LocationCode  string `json:"LocationCode"`
		Note          string `json:"Note"`
		Quantity     int  `json:"Quantity"`
		SaleID        string `json:"SaleId"`
		ScannedCode   string `json:"ScannedCode"`
		Sku           string `json:"Sku"`
		WarehouseID  int  `json:"WarehouseId"`
	} `json:"Items"`
}

// PickItemBulkResponse is a automatically generated struct from json provided by sku vault's api docs.
type PickItemBulkResponse struct {
	Errors  []interface{} `json:"Errors"`
	Results []string      `json:"Results"`
	Status  string        `json:"Status"`
}
