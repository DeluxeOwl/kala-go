package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("value"),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {

	return []ent.Edge{
		// The typeconfig has these permissions
		edge.From("typeconfig", TypeConfig.Type).
			Ref("permissions").
			Unique(),
		// Points to the relations in the value field for easier traversal
		edge.To("relations", Relation.Type),
	}
}
