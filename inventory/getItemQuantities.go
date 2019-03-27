package inventory

// GetItemQuantities is a automatically generated struct from json provided by sku vault's api docs.
type GetItemQuantities struct {
	ModifiedAfterDateTimeUtc  time.Time   `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc time.Time   `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int      `json:"PageNumber"`
	PageSize                  int      `json:"PageSize"`
	ProductCodes              []string `json:"ProductCodes"`
}

// GetItemQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetItemQuantitiesResponse struct {
	Items []struct {
		AvailableQuantity       int    `json:"AvailableQuantity"`
		Code                    string `json:"Code"`
		HeldQuantity            int    `json:"HeldQuantity"`
		LastModifiedDateTimeUtc time.Time `json:"LastModifiedDateTimeUtc"`
		PendingQuantity         int    `json:"PendingQuantity"`
		PickedQuantity          int    `json:"PickedQuantity"`
		Sku                     string `json:"Sku"`
		TotalOnHand             int    `json:"TotalOnHand"`
	} `json:"Items"`
}
