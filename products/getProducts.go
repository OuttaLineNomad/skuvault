package products

// GetProducts is a automatically generated struct from json provided by sku vault's api docs.
type GetProducts struct {
	ModifiedAfterDateTimeUtc  string   `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc string   `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int64    `json:"PageNumber"`
	PageSize                  int64    `json:"PageSize"`
	ProductCodes              []string `json:"ProductCodes"`
	ProductSKUs               []string `json:"ProductSKUs"`
}

// GetProductsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetProductsResponse struct {
	Errors   []interface{} `json:"Errors"`
	Products []struct {
		Attributes []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Attributes"`
		Brand               string   `json:"Brand"`
		Classification      string   `json:"Classification"`
		Code                string   `json:"Code"`
		Cost                int64    `json:"Cost"`
		CreatedDateUtc      string   `json:"CreatedDateUtc"`
		Description         string   `json:"Description"`
		DisableQuantitySync bool     `json:"DisableQuantitySync"`
		ID                  string   `json:"Id"`
		IncrementalQuantity int64    `json:"IncrementalQuantity"`
		IsAlternateCode     bool     `json:"IsAlternateCode"`
		IsAlternateSKU      bool     `json:"IsAlternateSKU"`
		LongDescription     string   `json:"LongDescription"`
		Moq                 int64    `json:"MOQ"`
		MOQInfo             string   `json:"MOQInfo"`
		ModifiedDateUtc     string   `json:"ModifiedDateUtc"`
		PartNumber          string   `json:"PartNumber"`
		Pictures            []string `json:"Pictures"`
		QuantityAvailable   int64    `json:"QuantityAvailable"`
		QuantityInStock     int64    `json:"QuantityInStock"`
		QuantityInbound     int64    `json:"QuantityInbound"`
		QuantityIncoming    int64    `json:"QuantityIncoming"`
		QuantityOnHand      int64    `json:"QuantityOnHand"`
		QuantityOnHold      int64    `json:"QuantityOnHold"`
		QuantityPending     int64    `json:"QuantityPending"`
		QuantityPicked      int64    `json:"QuantityPicked"`
		QuantityTotalFBA    int64    `json:"QuantityTotalFBA"`
		QuantityTransfer    int64    `json:"QuantityTransfer"`
		ReorderPoint        int64    `json:"ReorderPoint"`
		RetailPrice         int64    `json:"RetailPrice"`
		SalePrice           int64    `json:"SalePrice"`
		ShortDescription    string   `json:"ShortDescription"`
		Sku                 string   `json:"Sku"`
		Statuses            []string `json:"Statuses"`
		Supplier            string   `json:"Supplier"`
		SupplierInfo        []struct {
			Cost               int64  `json:"Cost"`
			IsActive           bool   `json:"IsActive"`
			IsPrimary          bool   `json:"IsPrimary"`
			LeadTime           int64  `json:"LeadTime"`
			SupplierName       string `json:"SupplierName"`
			SupplierPartNumber string `json:"SupplierPartNumber"`
		} `json:"SupplierInfo"`
		VariationParentSku string `json:"VariationParentSku"`
		WeightUnit         string `json:"WeightUnit"`
		WeightValue        string `json:"WeightValue"`
	} `json:"Products"`
}
