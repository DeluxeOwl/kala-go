package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return []ent.Field{
		// Has a unique name, TODO: should it be unique?
		field.String("name").
			Unique(),
	}
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return []ent.Edge{
		// Has a unique type
		edge.From("type", TypeConfig.Type).
			Ref("subjects").
			Unique(),
		// The subject has the following relations
		edge.From("relations", Relation.Type).
			Ref("subjects"),
	}
}
