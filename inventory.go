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
	WarehouseID               string
	TransactionType           string
	TransactionReasons        []string
	ExcludeTransactionReasons []string
	PageNumber                int
	PageSize                  int
}

// postGetTransactions payload sent to Sku Vault
type postGetTransactions struct {
	*GetTransactions
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

// GetWarehouseItemQuantity creates http request for this SKU vault endpoint
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

// GetTransactions creates http request for this SKU vault endpoint
// Look at your transaction history.
func (lc *ILoginCredentials) GetTransactions(pld *GetTransactions) *GetWarehouseItemQuantityResponse {
	credPld := &postGetTransactions{
		GetTransactions: pld,
		TenantToken:     lc.tenantToken,
		UserToken:       lc.userToken,
	}
	response := &GetWarehouseItemQuantityResponse{}
	inventoryCall(credPld, response, "getTransactions")
	return response

}

func inventoryCall(pld interface{}, response interface{}, endPoint string) {
	full := "inventory/" + endPoint
	do(pld, response, full)
}
