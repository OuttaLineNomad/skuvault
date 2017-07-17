package skuvault

// holds all api call endpoints with that are in the folder /Inventory

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// PostGetWarehouseItemQuantity global parts user needs to use for api call
type PostGetWarehouseItemQuantity struct {
	Sku         string
	WarehouseID int
}

// postGetWarehouseItemQuantity payload sent to Sku Vault
type postGetWarehouseItemQuantity struct {
	PostGetWarehouseItemQuantity
	TenantToken string `json:"tenantToken"`
	UserToken   string `json:"userToken"`
}

// ResponseGetWarehouseItemQuantity the response from SKU Vault endpoint
type ResponseGetWarehouseItemQuantity struct {
	Errors              []interface{}
	TotalQuantityOnHand int
}

// GetWarehouseItemQuantity creates http request for this SKU vault endpoint
// Returns the quantity for a specified SKU
func (lc *ILoginCredentials) GetWarehouseItemQuantity(pld *PostGetWarehouseItemQuantity) *ResponseGetWarehouseItemQuantity {
	fullURL := url + "inventory/getWarehouseItemQuantity"
	credPld := &postGetWarehouseItemQuantity{
		PostGetWarehouseItemQuantity: *pld,
		TenantToken:                  lc.tenantToken,
		UserToken:                    lc.userToken,
	}
	b, err := json.Marshal(credPld)
	if err != nil {
		log.Fatal(err)
	}
	payload := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPost, fullURL, payload)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	response := ResponseGetWarehouseItemQuantity{}
	ctr := &svController{
		request:    req,
		respStruct: &response,
	}
	do(ctr)
	return &response
}
