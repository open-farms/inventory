package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Annotations(entproto.Field(4)),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("vehicles", Vehicle.Type).Annotations(entproto.Field(5)),
		// edge.To("tools", Tool.Type).Annotations(entproto.Field(6)),
		// edge.To("implements", Implement.Type).Annotations(entproto.Field(7)),
		// edge.To("equipment", Equipment.Type).Annotations(entproto.Field(8)),
	}
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
