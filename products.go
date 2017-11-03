package skuvault

import "github.com/OuttaLineNomad/skuvault/products"

// holds all api call endpoints with that are in the folder /Products

// postGetWarehouseItemQuantity payload sent to Sku Vault
type postGetProducts struct {
	*products.GetProducts
	*PLoginCredentials
}

// postAddItem  payload sent to Sku Vault.
type postAddItem struct {
	*products.AddItem
	*PLoginCredentials
}

// AddItem creates http request for this SKU vault endpoint
// Moderate throttling
// Add quantity to a warehouse location. Bulk version available
func (lc *PLoginCredentials) AddItem(pld *products.AddItem) *products.AddItemResponse {
	credPld := &postAddItem{
		AddItem:           pld,
		PLoginCredentials: lc,
	}

	response := &products.AddItemResponse{}
	productsCall(credPld, response, "addItem")
	return response
}

// This call returns product(not kit) details. The date parameters include updating details as well as product
// creation. Details do not include quantity.
func (lc *PLoginCredentials) GetProducts(pld *products.GetProducts) *products.GetProductsResponse {
	credPld := &postGetProducts{
		GetProducts:       pld,
		PLoginCredentials: lc,
	}

	response := &products.GetProductsResponse{}
	productsCall(credPld, response, "getProducts")
	return response
}

// productsCall adds products/ to url for do() call.
func productsCall(pld interface{}, response interface{}, endPoint string) {
	full := "products/" + endPoint
	do(pld, response, full)
}
