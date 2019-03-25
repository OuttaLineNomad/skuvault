package skuvault

import "github.com/OuttaLineNomad/skuvault/integration"

// postGetIntegrations payload sent to Sku Vault.
type postGetIntegrations struct {
	*INLoginCredentials
}

// GetIntegrations creates http request for this SKU vault endpoint.
// Severe Throttle.
// Returns a list of your enabled channel accounts. No page parameters..
func (lc *INLoginCredentials) GetIntegrations() *integration.GetIntegrationsResponse {
	credPld := &postGetIntegrations{
		INLoginCredentials: lc,
	}

	response := &integration.GetIntegrationsResponse{}
	do(credPld, response, "integration/getIntegrations")
	return response
}
