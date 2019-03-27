package products

// UpdateProduct is a automatically generated struct from json provided by sku vault's api docs.
type UpdateProduct struct {
	Attributes struct {
		String string `json:"String"`
	} `json:"Attributes"`
	Brand                    string   `json:"Brand"`
	Classification           string   `json:"Classification"`
	Code                     string   `json:"Code"`
	Cost                     float64  `json:"Cost"`
	Description              string   `json:"Description"`
	LongDescription          string   `json:"LongDescription"`
	MinimumOrderQuantity     int      `json:"MinimumOrderQuantity"`
	MinimumOrderQuantityInfo string   `json:"MinimumOrderQuantityInfo"`
	Note                     string   `json:"Note"`
	PartNumber               string   `json:"PartNumber"`
	Pictures                 []string `json:"Pictures"`
	ReorderPoint             float64  `json:"ReorderPoint"`
	RetailPrice              float64  `json:"RetailPrice"`
	SalePrice                float64  `json:"SalePrice"`
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
	VariationParentSku string  `json:"VariationParentSku"`
	Weight             float64 `json:"Weight"`
	WeightUnit         string  `json:"WeightUnit"`
}

// UpdateProductResponse is a automatically generated struct from json provided by sku vault's api docs.
type UpdateProductResponse struct {
	Errors []interface{} `json:"Errors"`
	Status string        `json:"Status"`
}
