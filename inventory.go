package skuvault

import "github.com/OuttaLineNomad/skuvault/inventory"

// postAddItem payload sent to Sku Vault.
type postAddItem struct {
	*inventory.AddItem
	*ILoginCredentials
}

// AddItem creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Bulk version available.
func (lc *ILoginCredentials) AddItem(pld *inventory.AddItem) *inventory.AddItemResponse {
	credPld := &postAddItem{
		AddItem:           pld,
		ILoginCredentials: lc,
	}

	response := &inventory.AddItemResponse{}
	do(credPld, response, "inventory/addItem")
	return response
}

// postAddItemBulk payload sent to Sku Vault.
type postAddItemBulk struct {
	*inventory.AddItemBulk
	*ILoginCredentials
}

// AddItemBulk creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Add quantity to warehouse locations, 100 at a time..
func (lc *ILoginCredentials) AddItemBulk(pld *inventory.AddItemBulk) *inventory.AddItemBulkResponse {
	credPld := &postAddItemBulk{
		AddItemBulk:       pld,
		ILoginCredentials: lc,
	}

	response := &inventory.AddItemBulkResponse{}
	do(credPld, response, "inventory/addItemBulk")
	return response
}

// postGetAvailableQuantities payload sent to Sku Vault.
type postGetAvailableQuantities struct {
	*inventory.GetAvailableQuantities
	*ILoginCredentials
}

// GetAvailableQuantities creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Retrieve a list of SKUs and their total available quantities across all warehouses. Available Quantity is the quantity that is actually available to sell across all your sales channels..
func (lc *ILoginCredentials) GetAvailableQuantities(pld *inventory.GetAvailableQuantities) *inventory.GetAvailableQuantitiesResponse {
	credPld := &postGetAvailableQuantities{
		GetAvailableQuantities: pld,
		ILoginCredentials:      lc,
	}

	response := &inventory.GetAvailableQuantitiesResponse{}
	do(credPld, response, "inventory/getAvailableQuantities")
	return response
}

// postGetExternalWarehouseQuantities payload sent to Sku Vault.
type postGetExternalWarehouseQuantities struct {
	*inventory.GetExternalWarehouseQuantities
	*ILoginCredentials
}

// GetExternalWarehouseQuantities creates http request for this SKU vault endpoint.
// Heavy Throttle.
// This call gets the quantities in a designated External Warehouse. Please note these are different than ordinary warehouses..
func (lc *ILoginCredentials) GetExternalWarehouseQuantities(pld *inventory.GetExternalWarehouseQuantities) *inventory.GetExternalWarehouseQuantitiesResponse {
	credPld := &postGetExternalWarehouseQuantities{
		GetExternalWarehouseQuantities: pld,
		ILoginCredentials:              lc,
	}

	response := &inventory.GetExternalWarehouseQuantitiesResponse{}
	do(credPld, response, "inventory/getExternalWarehouseQuantities")
	return response
}

// postGetExternalWarehouses payload sent to Sku Vault.
type postGetExternalWarehouses struct {
	*ILoginCredentials
}

// GetExternalWarehouses creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Returns your external warehouses. No page parameters..
func (lc *ILoginCredentials) GetExternalWarehouses() *inventory.GetExternalWarehousesResponse {
	credPld := &postGetExternalWarehouses{
		ILoginCredentials: lc,
	}

	response := &inventory.GetExternalWarehousesResponse{}
	do(credPld, response, "inventory/getExternalWarehouses")
	return response
}

// postGetInventoryByLocation payload sent to Sku Vault.
type postGetInventoryByLocation struct {
	*inventory.GetInventoryByLocation
	*ILoginCredentials
}

// GetInventoryByLocation creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns location and warehouse per product..
func (lc *ILoginCredentials) GetInventoryByLocation(pld *inventory.GetInventoryByLocation) *inventory.GetInventoryByLocationResponse {
	credPld := &postGetInventoryByLocation{
		GetInventoryByLocation: pld,
		ILoginCredentials:      lc,
	}

	response := &inventory.GetInventoryByLocationResponse{}
	do(credPld, response, "inventory/getInventoryByLocation")
	return response
}

// postGetItemQuantities payload sent to Sku Vault.
type postGetItemQuantities struct {
	*inventory.GetItemQuantities
	*ILoginCredentials
}

