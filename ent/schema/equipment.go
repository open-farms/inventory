package schema

import (
	"errors"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

func validateName() func(s string) error {
	return func(s string) error {
		if strings.ToLower(s) == s {
			return errors.New("name must begin with uppercase")
		}
		return nil
	}
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("name").Validate(
			validateName(),
		),
		field.Strings("tags"),
		field.Enum("condition").
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

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return nil
}
