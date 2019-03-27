package inventory

// UpdateExternalWarehouseQuantities is a automatically generated struct from json provided by sku vault's api docs.
type UpdateExternalWarehouseQuantities struct {
	Quantities []struct {
		InStockQuantity     int  `json:"InStockQuantity"`
		InboundQuantity     int  `json:"InboundQuantity"`
		ReserveQuantity     int  `json:"ReserveQuantity"`
		Sku              string `json:"Sku"`
		TotalQuantity     int  `json:"TotalQuantity"`
		TransferQuantity     int  `json:"TransferQuantity"`
	} `json:"Quantities"`
	WarehouseID string `json:"WarehouseId"`
}

// UpdateExternalWarehouseQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type UpdateExternalWarehouseQuantitiesResponse struct {
	Errors []string `json:"Errors"`
	Status string   `json:"Status"`
}
