// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DeluxeOwl/kala-go/ent/permission"
	"github.com/DeluxeOwl/kala-go/ent/relation"
	"github.com/DeluxeOwl/kala-go/ent/subject"
	"github.com/DeluxeOwl/kala-go/ent/typeconfig"
)

// RelationCreate is the builder for creating a Relation entity.
type RelationCreate struct {
	config
	mutation *RelationMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *RelationCreate) SetName(s string) *RelationCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetValue sets the "value" field.
func (rc *RelationCreate) SetValue(s string) *RelationCreate {
	rc.mutation.SetValue(s)
	return rc
}

// AddSubjectIDs adds the "subjects" edge to the Subject entity by IDs.
func (rc *RelationCreate) AddSubjectIDs(ids ...int) *RelationCreate {
	rc.mutation.AddSubjectIDs(ids...)
	return rc
}

// AddSubjects adds the "subjects" edges to the Subject entity.
func (rc *RelationCreate) AddSubjects(s ...*Subject) *RelationCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return rc.AddSubjectIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (rc *RelationCreate) AddPermissionIDs(ids ...int) *RelationCreate {
	rc.mutation.AddPermissionIDs(ids...)
	return rc
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (rc *RelationCreate) AddPermissions(p ...*Permission) *RelationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPermissionIDs(ids...)
}

// SetTypeconfigID sets the "typeconfig" edge to the TypeConfig entity by ID.
func (rc *RelationCreate) SetTypeconfigID(id int) *RelationCreate {
	rc.mutation.SetTypeconfigID(id)
	return rc
}

// SetNillableTypeconfigID sets the "typeconfig" edge to the TypeConfig entity by ID if the given value is not nil.
func (rc *RelationCreate) SetNillableTypeconfigID(id *int) *RelationCreate {
	if id != nil {
		rc = rc.SetTypeconfigID(*id)
	}
	return rc
}

// SetTypeconfig sets the "typeconfig" edge to the TypeConfig entity.
func (rc *RelationCreate) SetTypeconfig(t *TypeConfig) *RelationCreate {
	return rc.SetTypeconfigID(t.ID)
}

// Mutation returns the RelationMutation object of the builder.
func (rc *RelationCreate) Mutation() *RelationMutation {
	return rc.mutation
}

// Save creates the Relation in the database.
func (rc *RelationCreate) Save(ctx context.Context) (*Relation, error) {
	var (
		err  error
		node *Relation
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RelationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RelationCreate) SaveX(ctx context.Context) *Relation {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RelationCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RelationCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RelationCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Relation.name"`)}
	}
	if _, ok := rc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Relation.value"`)}
	}
	return nil
}

func (rc *RelationCreate) sqlSave(ctx context.Context) (*Relation, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RelationCreate) createSpec() (*Relation, *sqlgraph.CreateSpec) {
	var (
		_node = &Relation{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: relation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: relation.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relation.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relation.FieldValue,
		})
		_node.Value = value
	}
	if nodes := rc.mutation.SubjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   relation.SubjectsTable,
			Columns: relation.SubjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   relation.PermissionsTable,
			Columns: relation.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.TypeconfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   relation.TypeconfigTable,
			Columns: []string{relation.TypeconfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typeconfig.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.type_config_relations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RelationCreateBulk is the builder for creating many Relation entities in bulk.
type RelationCreateBulk struct {
	config
	builders []*RelationCreate
}

// Save creates the Relation entities in the database.
func (rcb *RelationCreateBulk) Save(ctx context.Context) ([]*Relation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Relation, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RelationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RelationCreateBulk) SaveX(ctx context.Context) []*Relation {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RelationCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RelationCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
