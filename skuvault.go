package skuvault

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

// svController holds the struct to use to Unmarshal json and http request
type svController struct {
	respStruct interface{}
	request    *http.Request
}

// ILoginCredentials hold credentials to sign into SKU Vault API for inventory endpoints.
type ILoginCredentials struct {
	tenantToken string
	userToken   string
}

// PLoginCredentials hold credentials to sign into SKU Vault API for product endpoints.
type PLoginCredentials struct {
	tenantToken string
	userToken   string
}

// NewEnvCredSession takes tokens from systems enviomantal varables.
// TENANT_TOKEN and USER_TOKEN
func NewEnvCredSession() {
	NewSession(os.Getenv("TENANT_TOKEN"), os.Getenv("USER_TOKEN"))
}

// NewSession creates a new session sets credentails to make call
func NewSession(tTok, uTok string) *Ctr {
	iCreds := &ILoginCredentials{
		tenantToken: tTok,
		userToken:   uTok,
	}

	pCreds := &PLoginCredentials{
		tenantToken: tTok,
		userToken:   uTok,
	}

	return &Ctr{
		Inventory: iCreds,
		Products:  pCreds,
	}
}

// do internal makes calls based on information passed in from other Do calls for each endpoint
func do(sc *svController) {

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Do(sc.request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(b))
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &sc.respStruct)
	if err != nil {
		log.Fatal(err)
	}

}
