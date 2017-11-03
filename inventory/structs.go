package inventory

import "time"

// GetWarehouseItemQuantity global parts user needs to use for api call
type GetWarehouseItemQuantity struct {
	Sku         string
	WarehouseID int
}

// GetWarehouseItemQuantityResponse the response from SKU Vault endpoint
type GetWarehouseItemQuantityResponse struct {
	Errors              []interface{}
	TotalQuantityOnHand int
}

// GetTransactions global parts user needs to use for api call
type GetTransactions struct {
	FromDate                  string
	ToDate                    string
	WarehouseID               int
	TransactionType           string
	TransactionReasons        []string
	ExcludeTransactionReasons []string
	PageNumber                int
	PageSize                  int
}

// Transaction data from GetTrasaction for one transaction.
type Transaction struct {
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
	Transactions []Transaction
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
