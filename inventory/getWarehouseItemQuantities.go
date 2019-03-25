package inventory

// GetWarehouseItemQuantities is a automatically generated struct from json provided by sku vault's api docs.
type GetWarehouseItemQuantities struct {
	ModifiedAfterDateTimeUtc  string `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc string `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int64  `json:"PageNumber"`
	PageSize                  int64  `json:"PageSize"`
	WarehouseID               int64  `json:"WarehouseId"`
}

// GetWarehouseItemQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetWarehouseItemQuantitiesResponse struct {
	ItemQuantities []struct {
		Quantity int64  `json:"Quantity"`
		Sku      string `json:"Sku"`
	} `json:"ItemQuantities"`
}
