package inventory

// GetItemQuantities is a automatically generated struct from json provided by sku vault's api docs.
type GetItemQuantities struct {
	ModifiedAfterDateTimeUtc  string      `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc string      `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int64       `json:"PageNumber"`
	PageSize                  interface{} `json:"PageSize"`
	ProductCodes              []string    `json:"ProductCodes"`
}

// GetItemQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetItemQuantitiesResponse struct {
	Items []struct {
		AvailableQuantity       int64  `json:"AvailableQuantity"`
		Code                    string `json:"Code"`
		HeldQuantity            int64  `json:"HeldQuantity"`
		LastModifiedDateTimeUtc string `json:"LastModifiedDateTimeUtc"`
		PendingQuantity         int64  `json:"PendingQuantity"`
		PickedQuantity          int64  `json:"PickedQuantity"`
		Sku                     string `json:"Sku"`
		TotalOnHand             int64  `json:"TotalOnHand"`
	} `json:"Items"`
}
