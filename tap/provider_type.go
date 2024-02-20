/*
Package tap wraps a set of interfaces and object to provide a generic interface to a delegated authentication
proxy
*/
package tap

// ProviderType is a way of identifying whether a provider passes through or redirects
type ProviderType string

const (
	PASSTHROUGH_PROVIDER ProviderType = "passthrough"
	REDIRECT_PROVIDER    ProviderType = "redirect"
)
