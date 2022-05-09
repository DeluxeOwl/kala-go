package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Tuple holds the schema definition for the Tuple entity.
type Tuple struct {
	ent.Schema
}

// Fields of the Tuple.
func (Tuple) Fields() []ent.Field {
	return []ent.Field{
		field.Int("subject_id"),
		field.Int("relation_id"),
		field.Int("resource_id"),
	}
}

// Edges of the Tuple.
func (Tuple) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subject", Subject.Type).
			Field("subject_id").
			Unique().
			Required(),
		edge.To("relation", Relation.Type).
			Field("relation_id").
			Unique().
			Required(),
		edge.To("resource", Subject.Type).
			Field("resource_id").
			Unique().
			Required(),
	}
}

// Indexes of the Tuple.
func (Tuple) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("subject_id", "relation_id", "resource_id").
			Unique(),
	}
}
