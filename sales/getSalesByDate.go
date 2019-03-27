package sales

import "time"

// GetSalesByDate is a automatically generated struct from json provided by sku vault's api docs.
type GetSalesByDate struct {
	DateField  time.Time `json:"DateField"`
	FromDate   time.Time `json:"FromDate"`
	PageNumber int       `json:"PageNumber"`
	PageSize   int       `json:"PageSize"`
	ToDate     time.Time `json:"ToDate"`
}

// GetSalesByDateResponse is a automatically generated struct from json provided by sku vault's api docs.
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
	FulfilledItems []struct {
		Quantity  int     `json:"Quantity"`
		Sku       string  `json:"Sku"`
		UnitPrice float64 `json:"UnitPrice"`
	} `json:"FulfilledItems"`
	FulfilledKits   []interface{} `json:"FulfilledKits"`
	ID              string        `json:"Id"`
	LastPrintedDate time.Time     `json:"LastPrintedDate"`
	Marketplace     string        `json:"Marketplace"`
	MarketplaceID   string        `json:"MarketplaceId"`
	MerchantItems   []struct {
		PartNumber string `json:"PartNumber"`
		Quantity   int    `json:"Quantity"`
		Sku        string `json:"Sku"`
		UnitPrice  struct {
			A float64 `json:"a"`
			S string  `json:"s"`
		} `json:"UnitPrice"`
	} `json:"MerchantItems"`
	MerchantKits []struct {
		KitItems  struct{} `json:"KitItems"`
		Quantity  int      `json:"Quantity"`
		Sku       string   `json:"Sku"`
		UnitPrice struct {
			A float64 `json:"a"`
			S string  `json:"s"`
		} `json:"UnitPrice"`
	} `json:"MerchantKits"`
	Notes          string `json:"Notes"`
	PrintedStatus  bool   `json:"PrintedStatus"`
	ProcessedItems []struct {
		FailedQuantity int    `json:"FailedQuantity"`
		PassedQuantity int    `json:"PassedQuantity"`
		PickedQuantity int    `json:"PickedQuantity"`
		Sku            string `json:"Sku"`
	} `json:"ProcessedItems"`
	SaleDate        time.Time `json:"SaleDate"`
	ShippingCarrier string    `json:"ShippingCarrier"`
	ShippingCharge  struct {
		A float64 `json:"a"`
		S string  `json:"s"`
	} `json:"ShippingCharge"`
	ShippingClass string `json:"ShippingClass"`
	ShippingCost  struct {
		A float64 `json:"a"`
		S string  `json:"s"`
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
