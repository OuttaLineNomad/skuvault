package sales

// GetSalesByDate is a automatically generated struct from json provided by sku vault's api docs.
type GetSalesByDate struct {
	DateField  string `json:"DateField"`
	FromDate   string `json:"FromDate"`
	PageNumber int64  `json:"PageNumber"`
	PageSize   int64  `json:"PageSize"`
	ToDate     string `json:"ToDate"`
}

type GetSalesByDateResponse []struct {
	Channel     string `json:"Channel"`
	ChannelID   string `json:"ChannelId"`
	Client      string `json:"Client"`
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
		PartNumber string `json:"PartNumber"`
		Quantity   int64  `json:"Quantity"`
		Sku        string `json:"Sku"`
		UnitPrice  struct {
			A int64  `json:"a"`
			S string `json:"s"`
		} `json:"UnitPrice"`
	} `json:"MerchantItems"`
	MerchantKits   []interface{} `json:"MerchantKits"`
	Notes          string        `json:"Notes"`
	PrintedStatus  bool          `json:"PrintedStatus"`
	ProcessedItems []struct {
		FailedQuantity int64  `json:"FailedQuantity"`
		PassedQuantity int64  `json:"PassedQuantity"`
		PickedQuantity int64  `json:"PickedQuantity"`
		Sku            string `json:"Sku"`
	} `json:"ProcessedItems"`
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
}
