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
func (TypeConfig) Fields() []ent.Field {
	return []ent.Field{
		// The name of the type
		field.String("name").Unique(),
	}
}

// Edges of the TypeConfig.
func (TypeConfig) Edges() []ent.Edge {
	return []ent.Edge{
		// Points to these relations in the config
		edge.To("relations", Relation.Type),
		// and these permissions
		edge.To("permissions", Permission.Type),
		// Points to a list of subjects with this type
		edge.To("subjects", Subject.Type),
		// Points from related typeconfigs
		edge.From("rel_typeconfigs", Relation.Type).
			Ref("rel_typeconfigs"),
	}
}
