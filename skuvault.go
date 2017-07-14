package skuvault

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	// URL is a constant for all SKU Vault calls
	URL = "https://app.skuvault.com/api"
)

// svController holds the struct to use to Unmarshal json and http request
type svController struct {
	respStruct interface{}
	request    *http.Request
}

// LoginCredentials hold credentials to sign into SKU Vault API
type LoginCredentials struct {
	TenantToken string
	UserToken   string
}

// NewSession creates a new session sets credentails to make call
func NewSession(tTok, uTok string) *LoginCredentials {
	creds := &LoginCredentials{
		TenantToken: tTok,
		UserToken:   uTok,
	}

	return creds
}

// do internal makes calls based on information passed in from other Do calls for each endpoint
func (sc *svController) do() {

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Do(sc.request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &sc.respStruct)
	if err != nil {
		log.Fatal(err)
	}

}
