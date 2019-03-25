package inventory

// GetLocationsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetLocationsResponse struct {
	Errors []interface{} `json:"Errors"`
	Items  []struct {
		ContainerCode  string `json:"ContainerCode"`
		LocationCode   string `json:"LocationCode"`
		ParentLocation string `json:"ParentLocation"`
		TotalQuantity  int64  `json:"TotalQuantity"`
		WarehouseCode  string `json:"WarehouseCode"`
		WarehouseName  string `json:"WarehouseName"`
	} `json:"Items"`
	Status string `json:"Status"`
}
