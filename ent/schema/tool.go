package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Tool holds the schema definition for the Tool entity.
type Tool struct {
	ent.Schema
}

func (Tool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Tool.
func (Tool) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Annotations(entproto.Field(4)),
		field.Bool("powered").
			Default(false).
			Annotations(entproto.Field(5)),
	}
}

// Edges of the Tool.
func (Tool) Edges() []ent.Edge {
	return nil
}

func (Tool) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
