package purchaseorders

import "time"

// GetPOs is a automatically generated struct from json provided by sku vault's api docs.
type GetPOs struct {
	IncludeProducts           bool      `json:"IncludeProducts"`
	ModifiedAfterDateTimeUtc  time.Time `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc time.Time `json:"ModifiedBeforeDateTimeUtc"`
	PONumbers                 []string  `json:"PONumbers"`
	PageNumber                int       `json:"PageNumber"`
	Status                    string    `json:"Status"`
}

// GetPOsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetPOsResponse struct {
	PurchaseOrders []struct {
		ActualShippedDate time.Time `json:"ActualShippedDate"`
		ArrivalDueDate    time.Time `json:"ArrivalDueDate"`
		Costs             []float64 `json:"Costs"`
		CreatedDate       time.Time `json:"CreatedDate"`
		LineItems         []struct {
			Cost                  float64 `json:"Cost"`
			Identifier            string  `json:"Identifier"`
			PrivateNotes          string  `json:"PrivateNotes"`
			ProductID             string  `json:"ProductId"`
			PublicNotes           string  `json:"PublicNotes"`
			Quantity              int     `json:"Quantity"`
			QuantityTo3PL         int     `json:"QuantityTo3PL"`
			ReceivedQuantity      int     `json:"ReceivedQuantity"`
			ReceivedQuantityTo3PL int     `json:"ReceivedQuantityTo3PL"`
			RetailCost            float64 `json:"RetailCost"`
			Sku                   string  `json:"SKU"`
			Variant               string  `json:"Variant"`
		} `json:"LineItems"`
		OrderCancelDate      time.Time `json:"OrderCancelDate"`
		OrderDate            time.Time `json:"OrderDate"`
		PaymentStatus        string    `json:"PaymentStatus"`
		PoID                 string    `json:"PoId"`
		PoNumber             string    `json:"PoNumber"`
		PrivateNotes         string    `json:"PrivateNotes"`
		PublicNotes          string    `json:"PublicNotes"`
		RequestedShipDate    time.Time `json:"RequestedShipDate"`
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
	} `json:"PurchaseOrders"`
}
