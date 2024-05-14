package retrieve

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/mr-chelyshkin/hyperproof"
	"github.com/urfave/cli/v2"
)

var (
	flagServiceAccountPath = "service-account"
	flagProjectID          = "project-id"
	flagKeyName            = "key-name"

	name  = "retrieve"
	usage = "retrieve GCP token and collect it in MS Azure"
)

type tmplUsage struct{}

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
	err = tmpl.Execute(&usageText, tmplUsage{})
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
