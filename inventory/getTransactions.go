package inventory

// GetTransactions is a automatically generated struct from json provided by sku vault's api docs.
type GetTransactions struct {
	ExcludeTransactionReasons []string `json:"ExcludeTransactionReasons"`
	FromDate                  string   `json:"FromDate"`
	PageNumber                int64    `json:"PageNumber"`
	PageSize                  int64    `json:"PageSize"`
	SaleID                    string   `json:"SaleId"`
	ToDate                    string   `json:"ToDate"`
	TransactionReasons        []string `json:"TransactionReasons"`
	WarehouseID               string   `json:"WarehouseId"`
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
		Quantity          int64  `json:"Quantity"`
		QuantityAfter     int64  `json:"QuantityAfter"`
		QuantityBefore    int64  `json:"QuantityBefore"`
		ScannedCode       string `json:"ScannedCode"`
		Sku               string `json:"Sku"`
		Title             string `json:"Title"`
		TransactionDate   string `json:"TransactionDate"`
		TransactionNote   string `json:"TransactionNote"`
		TransactionReason string `json:"TransactionReason"`
		User              string `json:"User"`
	} `json:"Transactions"`
}
