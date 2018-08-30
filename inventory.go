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
func (lc *ILoginCredentials) GetWarehouseItemQuantity(pld *inventory.GetWarehouseItemQuantity) (*inventory.GetWarehouseItemQuantityResponse, error) {
	credPld := &postGetWarehouseItemQuantity{
		GetWarehouseItemQuantity: pld,
		ILoginCredentials:        lc,
	}

	response := &inventory.GetWarehouseItemQuantityResponse{}
	err := inventoryCall(credPld, response, "getWarehouseItemQuantity")
	return response, err
}

// GetTransactions creates http request for this SKU vault endpoint
// Heavy throttling
// Look at your transaction history.
func (lc *ILoginCredentials) GetTransactions(pld *inventory.GetTransactions) (*inventory.GetTransactionsResponse, error) {
	credPld := &postGetTransactions{
		GetTransactions:   pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetTransactionsResponse{}
	err := inventoryCall(credPld, response, "getTransactions")
	return response, err

}

// GetInventoryByLocation creates http request for this SKU vault endpoint
// Heavy throttling
// Returns location and warehouse per product.
func (lc *ILoginCredentials) GetInventoryByLocation(pld *inventory.GetInventoryByLocation) (*inventory.GetInventoryByLocationResponse, error) {
	credPld := &postGetInventoryByLocation{
		GetInventoryByLocation: pld,
		ILoginCredentials:      lc,
	}
	response := &inventory.GetInventoryByLocationResponse{}
	err := inventoryCall(credPld, response, "getInventoryByLocation")
	return response, err

}

// inventoryCalladds inventory/ to url for do call.
func inventoryCall(pld interface{}, response interface{}, endPoint string) error {
	full := "inventory/" + endPoint
	return do(pld, response, full)
}
