package commands

import (
	"fmt"

	"github.com/open-farms/inventory/internal/biz"
)

type MigrateCmd struct {
	Dry bool `name:"dry" help:"Print out changes before applying"`
}

func (m *MigrateCmd) Run(globals *Globals) error {
	if m.Dry {
		fmt.Printf("\n%s\n\n", "Dry run enabled - no changes will be committed.")
	}

	err := biz.Migrate(globals.Config, m.Dry)
	if err != nil {
		return err
	}
	return nil
}
