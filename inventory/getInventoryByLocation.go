package inventory

// GetInventoryByLocation is a automatically generated struct from json provided by sku vault's api docs.
type GetInventoryByLocation struct {
	IsReturnByCodes bool     `json:"IsReturnByCodes"`
	PageNumber      int      `json:"PageNumber"`
	ProductCodes    []string `json:"ProductCodes"`
	ProductSKUs     []string `json:"ProductSKUs"`
	Usertoken       string   `json:"Usertoken"`
}

// GetInventoryByLocationResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetInventoryByLocationResponse struct {
	Errors []string `json:"Errors"`
	Items  map[string]struct {
		LocationCode  string `json:"LocationCode"`
		Quantity      int    `json:"Quantity"`
		Reserve       bool   `json:"Reserve"`
		WarehouseCode string `json:"WarehouseCode"`
	} `json:"Items"`
}
