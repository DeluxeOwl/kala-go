package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tuple holds the schema definition for the Tuple entity.
type Tuple struct {
	ent.Schema
}

// Fields of the Tuple.
func (Tuple) Fields() []ent.Field {
	return []ent.Field{
		// TODO: ids or names?
		field.Int("subject_id"),
		field.Int("relation_id"),
		field.Int("resource_id"),
	}
}

// Edges of the Tuple.
func (Tuple) Edges() []ent.Edge {
	// TODO: should these be unique?
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
