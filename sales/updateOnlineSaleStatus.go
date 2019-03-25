package sales

// UpdateOnlineSaleStatus is a automatically generated struct from json provided by sku vault's api docs.
type UpdateOnlineSaleStatus struct {
	SaleID string `json:"SaleId"`
	Status string `json:"Status"`
}

// UpdateOnlineSaleStatusResponse is a automatically generated struct from json provided by sku vault's api docs.
type UpdateOnlineSaleStatusResponse struct {
	Errors []interface{} `json:"Errors"`
	Status string        `json:"Status"`
}
