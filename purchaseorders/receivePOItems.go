package purchaseorders

// ReceivePOItems is a automatically generated struct from json provided by sku vault's api docs.
type ReceivePOItems struct {
	LineItems []struct {
		Location      string `json:"Location"`
		Quantity      int64  `json:"Quantity"`
		QuantityTo3PL int64  `json:"QuantityTo3PL"`
		Sku           string `json:"SKU"`
	} `json:"LineItems"`
	PoNumber     string `json:"PoNumber"`
	ReceiptDate  string `json:"ReceiptDate"`
	SupplierName string `json:"SupplierName"`
	WarehouseID  int64  `json:"WarehouseId"`
}

// ReceivePOItemsResponse is a automatically generated struct from json provided by sku vault's api docs.
type ReceivePOItemsResponse struct {
	Errors []string `json:"Errors"`
	Status string   `json:"Status"`
}
