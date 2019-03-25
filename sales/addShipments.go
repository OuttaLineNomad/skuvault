package sales

// AddShipments is a automatically generated struct from json provided by sku vault's api docs.
type AddShipments struct {
	Shipments []struct {
		Carrier string `json:"Carrier"`
		Class   string `json:"Class"`
		Costs   []struct {
			Amount          int64  `json:"Amount"`
			CostType        string `json:"CostType"`
			CurrencyIsoCode string `json:"CurrencyIsoCode"`
		} `json:"Costs"`
		LandedCosts []struct {
			Amount          int64  `json:"Amount"`
			CurrencyIsoCode string `json:"CurrencyIsoCode"`
			Sku             string `json:"Sku"`
		} `json:"LandedCosts"`
		Parcels []struct {
			Items []struct {
				Quantity int64  `json:"Quantity"`
				Sku      string `json:"Sku"`
			} `json:"Items"`
			Note       string `json:"Note"`
			Number     int64  `json:"Number"`
			Weight     int64  `json:"Weight"`
			WeightUnit string `json:"WeightUnit"`
		} `json:"Parcels"`
		SaleID      string `json:"SaleId"`
		ShippedFrom struct {
			Address struct {
				Address1   string `json:"Address1"`
				City       string `json:"City"`
				Company    string `json:"Company"`
				Country    string `json:"Country"`
				Email      string `json:"Email"`
				FirstName  string `json:"FirstName"`
				LastName   string `json:"LastName"`
				PostalCode string `json:"PostalCode"`
				Region     string `json:"Region"`
			} `json:"Address"`
			ThreePL     bool   `json:"ThreePL"`
			WarehouseID string `json:"WarehouseId"`
		} `json:"ShippedFrom"`
		Source         string `json:"Source"`
		Status         string `json:"Status"`
		TrackingNumber string `json:"TrackingNumber"`
		TrackingURL    string `json:"TrackingUrl"`
		Type           string `json:"Type"`
	} `json:"Shipments"`
}

// AddShipmentsResponse is a automatically generated struct from json provided by sku vault's api docs.
type AddShipmentsResponse struct {
	Errors []interface{} `json:"Errors"`
	Status string        `json:"Status"`
}
