package skuvault

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	// url is a constant for all skuvault calls
	url = "https://app.skuvault.com/api/"
)

// Ctr is the controls the flow of endpoints.
type Ctr struct {
	Inventory      *ILoginCredentials
	Products       *PLoginCredentials
	Sales          *SLoginCredentials
	Purchaseorders *POLoginCredentials
	Integration    *INLoginCredentials
}

// LoginCredentials hold credentials to sign into SKU Vault API for endpoints.
type LoginCredentials struct {
	TenantToken string
	UserToken   string
}

// ILoginCredentials hold credentials to sign into SKU Vault API for inventory endpoints.
type ILoginCredentials struct {
	LoginCredentials
}

// PLoginCredentials hold credentials to sign into SKU Vault API for products endpoints.
type PLoginCredentials struct {
	LoginCredentials
}

// SLoginCredentials hold credentials to sign into SKU Vault API for sales endpoints.
type SLoginCredentials struct {
	LoginCredentials
}

// POLoginCredentials hold credentials to sign into SKU Vault API for purchaseorders endpoints.
type POLoginCredentials struct {
	LoginCredentials
}

// INLoginCredentials hold credentials to sign into SKU Vault API for integration endpoints.
type INLoginCredentials struct {
	LoginCredentials
}

// New takes tokens from systems environmental varables.
// TENANT_TOKEN and USER_TOKEN
func New() *Ctr {
	return NewSession(os.Getenv("SV_TENANT_TOKEN"), os.Getenv("SV_USER_TOKEN"))
}

// NewSession creates a new session sets credentails to make call
func NewSession(tTok, uTok string) *Ctr {
	svc := LoginCredentials{
		TenantToken: tTok,
		UserToken:   uTok,
	}

	return &Ctr{
		Inventory:      &ILoginCredentials{svc},
		Products:       &PLoginCredentials{svc},
		Sales:          &SLoginCredentials{svc},
		Purchaseorders: &POLoginCredentials{svc},
		Integration:    &INLoginCredentials{svc},
	}
}

// postGetTokens payload sent to Sku Vault.
type postGetTokens struct {
	*GetTokens
}

// GetTokensResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetTokensResponse struct {
	TenantToken string
	UserToken   string
}

// GetTokens is a automatically generated struct from json provided by sku vault's api docs.
type GetTokens struct {
	Email    string
	Password string
}

// GetTokensCall creates http request for this SKU vault endpoint.
// Very Light Throttle.
// Use this call to retrieve your API tokens from SkuVault using your login email and password..
func GetTokensCall(pld *GetTokens) *GetTokensResponse {
	credPld := &postGetTokens{
		GetTokens: pld,
	}

	response := &GetTokensResponse{}
	do(credPld, response, "skuvault/getTokens")
	return response
}

// do internal makes calls based on information passed in from other Do calls for each endpoint
func do(pld interface{}, response interface{}, endPoint string) error {
	fullURL := url + endPoint
	bt, err := json.Marshal(pld)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(bt)
	req, err := http.NewRequest(http.MethodPost, fullURL, payload)
	if err != nil {
		return err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(b))
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, response)
	if err != nil {
		return err
	}
	return nil
}

// TimeString converts time to proper formated string for Sku Vault
func TimeString(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.9999999Z")
}
