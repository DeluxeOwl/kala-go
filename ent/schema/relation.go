package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Relation holds the schema definition for the Relation entity.
type Relation struct {
	ent.Schema
}

// Fields of the Relation.
func (Relation) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("value"),
	}
}

// Edges of the Relation.
func (Relation) Edges() []ent.Edge {
	return []ent.Edge{
		// This type has this relation
		edge.From("typeconfig", TypeConfig.Type).
			Ref("relations").
			Unique(),
		// Points to these types in the value for easier traversal
		edge.To("rel_typeconfigs", TypeConfig.Type),
		// These permissions reference this relation
		edge.From("permissions", Permission.Type).
			Ref("relations"),
		// Tuples where this relation is referenced
		edge.From("tuples", Tuple.Type).
			Ref("relation"),
	}
}
