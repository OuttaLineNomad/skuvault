package inventory

// GetKitQuantities is a automatically generated struct from json provided by sku vault's api docs.
type GetKitQuantities struct {
	ModifiedAfterDateTimeUtc  string `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc string `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int64  `json:"PageNumber"`
}

// GetKitQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetKitQuantitiesResponse struct {
	Kits []struct {
		AvailableQuantity       int64  `json:"AvailableQuantity"`
		LastModifiedDateTimeUtc string `json:"LastModifiedDateTimeUtc"`
		Sku                     string `json:"Sku"`
	} `json:"Kits"`
}
