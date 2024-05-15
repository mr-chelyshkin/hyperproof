package retrieve

import (
	"fmt"
	"strings"

	"github.com/mr-chelyshkin/hyperproof"
	"github.com/urfave/cli/v2"
)

func getFlagGcpServiceAccountPath(ctx *cli.Context) string {
	return ctx.String(flagGcpServiceAccountPath)
}

func getFlagGcpProjectID(ctx *cli.Context) string {
	return ctx.String(flagGcpProjectID)
}

func getFlagKeyName(ctx *cli.Context) string {
	return ctx.String(flagKeyName)
}

func getFlagAzureVaultName(ctx *cli.Context) string {
	return ctx.String(flagAzureVaultName)
}

func getFlagAzureClientID(ctx *cli.Context) string {
	return ctx.String(flagAzureClientID)
}

func getFlagAzureClientSecret(ctx *cli.Context) string {
	return ctx.String(flagAzureClientSecret)
}

func getFlagAzureTenantID(ctx *cli.Context) string {
	return ctx.String(flagAzureTenantID)
}

func getFlagGcpKeyTargets(ctx *cli.Context) []string {
	return strings.Split(
		strings.ReplaceAll(ctx.String(flagGcpKeyTargets), " ", ""),
		",",
	)
}

func getFlagGcpKeyIPs(ctx *cli.Context) []string {
	return strings.Split(
		strings.ReplaceAll(ctx.String(flagGcpKeyIPs), " ", ""),
		",",
	)
}

func flags() []cli.Flag {
	return []cli.Flag{
		// required flags:
		&cli.StringFlag{
			Name:     flagGcpServiceAccountPath,
			Usage:    "path to service account JSON file",
			EnvVars:  []string{fmt.Sprintf("%s_GCP_SERVICE_ACCOUNT", hyperproof.EnvName)},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagGcpProjectID,
			Usage:    "GCP service ID",
			EnvVars:  []string{fmt.Sprintf("%s_GCP_PROJECT_ID", hyperproof.EnvName)},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagKeyName,
			Usage:    "GCP key name",
			EnvVars:  []string{fmt.Sprintf("%s_KEY_NAME", hyperproof.EnvName)},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagAzureVaultName,
			Usage:    "Azure vault name",
			EnvVars:  []string{fmt.Sprintf("%s_AZURE_VAULT_NAME", hyperproof.EnvName)},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagAzureClientID,
			Usage:    "Azure client id",
			EnvVars:  []string{fmt.Sprintf("%s_AZURE_CLIENT_ID", hyperproof.EnvName)},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagAzureClientSecret,
			Usage:    "Azure client secret",
			EnvVars:  []string{fmt.Sprintf("%s_AZURE_CLIENT_SECRET", hyperproof.EnvName)},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagAzureTenantID,
			Usage:    "Azure tenant id",
			EnvVars:  []string{fmt.Sprintf("%s_AZURE_TENANT_ID", hyperproof.EnvName)},
			Required: true,
		},
		// non-required flags:
		&cli.StringFlag{
			Name:     flagGcpKeyTargets,
			Usage:    "GCP api key targets restriction by comma",
			Required: false,
		},
		&cli.StringFlag{
			Name:     flagGcpKeyIPs,
			Usage:    "GCP api key IPs restriction by comma",
			Required: false,
		},
	}
}
