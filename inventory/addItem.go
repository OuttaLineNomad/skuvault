package inventory

// AddItem is a automatically generated struct from json provided by sku vault's api docs.
type AddItem struct {
	Code         string `json:"Code"`
	LocationCode string `json:"LocationCode"`
	Note         string `json:"Note"`
	Quantity     int    `json:"Quantity"`
	Reason       string `json:"Reason"`
	Sku          string `json:"Sku"`
	WarehouseID  int    `json:"WarehouseId"`
}

// AddItemResponse is a automatically generated struct from json provided by sku vault's api docs.
type AddItemResponse struct {
	AddItemStatus string `json:"AddItemStatus"`
}
