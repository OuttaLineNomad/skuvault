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
func (lc *PLoginCredentials) CreateBrands(pld *products.CreateBrands) *products.CreateBrandsResponse {
	credPld := &postCreateBrands{
		CreateBrands:      pld,
		PLoginCredentials: lc,
	}

	response := &products.CreateBrandsResponse{}
	do(credPld, response, "products/createBrands")
	return response
}

// postCreateKit payload sent to Sku Vault.
type postCreateKit struct {
	*products.CreateKit
	*PLoginCredentials
}

// CreateKit creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Using this call, users may create a kit inside of SkuVault..
func (lc *PLoginCredentials) CreateKit(pld *products.CreateKit) *products.CreateKitResponse {
	credPld := &postCreateKit{
		CreateKit:         pld,
		PLoginCredentials: lc,
	}

	response := &products.CreateKitResponse{}
	do(credPld, response, "products/createKit")
	return response
}

// postCreateProduct payload sent to Sku Vault.
type postCreateProduct struct {
	*products.CreateProduct
	*PLoginCredentials
}

// CreateProduct creates http request for this SKU vault endpoint.
// Throttling: Moderate Throttle.
// /createProducts.
func (lc *PLoginCredentials) CreateProduct(pld *products.CreateProduct) *products.CreateProductResponse {
	credPld := &postCreateProduct{
		CreateProduct:     pld,
		PLoginCredentials: lc,
	}

	response := &products.CreateProductResponse{}
	do(credPld, response, "products/createProduct")
	return response
}

// postCreateProducts payload sent to Sku Vault.
type postCreateProducts struct {
	*products.CreateProducts
	*PLoginCredentials
}

// CreateProducts creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Create 100 products per request..
func (lc *PLoginCredentials) CreateProducts(pld *products.CreateProducts) *products.CreateProductsResponse {
	credPld := &postCreateProducts{
		CreateProducts:    pld,
		PLoginCredentials: lc,
	}

	response := &products.CreateProductsResponse{}
	do(credPld, response, "products/createProducts")
	return response
}

// postCreateSuppliers payload sent to Sku Vault.
type postCreateSuppliers struct {
	*products.CreateSuppliers
	*PLoginCredentials
}

// CreateSuppliers creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Returns the list of current Suppliers in a SkuVault account..
func (lc *PLoginCredentials) CreateSuppliers(pld *products.CreateSuppliers) *products.CreateSuppliersResponse {
	credPld := &postCreateSuppliers{
		CreateSuppliers:   pld,
		PLoginCredentials: lc,
	}

	response := &products.CreateSuppliersResponse{}
	do(credPld, response, "products/createSuppliers")
	return response
}

// postGetBrands payload sent to Sku Vault.
type postGetBrands struct {
	*products.GetBrands
	*PLoginCredentials
}

// GetBrands creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns a list of Brands created in SkuVault..
func (lc *PLoginCredentials) GetBrands(pld *products.GetBrands) *products.GetBrandsResponse {
	credPld := &postGetBrands{
		GetBrands:         pld,
		PLoginCredentials: lc,
	}

	response := &products.GetBrandsResponse{}
	do(credPld, response, "products/getBrands")
	return response
}

// postGetClassifications payload sent to Sku Vault.
type postGetClassifications struct {
	*products.GetClassifications
	*PLoginCredentials
}

// GetClassifications creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns classifications and, if they exist, their named attributes..
func (lc *PLoginCredentials) GetClassifications(pld *products.GetClassifications) *products.GetClassificationsResponse {
	credPld := &postGetClassifications{
		GetClassifications: pld,
		PLoginCredentials:  lc,
	}

	response := &products.GetClassificationsResponse{}
	do(credPld, response, "products/getClassifications")
	return response
}

// postGetHandlingTime payload sent to Sku Vault.
type postGetHandlingTime struct {
	*PLoginCredentials
}

// GetHandlingTime creates http request for this SKU vault endpoint.
// No Throttle Throttle.
// No Info.
func (lc *PLoginCredentials) GetHandlingTime() *products.GetHandlingTimeResponse {
	credPld := &postGetHandlingTime{
		PLoginCredentials: lc,
	}

	response := &products.GetHandlingTimeResponse{}
	do(credPld, response, "products/getHandlingTime")
	return response
}

