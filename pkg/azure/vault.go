package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
	"github.com/pkg/errors"
)

// Vault object for manage Microsoft Azure vault keys.
type Vault interface {
	// Put key/value data to vault storage.
	Put(ctx context.Context, key, value string, opts ...VaultPutOpt) error

	// Get value by key from vault storage, return error if key not found.
	Get(ctx context.Context, key string) (string, error)

	// Delete key/value data from vault storage by key.
	Delete(ctx context.Context, key string) error

	setApiScheme(ApiScheme)
	setApiDomain(string)
}

type vault struct {
	client   *azsecrets.Client
	vaultUrl apiURL
}

// NewVaultWithClientSecret initialize client object for manage vault data.
func NewVaultWithClientSecret(vaultName, clientID, clientSecret, tenantID string, opts ...VaultOpt) (Vault, error) {
	v := &vault{
		client:   nil,
		vaultUrl: newApiUrl(ServiceVault),
	}
	for _, opt := range opts {
		opt(v)
	}

	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		return nil, err
	}
	client, err := azsecrets.NewClient(v.vaultUrl.get(vaultName), cred, nil)
	if err != nil {
		return nil, err
	}

	v.client = client
	return v, nil
}

// Put key/value in vault storage.
func (v *vault) Put(ctx context.Context, key, value string, opts ...VaultPutOpt) (err error) {
	p := vaultPutParams{
		params: azsecrets.SetSecretParameters{
			Value: &value,
		},
	}
	for _, opt := range opts {
		opt(p)
	}
	_, err = v.client.SetSecret(ctx, key, p.params, nil)
	return
}

// Get value from vault storage by key.
func (v *vault) Get(_ context.Context, _ string) (string, error) {
	return "", errors.New("Method not implemented")
}

// Delete key/value data from vault storage.
func (v *vault) Delete(_ context.Context, _ string) error {
	return errors.New("Method not implemented")
}

func (v *vault) setApiScheme(scheme ApiScheme) {
	v.vaultUrl.setScheme(scheme)
}

func (v *vault) setApiDomain(domain string) {
	v.vaultUrl.setDomain(domain)
}
