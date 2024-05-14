package commands

import (
	"github.com/mr-chelyshkin/hyperproof/commands/retrieve"
	"github.com/urfave/cli/v2"
)

func Commands() []*cli.Command {
	return []*cli.Command{
		retrieve.Command(),
	}
}
