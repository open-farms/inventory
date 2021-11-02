package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Vehicle holds the schema definition for the Vehicle entity.
type Vehicle struct {
	ent.Schema
}

// Fields of the Vehicle.
func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Positive(),
		field.String("make"),
		field.String("model"),
		field.String("year"),
		field.Bool("active"),
		field.Strings("tags"),
		field.Enum("condition").
			Values(
				"UNSPECIFIED",
				"MINT",
				"GOOD",
				"POOR",
				"BROKEN",
			),
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL:    "datetime",
				dialect.Postgres: "timestamp",
			}),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL:    "datetime",
				dialect.Postgres: "timestamp",
			}),
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return nil
}
