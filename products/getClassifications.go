package products

// GetClassifications is a automatically generated struct from json provided by sku vault's api docs.
type GetClassifications struct {
	PageNumber int64 `json:"PageNumber"`
}

// GetClassificationsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetClassificationsResponse struct {
	Classifications []struct {
		Attributes []struct {
			IsEnabled  bool     `json:"IsEnabled"`
			IsRequired bool     `json:"IsRequired"`
			Name       string   `json:"Name"`
			Values     []string `json:"Values"`
		} `json:"Attributes"`
		IsEnabled bool   `json:"IsEnabled"`
		Name      string `json:"Name"`
	} `json:"Classifications"`
}
