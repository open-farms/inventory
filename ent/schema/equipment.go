package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Annotations(entproto.Field(2)),
		field.String("condition").Annotations(entproto.Field(3)),
		field.Time("create_time").
			Default(time.Now).
			UpdateDefault(time.Now).
			Immutable().Annotations(entproto.Field(4)),
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now).Annotations(entproto.Field(5)),
	}
}

func (Equipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("administered", Category.Type).
			Ref("admin").
			Annotations(entproto.Field(5)),
	}
}
