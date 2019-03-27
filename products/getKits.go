package products

import "time"

// GetKits is a automatically generated struct from json provided by sku vault's api docs.
type GetKits struct {
	AvailableQuantityModifiedAfterDateTimeUtc  time.Time `json:"AvailableQuantityModifiedAfterDateTimeUtc"`
	AvailableQuantityModifiedBeforeDateTimeUtc time.Time `json:"AvailableQuantityModifiedBeforeDateTimeUtc"`
	GetAvailableQuantity                       bool      `json:"GetAvailableQuantity"`
	IncludeKitCost                             bool      `json:"IncludeKitCost"`
	KitSKUs                                    []string  `json:"KitSKUs"`
	ModifiedAfterDateTimeUtc                   time.Time `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc                  time.Time `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                                 int       `json:"PageNumber"`
}

// GetKitsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetKitsResponse struct {
	Errors []string `json:"Errors"`
	Kits   []struct {
		AvailableQuantity                        int       `json:"AvailableQuantity"`
		AvailableQuantityLastModifiedDateTimeUtc time.Time `json:"AvailableQuantityLastModifiedDateTimeUtc"`
		Code                                     string    `json:"Code"`
		Cost                                     float64   `json:"Cost"`
		Description                              string    `json:"Description"`
		KitLines                                 []struct {
			Combine int `json:"Combine"`
			Items   []struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
				Sku         string `json:"SKU"`
			} `json:"Items"`
			LineName string `json:"LineName"`
			Quantity int    `json:"Quantity"`
		} `json:"KitLines"`
		LastModifiedDateTimeUtc time.Time `json:"LastModifiedDateTimeUtc"`
		Sku                     string    `json:"SKU"`
		Statuses                []string  `json:"Statuses"`
	} `json:"Kits"`
}
