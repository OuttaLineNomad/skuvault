package inventory

// GetWarehouses is a automatically generated struct from json provided by sku vault's api docs.
type GetWarehouses struct {
	PageNumber int64 `json:"PageNumber"`
}

// GetWarehousesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetWarehousesResponse struct {
	Warehouses []struct {
		Code string `json:"Code"`
		ID   string `json:"Id"`
	} `json:"Warehouses"`
}
