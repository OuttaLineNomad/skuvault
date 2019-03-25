package products

// GetHandlingTimeResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetHandlingTimeResponse struct {
	Errors []interface{} `json:"Errors"`
	Items  []struct {
		AccountID string `json:"AccountId"`
		Kits      []struct {
			Quantity int64  `json:"Quantity"`
			Sku      string `json:"Sku"`
		} `json:"Kits"`
		Products []struct {
			Quantity int64  `json:"Quantity"`
			Sku      string `json:"Sku"`
		} `json:"Products"`
	} `json:"Items"`
	Status string `json:"Status"`
}
