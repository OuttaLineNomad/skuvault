package skuvault

import "github.com/OuttaLineNomad/skuvault/sales"

// postAddShipments payload sent to Sku Vault.
type postAddShipments struct {
	*sales.AddShipments
	*SLoginCredentials
}

// AddShipments creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Using this call, users can add shipments to a sale..
func (lc *SLoginCredentials) AddShipments(pld *sales.AddShipments) *sales.AddShipmentsResponse {
	credPld := &postAddShipments{
		AddShipments:      pld,
		SLoginCredentials: lc,
	}

	response := &sales.AddShipmentsResponse{}
	do(credPld, response, "sales/addShipments")
	return response
}

// postCreateHolds payload sent to Sku Vault.
type postCreateHolds struct {
	*sales.CreateHolds
	*SLoginCredentials
}

// CreateHolds creates http request for this SKU vault endpoint.
// Moderate Throttle.
//  throttling.
func (lc *SLoginCredentials) CreateHolds(pld *sales.CreateHolds) *sales.CreateHoldsResponse {
	credPld := &postCreateHolds{
		CreateHolds:       pld,
		SLoginCredentials: lc,
	}

	response := &sales.CreateHoldsResponse{}
	do(credPld, response, "sales/createHolds")
	return response
}

// postGetOnlineSaleStatus payload sent to Sku Vault.
type postGetOnlineSaleStatus struct {
	*sales.GetOnlineSaleStatus
	*SLoginCredentials
}

// GetOnlineSaleStatus creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns a list of sales and their statuses..
func (lc *SLoginCredentials) GetOnlineSaleStatus(pld *sales.GetOnlineSaleStatus) *sales.GetOnlineSaleStatusResponse {
	credPld := &postGetOnlineSaleStatus{
		GetOnlineSaleStatus: pld,
		SLoginCredentials:   lc,
	}

	response := &sales.GetOnlineSaleStatusResponse{}
	do(credPld, response, "sales/getOnlineSaleStatus")
	return response
}

// postGetSaleItemCost payload sent to Sku Vault.
type postGetSaleItemCost struct {
	*sales.GetSaleItemCost
	*SLoginCredentials
}

// GetSaleItemCost creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Retrieves the Item Cost for a Product for a Specific Sale..
func (lc *SLoginCredentials) GetSaleItemCost(pld *sales.GetSaleItemCost) *sales.GetSaleItemCostResponse {
	credPld := &postGetSaleItemCost{
		GetSaleItemCost:   pld,
		SLoginCredentials: lc,
	}

	response := &sales.GetSaleItemCostResponse{}
	do(credPld, response, "sales/getSaleItemCost")
	return response
}

// postGetSales payload sent to Sku Vault.
type postGetSales struct {
	*sales.GetSales
	*SLoginCredentials
}

// GetSales creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Use this call to retrieve a list of sales from SkuVault. 10,000 sales are returned per page..
func (lc *SLoginCredentials) GetSales(pld *sales.GetSales) *sales.GetSalesResponse {
	credPld := &postGetSales{
		GetSales:          pld,
		SLoginCredentials: lc,
	}

	response := &sales.GetSalesResponse{}
	do(credPld, response, "sales/getSales")
	return response
}

// postGetSalesByDate payload sent to Sku Vault.
type postGetSalesByDate struct {
	*sales.GetSalesByDate
	*SLoginCredentials
}

// GetSalesByDate creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns sales based on a date range. 10,000 sales are returned per page..
func (lc *SLoginCredentials) GetSalesByDate(pld *sales.GetSalesByDate) *sales.GetSalesByDateResponse {
	credPld := &postGetSalesByDate{
		GetSalesByDate:    pld,
		SLoginCredentials: lc,
	}

	response := &sales.GetSalesByDateResponse{}
	do(credPld, response, "sales/getSalesByDate")
	return response
}

// postGetShipments payload sent to Sku Vault.
type postGetShipments struct {
	*sales.GetShipments
	*SLoginCredentials
}

// GetShipments creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Using this call, users can get current shipment information..
func (lc *SLoginCredentials) GetShipments(pld *sales.GetShipments) *sales.GetShipmentsResponse {
	credPld := &postGetShipments{
		GetShipments:      pld,
		SLoginCredentials: lc,
	}

	response := &sales.GetShipmentsResponse{}
	do(credPld, response, "sales/getShipments")
	return response
}

// postGetSoldItems payload sent to Sku Vault.
type postGetSoldItems struct {
	*sales.GetSoldItems
	*SLoginCredentials
}

// GetSoldItems creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns a list of sold items filtered by date. 10,000 sales are returned per page..
func (lc *SLoginCredentials) GetSoldItems(pld *sales.GetSoldItems) *sales.GetSoldItemsResponse {
	credPld := &postGetSoldItems{
		GetSoldItems:      pld,
		SLoginCredentials: lc,
	}

	response := &sales.GetSoldItemsResponse{}
	do(credPld, response, "sales/getSoldItems")
	return response
}

// postReleaseHeldQuantities payload sent to Sku Vault.
type postReleaseHeldQuantities struct {
	*sales.ReleaseHeldQuantities
	*SLoginCredentials
}

// ReleaseHeldQuantities creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Release holds before their expiration date expires..
func (lc *SLoginCredentials) ReleaseHeldQuantities(pld *sales.ReleaseHeldQuantities) *sales.ReleaseHeldQuantitiesResponse {
	credPld := &postReleaseHeldQuantities{
		ReleaseHeldQuantities: pld,
		SLoginCredentials:     lc,
	}

	response := &sales.ReleaseHeldQuantitiesResponse{}
	do(credPld, response, "sales/releaseHeldQuantities")
	return response
}

