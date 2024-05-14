package main

import (
	"os"

	"github.com/mr-chelyshkin/hyperproof"
	"github.com/mr-chelyshkin/hyperproof/commands"
	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

func init() {
	hyperproof.Logger.SetOutput(os.Stdout)
	hyperproof.Logger.SetLevel(logrus.InfoLevel)
	hyperproof.Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

func main() {
	app := &cli.App{
		Name:     hyperproof.Name,
		Usage:    hyperproof.Usage,
		Commands: commands.Commands(),
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
