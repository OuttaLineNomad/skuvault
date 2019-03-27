package products

import "time"

// GetProducts is a automatically generated struct from json provided by sku vault's api docs.
type GetProducts struct {
	ModifiedAfterDateTimeUtc  time.Time `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc time.Time `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int       `json:"PageNumber"`
	PageSize                  int       `json:"PageSize"`
	ProductCodes              []string  `json:"ProductCodes"`
	ProductSKUs               []string  `json:"ProductSKUs"`
}

// GetProductsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetProductsResponse struct {
	Errors   []interface{} `json:"Errors"`
	Products []struct {
		Attributes []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Attributes"`
		Brand               string    `json:"Brand"`
		Classification      string    `json:"Classification"`
		Code                string    `json:"Code"`
		Cost                float64   `json:"Cost"`
		CreatedDateUtc      time.Time `json:"CreatedDateUtc"`
		Description         string    `json:"Description"`
		DisableQuantitySync bool      `json:"DisableQuantitySync"`
		ID                  string    `json:"Id"`
		IncrementalQuantity int       `json:"IncrementalQuantity"`
		IsAlternateCode     bool      `json:"IsAlternateCode"`
		IsAlternateSKU      bool      `json:"IsAlternateSKU"`
		LongDescription     string    `json:"LongDescription"`
		MOQ                 int       `json:"MOQ"`
		MOQInfo             string    `json:"MOQInfo"`
		ModifiedDateUtc     time.Time `json:"ModifiedDateUtc"`
		PartNumber          string    `json:"PartNumber"`
		Pictures            []string  `json:"Pictures"`
		QuantityAvailable   int       `json:"QuantityAvailable"`
		QuantityInStock     int       `json:"QuantityInStock"`
		QuantityInbound     int       `json:"QuantityInbound"`
		QuantityIncoming    int       `json:"QuantityIncoming"`
		QuantityOnHand      int       `json:"QuantityOnHand"`
		QuantityOnHold      int       `json:"QuantityOnHold"`
		QuantityPending     int       `json:"QuantityPending"`
		QuantityPicked      int       `json:"QuantityPicked"`
		QuantityTotalFBA    int       `json:"QuantityTotalFBA"`
		QuantityTransfer    int       `json:"QuantityTransfer"`
		ReorderPoint        int       `json:"ReorderPoint"`
		RetailPrice         float64   `json:"RetailPrice"`
		SalePrice           float64   `json:"SalePrice"`
		ShortDescription    string    `json:"ShortDescription"`
		Sku                 string    `json:"Sku"`
		Statuses            []string  `json:"Statuses"`
		Supplier            string    `json:"Supplier"`
		SupplierInfo        []struct {
			Cost               float64 `json:"Cost"`
			IsActive           bool    `json:"IsActive"`
			IsPrimary          bool    `json:"IsPrimary"`
			LeadTime           float64 `json:"LeadTime"`
			SupplierName       string  `json:"SupplierName"`
			SupplierPartNumber string  `json:"SupplierPartNumber"`
		} `json:"SupplierInfo"`
		VariationParentSku string `json:"VariationParentSku"`
		WeightUnit         string `json:"WeightUnit"`
		WeightValue        string `json:"WeightValue"`
	} `json:"Products"`
}
