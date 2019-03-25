package products

// CreateBrands is a automatically generated struct from json provided by sku vault's api docs.
type CreateBrands struct {
	Brands []struct {
		Name string `json:"Name"`
	} `json:"Brands"`
}

// CreateBrandsResponse is a automatically generated struct from json provided by sku vault's api docs.
type CreateBrandsResponse struct {
	Errors []interface{} `json:"Errors"`
	Status string        `json:"Status"`
}
