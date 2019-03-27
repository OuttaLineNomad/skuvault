package inventory

import "time"

// GetWarehouseItemQuantities is a automatically generated struct from json provided by sku vault's api docs.
type GetWarehouseItemQuantities struct {
	ModifiedAfterDateTimeUtc  time.Time `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc time.Time `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int       `json:"PageNumber"`
	PageSize                  int       `json:"PageSize"`
	WarehouseID               int       `json:"WarehouseId"`
}

// GetWarehouseItemQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetWarehouseItemQuantitiesResponse struct {
	ItemQuantities []struct {
		Quantity int    `json:"Quantity"`
		Sku      string `json:"Sku"`
	} `json:"ItemQuantities"`
}
