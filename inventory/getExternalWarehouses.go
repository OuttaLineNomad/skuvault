package inventory

// GetExternalWarehousesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetExternalWarehousesResponse struct {
	Warehouses []struct {
		Code string `json:"Code"`
		ID   string `json:"Id"`
	} `json:"Warehouses"`
}
