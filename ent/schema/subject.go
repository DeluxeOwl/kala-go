package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return []ent.Field{
		// unique would mean across all subjects
		field.String("name"),
	}
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return []ent.Edge{
		// Has a unique type
		edge.From("type", TypeConfig.Type).
			Ref("subjects").
			Unique(),
	}
}

// Indexes of the Subject.
func (Subject) Indexes() []ent.Index {
	// subjects are unique under a type
	return []ent.Index{
		index.Fields("name").
			Edges("type").
			Unique(),
	}
}
