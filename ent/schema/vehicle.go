package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Vehicle holds the schema definition for the Vehicle entity.
type Vehicle struct {
	ent.Schema
}

// Fields of the Vehicle.
func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("make").Optional(),
		field.String("model").Optional(),
		field.Int64("miles").Optional(),
		field.Int64("mpg").Optional(),
		field.String("owner").Optional(),
		field.String("year").Optional(),
		field.Bool("active").Optional(),
		field.Strings("tags").Optional(),
		field.Enum("condition").Optional().
			Values(
				"UNSPECIFIED",
				"MINT",
				"GOOD",
				"POOR",
				"BROKEN",
			),
		field.Time("create_time").
			Default(time.Now).
			UpdateDefault(time.Now).
			Immutable(),
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return nil
}
