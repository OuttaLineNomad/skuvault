package skuvault

import "github.com/OuttaLineNomad/skuvault/integration"

// postGetIntegrations payload sent to Sku Vault.
type postGetIntegrations struct {
	*INLoginCredentials
}

// GetIntegrations creates http request for this SKU vault endpoint.
// Severe Throttle.
// Returns a list of your enabled channel accounts. No page parameters..
func (lc *INLoginCredentials) GetIntegrations() (*integration.GetIntegrationsResponse, error) {
	credPld := &postGetIntegrations{
		INLoginCredentials: lc,
	}

	response := &integration.GetIntegrationsResponse{}
	err := do(credPld, response, "integration/getIntegrations")
	if err != nil {
		return nil, &Error{"GetIntegrations()", err}
	}

	return response, nil
}
