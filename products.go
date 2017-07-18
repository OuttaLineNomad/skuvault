package skuvault

// holds all api call endpoints with that are in the folder /Products

import (
	"time"
)

// PostGetProducts global parts user needs to use for api call
type PostGetProducts struct {
	ModifiedAfterDateTimeUtc  time.Time `json:"ModifiedAfterDateTimeUtc"`
	ModifiedBeforeDateTimeUtc time.Time `json:"ModifiedBeforeDateTimeUtc"`
	PageNumber                int       `json:"PageNumber"`
	PageSize                  int       `json:"PageSize"`
	ProductSKUs               []string  `json:"ProductSKUs"`
}

// postGetWarehouseItemQuantity payload sent to Sku Vault
type postGetProducts struct {
	*PostGetProducts
	TenantToken string `json:"TenantToken"`
	UserToken   string `json:"UserToken"`
}

// products all data on one product
type product struct {
	ID               string  `json:"Id"`
	Code             string  `json:"Code"`
	Sku              string  `json:"Sku"`
	PartNumber       string  `json:"PartNumber"`
	Description      string  `json:"Description"`
	ShortDescription string  `json:"ShortDescription"`
	LongDescription  string  `json:"LongDescription"`
	Cost             float64 `json:"Cost"`
	RetailPrice      float64 `json:"RetailPrice"`
	SalePrice        float64 `json:"SalePrice"`
	WeightValue      string  `json:"WeightValue"`
	WeightUnit       string  `json:"WeightUnit"`
	ReorderPoint     int     `json:"ReorderPoint"`
	Brand            string  `json:"Brand"`
	Supplier         string  `json:"Supplier"`
	SupplierInfo     []struct {
		SupplierName       string  `json:"SupplierName"`
		SupplierPartNumber string  `json:"SupplierPartNumber"`
		Cost               float64 `json:"Cost"`
		LeadTime           int     `json:"LeadTime"`
		IsActive           bool    `json:"IsActive"`
		IsPrimary          bool    `json:"IsPrimary"`
	} `json:"SupplierInfo"`
	Classification    string    `json:"Classification"`
	QuantityOnHand    int       `json:"QuantityOnHand"`
	QuantityOnHold    int       `json:"QuantityOnHold"`
	QuantityPicked    int       `json:"QuantityPicked"`
	QuantityPending   int       `json:"QuantityPending"`
	QuantityAvailable int       `json:"QuantityAvailable"`
	QuantityIncoming  int       `json:"QuantityIncoming"`
	QuantityInbound   int       `json:"QuantityInbound"`
	QuantityTransfer  int       `json:"QuantityTransfer"`
	QuantityInStock   int       `json:"QuantityInStock"`
	QuantityTotalFBA  int       `json:"QuantityTotalFBA"`
	CreatedDateUtc    time.Time `json:"CreatedDateUtc"`
	ModifiedDateUtc   time.Time `json:"ModifiedDateUtc"`
	Pictures          []string  `json:"Pictures"`
	Attributes        []struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	} `json:"Attributes"`
	VariationParentSku  string        `json:"VariationParentSku"`
	IsAlternateSKU      bool          `json:"IsAlternateSKU"`
	MOQ                 int           `json:"MOQ"`
	MOQInfo             string        `json:"MOQInfo"`
	IncrementalQuantity int           `json:"IncrementalQuantity"`
	DisableQuantitySync bool          `json:"DisableQuantitySync"`
	Statuses            []interface{} `json:"Statuses"`
}

// ResponseGetProducts the response from SKU Vault endpoint
type ResponseGetProducts struct {
	Products []product     `json:"Products"`
	Errors   []interface{} `json:"Errors"`
}

// GetProducts creates http request for this SKU vault endpoint
// Heavy throttling
// This call returns product(not kit) details. The date parameters include updating details as well as product
// creation. Details do not include quantity.
func (lc *PLoginCredentials) GetProducts(pld *PostGetProducts) *ResponseGetProducts {
	credPld := &postGetProducts{
		PostGetProducts: pld,
		TenantToken:     lc.tenantToken,
		UserToken:       lc.userToken,
	}

	response := &ResponseGetProducts{}
	productsCall(credPld, response, "getProducts")
	return response
}

// productsCall adds products/ to url for do() call.
func productsCall(pld interface{}, response interface{}, endPoint string) {
	full := "products/" + endPoint
	do(pld, response, full)
}
