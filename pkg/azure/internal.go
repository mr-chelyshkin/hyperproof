package azure

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
)

type apiURL struct {
	service AzureService
	scheme  ApiScheme
	domain  string
}

func newApiUrl(service AzureService) apiURL {
	return apiURL{
		domain:  "azure.net",
		service: service,
		scheme:  HTTPS,
	}
}

func (a *apiURL) get(subdomain string) string {
	return fmt.Sprintf("%s://%s.%s.%s/", a.scheme, subdomain, a.service, a.domain)
}

func (a *apiURL) setScheme(scheme ApiScheme) {
	a.scheme = scheme
}

func (a *apiURL) setDomain(domain string) {
	a.domain = domain
}

type vaultPutParams struct {
	params azsecrets.SetSecretParameters
}
