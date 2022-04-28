package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TypeConfig holds the schema definition for the TypeConfig entity.
type TypeConfig struct {
	ent.Schema
}

// Fields of the TypeConfig.
// TODO: validation
func (TypeConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the TypeConfig.
func (TypeConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("relations", Relation.Type),
		edge.To("permissions", Permission.Type),
		edge.To("subjects", Subject.Type),
	}
}
