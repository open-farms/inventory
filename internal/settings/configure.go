package settings

import (
	"context"
	"os"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/open-farms/inventory/ent"
)

type Auth struct {
	Enabled bool   `default:"false"`
	Token   string `default:"inventory"`
}

type Config struct {
	HTTP struct {
		Addr    string        `default:"0.0.0.0:8000"`
		Timeout time.Duration `default:"5s"`
		Auth    Auth
	}
	GRPC struct {
		Addr    string        `default:"0.0.0.0:9000"`
		Timeout time.Duration `default:"5s"`
	}
	Storage struct {
		Driver  string `default:"postgres"`
		URI     string `required:"true"`
		Migrate bool   `default:"true"`
		Debug   bool   `default:"false"`
	}
}

func Migrate(dry bool) error {

	// Initialize configuration
	c, err := Configure()
	if err != nil {
		return err
	}

	// Setup client for database
	client, err := ent.Open(c.Storage.Driver, c.Storage.URI)
	if err != nil {
		return err
	}
	defer client.Close()

	if dry {
		err := client.Schema.WriteTo(context.Background(), os.Stdout)
		if err != nil {
			return err
		}
	}

	// Run the auto migration tool
	ctx := context.Background()
	if c.Storage.Debug {
		err = client.Debug().Schema.Create(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	err = client.Schema.Create(ctx,
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	)
	if err != nil {
		return err
	}

	return nil
}

func Configure() (*Config, error) {
	var c Config
	err := envconfig.Process("INVENTORY", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
