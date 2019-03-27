package sales

import "time"

// GetShipments is a automatically generated struct from json provided by sku vault's api docs.
type GetShipments struct {
	SaleIds []string `json:"SaleIds"`
}

// GetShipmentsResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetShipmentsResponse struct {
	Errors    []interface{} `json:"Errors"`
	Shipments []struct {
		AlternateID string `json:"AlternateId"`
		Carrier     string `json:"Carrier"`
		Class       string `json:"Class"`
		Costs       []struct {
			Cost struct {
				A float64 `json:"a"`
				S string  `json:"s"`
			} `json:"Cost"`
			CostType string `json:"CostType"`
		} `json:"Costs"`
		CreatedDate           time.Time  `json:"CreatedDate"`
		EstimatedDeliveryDate time.Time  `json:"EstimatedDeliveryDate"`
		EstimatedShipDate     time.Time  `json:"EstimatedShipDate"`
		FileIds               []struct{} `json:"FileIds"`
		LandedCosts           []struct {
			Cost struct {
				A float64 `json:"a"`
				S string  `json:"s"`
			} `json:"Cost"`
			Sku string `json:"Sku"`
		} `json:"LandedCosts"`
		ManifestID string `json:"ManifestId"`
		Note       string `json:"Note"`
		Parcels    []struct {
			Dimensions struct {
				Height int     `json:"Height"`
				Length int     `json:"Length"`
				Unit   string  `json:"Unit"`
				Width  float64 `json:"Width"`
			} `json:"Dimensions"`
			Items []struct {
				ID       string `json:"Id"`
				Quantity int    `json:"Quantity"`
				Sku      string `json:"Sku"`
			} `json:"Items"`
			Kits []struct {
				ID    string `json:"Id"`
				Items []struct {
					ID       string `json:"Id"`
					Quantity int    `json:"Quantity"`
				} `json:"Items"`
				Quantity int    `json:"Quantity"`
				Sku      string `json:"Sku"`
			} `json:"Kits"`
			Note       string  `json:"Note"`
			Number     int     `json:"Number"`
			Weight     float64 `json:"Weight"`
			WeightUnit string  `json:"WeightUnit"`
		} `json:"Parcels"`
		SaleID      string `json:"SaleId"`
		ShippedFrom struct {
			Address struct {
				Address1   string `json:"Address1"`
				Address2   string `json:"Address2"`
				City       string `json:"City"`
				Company    string `json:"Company"`
				Country    string `json:"Country"`
				Email      string `json:"Email"`
				FirstName  string `json:"FirstName"`
				LastName   string `json:"LastName"`
				MiddleName string `json:"MiddleName"`
				PostalCode string `json:"PostalCode"`
				Region     string `json:"Region"`
			} `json:"Address"`
			ThreePL       bool   `json:"ThreePL"`
			WarehouseCode string `json:"WarehouseCode"`
		} `json:"ShippedFrom"`
		Source         string  `json:"Source"`
		Status         string  `json:"Status"`
		TotalWeight    float64 `json:"TotalWeight"`
		TrackingNumber string  `json:"TrackingNumber"`
		TrackingURL    string  `json:"TrackingUrl"`
		Type           string  `json:"Type"`
		WeightUnit     string  `json:"WeightUnit"`
	} `json:"Shipments"`
	Status string `json:"Status"`
}
