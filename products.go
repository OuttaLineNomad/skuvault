package skuvault

import "github.com/OuttaLineNomad/skuvault/products"

// holds all api call endpoints with that are in the folder /Products

// postGetWarehouseItemQuantity payload sent to Sku Vault
type postGetProducts struct {
	*products.GetProducts
	*PLoginCredentials
}

// postAddItem  payload sent to Sku Vault.
type postAddItem struct {
	*products.AddItem
	*PLoginCredentials
}

// AddItem creates http request for this SKU vault endpoint
// Moderate throttling
// Add quantity to a warehouse location. Bulk version available
func (lc *PLoginCredentials) AddItem(pld *products.AddItem) *products.AddItemResponse {
	credPld := &postAddItem{
		AddItem:           pld,
		PLoginCredentials: lc,
	}

	response := &products.AddItemResponse{}
	productsCall(credPld, response, "addItem")
	return response
}

// GetProducts creates http request for this SKU vault endpoint
// Heavy throttling{
    /*
    	// Place your snippets for Go here. Each snippet is defined under a snippet name and has a prefix, body and 
    	// description. The prefix is what is used to trigger the snippet and the body will be expanded and inserted. Possible variables are:
    	// $1, $2 for tab stops, $0 for the final cursor position, and ${1:label}, ${2:another} for placeholders. Placeholders with the 
    	// same ids are connected.
    	// Example:
    	"Print to console": {
    		"prefix": "log",
    		"body": [
    			"console.log('$1');",
    			"$2"
    		],
    		"description": "Log output to console"
    	}
    */

    "SV Func": {
        "prefix": "SV func",
        "body": [
            "// ${1} creates http request for this SKU vault endpoint",
            "//${2}",
            "//${3}",
            "func (lc *PLoginCredentials) ${1}(pld *${4}.${1}) *${4}.${1}Response {",
            "credPld := &post${1}{",
            "${1}:       pld,",
            "${5}LoginCredentials: lc,",
            "}",
            "",
            "response := &${4}.${1}Response{}",
            "${4}Call(credPld, response, ${6})",
            "return response",
            "}"
        ],
        "description": "this is a sku vault api func"
    },

    "SV Global": {
        "prefix": "SV Global",
        "body": [
            "// ${1} global parts user needs to use for api call.",
            "type ${1} struct{",
            "${2}",
            "}"
        ]
    },
    "SV Post": {
        "prefix": "SV Post",
        "body": [
            "// post${1}  payload sent to Sku Vault.",
            "type post${1} struct {",
            "*${2}.${1}",
            "*${3}LoginCredentials",
            "}"
        ]
    },
    "SV Resp": {
        "prefix": "SV Resp",
        "body": [
            "// ${1}Response the response from SKU Vault endpoint.",
            "type ${1}Response struct{",
            "${2}",
            "}"
        ]
    }
}
// This call returns product(not kit) details. The date parameters include updating details as well as product
// creation. Details do not include quantity.
func (lc *PLoginCredentials) GetProducts(pld *products.GetProducts) *products.GetProductsResponse {
	credPld := &postGetProducts{
		GetProducts:       pld,
		PLoginCredentials: lc,
	}

	response := &products.GetProductsResponse{}
	productsCall(credPld, response, "getProducts")
	return response
}

// productsCall adds products/ to url for do() call.
func productsCall(pld interface{}, response interface{}, endPoint string) {
	full := "products/" + endPoint
	do(pld, response, full)
}
