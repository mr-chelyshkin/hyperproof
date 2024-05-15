package retrieve

import (
	"github.com/mr-chelyshkin/hyperproof"
	"github.com/mr-chelyshkin/hyperproof/pkg/azure"
	"github.com/mr-chelyshkin/hyperproof/pkg/google"
	"github.com/urfave/cli/v2"
	"sync"
)

func action(ctx *cli.Context) error {
	gcpCli, err := google.NewGCPCliWithCredFromFile(ctx.Context, getFlagGcpServiceAccountPath(ctx))
	if err != nil {
		return err
	}
	azrCli, err := azure.NewVaultWithClientSecret(
		getFlagAzureVaultName(ctx),
		getFlagAzureClientID(ctx),
		getFlagAzureClientSecret(ctx),
		getFlagAzureTenantID(ctx),
	)
	if err != nil {
		return err
	}

	existedKeys, err := gcpCli.KeysList(getFlagGcpProjectID(ctx), getFlagKeyName(ctx))
	if err != nil {
		return err
	}
	newKey, err := gcpCli.KeysCreate(
		getFlagGcpProjectID(ctx),
		getFlagKeyName(ctx),

		google.WithApiKeysCreateTargetsRestrictions(getFlagGcpKeyTargets(ctx)),
		google.WithApiKeysCreateIPRestrictions(getFlagGcpKeyIPs(ctx)),
	)
	if err != nil {
		return err
	}
	hyperproof.Logger.Println("GCP: new API key was generated: ", newKey.Mask())

	if err := azrCli.Put(ctx.Context, getFlagKeyName(ctx), newKey.Token); err != nil {
		return err
	}
	hyperproof.Logger.Println("VAULT: key was updated: ", getFlagKeyName(ctx))

	var wg sync.WaitGroup
	var errMutex sync.Mutex
	var errors []error

	for _, oldKey := range existedKeys {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()

			if err := gcpCli.KeysDelete(key); err != nil {
				errMutex.Lock()
				errors = append(errors, err)
				errMutex.Unlock()
			}
		}(oldKey.Name)
	}
	wg.Wait()
	for _, e := range errors {
		hyperproof.Logger.Warn(e.Error())
	}
	return nil
}
