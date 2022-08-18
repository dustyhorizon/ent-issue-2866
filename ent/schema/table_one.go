package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// One holds the schema definition for the One entity.
type One struct {
	ent.Schema
}

func (One) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "table_one"},
	}
}

// Fields of the One.
func (One) Fields() []ent.Field {
	return []ent.Field{
		field.String("something").
			Unique().
			NotEmpty(),
	}
}

// Edges of the One.
func (One) Edges() []ent.Edge {
	return nil
}

func (One) Indexes() []ent.Index {
	return nil
}

func (One) Mixin() []ent.Mixin {
	return nil
}

func (One) Hooks() []ent.Hook {
	return nil
}
