package integration

// GetIntegrationsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetIntegrationsResponse struct {
	Accounts []struct {
		ID     string `json:"Id"`
		LongID string `json:"LongId"`
		Name   string `json:"Name"`
		Type   string `json:"Type"`
	} `json:"Accounts"`
}
