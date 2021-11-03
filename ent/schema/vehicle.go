package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Vehicle holds the schema definition for the Vehicle entity.
type Vehicle struct {
	ent.Schema
}

func (Vehicle) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Vehicle.
func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.String("make").
			Annotations(entproto.Field(4)),
		field.String("model").
			Annotations(entproto.Field(5)),
		field.Int64("miles").
			Optional().
			Annotations(entproto.Field(6)),
		field.Int64("mpg").
			Optional().
			Annotations(entproto.Field(7)),
		field.String("owner").
			Optional().
			Annotations(entproto.Field(8)),
		field.String("year").
			Optional().
			Annotations(entproto.Field(9)),
		field.Bool("active").
			Optional().
			Annotations(entproto.Field(10)),
		field.String("condition").
			Optional().
			Match(conditionPattern).
			Annotations(entproto.Field(11)),
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return nil
}

func (Vehicle) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