// postSetShipmentFile payload sent to Sku Vault.
type postSetShipmentFile struct {
	*sales.SetShipmentFile
	*SLoginCredentials
}

// SetShipmentFile creates http request for this SKU vault endpoint.
// Severe Throttle.
// Using this call, users can attach a base64 PDF file to Shipment Tracking Number..
func (lc *SLoginCredentials) SetShipmentFile(pld *sales.SetShipmentFile) *sales.SetShipmentFileResponse {
	credPld := &postSetShipmentFile{
		SetShipmentFile:   pld,
		SLoginCredentials: lc,
	}

	response := &sales.SetShipmentFileResponse{}
	do(credPld, response, "sales/setShipmentFile")
	return response
}

// postSyncOnlineSale payload sent to Sku Vault.
type postSyncOnlineSale struct {
	*sales.SyncOnlineSale
	*SLoginCredentials
}

// SyncOnlineSale creates http request for this SKU vault endpoint.
// Moderate Throttle.
// /syncOnlineSales.
func (lc *SLoginCredentials) SyncOnlineSale(pld *sales.SyncOnlineSale) *sales.SyncOnlineSaleResponse {
	credPld := &postSyncOnlineSale{
		SyncOnlineSale:    pld,
		SLoginCredentials: lc,
	}

	response := &sales.SyncOnlineSaleResponse{}
	do(credPld, response, "sales/syncOnlineSale")
	return response
}

// postSyncOnlineSales payload sent to Sku Vault.
type postSyncOnlineSales struct {
	*sales.SyncOnlineSales
	*SLoginCredentials
}

// SyncOnlineSales creates http request for this SKU vault endpoint.
// Severe Throttle.
//  Can make this call 2x per minute, 100 sales max.
func (lc *SLoginCredentials) SyncOnlineSales(pld *sales.SyncOnlineSales) *sales.SyncOnlineSalesResponse {
	credPld := &postSyncOnlineSales{
		SyncOnlineSales:   pld,
		SLoginCredentials: lc,
	}

	response := &sales.SyncOnlineSalesResponse{}
	do(credPld, response, "sales/syncOnlineSales")
	return response
}

// postSyncShippedSaleAndRemoveItems payload sent to Sku Vault.
type postSyncShippedSaleAndRemoveItems struct {
	*sales.SyncShippedSaleAndRemoveItems
	*SLoginCredentials
}

// SyncShippedSaleAndRemoveItems creates http request for this SKU vault endpoint.
// Moderate Throttle.
// This method syncs a shipped sale and auto-removes quantity. If the product is in multiple locations, then the remove will occur at the location that is first in alphanumerical order. Quantity can be removed from normal, Reserved, and Backstock locations..
func (lc *SLoginCredentials) SyncShippedSaleAndRemoveItems(pld *sales.SyncShippedSaleAndRemoveItems) *sales.SyncShippedSaleAndRemoveItemsResponse {
	credPld := &postSyncShippedSaleAndRemoveItems{
		SyncShippedSaleAndRemoveItems: pld,
		SLoginCredentials:             lc,
	}

	response := &sales.SyncShippedSaleAndRemoveItemsResponse{}
	do(credPld, response, "sales/syncShippedSaleAndRemoveItems")
	return response
}

// postSyncShippedSaleAndRemoveItemsBulk payload sent to Sku Vault.
type postSyncShippedSaleAndRemoveItemsBulk struct {
	*sales.SyncShippedSaleAndRemoveItemsBulk
	*SLoginCredentials
}

// SyncShippedSaleAndRemoveItemsBulk creates http request for this SKU vault endpoint.
// Moderate Throttle.
// This method can sync multiple shipped sales and auto-removes quantity. Up to 100 sales per call. .
func (lc *SLoginCredentials) SyncShippedSaleAndRemoveItemsBulk(pld *sales.SyncShippedSaleAndRemoveItemsBulk) *sales.SyncShippedSaleAndRemoveItemsBulkResponse {
	credPld := &postSyncShippedSaleAndRemoveItemsBulk{
		SyncShippedSaleAndRemoveItemsBulk: pld,
		SLoginCredentials:                 lc,
	}

	response := &sales.SyncShippedSaleAndRemoveItemsBulkResponse{}
	do(credPld, response, "sales/syncShippedSaleAndRemoveItemsBulk")
	return response
}

// postUpdateOnlineSaleStatus payload sent to Sku Vault.
type postUpdateOnlineSaleStatus struct {
	*sales.UpdateOnlineSaleStatus
	*SLoginCredentials
}

// UpdateOnlineSaleStatus creates http request for this SKU vault endpoint.
// Light Throttle.
// Update the status of a sale..
func (lc *SLoginCredentials) UpdateOnlineSaleStatus(pld *sales.UpdateOnlineSaleStatus) *sales.UpdateOnlineSaleStatusResponse {
	credPld := &postUpdateOnlineSaleStatus{
		UpdateOnlineSaleStatus: pld,
		SLoginCredentials:      lc,
	}

	response := &sales.UpdateOnlineSaleStatusResponse{}
	do(credPld, response, "sales/updateOnlineSaleStatus")
	return response
}

// postUpdateShipments payload sent to Sku Vault.
type postUpdateShipments struct {
	*sales.UpdateShipments
	*SLoginCredentials
}

// UpdateShipments creates http request for this SKU vault endpoint.
// Severe Throttle.
// Using this call, users can update shipments to a sale..
func (lc *SLoginCredentials) UpdateShipments(pld *sales.UpdateShipments) *sales.UpdateShipmentsResponse {
	credPld := &postUpdateShipments{
		UpdateShipments:   pld,
		SLoginCredentials: lc,
	}

	response := &sales.UpdateShipmentsResponse{}
	do(credPld, response, "sales/updateShipments")
	return response
}
