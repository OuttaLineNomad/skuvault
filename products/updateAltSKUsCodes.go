package products

// UpdateAltSKUsCodes is a automatically generated struct from json provided by sku vault's api docs.
type UpdateAltSKUsCodes struct {
	Action string `json:"Action"`
	Items  []struct {
		AltCodes []string `json:"AltCodes"`
		AltSKUs  []string `json:"AltSKUs"`
		Sku      string   `json:"Sku"`
	} `json:"Items"`
}

// UpdateAltSKUsCodesResponse is a automatically generated struct from json provided by sku vault's api docs.
type UpdateAltSKUsCodesResponse struct {
	Errors []interface{} `json:"Errors"`
	Status string        `json:"Status"`
}
