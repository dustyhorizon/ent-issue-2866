package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Two holds the schema definition for the Two entity.
type Two struct {
	ent.Schema
}

func (Two) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "table_two"},
	}
}

// Fields of the Two.
func (Two) Fields() []ent.Field {
	return []ent.Field{
		field.String("something").
			Unique().
			NotEmpty(),
	}
}

// Edges of the Two.
func (Two) Edges() []ent.Edge {
	return nil
}

func (Two) Indexes() []ent.Index {
	return nil
}

func (Two) Mixin() []ent.Mixin {
	return nil
}

func (Two) Hooks() []ent.Hook {
	return nil
}
