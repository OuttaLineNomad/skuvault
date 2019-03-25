package sales

// SetShipmentFile is a automatically generated struct from json provided by sku vault's api docs.
type SetShipmentFile struct {
	Shipments []struct {
		Carrier        string `json:"Carrier"`
		FilesData      string `json:"FilesData"`
		SaleID         string `json:"SaleId"`
		TrackingNumber string `json:"TrackingNumber"`
	} `json:"Shipments"`
}

// SetShipmentFileResponse is a automatically generated struct from json provided by sku vault's api docs.
type SetShipmentFileResponse struct {
	Errors  []interface{} `json:"Errors"`
	FileIds []struct{}    `json:"FileIds"`
	Status  string        `json:"Status"`
}
