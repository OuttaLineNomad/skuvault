package sales

// SyncOnlineSales is a automatically generated struct from json provided by sku vault's api docs.
type SyncOnlineSales struct {
	Sales []struct {
		ChannelAccountID string `json:"ChannelAccountId"`
		ChannelType      string `json:"ChannelType"`
		CheckoutStatus   string `json:"CheckoutStatus"`
		FulfilledItems   []struct {
			Quantity  int64  `json:"Quantity"`
			Sku       string `json:"Sku"`
			UnitPrice int64  `json:"UnitPrice"`
		} `json:"FulfilledItems"`
		ItemSkus []struct {
			Quantity  int64  `json:"Quantity"`
			Sku       string `json:"Sku"`
			UnitPrice int64  `json:"UnitPrice"`
		} `json:"ItemSkus"`
		MarketplaceID string `json:"MarketplaceId"`
		Notes         string `json:"Notes"`
		OrderDateUtc  string `json:"OrderDateUtc"`
		OrderID       string `json:"OrderId"`
		OrderTotal    int64  `json:"OrderTotal"`
		PaymentStatus string `json:"PaymentStatus"`
		SaleState     string `json:"SaleState"`
		ShippingInfo  struct {
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
			ShippingStatus  string `json:"ShippingStatus"`
		} `json:"ShippingInfo"`
	} `json:"Sales"`
	Usertoken string `json:"Usertoken"`
}

// SyncOnlineSalesResponse is a automatically generated struct from json provided by sku vault's api docs.
type SyncOnlineSalesResponse struct {
	Errors []struct {
		ErrorMessages []string `json:"ErrorMessages"`
		OrderID       string   `json:"OrderId"`
	} `json:"Errors"`
	Status string `json:"Status"`
}
