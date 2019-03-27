package skuvault

import "github.com/OuttaLineNomad/skuvault/products"

// postCreateBrands payload sent to Sku Vault.
type postCreateBrands struct {
	*products.CreateBrands
	*PLoginCredentials
}

// CreateBrands creates http request for this SKU vault endpoint.
// Moderate Throttle.
//  throttling.
func (lc *PLoginCredentials) CreateBrands(pld *products.CreateBrands) (*products.CreateBrandsResponse, error) {
	credPld := &postCreateBrands{
		CreateBrands:      pld,
		PLoginCredentials: lc,
	}

	response := &products.CreateBrandsResponse{}
	err := do(credPld, response, "products/createBrands")
	if err != nil {
		return nil, &Error{"CreateBrands()", err}
	}

	return response, nil
}

// postCreateKit payload sent to Sku Vault.
type postCreateKit struct {
	*products.CreateKit
	*PLoginCredentials
}

// CreateKit creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Using this call, users may create a kit inside of SkuVault..
func (lc *PLoginCredentials) CreateKit(pld *products.CreateKit) (*products.CreateKitResponse, error) {
	credPld := &postCreateKit{
		CreateKit:         pld,
		PLoginCredentials: lc,
	}

	response := &products.CreateKitResponse{}
	err := do(credPld, response, "products/createKit")
	if err != nil {
		return nil, &Error{"CreateKit()", err}
	}

	return response, nil
}

// postCreateProduct payload sent to Sku Vault.
type postCreateProduct struct {
	*products.CreateProduct
	*PLoginCredentials
}

// CreateProduct creates http request for this SKU vault endpoint.
// Throttling: Moderate Throttle.
// /createProducts.
func (lc *PLoginCredentials) CreateProduct(pld *products.CreateProduct) (*products.CreateProductResponse, error) {
	credPld := &postCreateProduct{
		CreateProduct:     pld,
		PLoginCredentials: lc,
	}

	response := &products.CreateProductResponse{}
	err := do(credPld, response, "products/createProduct")
	if err != nil {
		return nil, &Error{"CreateProduct()", err}
	}

	return response, nil
}

// postCreateProducts payload sent to Sku Vault.
type postCreateProducts struct {
	*products.CreateProducts
	*PLoginCredentials
}

// CreateProducts creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Create 100 products per request..
func (lc *PLoginCredentials) CreateProducts(pld *products.CreateProducts) (*products.CreateProductsResponse, error) {
	credPld := &postCreateProducts{
		CreateProducts:    pld,
		PLoginCredentials: lc,
	}

	response := &products.CreateProductsResponse{}
	err := do(credPld, response, "products/createProducts")
	if err != nil {
		return nil, &Error{"CreateProducts()", err}
	}

	return response, nil
}

// postCreateSuppliers payload sent to Sku Vault.
type postCreateSuppliers struct {
	*products.CreateSuppliers
	*PLoginCredentials
}

// CreateSuppliers creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Returns the list of current Suppliers in a SkuVault account..
func (lc *PLoginCredentials) CreateSuppliers(pld *products.CreateSuppliers) (*products.CreateSuppliersResponse, error) {
	credPld := &postCreateSuppliers{
		CreateSuppliers:   pld,
		PLoginCredentials: lc,
	}

	response := &products.CreateSuppliersResponse{}
	err := do(credPld, response, "products/createSuppliers")
	if err != nil {
		return nil, &Error{"CreateSuppliers()", err}
	}

	return response, nil
}

// postGetBrands payload sent to Sku Vault.
type postGetBrands struct {
	*products.GetBrands
	*PLoginCredentials
}

// GetBrands creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns a list of Brands created in SkuVault..
func (lc *PLoginCredentials) GetBrands(pld *products.GetBrands) (*products.GetBrandsResponse, error) {
	credPld := &postGetBrands{
		GetBrands:         pld,
		PLoginCredentials: lc,
	}

	response := &products.GetBrandsResponse{}
	err := do(credPld, response, "products/getBrands")
	if err != nil {
		return nil, &Error{"GetBrands()", err}
	}

	return response, nil
}

// postGetClassifications payload sent to Sku Vault.
type postGetClassifications struct {
	*products.GetClassifications
	*PLoginCredentials
}

// GetClassifications creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns classifications and, if they exist, their named attributes..
func (lc *PLoginCredentials) GetClassifications(pld *products.GetClassifications) (*products.GetClassificationsResponse, error) {
	credPld := &postGetClassifications{
		GetClassifications: pld,
		PLoginCredentials:  lc,
	}

	response := &products.GetClassificationsResponse{}
	err := do(credPld, response, "products/getClassifications")
	if err != nil {
		return nil, &Error{"GetClassifications()", err}
	}

	return response, nil
}

// postGetHandlingTime payload sent to Sku Vault.
type postGetHandlingTime struct {
	*PLoginCredentials
}

// GetHandlingTime creates http request for this SKU vault endpoint.
// No Throttle Throttle.
// No Info.
func (lc *PLoginCredentials) GetHandlingTime() (*products.GetHandlingTimeResponse, error) {
	credPld := &postGetHandlingTime{
		PLoginCredentials: lc,
	}

	response := &products.GetHandlingTimeResponse{}
	err := do(credPld, response, "products/getHandlingTime")
	if err != nil {
		return nil, &Error{"GetHandlingTime()", err}
	}

	return response, nil
}

