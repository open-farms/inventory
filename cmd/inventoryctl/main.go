package main

import (
	"github.com/alecthomas/kong"
	"github.com/open-farms/inventory/cmd/inventoryctl/commands"
)

var cli struct {
	Migrate commands.MigrateCmd `cmd:"" help:"Perform database migrations."`
}

func main() {
	ctx := kong.Parse(&cli)

	err := ctx.Run(&commands.Context{})
	ctx.FatalIfErrorf(err)
}
