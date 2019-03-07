package products

import "time"

// AddItem global parts user needs to use for api call.
type AddItem struct {
	Sku          string
	Code         string
	WarehouseID  int
	LocationCode string
	Quantity     int
	Reason       string
}

// AddItemResponse the response from SKU Vault endpoint.
type AddItemResponse struct {
	AddItemStatus string
}

// GetProducts global parts user needs to use for api call
type GetProducts struct {
	ModifiedAfterDateTimeUtc  time.Time
	ModifiedBeforeDateTimeUtc time.Time
	PageNumber                int
	PageSize                  int
	ProductSKUs               []string
	ProductCodes              []string
}

// Product all data on one product
type Product struct {
	ID               interface{}
	Code             string
	Sku              string
	PartNumber       string
	Description      string
	ShortDescription string
	LongDescription  string
	Cost             float64
	RetailPrice      float64
	SalePrice        float64
	WeightValue      string
	WeightUnit       string
	ReorderPoint     int
	Brand            string
	Supplier         string
	SupplierInfo     []struct {
		SupplierName       string
		SupplierPartNumber string
		Cost               float64
		LeadTime           int
		IsActive           bool
		IsPrimary          bool
	}
	Classification    string
	QuantityOnHand    int
	QuantityOnHold    int
	QuantityPicked    int
	QuantityPending   int
	QuantityAvailable int
	QuantityIncoming  int
	QuantityInbound   int
	QuantityTransfer  int
	QuantityInStock   int
	QuantityTotalFBA  int
	CreatedDateUtc    time.Time
	ModifiedDateUtc   time.Time
	Pictures          []string
	Attributes        []struct {
		Name  string
		Value string
	}
	VariationParentSku  string
	IsAlternateSKU      bool
	MOQ                 int
	MOQInfo             string
	IncrementalQuantity int
	DisableQuantitySync bool
	Statuses            []interface{}
}

// GetProductsResponse the response from SKU Vault endpoint
type GetProductsResponse struct {
	Products []Product
	Errors   []interface{}
}
