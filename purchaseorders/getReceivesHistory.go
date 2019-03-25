package purchaseorders

// GetReceivesHistory is a automatically generated struct from json provided by sku vault's api docs.
type GetReceivesHistory struct {
	PageNumber                int64    `json:"PageNumber"`
	PageSize                  int64    `json:"PageSize"`
	PoNumberFilter            []string `json:"PoNumberFilter"`
	ReceivedAfterDateTimeUTC  string   `json:"ReceivedAfterDateTimeUTC"`
	ReceivedBeforeDateTimeUTC string   `json:"ReceivedBeforeDateTimeUTC"`
	WarehouseFilter           []string `json:"WarehouseFilter"`
}

// GetReceivesHistoryResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetReceivesHistoryResponse struct {
	Corrections []struct {
		Code                  string `json:"Code"`
		CorrectedDate         string `json:"CorrectedDate"`
		NewQuantity           int64  `json:"NewQuantity"`
		NewQuantity3pl        int64  `json:"NewQuantity3pl"`
		OldQuantity           int64  `json:"OldQuantity"`
		OldQuantity3pl        int64  `json:"OldQuantity3pl"`
		PONumber              string `json:"PONumber"`
		PartNumber            string `json:"PartNumber"`
		QuantityByReceiptDate []struct {
			NewQuantity int64  `json:"NewQuantity"`
			OldQuantity int64  `json:"OldQuantity"`
			ReceiptDate string `json:"ReceiptDate"`
		} `json:"QuantityByReceiptDate"`
		ReceivedDate string `json:"ReceivedDate"`
		Sku          string `json:"SKU"`
		Username     string `json:"Username"`
	} `json:"Corrections"`
	Receives []struct {
		Code               string `json:"Code"`
		Location           string `json:"Location"`
		PONumber           string `json:"PONumber"`
		PartNumber         string `json:"PartNumber"`
		Quantity           int64  `json:"Quantity"`
		Quantity3pl        int64  `json:"Quantity3pl"`
		QuantityToLocation int64  `json:"QuantityToLocation"`
		ReceiptDate        string `json:"ReceiptDate"`
		ReceivedDate       string `json:"ReceivedDate"`
		Sku                string `json:"SKU"`
		Username           string `json:"Username"`
		Warehouse          string `json:"Warehouse"`
	} `json:"Receives"`
}
