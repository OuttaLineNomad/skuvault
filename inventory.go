package skuvault

// holds all api call endpoints with that are in the folder /Inventory

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// getWarehouseItemQuantity holds requst for same endpoint
type getWarehouseItemQuantity struct {
	request *http.Request
}

// PostGetWarehouseItemQuantity global parts user needs to use for api call
type PostGetWarehouseItemQuantity struct {
	Sku         string
	WarehouseID int
}

// postGetWarehouseItemQuantity payload sent to Sku Vault
type postGetWarehouseItemQuantity struct {
	Sku         string `json:"sku"`
	TenantToken string `json:"tenantToken"`
	UserToken   string `json:"userToken"`
	WarehouseID int    `json:"warehouseId"`
}

// ResponseGetWarehouseItemQuantity the response from SKU Vault endpoint
type ResponseGetWarehouseItemQuantity struct {
	Errors              []string
	TotalQuantityOnHand int
}

// GetWarehouseItemQuantity creates http request for this SKU vault endpoint
func (lc *LoginCredentials) GetWarehouseItemQuantity(pld *PostGetWarehouseItemQuantity) *getWarehouseItemQuantity {
	fullURL := URL + "/inventory/getWarehouseItemQuantity"
	credPld := &postGetWarehouseItemQuantity{
		Sku:         pld.Sku,
		TenantToken: lc.TenantToken,
		UserToken:   lc.UserToken,
		WarehouseID: pld.WarehouseID,
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

	return &getWarehouseItemQuantity{
		request: req,
	}
}

// Do makes the request to the attached method endpoint
func (ct *getWarehouseItemQuantity) Do() *ResponseGetWarehouseItemQuantity {
	response := ResponseGetWarehouseItemQuantity{}
	ctr := &svController{
		request:    ct.request,
		respStruct: &response,
	}
	ctr.do()
	return &response
}
