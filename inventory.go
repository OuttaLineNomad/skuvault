package skuvault

import "github.com/OuttaLineNomad/skuvault/inventory"

// holds all api call endpoints for calles to SKU Vault API under /Inventory.

// postGetWarehouseItemQuantity payload sent to Sku Vault
type postGetWarehouseItemQuantity struct {
	*inventory.GetWarehouseItemQuantity
	*ILoginCredentials
}

// postGetTransactions payload sent to Sku Vault
type postGetTransactions struct {
	*inventory.GetTransactions
	*ILoginCredentials
}

// postGetInventoryByLocation payload sent to Sku Vault
type postGetInventoryByLocation struct {
	*inventory.GetInventoryByLocation
	*ILoginCredentials
}

// GetWarehouseItemQuantity creates http request for this SKU vault endpoint
// Heavy throttling
// Returns the quantity for a specified SKU
func (lc *ILoginCredentials) GetWarehouseItemQuantity(pld *inventory.GetWarehouseItemQuantity) *inventory.GetWarehouseItemQuantityResponse {
	credPld := &postGetWarehouseItemQuantity{
		GetWarehouseItemQuantity: pld,
		ILoginCredentials:        lc,
	}

	response := &inventory.GetWarehouseItemQuantityResponse{}
	inventoryCall(credPld, response, "getWarehouseItemQuantity")
	return response
}

// GetTransactions creates http request for this SKU vault endpoint
// Heavy throttling
// Look at your transaction history.
func (lc *ILoginCredentials) GetTransactions(pld *inventory.GetTransactions) *inventory.GetTransactionsResponse {
	credPld := &postGetTransactions{
		GetTransactions:   pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetTransactionsResponse{}
	inventoryCall(credPld, response, "getTransactions")
	return response

}

// GetInventoryByLocation creates http request for this SKU vault endpoint
// Heavy throttling
// Returns location and warehouse per product.
func (lc *ILoginCredentials) GetInventoryByLocation(pld *inventory.GetInventoryByLocation) *inventory.GetInventoryByLocationResponse {
	credPld := &postGetInventoryByLocation{
		GetInventoryByLocation: pld,
		ILoginCredentials:      lc,
	}
	response := &inventory.GetInventoryByLocationResponse{}
	inventoryCall(credPld, response, "getInventoryByLocation")
	return response

}

// inventoryCalladds inventory/ to url for do call.
func inventoryCall(pld interface{}, response interface{}, endPoint string) {
	full := "inventory/" + endPoint
	do(pld, response, full)
}
