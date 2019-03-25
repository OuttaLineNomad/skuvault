package inventory

// PickItem is a automatically generated struct from json provided by sku vault's api docs.
type PickItem struct {
	Code          string `json:"Code"`
	IsExpressPick bool   `json:"IsExpressPick"`
	LocationCode  string `json:"LocationCode"`
	Note          string `json:"Note"`
	Quantity      int64  `json:"Quantity"`
	Sku           string `json:"Sku"`
	WarehouseID   int64  `json:"WarehouseId"`
}

// PickItemResponse is a automatically generated struct from json provided by sku vault's api docs.
type PickItemResponse struct {
	PickItemStatus string `json:"PickItemStatus"`
}
