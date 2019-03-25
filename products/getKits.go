package products

// GetKits is a automatically generated struct from json provided by sku vault's api docs.
type GetKits struct {
	AvailableQuantityModifiedAfterDateTimeUtc  string   `json:"AvailableQuantityModifiedAfterDateTimeUtc"`
	AvailableQuantityModifiedBeforeDateTimeUtc string   `json:"AvailableQuantityModifiedBeforeDateTimeUtc"`
	GetAvailableQuantity                       bool     `json:"GetAvailableQuantity"`
	IncludeKitCost                             bool     `json:"IncludeKitCost"`
	KitSKUs                                    []string `json:"KitSKUs"`
	ModifiedAfterDateTimeUtc                   string   `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc                  string   `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                                 int64    `json:"PageNumber"`
}

// GetKitsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetKitsResponse struct {
	Errors []string `json:"Errors"`
	Kits   []struct {
		AvailableQuantity                        int64  `json:"AvailableQuantity"`
		AvailableQuantityLastModifiedDateTimeUtc string `json:"AvailableQuantityLastModifiedDateTimeUtc"`
		Code                                     string `json:"Code"`
		Cost                                     int64  `json:"Cost"`
		Description                              string `json:"Description"`
		KitLines                                 []struct {
			Combine int64 `json:"Combine"`
			Items   []struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
				Sku         string `json:"SKU"`
			} `json:"Items"`
			LineName string `json:"LineName"`
			Quantity int64  `json:"Quantity"`
		} `json:"KitLines"`
		LastModifiedDateTimeUtc string   `json:"LastModifiedDateTimeUtc"`
		Sku                     string   `json:"SKU"`
		Statuses                []string `json:"Statuses"`
	} `json:"Kits"`
}
