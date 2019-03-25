package products

// CreateSuppliers is a automatically generated struct from json provided by sku vault's api docs.
type CreateSuppliers struct {
	Suppliers []struct {
		EmailTemplateMessage string   `json:"EmailTemplateMessage"`
		EmailTemplateSubject string   `json:"EmailTemplateSubject"`
		Emails               []string `json:"Emails"`
		Name                 string   `json:"Name"`
	} `json:"Suppliers"`
}

// CreateSuppliersResponse is a automatically generated struct from json provided by sku vault's api docs.
type CreateSuppliersResponse struct {
	Errors []string `json:"Errors"`
}
