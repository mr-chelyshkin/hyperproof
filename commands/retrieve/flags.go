package retrieve

import (
	"fmt"

	"github.com/mr-chelyshkin/hyperproof"
	"github.com/urfave/cli/v2"
)

func getFlagServiceAccountPath(ctx *cli.Context) string {
	return ctx.String(flagServiceAccountPath)
}

func getFlagProjectID(ctx *cli.Context) string {
	return ctx.String(flagProjectID)
}

func getFlagKeyName(ctx *cli.Context) string {
	return ctx.String(flagKeyName)
}

func flags() []cli.Flag {
	return []cli.Flag{
		// required flags:
		&cli.StringFlag{
			Name:     flagServiceAccountPath,
			Usage:    "path to service account JSON file",
			EnvVars:  []string{fmt.Sprintf("%s_SERVICE_ACCOUNT", hyperproof.EnvName)},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagProjectID,
			Usage:    "GCP service ID",
			EnvVars:  []string{fmt.Sprintf("%s_PROJECT_ID", hyperproof.EnvName)},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagKeyName,
			Usage:    "GCP key name",
			EnvVars:  []string{fmt.Sprintf("%s_KEY_NAME", hyperproof.EnvName)},
			Required: true,
		},
	}
}
