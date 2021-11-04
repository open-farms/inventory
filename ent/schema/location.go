package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

func (Location) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Annotations(entproto.Field(4)),
		field.Int32("zone").
			Unique().
			Annotations(entproto.Field(5)),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vehicle", Vehicle.Type).
			Annotations(entproto.Field(6)),
	}
}

func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
