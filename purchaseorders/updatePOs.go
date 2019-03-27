package purchaseorders

import "time"

// UpdatePOs is a automatically generated struct from json provided by sku vault's api docs.
type UpdatePOs struct {
	POs []struct {
		ActualShippedDate time.Time `json:"ActualShippedDate"`
		ArrivalDueDate    time.Time `json:"ArrivalDueDate"`
		LineItems         []struct {
			Cost          float64 `json:"Cost"`
			Identifier    string  `json:"Identifier"`
			PrivateNotes  string  `json:"PrivateNotes"`
			PublicNotes   string  `json:"PublicNotes"`
			Quantity      int     `json:"Quantity"`
			QuantityTo3PL int     `json:"QuantityTo3PL"`
			Sku           string  `json:"SKU"`
			Variant       string  `json:"Variant"`
		} `json:"LineItems"`
		OrderCancelDate time.Time `json:"OrderCancelDate"`
		OrderDate       time.Time `json:"OrderDate"`
		PaymentStatus   string    `json:"PaymentStatus"`
		Payments        []struct {
			Amount      float64 `json:"Amount"`
			Note        string  `json:"Note"`
			PaymentName string  `json:"PaymentName"`
		} `json:"Payments"`
		PoNumber             string    `json:"PoNumber"`
		PrivateNotes         string    `json:"PrivateNotes"`
		PublicNotes          string    `json:"PublicNotes"`
		PurchaseOrderID      string    `json:"PurchaseOrderId"`
		RequestedShipDate    time.Time `json:"RequestedShipDate"`
		SentStatus           string    `json:"SentStatus"`
		ShipToAddress        string    `json:"ShipToAddress"`
		ShipToWarehouse      string    `json:"ShipToWarehouse"`
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
