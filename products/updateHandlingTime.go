package products

// UpdateHandlingTime is a automatically generated struct from json provided by sku vault's api docs.
type UpdateHandlingTime struct {
	Items []struct {
		HandlingTime []struct {
			AccountID string `json:"AccountId"`
			Quantity  int64  `json:"Quantity"`
		} `json:"HandlingTime"`
		Sku string `json:"Sku"`
	} `json:"Items"`
}

// UpdateHandlingTimeResponse is a automatically generated struct from json provided by sku vault's api docs.
type UpdateHandlingTimeResponse struct {
	Errors []interface{} `json:"Errors"`
	Status string        `json:"Status"`
}
