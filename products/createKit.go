package products

// CreateKit is a automatically generated struct from json provided by sku vault's api docs.
type CreateKit struct {
	AllowCreateAp bool   `json:"AllowCreateAp"`
	Code          string `json:"Code"`
	KitLines      []struct {
		Combine  int64    `json:"Combine"`
		Items    []string `json:"Items"`
		LineName string   `json:"LineName"`
		Quantity int64    `json:"Quantity"`
	} `json:"KitLines"`
	Sku   string `json:"Sku"`
	Title string `json:"Title"`
}

// CreateKitResponse is a automatically generated struct from json provided by sku vault's api docs.
type CreateKitResponse struct {
	Errors  []string `json:"Errors"`
	Success string   `json:"Success"`
}
