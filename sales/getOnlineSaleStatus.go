package sales

import "time"

// GetOnlineSaleStatus is a automatically generated struct from json provided by sku vault's api docs.
type GetOnlineSaleStatus struct {
	OrderIds []string `json:"OrderIds"`
}

// GetOnlineSaleStatusResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetOnlineSaleStatusResponse struct {
	Sales []struct {
		ID    string `json:"Id"`
		Items []struct {
			OnlineSaleItemStatus string `json:"OnlineSaleItemStatus"`
			Quantity             int    `json:"Quantity"`
			Sku                  string `json:"Sku"`
		} `json:"Items"`
		LastPrintedDate time.Time `json:"LastPrintedDate"`
		Notes           string    `json:"Notes"`
		PrintedStatus   bool      `json:"PrintedStatus"`
		Status          string    `json:"Status"`
	} `json:"Sales"`
}
