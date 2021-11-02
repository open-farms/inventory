package main

import (
	"github.com/alecthomas/kong"
	"github.com/open-farms/inventory/cmd/inventoryctl/commands"
)

type CLI struct {
	commands.Globals
	Migrate commands.MigrateCmd `cmd:"" help:"Perform database migrations."`
}

func main() {

	cli := CLI{
		Globals: commands.Globals{},
	}

	ctx := kong.Parse(
		&cli,
		kong.Name("inventoryctl"),
		kong.Description("Administration command line for the inventory agriculture service"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)

	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
