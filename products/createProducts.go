package products

// CreateProducts is a automatically generated struct from json provided by sku vault's api docs.
type CreateProducts struct {
	Items []struct {
		AllowCreateAp bool `json:"AllowCreateAp"`
		Attributes    struct {
			String string `json:"String"`
		} `json:"Attributes"`
		Brand                    string   `json:"Brand"`
		Classification           string   `json:"Classification"`
		Code                     string   `json:"Code"`
		Cost                     int64    `json:"Cost"`
		Description              string   `json:"Description"`
		LongDescription          string   `json:"LongDescription"`
		MinimumOrderQuantity     int64    `json:"MinimumOrderQuantity"`
		MinimumOrderQuantityInfo string   `json:"MinimumOrderQuantityInfo"`
		Note                     string   `json:"Note"`
		PartNumber               string   `json:"PartNumber"`
		Pictures                 []string `json:"Pictures"`
		ReorderPoint             int64    `json:"ReorderPoint"`
		RetailPrice              int64    `json:"RetailPrice"`
		SalePrice                int64    `json:"SalePrice"`
		ShortDescription         string   `json:"ShortDescription"`
		Sku                      string   `json:"Sku"`
		Statuses                 []string `json:"Statuses"`
		Supplier                 string   `json:"Supplier"`
		SupplierInfo             []struct {
			Cost               string `json:"Cost"`
			IsActive           string `json:"IsActive"`
			IsPrimary          string `json:"IsPrimary"`
			LeadTime           string `json:"LeadTime"`
			SupplierName       string `json:"SupplierName"`
			SupplierPartNumber string `json:"SupplierPartNumber"`
		} `json:"SupplierInfo"`
		VariationParentSku string `json:"VariationParentSku"`
		Weight             int64  `json:"Weight"`
		WeightUnit         string `json:"WeightUnit"`
	} `json:"Items"`
}

// CreateProductsResponse is a automatically generated struct from json provided by sku vault's api docs.
type CreateProductsResponse struct {
	Errors []struct {
		ErrorMessages []string `json:"ErrorMessages"`
		Sku           string   `json:"Sku"`
	} `json:"Errors"`
	Status string `json:"Status"`
}
