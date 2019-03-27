package inventory

import "time"

// GetTransactions is a automatically generated struct from json provided by sku vault's api docs.
type GetTransactions struct {
	ExcludeTransactionReasons []string  `json:"ExcludeTransactionReasons"`
	FromDate                  time.Time `json:"FromDate"`
	PageNumber                int       `json:"PageNumber"`
	PageSize                  int       `json:"PageSize"`
	SaleID                    string    `json:"SaleId"`
	ToDate                    time.Time `json:"ToDate"`
	TransactionReasons        []string  `json:"TransactionReasons"`
	WarehouseID               string    `json:"WarehouseId"`
}

// GetTransactionsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetTransactionsResponse struct {
	Transactions []struct {
		Code    string `json:"Code"`
		Context struct {
			ID   string `json:"ID"`
			Type string `json:"Type"`
		} `json:"Context"`
		Location          string `json:"Location"`
		Quantity          int    `json:"Quantity"`
		QuantityAfter     int    `json:"QuantityAfter"`
		QuantityBefore    int    `json:"QuantityBefore"`
		ScannedCode       string `json:"ScannedCode"`
		Sku               string `json:"Sku"`
		Title             string `json:"Title"`
		TransactionDate   time.Time `json:"TransactionDate"`
		TransactionNote   string `json:"TransactionNote"`
		TransactionReason string `json:"TransactionReason"`
		User              string `json:"User"`
	} `json:"Transactions"`
}
