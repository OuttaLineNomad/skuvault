package inventory

// GetKitQuantities is a automatically generated struct from json provided by sku vault's api docs.
type GetKitQuantities struct {
	ModifiedAfterDateTimeUtc  time.Time `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc time.Time `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int  `json:"PageNumber"`
}

// GetKitQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetKitQuantitiesResponse struct {
	Kits []struct {
		AvailableQuantity     int  `json:"AvailableQuantity"`
		LastModifiedDateTimeUtc time.Time `json:"LastModifiedDateTimeUtc"`
		Sku                     string `json:"Sku"`
	} `json:"Kits"`
}