// GetItemQuantities creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns product quantities..
func (lc *ILoginCredentials) GetItemQuantities(pld *inventory.GetItemQuantities) *inventory.GetItemQuantitiesResponse {
	credPld := &postGetItemQuantities{
		GetItemQuantities: pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetItemQuantitiesResponse{}
	do(credPld, response, "inventory/getItemQuantities")
	return response
}

// postGetKitQuantities payload sent to Sku Vault.
type postGetKitQuantities struct {
	*inventory.GetKitQuantities
	*ILoginCredentials
}

// GetKitQuantities creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns kit quantities..
func (lc *ILoginCredentials) GetKitQuantities(pld *inventory.GetKitQuantities) *inventory.GetKitQuantitiesResponse {
	credPld := &postGetKitQuantities{
		GetKitQuantities:  pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetKitQuantitiesResponse{}
	do(credPld, response, "inventory/getKitQuantities")
	return response
}

// postGetLocations payload sent to Sku Vault.
type postGetLocations struct {
	*ILoginCredentials
}

// GetLocations creates http request for this SKU vault endpoint.
// Severe Throttle.
//  throttling. Returns all your locations in enabled warehouses..
func (lc *ILoginCredentials) GetLocations() *inventory.GetLocationsResponse {
	credPld := &postGetLocations{
		ILoginCredentials: lc,
	}

	response := &inventory.GetLocationsResponse{}
	do(credPld, response, "inventory/getLocations")
	return response
}

// postGetTransactions payload sent to Sku Vault.
type postGetTransactions struct {
	*inventory.GetTransactions
	*ILoginCredentials
}

// GetTransactions creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Look at your transaction history..
func (lc *ILoginCredentials) GetTransactions(pld *inventory.GetTransactions) *inventory.GetTransactionsResponse {
	credPld := &postGetTransactions{
		GetTransactions:   pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetTransactionsResponse{}
	do(credPld, response, "inventory/getTransactions")
	return response
}

// postGetWarehouseItemQuantities payload sent to Sku Vault.
type postGetWarehouseItemQuantities struct {
	*inventory.GetWarehouseItemQuantities
	*ILoginCredentials
}

// GetWarehouseItemQuantities creates http request for this SKU vault endpoint.
// Heavy Throttle.
// This call returns SKUs and quantities from a specified warehouse. 10,000 SKUs returned per page..
func (lc *ILoginCredentials) GetWarehouseItemQuantities(pld *inventory.GetWarehouseItemQuantities) *inventory.GetWarehouseItemQuantitiesResponse {
	credPld := &postGetWarehouseItemQuantities{
		GetWarehouseItemQuantities: pld,
		ILoginCredentials:          lc,
	}

	response := &inventory.GetWarehouseItemQuantitiesResponse{}
	do(credPld, response, "inventory/getWarehouseItemQuantities")
	return response
}

// postGetWarehouseItemQuantity payload sent to Sku Vault.
type postGetWarehouseItemQuantity struct {
	*inventory.GetWarehouseItemQuantity
	*ILoginCredentials
}

// GetWarehouseItemQuantity creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns the quantity for a specified SKU..
func (lc *ILoginCredentials) GetWarehouseItemQuantity(pld *inventory.GetWarehouseItemQuantity) *inventory.GetWarehouseItemQuantityResponse {
	credPld := &postGetWarehouseItemQuantity{
		GetWarehouseItemQuantity: pld,
		ILoginCredentials:        lc,
	}

	response := &inventory.GetWarehouseItemQuantityResponse{}
	do(credPld, response, "inventory/getWarehouseItemQuantity")
	return response
}

// postGetWarehouses payload sent to Sku Vault.
type postGetWarehouses struct {
	*inventory.GetWarehouses
	*ILoginCredentials
}

// GetWarehouses creates http request for this SKU vault endpoint.
// Severe Throttle.
//  throttling. Returns all your regular warehouses..
func (lc *ILoginCredentials) GetWarehouses(pld *inventory.GetWarehouses) *inventory.GetWarehousesResponse {
	credPld := &postGetWarehouses{
		GetWarehouses:     pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetWarehousesResponse{}
	do(credPld, response, "inventory/getWarehouses")
	return response
}

// postPickItem payload sent to Sku Vault.
type postPickItem struct {
	*inventory.PickItem
	*ILoginCredentials
}

// PickItem creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Perform a pick transaction through the API. [Bulk version available].
func (lc *ILoginCredentials) PickItem(pld *inventory.PickItem) *inventory.PickItemResponse {
	credPld := &postPickItem{
		PickItem:          pld,
		ILoginCredentials: lc,
	}

	response := &inventory.PickItemResponse{}
	do(credPld, response, "inventory/pickItem")
	return response
}

// postPickItemBulk payload sent to Sku Vault.
type postPickItemBulk struct {
	*inventory.PickItemBulk
	*ILoginCredentials
}

// PickItemBulk creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Perform a bulk pick transaction through the API..
func (lc *ILoginCredentials) PickItemBulk(pld *inventory.PickItemBulk) *inventory.PickItemBulkResponse {
	credPld := &postPickItemBulk{
		PickItemBulk:      pld,
		ILoginCredentials: lc,
	}

	response := &inventory.PickItemBulkResponse{}
	do(credPld, response, "inventory/pickItemBulk")
	return response
}

// postRemoveItem payload sent to Sku Vault.
type postRemoveItem struct {
	*inventory.RemoveItem
	*ILoginCredentials
}

// RemoveItem creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Bulk version available.
func (lc *ILoginCredentials) RemoveItem(pld *inventory.RemoveItem) *inventory.RemoveItemResponse {
	credPld := &postRemoveItem{
		RemoveItem:        pld,
		ILoginCredentials: lc,
	}

	response := &inventory.RemoveItemResponse{}
	do(credPld, response, "inventory/removeItem")
	return response
}

// postRemoveItemBulk payload sent to Sku Vault.
type postRemoveItemBulk struct {
	*inventory.RemoveItemBulk
	*ILoginCredentials
}

// RemoveItemBulk creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Remove quantity from warehouse locations, 100 at a time..
func (lc *ILoginCredentials) RemoveItemBulk(pld *inventory.RemoveItemBulk) *inventory.RemoveItemBulkResponse {
	credPld := &postRemoveItemBulk{
		RemoveItemBulk:    pld,
		ILoginCredentials: lc,
	}

	response := &inventory.RemoveItemBulkResponse{}
	do(credPld, response, "inventory/removeItemBulk")
	return response
}

// postSetItemQuantities payload sent to Sku Vault.
type postSetItemQuantities struct {
	*inventory.SetItemQuantities
	*ILoginCredentials
}

// SetItemQuantities creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Sets quantity for multiple products in one request..
func (lc *ILoginCredentials) SetItemQuantities(pld *inventory.SetItemQuantities) *inventory.SetItemQuantitiesResponse {
	credPld := &postSetItemQuantities{
		SetItemQuantities: pld,
		ILoginCredentials: lc,
	}

	response := &inventory.SetItemQuantitiesResponse{}
	do(credPld, response, "inventory/setItemQuantities")
	return response
}

// postSetItemQuantity payload sent to Sku Vault.
type postSetItemQuantity struct {
	*inventory.SetItemQuantity
	*ILoginCredentials
}

// SetItemQuantity creates http request for this SKU vault endpoint.
// Moderate Throttle.
// This lets you explicitly set quantity for an item in a warehouse&#39;s location..
func (lc *ILoginCredentials) SetItemQuantity(pld *inventory.SetItemQuantity) *inventory.SetItemQuantityResponse {
	credPld := &postSetItemQuantity{
		SetItemQuantity:   pld,
		ILoginCredentials: lc,
	}

	response := &inventory.SetItemQuantityResponse{}
	do(credPld, response, "inventory/setItemQuantity")
	return response
}

// postUpdateExternalWarehouseQuantities payload sent to Sku Vault.
type postUpdateExternalWarehouseQuantities struct {
	*inventory.UpdateExternalWarehouseQuantities
	*ILoginCredentials
}

// UpdateExternalWarehouseQuantities creates http request for this SKU vault endpoint.
// Severe Throttle.
// Set the quantity of SKUs in a specified external warehouse. Updating an external warehouse&#39;s quantities will overwrite any existing quantities. The limit is 200,000 SKUs per call..
func (lc *ILoginCredentials) UpdateExternalWarehouseQuantities(pld *inventory.UpdateExternalWarehouseQuantities) *inventory.UpdateExternalWarehouseQuantitiesResponse {
	credPld := &postUpdateExternalWarehouseQuantities{
		UpdateExternalWarehouseQuantities: pld,
		ILoginCredentials:                 lc,
	}

	response := &inventory.UpdateExternalWarehouseQuantitiesResponse{}
	do(credPld, response, "inventory/updateExternalWarehouseQuantities")
	return response
}
