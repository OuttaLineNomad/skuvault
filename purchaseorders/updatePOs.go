package purchaseorders

// UpdatePOs is a automatically generated struct from json provided by sku vault's api docs.
type UpdatePOs struct {
	POs []struct {
		ActualShippedDate string `json:"ActualShippedDate"`
		ArrivalDueDate    string `json:"ArrivalDueDate"`
		LineItems         []struct {
			Cost          int64  `json:"Cost"`
			Identifier    string `json:"Identifier"`
			PrivateNotes  string `json:"PrivateNotes"`
			PublicNotes   string `json:"PublicNotes"`
			Quantity      int64  `json:"Quantity"`
			QuantityTo3PL int64  `json:"QuantityTo3PL"`
			Sku           string `json:"SKU"`
			Variant       string `json:"Variant"`
		} `json:"LineItems"`
		OrderCancelDate string `json:"OrderCancelDate"`
		OrderDate       string `json:"OrderDate"`
		PaymentStatus   string `json:"PaymentStatus"`
		Payments        []struct {
			Amount      int64  `json:"Amount"`
			Note        string `json:"Note"`
			PaymentName string `json:"PaymentName"`
		} `json:"Payments"`
		PoNumber             string `json:"PoNumber"`
		PrivateNotes         string `json:"PrivateNotes"`
		PublicNotes          string `json:"PublicNotes"`
		PurchaseOrderID      string `json:"PurchaseOrderId"`
		RequestedShipDate    string `json:"RequestedShipDate"`
		SentStatus           string `json:"SentStatus"`
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
	} `json:"POs"`
}

// UpdatePOsResponse is a automatically generated struct from json provided by sku vault's api docs.
type UpdatePOsResponse struct {
	Errors []interface{} `json:"Errors"`
	Status string        `json:"Status"`
}
