package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TypeConfig holds the schema definition for the TypeConfig entity.
type TypeConfig struct {
	ent.Schema
}

type Relations map[string]string
type Permissions map[string]string

// Fields of the TypeConfig.
// TODO: validation
func (TypeConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.JSON("relations", &Relations{}),
		field.JSON("permissions", &Permissions{}),
	}
}

// Edges of the TypeConfig.
func (TypeConfig) Edges() []ent.Edge {
	return nil
}
