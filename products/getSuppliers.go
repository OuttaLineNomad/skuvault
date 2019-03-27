package products

// GetSuppliers is a automatically generated struct from json provided by sku vault's api docs.
type GetSuppliers struct {
	PageNumber                int `json:"PageNumber"`
}

// GetSuppliersResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetSuppliersResponse struct {
	Supplier []struct {
		IsEnabled bool   `json:"IsEnabled"`
		Name      string `json:"Name"`
	} `json:"Supplier"`
}
