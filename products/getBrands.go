package products

// GetBrands is a automatically generated struct from json provided by sku vault's api docs.
type GetBrands struct {
	PageNumber                int `json:"PageNumber"`
}

// GetBrandsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetBrandsResponse struct {
	Brands []struct {
		IsEnabled bool   `json:"IsEnabled"`
		Name      string `json:"Name"`
	} `json:"Brands"`
}
