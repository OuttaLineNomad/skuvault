package sales

// UpdateShipments is a automatically generated struct from json provided by sku vault's api docs.
type UpdateShipments struct {
	Shipments []struct {
		Carrier     string `json:"Carrier"`
		LandedCosts []struct {
			Amount          int64  `json:"Amount"`
			CurrencyIsoCode string `json:"CurrencyIsoCode"`
			Sku             string `json:"Sku"`
		} `json:"LandedCosts"`
		SaleID         string `json:"SaleId"`
		Status         string `json:"Status"`
		TrackingNumber string `json:"TrackingNumber"`
	} `json:"Shipments"`
}

// UpdateShipmentsResponse is a automatically generated struct from json provided by sku vault's api docs.
type UpdateShipmentsResponse struct {
	Errors []interface{} `json:"Errors"`
	Status string        `json:"Status"`
}
