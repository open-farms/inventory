package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Positive(),
		field.String("name").
			Default("unknown"),
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

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return nil
}