// postGetKits payload sent to Sku Vault.
type postGetKits struct {
	*products.GetKits
	*PLoginCredentials
}

// GetKits creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns kit details..
func (lc *PLoginCredentials) GetKits(pld *products.GetKits) *products.GetKitsResponse {
	credPld := &postGetKits{
		GetKits:           pld,
		PLoginCredentials: lc,
	}

	response := &products.GetKitsResponse{}
	do(credPld, response, "products/getKits")
	return response
}

// postGetProducts payload sent to Sku Vault.
type postGetProducts struct {
	*products.GetProducts
	*PLoginCredentials
}

// GetProducts creates http request for this SKU vault endpoint.
// Heavy Throttle.
// This call returns product (not kit) details. In addition to product information, the response also includes product quantities..
func (lc *PLoginCredentials) GetProducts(pld *products.GetProducts) *products.GetProductsResponse {
	credPld := &postGetProducts{
		GetProducts:       pld,
		PLoginCredentials: lc,
	}

	response := &products.GetProductsResponse{}
	do(credPld, response, "products/getProducts")
	return response
}

// postGetSuppliers payload sent to Sku Vault.
type postGetSuppliers struct {
	*products.GetSuppliers
	*PLoginCredentials
}

// GetSuppliers creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns the list of current Suppliers in a SkuVault account..
func (lc *PLoginCredentials) GetSuppliers(pld *products.GetSuppliers) *products.GetSuppliersResponse {
	credPld := &postGetSuppliers{
		GetSuppliers:      pld,
		PLoginCredentials: lc,
	}

	response := &products.GetSuppliersResponse{}
	do(credPld, response, "products/getSuppliers")
	return response
}

// postUpdateAltSKUsCodes payload sent to Sku Vault.
type postUpdateAltSKUsCodes struct {
	*products.UpdateAltSKUsCodes
	*PLoginCredentials
}

// UpdateAltSKUsCodes creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Update the Alternate SKUs and Codes for your products..
func (lc *PLoginCredentials) UpdateAltSKUsCodes(pld *products.UpdateAltSKUsCodes) *products.UpdateAltSKUsCodesResponse {
	credPld := &postUpdateAltSKUsCodes{
		UpdateAltSKUsCodes: pld,
		PLoginCredentials:  lc,
	}

	response := &products.UpdateAltSKUsCodesResponse{}
	do(credPld, response, "products/updateAltSKUsCodes")
	return response
}

// postUpdateHandlingTime payload sent to Sku Vault.
type postUpdateHandlingTime struct {
	*products.UpdateHandlingTime
	*PLoginCredentials
}

// UpdateHandlingTime creates http request for this SKU vault endpoint.
// Severe Throttle.
// Update the handling time for each product per Amazon channel account, 500 at a time..
func (lc *PLoginCredentials) UpdateHandlingTime(pld *products.UpdateHandlingTime) *products.UpdateHandlingTimeResponse {
	credPld := &postUpdateHandlingTime{
		UpdateHandlingTime: pld,
		PLoginCredentials:  lc,
	}

	response := &products.UpdateHandlingTimeResponse{}
	do(credPld, response, "products/updateHandlingTime")
	return response
}

// postUpdateProduct payload sent to Sku Vault.
type postUpdateProduct struct {
	*products.UpdateProduct
	*PLoginCredentials
}

// UpdateProduct creates http request for this SKU vault endpoint.
// Moderate Throttle.
// /updateProducts.
func (lc *PLoginCredentials) UpdateProduct(pld *products.UpdateProduct) *products.UpdateProductResponse {
	credPld := &postUpdateProduct{
		UpdateProduct:     pld,
		PLoginCredentials: lc,
	}

	response := &products.UpdateProductResponse{}
	do(credPld, response, "products/updateProduct")
	return response
}

// postUpdateProducts payload sent to Sku Vault.
type postUpdateProducts struct {
	*products.UpdateProducts
	*PLoginCredentials
}

// UpdateProducts creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Update products in SkuVault, 100 at a time..
func (lc *PLoginCredentials) UpdateProducts(pld *products.UpdateProducts) *products.UpdateProductsResponse {
	credPld := &postUpdateProducts{
		UpdateProducts:    pld,
		PLoginCredentials: lc,
	}

	response := &products.UpdateProductsResponse{}
	do(credPld, response, "products/updateProducts")
	return response
}
