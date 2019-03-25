package inventory

// GetWarehouseItemQuantity is a automatically generated struct from json provided by sku vault's api docs.
type GetWarehouseItemQuantity struct {
	Sku         string `json:"Sku"`
	WarehouseID int64  `json:"WarehouseId"`
}

// GetWarehouseItemQuantityResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetWarehouseItemQuantityResponse struct {
	Errors              []string `json:"Errors"`
	TotalQuantityOnHand int64    `json:"TotalQuantityOnHand"`
}
