package purchaseorders

// GetPOs is a automatically generated struct from json provided by sku vault's api docs.
type GetPOs struct {
	IncludeProducts           bool     `json:"IncludeProducts"`
	ModifiedAfterDateTimeUtc  string   `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc string   `json:"ModifiedBeforeDateTimeUtc"`
	PONumbers                 []string `json:"PONumbers"`
	PageNumber                int64    `json:"PageNumber"`
	Status                    string   `json:"Status"`
}

// GetPOsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetPOsResponse struct {
	PurchaseOrders []struct {
		ActualShippedDate string        `json:"ActualShippedDate"`
		ArrivalDueDate    string        `json:"ArrivalDueDate"`
		Costs             []interface{} `json:"Costs"`
		CreatedDate       string        `json:"CreatedDate"`
		LineItems         []struct {
			Cost                  int64  `json:"Cost"`
			Identifier            string `json:"Identifier"`
			PrivateNotes          string `json:"PrivateNotes"`
			ProductID             string `json:"ProductId"`
			PublicNotes           string `json:"PublicNotes"`
			Quantity              int64  `json:"Quantity"`
			QuantityTo3PL         int64  `json:"QuantityTo3PL"`
			ReceivedQuantity      int64  `json:"ReceivedQuantity"`
			ReceivedQuantityTo3PL int64  `json:"ReceivedQuantityTo3PL"`
			RetailCost            int64  `json:"RetailCost"`
			Sku                   string `json:"SKU"`
			Variant               string `json:"Variant"`
		} `json:"LineItems"`
		OrderCancelDate      string `json:"OrderCancelDate"`
		OrderDate            string `json:"OrderDate"`
		PaymentStatus        string `json:"PaymentStatus"`
		PoID                 string `json:"PoId"`
		PoNumber             string `json:"PoNumber"`
		PrivateNotes         string `json:"PrivateNotes"`
		PublicNotes          string `json:"PublicNotes"`
		RequestedShipDate    string `json:"RequestedShipDate"`
		ShipToAddress        string `json:"ShipToAddress"`
		ShipToWarehouse      string `json:"ShipToWarehouse"`
		ShippingCarrierClass struct {
			CarrierName string `json:"CarrierName"`
			ClassName   string `json:"ClassName"`
		} `json:"ShippingCarrierClass"`
		Status       string `json:"Status"`
		SupplierName string `json:"SupplierName"`
		TermsName    string `json:"TermsName"`
		TrackingInfo string `json:"TrackingInfo"`
	} `json:"PurchaseOrders"`
}
