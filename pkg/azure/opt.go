package azure

// VaultOpt is a common opts for initialize Azure Vault client.
type VaultOpt func(Vault)

// WithVaultApiScheme modify MS Azure api scheme.
func WithVaultApiScheme(scheme ApiScheme) VaultOpt {
	return func(v Vault) {
		v.setApiScheme(scheme)
	}
}

// WithVaultApiDomain modify MS Azure domain url.
func WithVaultApiDomain(domain string) VaultOpt {
	return func(v Vault) {
		v.setApiDomain(domain)
	}
}

// VaultPutOpt is a common opts which extends azsecrets.SetSecretParameters.
type VaultPutOpt func(vaultPutParams)

// WithVaultPutContentType add Content Type for vault put request api.
func WithVaultPutContentType(contentType string) VaultPutOpt {
	return func(p vaultPutParams) {
		p.params.ContentType = &contentType
	}
}

// WithVaultPutTags add additional tags for key/value in vault storage.
func WithVaultPutTags(tags map[string]*string) VaultPutOpt {
	return func(p vaultPutParams) {
		p.params.Tags = tags
	}
}