// postGetKits payload sent to Sku Vault.
type postGetKits struct {
	*products.GetKits
	*PLoginCredentials
}

// GetKits creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns kit details..
func (lc *PLoginCredentials) GetKits(pld *products.GetKits) (*products.GetKitsResponse, error) {
	credPld := &postGetKits{
		GetKits:           pld,
		PLoginCredentials: lc,
	}

	response := &products.GetKitsResponse{}
	err := do(credPld, response, "products/getKits")
	if err != nil {
		return nil, &Error{"GetKits()", err}
	}

	return response, nil
}

// postGetProducts payload sent to Sku Vault.
type postGetProducts struct {
	*products.GetProducts
	*PLoginCredentials
}

// GetProducts creates http request for this SKU vault endpoint.
// Heavy Throttle.
// This call returns product (not kit) details. In addition to product information, the response also includes product quantities..
func (lc *PLoginCredentials) GetProducts(pld *products.GetProducts) (*products.GetProductsResponse, error) {
	credPld := &postGetProducts{
		GetProducts:       pld,
		PLoginCredentials: lc,
	}

	response := &products.GetProductsResponse{}
	err := do(credPld, response, "products/getProducts")
	if err != nil {
		return nil, &Error{"GetProducts()", err}
	}

	return response, nil
}

// postGetSuppliers payload sent to Sku Vault.
type postGetSuppliers struct {
	*products.GetSuppliers
	*PLoginCredentials
}

// GetSuppliers creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns the list of current Suppliers in a SkuVault account..
func (lc *PLoginCredentials) GetSuppliers(pld *products.GetSuppliers) (*products.GetSuppliersResponse, error) {
	credPld := &postGetSuppliers{
		GetSuppliers:      pld,
		PLoginCredentials: lc,
	}

	response := &products.GetSuppliersResponse{}
	err := do(credPld, response, "products/getSuppliers")
	if err != nil {
		return nil, &Error{"GetSuppliers()", err}
	}

	return response, nil
}

// postUpdateAltSKUsCodes payload sent to Sku Vault.
type postUpdateAltSKUsCodes struct {
	*products.UpdateAltSKUsCodes
	*PLoginCredentials
}

// UpdateAltSKUsCodes creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Update the Alternate SKUs and Codes for your products..
func (lc *PLoginCredentials) UpdateAltSKUsCodes(pld *products.UpdateAltSKUsCodes) (*products.UpdateAltSKUsCodesResponse, error) {
	credPld := &postUpdateAltSKUsCodes{
		UpdateAltSKUsCodes: pld,
		PLoginCredentials:  lc,
	}

	response := &products.UpdateAltSKUsCodesResponse{}
	err := do(credPld, response, "products/updateAltSKUsCodes")
	if err != nil {
		return nil, &Error{"UpdateAltSKUsCodes()", err}
	}

	return response, nil
}

// postUpdateHandlingTime payload sent to Sku Vault.
type postUpdateHandlingTime struct {
	*products.UpdateHandlingTime
	*PLoginCredentials
}

// UpdateHandlingTime creates http request for this SKU vault endpoint.
// Severe Throttle.
// Update the handling time for each product per Amazon channel account, 500 at a time..
func (lc *PLoginCredentials) UpdateHandlingTime(pld *products.UpdateHandlingTime) (*products.UpdateHandlingTimeResponse, error) {
	credPld := &postUpdateHandlingTime{
		UpdateHandlingTime: pld,
		PLoginCredentials:  lc,
	}

	response := &products.UpdateHandlingTimeResponse{}
	err := do(credPld, response, "products/updateHandlingTime")
	if err != nil {
		return nil, &Error{"UpdateHandlingTime()", err}
	}

	return response, nil
}

// postUpdateProduct payload sent to Sku Vault.
type postUpdateProduct struct {
	*products.UpdateProduct
	*PLoginCredentials
}

// UpdateProduct creates http request for this SKU vault endpoint.
// Moderate Throttle.
// /updateProducts.
func (lc *PLoginCredentials) UpdateProduct(pld *products.UpdateProduct) (*products.UpdateProductResponse, error) {
	credPld := &postUpdateProduct{
		UpdateProduct:     pld,
		PLoginCredentials: lc,
	}

	response := &products.UpdateProductResponse{}
	err := do(credPld, response, "products/updateProduct")
	if err != nil {
		return nil, &Error{"UpdateProduct()", err}
	}

	return response, nil
}

// postUpdateProducts payload sent to Sku Vault.
type postUpdateProducts struct {
	*products.UpdateProducts
	*PLoginCredentials
}

// UpdateProducts creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Update products in SkuVault, 100 at a time..
func (lc *PLoginCredentials) UpdateProducts(pld *products.UpdateProducts) (*products.UpdateProductsResponse, error) {
	credPld := &postUpdateProducts{
		UpdateProducts:    pld,
		PLoginCredentials: lc,
	}

	response := &products.UpdateProductsResponse{}
	err := do(credPld, response, "products/updateProducts")
	if err != nil {
		return nil, &Error{"UpdateProducts()", err}
	}

	return response, nil
}
