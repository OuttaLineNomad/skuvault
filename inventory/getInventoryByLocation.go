package inventory

// GetInventoryByLocation is a automatically generated struct from json provided by sku vault's api docs.
type GetInventoryByLocation struct {
	IsReturnByCodes bool          `json:"IsReturnByCodes"`
	PageNumber      int64         `json:"PageNumber"`
	ProductCodes    []interface{} `json:"ProductCodes"`
	ProductSKUs     []interface{} `json:"ProductSKUs"`
	Usertoken       string        `json:"Usertoken"`
}

// GetInventoryByLocationResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetInventoryByLocationResponse struct {
	Errors []string `json:"Errors"`
	Items  struct {
		SKU_as_Key []struct {
			LocationCode  string `json:"LocationCode"`
			Quantity      int64  `json:"Quantity"`
			Reserve       bool   `json:"Reserve"`
			WarehouseCode string `json:"WarehouseCode"`
		} `json:"SKU-as-Key"`
	} `json:"Items"`
}
