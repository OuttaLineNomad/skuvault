package purchaseorders

import "time"

// ReceivePOItems is a automatically generated struct from json provided by sku vault's api docs.
type ReceivePOItems struct {
	LineItems []struct {
		Location      string `json:"Location"`
		Quantity      int    `json:"Quantity"`
		QuantityTo3PL int    `json:"QuantityTo3PL"`
		Sku           string `json:"SKU"`
	} `json:"LineItems"`
	PoNumber     string    `json:"PoNumber"`
	ReceiptDate  time.Time `json:"ReceiptDate"`
	SupplierName string    `json:"SupplierName"`
	WarehouseID  int       `json:"WarehouseId"`
}

// ReceivePOItemsResponse is a automatically generated struct from json provided by sku vault's api docs.
type ReceivePOItemsResponse struct {
	Errors []string `json:"Errors"`
	Status string   `json:"Status"`
}
