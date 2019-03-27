package purchaseorders

import "time"

// GetReceivesHistory is a automatically generated struct from json provided by sku vault's api docs.
type GetReceivesHistory struct {
	PageNumber                int       `json:"PageNumber"`
	PageSize                  int       `json:"PageSize"`
	PoNumberFilter            []string  `json:"PoNumberFilter"`
	ReceivedAfterDateTimeUtc  time.Time `json:"ReceivedAfterDateTimeUTC"`
	ReceivedBeforeDateTimeUtc time.Time `json:"ReceivedBeforeDateTimeUTC"`
	WarehouseFilter           []string  `json:"WarehouseFilter"`
}

// GetReceivesHistoryResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetReceivesHistoryResponse struct {
	Corrections []struct {
		Code                  string    `json:"Code"`
		CorrectedDate         time.Time `json:"CorrectedDate"`
		NewQuantity           int       `json:"NewQuantity"`
		NewQuantity3pl        int       `json:"NewQuantity3pl"`
		OldQuantity           int       `json:"OldQuantity"`
		OldQuantity3pl        int       `json:"OldQuantity3pl"`
		PONumber              string    `json:"PONumber"`
		PartNumber            string    `json:"PartNumber"`
		QuantityByReceiptDate []struct {
			NewQuantity int       `json:"NewQuantity"`
			OldQuantity int       `json:"OldQuantity"`
			ReceiptDate time.Time `json:"ReceiptDate"`
		} `json:"QuantityByReceiptDate"`
		ReceivedDate time.Time `json:"ReceivedDate"`
		Sku          string    `json:"SKU"`
		Username     string    `json:"Username"`
	} `json:"Corrections"`
	Receives []struct {
		Code               string    `json:"Code"`
		Location           string    `json:"Location"`
		PONumber           string    `json:"PONumber"`
		PartNumber         string    `json:"PartNumber"`
		Quantity           int       `json:"Quantity"`
		Quantity3pl        int       `json:"Quantity3pl"`
		QuantityToLocation int       `json:"QuantityToLocation"`
		ReceiptDate        time.Time `json:"ReceiptDate"`
		ReceivedDate       time.Time `json:"ReceivedDate"`
		Sku                string    `json:"SKU"`
		Username           string    `json:"Username"`
		Warehouse          string    `json:"Warehouse"`
	} `json:"Receives"`
}
