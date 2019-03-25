package sales

// GetSales is a automatically generated struct from json provided by sku vault's api docs.
type GetSales struct {
	OrderIds []string `json:"OrderIds"`
	Status   string   `json:"Status"`
}

// GetSalesResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetSalesResponse struct {
	Sales []struct {
		ChannelID   string `json:"ChannelId"`
		ContactInfo struct {
			Company   string `json:"Company"`
			Email     string `json:"Email"`
			FirstName string `json:"FirstName"`
			LastName  string `json:"LastName"`
			Phone     string `json:"Phone"`
		} `json:"ContactInfo"`
		FulfilledItems  []interface{} `json:"FulfilledItems"`
		FulfilledKits   []interface{} `json:"FulfilledKits"`
		ID              string        `json:"Id"`
		LastPrintedDate string        `json:"LastPrintedDate"`
		Marketplace     string        `json:"Marketplace"`
		MarketplaceID   string        `json:"MarketplaceId"`
		MerchantItems   []struct {
			Quantity  int64  `json:"Quantity"`
			Sku       string `json:"Sku"`
			UnitPrice struct {
				A int64  `json:"a"`
				S string `json:"s"`
			} `json:"UnitPrice"`
		} `json:"MerchantItems"`
		MerchantKits []struct {
			KitItems  struct{} `json:"KitItems"`
			Quantity  int64    `json:"Quantity"`
			Sku       string   `json:"Sku"`
			UnitPrice struct {
				A int64  `json:"a"`
				S string `json:"s"`
			} `json:"UnitPrice"`
		} `json:"MerchantKits"`
		Notes           string `json:"Notes"`
		PrintedStatus   bool   `json:"PrintedStatus"`
		SaleDate        string `json:"SaleDate"`
		ShippingCarrier string `json:"ShippingCarrier"`
		ShippingCharge  struct {
			A int64  `json:"a"`
			S string `json:"s"`
		} `json:"ShippingCharge"`
		ShippingClass string `json:"ShippingClass"`
		ShippingCost  struct {
			A int64  `json:"a"`
			S string `json:"s"`
		} `json:"ShippingCost"`
		ShippingInfo struct {
			Address1   string `json:"Address1"`
			Address2   string `json:"Address2"`
			City       string `json:"City"`
			Country    string `json:"Country"`
			PostalCode string `json:"PostalCode"`
			Region     string `json:"Region"`
		} `json:"ShippingInfo"`
		Status string `json:"Status"`
	} `json:"Sales"`
}
