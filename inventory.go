package skuvault

// holds all api call endpoints with that are in the folder /Inventory

import (
	"time"
)

// GetWarehouseItemQuantity global parts user needs to use for api call
type GetWarehouseItemQuantity struct {
	Sku         string
	WarehouseID int
}

// postGetWarehouseItemQuantity payload sent to Sku Vault
type postGetWarehouseItemQuantity struct {
	*GetWarehouseItemQuantity
	TenantToken string
	UserToken   string
}

// GetWarehouseItemQuantityResponse the response from SKU Vault endpoint
type GetWarehouseItemQuantityResponse struct {
	Errors              []interface{}
	TotalQuantityOnHand int
}

// GetTransactions global parts user needs to use for api call
type GetTransactions struct {
	FromDate                  time.Time
	ToDate                    time.Time
	WarehouseID               int
	TransactionType           string
	TransactionReasons        []string
	ExcludeTransactionReasons []string
	PageNumber                int
	PageSize                  int
}

// postGetTransactions payload sent to Sku Vault
type postGetTransactions struct {
	*postGetConvert
	TenantToken string
	UserToken   string
}

type transaction struct {
	User              string
	Sku               string
	Code              string
	ScannedCode       string
	Title             string
	Quantity          int
	QuantityBefore    int
	QuantityAfter     int
	Location          string
	TransactionType   string
	TransactionReason string
	TransactionNote   string
	TransactionDate   time.Time
}

// GetTransactionsResponse the response from SKU Vault endpoint
type GetTransactionsResponse struct {
	Transactions []transaction
	Status       string
	Errors       []interface{}
}

// GetInventoryByLocation global parts user needs to use for api call
type GetInventoryByLocation struct {
	IsReturnByCodes bool
	PageNumber      int
	PageSize        int
	ProductCodes    []string
	ProductSKUs     []string
}

// postGetInventoryByLocation payload sent to Sku Vault
type postGetInventoryByLocation struct {
	*GetInventoryByLocation
	TenantToken string
	Usertoken   string
}

// SkuLocations list of locations of one SKU.
type SkuLocations struct {
	WarehouseCode string
	LocationCode  string
	Quantity      int
	Reserve       bool
}

// GetInventoryByLocationResponse the response from SKU Vault endpoint
type GetInventoryByLocationResponse struct {
	Items  map[string][]SkuLocations
	Errors []interface{}
}

// GetWarehouseItemQuantity creates http request for this SKU vault endpoint
// Heavy throttling
// Returns the quantity for a specified SKU
func (lc *ILoginCredentials) GetWarehouseItemQuantity(pld *GetWarehouseItemQuantity) *GetWarehouseItemQuantityResponse {
	credPld := &postGetWarehouseItemQuantity{
		GetWarehouseItemQuantity: pld,
		TenantToken:              lc.tenantToken,
		UserToken:                lc.userToken,
	}

	response := &GetWarehouseItemQuantityResponse{}
	inventoryCall(credPld, response, "getWarehouseItemQuantity")
	return response
}

type postGetConvert struct {
	FromDate                  string
	ToDate                    string
	WarehouseID               int
	TransactionType           string
	TransactionReasons        []string
	ExcludeTransactionReasons []string
	PageNumber                int
	PageSize                  int
}

// GetTransactions creates http request for this SKU vault endpoint
// Heavy throttling
// Look at your transaction history.
func (lc *ILoginCredentials) GetTransactions(pld *GetTransactions) *GetTransactionsResponse {
	convPld := &postGetConvert{
		FromDate:                  pld.FromDate.Format("2006-01-02T15:04:05.9999999Z"),
		ToDate:                    pld.ToDate.Format("2006-01-02T15:04:05.9999999Z"),
		WarehouseID:               pld.WarehouseID,
		TransactionType:           pld.TransactionType,
		TransactionReasons:        pld.TransactionReasons,
		ExcludeTransactionReasons: pld.ExcludeTransactionReasons,
		PageNumber:                pld.PageNumber,
		PageSize:                  pld.PageSize,
	}
	credPld := &postGetTransactions{
		postGetConvert: convPld,
		TenantToken:          lc.tenantToken,
		UserToken:            lc.userToken,
	}

	response := &GetTransactionsResponse{}
	inventoryCall(credPld, response, "getTransactions")
	return response

}

// GetInventoryByLocation creates http request for this SKU vault endpoint
// Heavy throttling
// Returns location and warehouse per product.
func (lc *ILoginCredentials) GetInventoryByLocation(pld *GetInventoryByLocation) *GetInventoryByLocationResponse {
	credPld := &postGetInventoryByLocation{
		GetInventoryByLocation: pld,
		TenantToken:            lc.tenantToken,
		Usertoken:              lc.userToken,
	}
	response := &GetInventoryByLocationResponse{}
	inventoryCall(credPld, response, "getInventoryByLocation")
	return response

}

// inventoryCalladds inventory/ to url for do() call.
func inventoryCall(pld interface{}, response interface{}, endPoint string) {
	full := "inventory/" + endPoint
	do(pld, response, full)
}
