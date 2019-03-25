package inventory

// UpdateExternalWarehouseQuantities is a automatically generated struct from json provided by sku vault's api docs.
type UpdateExternalWarehouseQuantities struct {
	Quantities []struct {
		InStockQuantity  int64  `json:"InStockQuantity"`
		InboundQuantity  int64  `json:"InboundQuantity"`
		ReserveQuantity  int64  `json:"ReserveQuantity"`
		Sku              string `json:"Sku"`
		TotalQuantity    int64  `json:"TotalQuantity"`
		TransferQuantity int64  `json:"TransferQuantity"`
	} `json:"Quantities"`
	WarehouseID string `json:"WarehouseId"`
}

// UpdateExternalWarehouseQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type UpdateExternalWarehouseQuantitiesResponse struct {
	Errors []string `json:"Errors"`
	Status string   `json:"Status"`
}
