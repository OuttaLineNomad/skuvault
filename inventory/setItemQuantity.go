package inventory

// SetItemQuantity is a automatically generated struct from json provided by sku vault's api docs.
type SetItemQuantity struct {
	Code         string `json:"Code"`
	LocationCode string `json:"LocationCode"`
	Quantity     int  `json:"Quantity"`
	Sku          string `json:"Sku"`
	WarehouseID  int  `json:"WarehouseId"`
}

// SetItemQuantityResponse is a automatically generated struct from json provided by sku vault's api docs.
type SetItemQuantityResponse struct {
	SetItemQuantityStatusEnum string `json:"SetItemQuantityStatusEnum"`
}
