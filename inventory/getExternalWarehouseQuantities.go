package inventory

// GetExternalWarehouseQuantities is a automatically generated struct from json provided by sku vault's api docs.
type GetExternalWarehouseQuantities struct {
	GetTotalFbaQuantities bool        `json:"GetTotalFbaQuantities"`
	PageNumber            int64       `json:"PageNumber"`
	PageSize              interface{} `json:"PageSize"`
	Usertoken             string      `json:"Usertoken"`
	WarehouseID           string      `json:"WarehouseId"`
}

// GetExternalWarehouseQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetExternalWarehouseQuantitiesResponse struct {
	Errors     []string `json:"Errors"`
	Quantities []struct {
		InStockQuantity  int64  `json:"InStockQuantity"`
		InboundQuantity  int64  `json:"InboundQuantity"`
		ReserveQuantity  int64  `json:"ReserveQuantity"`
		Sku              string `json:"Sku"`
		TotalQuantity    int64  `json:"TotalQuantity"`
		TransferQuantity int64  `json:"TransferQuantity"`
	} `json:"Quantities"`
}
