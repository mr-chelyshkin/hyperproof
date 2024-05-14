package retrieve

import (
	"github.com/mr-chelyshkin/hyperproof/pkg/azure"
	"github.com/mr-chelyshkin/hyperproof/pkg/google"
	"github.com/urfave/cli/v2"
)

func action(ctx *cli.Context) error {

	///
	apiKeysCli, err := google.NewApiKeysWithCredFromFile(ctx.Context, getFlagServiceAccountPath(ctx))
	if err != nil {
		return err
	}
	vaultCli, err := azure.NewVaultWithClientSecret(vaultName, clientID, clientSecret, tenantID)
	if err != nil {
		return err
	}

	existMapApiKeys, err := apiKeysCli.List(getFlagProjectID(ctx), getFlagKeyName(ctx))
	if err != nil {
		return err
	}
	newKey, err := apiKeysCli.Create(getFlagProjectID(ctx), getFlagKeyName(ctx),
		google.WithApikeysTargetsRestrictions([]string{"example.googleapis.com"}),
		google.WithApikeysIPRestrictions([]string{"127.0.0.1"}))
	if err != nil {
		return err
	}
	if err := vaultCli.Put(ctx.Context, getFlagKeyName(ctx), newKey.Token); err != nil {
		return err
	}
	for _, oldKey := range existMapApiKeys {
		_ = apiKeysCli.Delete(oldKey.Name)
	}
	return nil
}
