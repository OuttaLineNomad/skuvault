package inventory

// GetAvailableQuantities is a automatically generated struct from json provided by sku vault's api docs.
type GetAvailableQuantities struct {
	ExpandAlternateSkus       bool        `json:"ExpandAlternateSkus"`
	ModifiedAfterDateTimeUtc  string      `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc string      `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int64       `json:"PageNumber"`
	PageSize                  interface{} `json:"PageSize"`
}

// GetAvailableQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetAvailableQuantitiesResponse struct {
	Items []struct {
		AvailableQuantity       int64  `json:"AvailableQuantity"`
		IsAlternateSku          bool   `json:"IsAlternateSku"`
		LastModifiedDateTimeUtc string `json:"LastModifiedDateTimeUtc"`
		Sku                     string `json:"Sku"`
	} `json:"Items"`
}
