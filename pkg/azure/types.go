package azure

type ApiScheme string
type AzureService string

var (
	HTTPS ApiScheme = "https"
	HTTP  ApiScheme = "http"
)

var (
	ServiceVault AzureService = "vault"
)
