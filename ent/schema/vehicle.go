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
		field.Strings("tags"),
		field.Enum("condition").
			Values(
				"CONDITION_UNSPECIFIED",
				"CONDITION_MINT",
				"CONDITION_GOOD",
				"CONDITION_POOR",
				"CONDITION_BROKEN",
			),
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL:    "datetime",
				dialect.Postgres: "time with time zone",
			}),
		field.Time("updated_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL:    "datetime",
				dialect.Postgres: "time with time zone",
			}),
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return nil
}
