package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Implement holds the schema definition for the Implement entity.
type Implement struct {
	ent.Schema
}

func (Implement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Implement.
func (Implement) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entproto.Field(4)),
	}
}

// Edges of the Implement.
func (Implement) Edges() []ent.Edge {
	return nil
}

func (Implement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
