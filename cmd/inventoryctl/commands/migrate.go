package commands

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/open-farms/inventory/ent"
)

type MigrateCmd struct {
	Driver string `name:"driver" required:"" env:"DB_DRIVER" default:"sqlite3" help:"Name of database driver"`
	URI    string `name:"uri" required:"" env:"DB_URI" default:"file:inventory?mode=memory&cache=shared&_fk=1" help:"Database connection string"`
	Dry    bool   `name:"dry" help:"Print out changes before applying"`
}

func (m *MigrateCmd) Run(ctx *Context) error {
	client, err := ent.Open(m.Driver, m.URI)
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	defer client.Close()

	if m.Dry {
		fmt.Printf("\n%s\n\n", "Dry run enabled - no changes will be committed.")
		if err := client.Schema.WriteTo(context.Background(), os.Stdout); err != nil {
			log.Fatalf("failed printing schema changes: %v", err)
		}
		return nil
	}

	// Run migration.
	err = client.Debug().Schema.Create(
		context.Background(),
	)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
