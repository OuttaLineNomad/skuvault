package sales

import "time"

// SyncOnlineSale is a automatically generated struct from json provided by sku vault's api docs.
type SyncOnlineSale struct {
	ChannelAccountID string `json:"ChannelAccountId"`
	ChannelType      string `json:"ChannelType"`
	CheckoutStatus   string `json:"CheckoutStatus"`
	FulfilledItems   []struct {
		Quantity  int     `json:"Quantity"`
		Sku       string  `json:"Sku"`
		UnitPrice float64 `json:"UnitPrice"`
	} `json:"FulfilledItems"`
	ItemSkus []struct {
		Quantity  int     `json:"Quantity"`
		Sku       string  `json:"Sku"`
		UnitPrice float64 `json:"UnitPrice"`
	} `json:"ItemSkus"`
	MarketplaceID string    `json:"MarketplaceId"`
	Notes         string    `json:"Notes"`
	OrderDateUtc  time.Time `json:"OrderDateUtc"`
	OrderID       string    `json:"OrderId"`
	OrderTotal    float64   `json:"OrderTotal"`
	PaymentStatus string    `json:"PaymentStatus"`
	SaleState     string    `json:"SaleState"`
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
}

// SyncOnlineSaleResponse is a automatically generated struct from json provided by sku vault's api docs.
type SyncOnlineSaleResponse struct {
	OrderID string `json:"OrderId"`
	Status  string `json:"Status"`
}
