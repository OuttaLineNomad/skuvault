package skuvault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	// url is a constant for all SKU Vault calls
	url = "https://app.skuvault.com/api/"
)

// Ctr is the controls the flow of endpoints.
type Ctr struct {
	Inventory *ILoginCredentials
	Products  *PLoginCredentials
}

// ILoginCredentials hold credentials to sign into SKU Vault API for inventory endpoints.
type ILoginCredentials struct {
	TenantToken string
	UserToken   string
}

// PLoginCredentials hold credentials to sign into SKU Vault API for product endpoints.
type PLoginCredentials struct {
	TenantToken string
	UserToken   string
}

// Error struct to share errors from package
type Error struct {
	Func string
	Err  error
}

// Error func to customize errors from package.
func (er *Error) Error() string {
	return `skuvault: ` + er.Func + `: ` + er.Err.Error()
}

// NewEnvCredSession takes tokens from systems enviomantal varables.
// TENANT_TOKEN and USER_TOKEN
func NewEnvCredSession() *Ctr {
	return NewSession(os.Getenv("SV_TENANT_TOKEN"), os.Getenv("SV_USER_TOKEN"))
}

// NewSession creates a new session sets credentails to make call
func NewSession(tTok, uTok string) *Ctr {
	iCreds := &ILoginCredentials{
		TenantToken: tTok,
		UserToken:   uTok,
	}

	pCreds := &PLoginCredentials{
		TenantToken: tTok,
		UserToken:   uTok,
	}

	return &Ctr{
		Inventory: iCreds,
		Products:  pCreds,
	}
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
		fmt.Println("sv response with caused error: \n", string(b))
		return err
	}
	return nil

}

// TimeString converts time to proper formated string for Sku Vault
func TimeString(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.9999999Z")
}
