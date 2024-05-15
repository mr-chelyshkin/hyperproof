package retrieve

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/mr-chelyshkin/hyperproof"
	"github.com/urfave/cli/v2"
)

var (
	flagAzureVaultName        = "azure-vault-name"
	flagAzureClientID         = "azure-vault-client-id"
	flagAzureClientSecret     = "azure-vault-client-secret"
	flagAzureTenantID         = "azure-vault-tenant-id"
	flagGcpServiceAccountPath = "gcp-service-account"
	flagGcpKeyTargets         = "gcp-key-targets"
	flagGcpKeyIPs             = "gcp-key-ips"
	flagGcpProjectID          = "gcp-project-id"
	flagKeyName               = "key-name"

	name  = "retrieve"
	usage = "retrieve GCP token and collect it in MS Azure"
)

type tmplUsage struct {
	FlagGcpServiceAccount string
	FlagGcpProjectID      string
	FlagKeyName           string
	FlagAzureVaultName    string
	FlagAzureClientID     string
	FlagAzureClientSecret string
	FlagAzureTenantID     string
	FlagGcpKeyTargets     string
	FlagGcpKeyIPs         string
}

func Command() *cli.Command {
	tmpl, err := template.New("usage").Funcs(
		template.FuncMap{
			"Join": strings.Join,
		},
	).Parse(usageTemplate)
	if err != nil {
		hyperproof.Logger.Fatal(err)
	}

	var usageText bytes.Buffer
	err = tmpl.Execute(&usageText, tmplUsage{
		FlagGcpServiceAccount: flagGcpServiceAccountPath,
		FlagGcpProjectID:      flagGcpProjectID,
		FlagKeyName:           flagKeyName,
		FlagAzureVaultName:    flagAzureVaultName,
		FlagAzureClientID:     flagAzureClientID,
		FlagAzureClientSecret: flagAzureClientSecret,
		FlagAzureTenantID:     flagAzureTenantID,
		FlagGcpKeyTargets:     flagGcpKeyTargets,
		FlagGcpKeyIPs:         flagGcpKeyIPs,
	})
	if err != nil {
		hyperproof.Logger.Fatal(err)
	}
	return &cli.Command{
		Name:      name,
		Usage:     usage,
		Flags:     flags(),
		Action:    action,
		UsageText: usageText.String(),
	}
}
