package sales

// GetSaleItemCost is a automatically generated struct from json provided by sku vault's api docs.
type GetSaleItemCost struct {
	SaleIds []string `json:"SaleIds"`
}

// GetSaleItemCostResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetSaleItemCostResponse struct {
	Errors []interface{} `json:"Errors"`
	Sales  []struct {
		Items []struct {
			Cost struct {
				A int64  `json:"a"`
				S string `json:"s"`
			} `json:"Cost"`
			Sku string `json:"Sku"`
		} `json:"Items"`
		SaleID string `json:"SaleId"`
	} `json:"Sales"`
	Status string `json:"Status"`
}
