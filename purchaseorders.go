package skuvault

import "github.com/OuttaLineNomad/skuvault/purchaseorders"

// postCreatePO payload sent to Sku Vault.
type postCreatePO struct {
	*purchaseorders.CreatePO
	*POLoginCredentials
}

// CreatePO creates http request for this SKU vault endpoint.
// Moderate Throttle.
// This call let&#39;s you create a PO using our API..
func (lc *POLoginCredentials) CreatePO(pld *purchaseorders.CreatePO) (*purchaseorders.CreatePOResponse, error) {
	credPld := &postCreatePO{
		CreatePO:           pld,
		POLoginCredentials: lc,
	}

	response := &purchaseorders.CreatePOResponse{}
	err := do(credPld, response, "purchaseorders/createPO")
	return response, err
}

// postGetIncomingItems payload sent to Sku Vault.
type postGetIncomingItems struct {
	*purchaseorders.GetIncomingItems
	*POLoginCredentials
}

// GetIncomingItems creates http request for this SKU vault endpoint.
// Get incoming items for incomplete purchase orders Throttle.
// Get incoming items for incomplete purchase orders.
func (lc *POLoginCredentials) GetIncomingItems(pld *purchaseorders.GetIncomingItems) (*purchaseorders.GetIncomingItemsResponse, error) {
	credPld := &postGetIncomingItems{
		GetIncomingItems:   pld,
		POLoginCredentials: lc,
	}

	response := &purchaseorders.GetIncomingItemsResponse{}
	err := do(credPld, response, "purchaseorders/getIncomingItems")
	return response, err
}

// postGetPOs payload sent to Sku Vault.
type postGetPOs struct {
	*purchaseorders.GetPOs
	*POLoginCredentials
}

// GetPOs creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns a list of purchase orders..
func (lc *POLoginCredentials) GetPOs(pld *purchaseorders.GetPOs) (*purchaseorders.GetPOsResponse, error) {
	credPld := &postGetPOs{
		GetPOs:             pld,
		POLoginCredentials: lc,
	}

	response := &purchaseorders.GetPOsResponse{}
	err := do(credPld, response, "purchaseorders/getPOs")
	return response, err
}

// postGetReceivesHistory payload sent to Sku Vault.
type postGetReceivesHistory struct {
	*purchaseorders.GetReceivesHistory
	*POLoginCredentials
}

// GetReceivesHistory creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns a list of purchase order receives and receipts..
func (lc *POLoginCredentials) GetReceivesHistory(pld *purchaseorders.GetReceivesHistory) (*purchaseorders.GetReceivesHistoryResponse, error) {
	credPld := &postGetReceivesHistory{
		GetReceivesHistory: pld,
		POLoginCredentials: lc,
	}

	response := &purchaseorders.GetReceivesHistoryResponse{}
	err := do(credPld, response, "purchaseorders/getReceivesHistory")
	return response, err
}

// postReceivePOItems payload sent to Sku Vault.
type postReceivePOItems struct {
	*purchaseorders.ReceivePOItems
	*POLoginCredentials
}

// ReceivePOItems creates http request for this SKU vault endpoint.
// Moderate Throttle.
//  throttling.
func (lc *POLoginCredentials) ReceivePOItems(pld *purchaseorders.ReceivePOItems) (*purchaseorders.ReceivePOItemsResponse, error) {
	credPld := &postReceivePOItems{
		ReceivePOItems:     pld,
		POLoginCredentials: lc,
	}

	response := &purchaseorders.ReceivePOItemsResponse{}
	err := do(credPld, response, "purchaseorders/receivePOItems")
	return response, err
}

// postUpdatePOs payload sent to Sku Vault.
type postUpdatePOs struct {
	*purchaseorders.UpdatePOs
	*POLoginCredentials
}

// UpdatePOs creates http request for this SKU vault endpoint.
// Heavy Throttle.
// /updatePOs.
func (lc *POLoginCredentials) UpdatePOs(pld *purchaseorders.UpdatePOs) (*purchaseorders.UpdatePOsResponse, error) {
	credPld := &postUpdatePOs{
		UpdatePOs:          pld,
		POLoginCredentials: lc,
	}

	response := &purchaseorders.UpdatePOsResponse{}
	err := do(credPld, response, "purchaseorders/updatePOs")
	return response, err
}
