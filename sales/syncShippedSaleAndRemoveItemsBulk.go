package sales

// SyncShippedSaleAndRemoveItemsBulk is a automatically generated struct from json provided by sku vault's api docs.
type SyncShippedSaleAndRemoveItemsBulk struct {
	Sales []struct {
		FulfilledItems []struct {
			Quantity  int     `json:"Quantity"`
			Sku       string  `json:"Sku"`
			UnitPrice float64 `json:"UnitPrice"`
		} `json:"FulfilledItems"`
		ItemSkus []struct {
			Quantity  int     `json:"Quantity"`
			Sku       string  `json:"Sku"`
			UnitPrice float64 `json:"UnitPrice"`
		} `json:"ItemSkus"`
		Notes        string  `json:"Notes"`
		OrderDateUtc time.Time  `json:"OrderDateUtc"`
		OrderID      string  `json:"OrderId"`
		OrderTotal   float64 `json:"OrderTotal"`
		ShippingInfo struct {
			City            string `json:"City"`
			CompanyName     string `json:"CompanyName"`
			Country         string `json:"Country"`
			Email           string `json:"Email"`
			FirstName       string `json:"FirstName"`
			LastName        string `json:"LastName"`
			Line1           string `json:"Line1"`
			Line2           string `json:"Line2"`
			PhoneNumber     string `json:"PhoneNumber"`
			Postal          string `json:"Postal"`
			Region          string `json:"Region"`
			ShippingCarrier string `json:"ShippingCarrier"`
			ShippingClass   string `json:"ShippingClass"`
		} `json:"ShippingInfo"`
		WarehouseID int `json:"WarehouseId"`
	} `json:"Sales"`
}

// SyncShippedSaleAndRemoveItemsBulkResponse is a automatically generated struct from json provided by sku vault's api docs.
type SyncShippedSaleAndRemoveItemsBulkResponse struct {
	Errors []interface{} `json:"Errors"`
	Sales  []interface{} `json:"Sales"`
	Status string        `json:"Status"`
}
