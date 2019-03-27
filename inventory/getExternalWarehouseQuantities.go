package inventory

// GetExternalWarehouseQuantities is a automatically generated struct from json provided by sku vault's api docs.
type GetExternalWarehouseQuantities struct {
	GetTotalFbaQuantities bool   `json:"GetTotalFbaQuantities"`
	PageNumber            int    `json:"PageNumber"`
	PageSize              int    `json:"PageSize"`
	Usertoken             string `json:"Usertoken"`
	WarehouseID           string `json:"WarehouseId"`
}

// GetExternalWarehouseQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetExternalWarehouseQuantitiesResponse struct {
	Errors     []string `json:"Errors"`
	Quantities []struct {
		InStockQuantity  int    `json:"InStockQuantity"`
		InboundQuantity  int    `json:"InboundQuantity"`
		ReserveQuantity  int    `json:"ReserveQuantity"`
		Sku              string `json:"Sku"`
		TotalQuantity    int    `json:"TotalQuantity"`
		TransferQuantity int    `json:"TransferQuantity"`
	} `json:"Quantities"`
}
