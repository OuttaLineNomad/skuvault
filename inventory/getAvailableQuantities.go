package inventory

import "time"

// GetAvailableQuantities is a automatically generated struct from json provided by sku vault's api docs.
type GetAvailableQuantities struct {
	ExpandAlternateSkus       bool      `json:"ExpandAlternateSkus"`
	ModifiedAfterDateTimeUtc  time.Time `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc time.Time `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int       `json:"PageNumber"`
	PageSize                  int       `json:"PageSize"`
}

// GetAvailableQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetAvailableQuantitiesResponse struct {
	Items []struct {
		AvailableQuantity       int       `json:"AvailableQuantity"`
		IsAlternateSku          bool      `json:"IsAlternateSku"`
		LastModifiedDateTimeUtc time.Time `json:"LastModifiedDateTimeUtc"`
		Sku                     string    `json:"Sku"`
	} `json:"Items"`
}
