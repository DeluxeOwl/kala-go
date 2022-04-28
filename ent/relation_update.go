// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DeluxeOwl/kala-go/ent/permission"
	"github.com/DeluxeOwl/kala-go/ent/predicate"
	"github.com/DeluxeOwl/kala-go/ent/relation"
	"github.com/DeluxeOwl/kala-go/ent/subject"
	"github.com/DeluxeOwl/kala-go/ent/typeconfig"
)

// RelationUpdate is the builder for updating Relation entities.
type RelationUpdate struct {
	config
	hooks    []Hook
	mutation *RelationMutation
}

// Where appends a list predicates to the RelationUpdate builder.
func (ru *RelationUpdate) Where(ps ...predicate.Relation) *RelationUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RelationUpdate) SetName(s string) *RelationUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetValue sets the "value" field.
func (ru *RelationUpdate) SetValue(s string) *RelationUpdate {
	ru.mutation.SetValue(s)
	return ru
}

// AddSubjectIDs adds the "subjects" edge to the Subject entity by IDs.
func (ru *RelationUpdate) AddSubjectIDs(ids ...int) *RelationUpdate {
	ru.mutation.AddSubjectIDs(ids...)
	return ru
}

// AddSubjects adds the "subjects" edges to the Subject entity.
func (ru *RelationUpdate) AddSubjects(s ...*Subject) *RelationUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.AddSubjectIDs(ids...)
}

// AddRelTypeconfigIDs adds the "rel_typeconfigs" edge to the TypeConfig entity by IDs.
func (ru *RelationUpdate) AddRelTypeconfigIDs(ids ...int) *RelationUpdate {
	ru.mutation.AddRelTypeconfigIDs(ids...)
	return ru
}

// AddRelTypeconfigs adds the "rel_typeconfigs" edges to the TypeConfig entity.
func (ru *RelationUpdate) AddRelTypeconfigs(t ...*TypeConfig) *RelationUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.AddRelTypeconfigIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (ru *RelationUpdate) AddPermissionIDs(ids ...int) *RelationUpdate {
	ru.mutation.AddPermissionIDs(ids...)
	return ru
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (ru *RelationUpdate) AddPermissions(p ...*Permission) *RelationUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPermissionIDs(ids...)
}

// SetTypeconfigID sets the "typeconfig" edge to the TypeConfig entity by ID.
func (ru *RelationUpdate) SetTypeconfigID(id int) *RelationUpdate {
	ru.mutation.SetTypeconfigID(id)
	return ru
}

// SetNillableTypeconfigID sets the "typeconfig" edge to the TypeConfig entity by ID if the given value is not nil.
func (ru *RelationUpdate) SetNillableTypeconfigID(id *int) *RelationUpdate {
	if id != nil {
		ru = ru.SetTypeconfigID(*id)
	}
	return ru
}

// SetTypeconfig sets the "typeconfig" edge to the TypeConfig entity.
func (ru *RelationUpdate) SetTypeconfig(t *TypeConfig) *RelationUpdate {
	return ru.SetTypeconfigID(t.ID)
}

// Mutation returns the RelationMutation object of the builder.
func (ru *RelationUpdate) Mutation() *RelationMutation {
	return ru.mutation
}

// ClearSubjects clears all "subjects" edges to the Subject entity.
func (ru *RelationUpdate) ClearSubjects() *RelationUpdate {
	ru.mutation.ClearSubjects()
	return ru
}

// RemoveSubjectIDs removes the "subjects" edge to Subject entities by IDs.
func (ru *RelationUpdate) RemoveSubjectIDs(ids ...int) *RelationUpdate {
	ru.mutation.RemoveSubjectIDs(ids...)
	return ru
}

// RemoveSubjects removes "subjects" edges to Subject entities.
func (ru *RelationUpdate) RemoveSubjects(s ...*Subject) *RelationUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.RemoveSubjectIDs(ids...)
}

// ClearRelTypeconfigs clears all "rel_typeconfigs" edges to the TypeConfig entity.
func (ru *RelationUpdate) ClearRelTypeconfigs() *RelationUpdate {
	ru.mutation.ClearRelTypeconfigs()
	return ru
}

// RemoveRelTypeconfigIDs removes the "rel_typeconfigs" edge to TypeConfig entities by IDs.
func (ru *RelationUpdate) RemoveRelTypeconfigIDs(ids ...int) *RelationUpdate {
	ru.mutation.RemoveRelTypeconfigIDs(ids...)
	return ru
}

// RemoveRelTypeconfigs removes "rel_typeconfigs" edges to TypeConfig entities.
func (ru *RelationUpdate) RemoveRelTypeconfigs(t ...*TypeConfig) *RelationUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.RemoveRelTypeconfigIDs(ids...)
}

// ClearPermissions clears all "permissions" edges to the Permission entity.
func (ru *RelationUpdate) ClearPermissions() *RelationUpdate {
	ru.mutation.ClearPermissions()
	return ru
}

// RemovePermissionIDs removes the "permissions" edge to Permission entities by IDs.
func (ru *RelationUpdate) RemovePermissionIDs(ids ...int) *RelationUpdate {
	ru.mutation.RemovePermissionIDs(ids...)
	return ru
}

// RemovePermissions removes "permissions" edges to Permission entities.
func (ru *RelationUpdate) RemovePermissions(p ...*Permission) *RelationUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePermissionIDs(ids...)
}

// ClearTypeconfig clears the "typeconfig" edge to the TypeConfig entity.
func (ru *RelationUpdate) ClearTypeconfig() *RelationUpdate {
	ru.mutation.ClearTypeconfig()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RelationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RelationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RelationUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RelationUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RelationUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RelationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   relation.Table,
			Columns: relation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: relation.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relation.FieldName,
		})
	}
	if value, ok := ru.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relation.FieldValue,
		})
	}
	if ru.mutation.SubjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedSubjectsIDs(); len(nodes) > 0 && !ru.mutation.SubjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.SubjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RelTypeconfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relation.RelTypeconfigsTable,
			Columns: []string{relation.RelTypeconfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typeconfig.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRelTypeconfigsIDs(); len(nodes) > 0 && !ru.mutation.RelTypeconfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relation.RelTypeconfigsTable,
			Columns: []string{relation.RelTypeconfigsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RelTypeconfigsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relation.RelTypeconfigsTable,
			Columns: []string{relation.RelTypeconfigsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.PermissionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !ru.mutation.PermissionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PermissionsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.TypeconfigCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.TypeconfigIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{relation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RelationUpdateOne is the builder for updating a single Relation entity.
type RelationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RelationMutation
}

// SetName sets the "name" field.
func (ruo *RelationUpdateOne) SetName(s string) *RelationUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetValue sets the "value" field.
func (ruo *RelationUpdateOne) SetValue(s string) *RelationUpdateOne {
	ruo.mutation.SetValue(s)
	return ruo
}

// AddSubjectIDs adds the "subjects" edge to the Subject entity by IDs.
func (ruo *RelationUpdateOne) AddSubjectIDs(ids ...int) *RelationUpdateOne {
	ruo.mutation.AddSubjectIDs(ids...)
	return ruo
}

// AddSubjects adds the "subjects" edges to the Subject entity.
func (ruo *RelationUpdateOne) AddSubjects(s ...*Subject) *RelationUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.AddSubjectIDs(ids...)
}

// AddRelTypeconfigIDs adds the "rel_typeconfigs" edge to the TypeConfig entity by IDs.
func (ruo *RelationUpdateOne) AddRelTypeconfigIDs(ids ...int) *RelationUpdateOne {
	ruo.mutation.AddRelTypeconfigIDs(ids...)
	return ruo
}

// AddRelTypeconfigs adds the "rel_typeconfigs" edges to the TypeConfig entity.
func (ruo *RelationUpdateOne) AddRelTypeconfigs(t ...*TypeConfig) *RelationUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.AddRelTypeconfigIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (ruo *RelationUpdateOne) AddPermissionIDs(ids ...int) *RelationUpdateOne {
	ruo.mutation.AddPermissionIDs(ids...)
	return ruo
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (ruo *RelationUpdateOne) AddPermissions(p ...*Permission) *RelationUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPermissionIDs(ids...)
}

// SetTypeconfigID sets the "typeconfig" edge to the TypeConfig entity by ID.
func (ruo *RelationUpdateOne) SetTypeconfigID(id int) *RelationUpdateOne {
	ruo.mutation.SetTypeconfigID(id)
	return ruo
}

// SetNillableTypeconfigID sets the "typeconfig" edge to the TypeConfig entity by ID if the given value is not nil.
func (ruo *RelationUpdateOne) SetNillableTypeconfigID(id *int) *RelationUpdateOne {
	if id != nil {
		ruo = ruo.SetTypeconfigID(*id)
	}
	return ruo
}

// SetTypeconfig sets the "typeconfig" edge to the TypeConfig entity.
func (ruo *RelationUpdateOne) SetTypeconfig(t *TypeConfig) *RelationUpdateOne {
	return ruo.SetTypeconfigID(t.ID)
}

// Mutation returns the RelationMutation object of the builder.
func (ruo *RelationUpdateOne) Mutation() *RelationMutation {
	return ruo.mutation
}

// ClearSubjects clears all "subjects" edges to the Subject entity.
func (ruo *RelationUpdateOne) ClearSubjects() *RelationUpdateOne {
	ruo.mutation.ClearSubjects()
	return ruo
}

// RemoveSubjectIDs removes the "subjects" edge to Subject entities by IDs.
func (ruo *RelationUpdateOne) RemoveSubjectIDs(ids ...int) *RelationUpdateOne {
	ruo.mutation.RemoveSubjectIDs(ids...)
	return ruo
}

// RemoveSubjects removes "subjects" edges to Subject entities.
func (ruo *RelationUpdateOne) RemoveSubjects(s ...*Subject) *RelationUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.RemoveSubjectIDs(ids...)
}

// ClearRelTypeconfigs clears all "rel_typeconfigs" edges to the TypeConfig entity.
func (ruo *RelationUpdateOne) ClearRelTypeconfigs() *RelationUpdateOne {
	ruo.mutation.ClearRelTypeconfigs()
	return ruo
}

// RemoveRelTypeconfigIDs removes the "rel_typeconfigs" edge to TypeConfig entities by IDs.
func (ruo *RelationUpdateOne) RemoveRelTypeconfigIDs(ids ...int) *RelationUpdateOne {
	ruo.mutation.RemoveRelTypeconfigIDs(ids...)
	return ruo
}

// RemoveRelTypeconfigs removes "rel_typeconfigs" edges to TypeConfig entities.
func (ruo *RelationUpdateOne) RemoveRelTypeconfigs(t ...*TypeConfig) *RelationUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.RemoveRelTypeconfigIDs(ids...)
}

// ClearPermissions clears all "permissions" edges to the Permission entity.
func (ruo *RelationUpdateOne) ClearPermissions() *RelationUpdateOne {
	ruo.mutation.ClearPermissions()
	return ruo
}

// RemovePermissionIDs removes the "permissions" edge to Permission entities by IDs.
func (ruo *RelationUpdateOne) RemovePermissionIDs(ids ...int) *RelationUpdateOne {
	ruo.mutation.RemovePermissionIDs(ids...)
	return ruo
}

// RemovePermissions removes "permissions" edges to Permission entities.
func (ruo *RelationUpdateOne) RemovePermissions(p ...*Permission) *RelationUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePermissionIDs(ids...)
}

// ClearTypeconfig clears the "typeconfig" edge to the TypeConfig entity.
func (ruo *RelationUpdateOne) ClearTypeconfig() *RelationUpdateOne {
	ruo.mutation.ClearTypeconfig()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RelationUpdateOne) Select(field string, fields ...string) *RelationUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Relation entity.
func (ruo *RelationUpdateOne) Save(ctx context.Context) (*Relation, error) {
	var (
		err  error
		node *Relation
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RelationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RelationUpdateOne) SaveX(ctx context.Context) *Relation {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RelationUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RelationUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RelationUpdateOne) sqlSave(ctx context.Context) (_node *Relation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   relation.Table,
			Columns: relation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: relation.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Relation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, relation.FieldID)
		for _, f := range fields {
			if !relation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != relation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relation.FieldName,
		})
	}
	if value, ok := ruo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: relation.FieldValue,
		})
	}
	if ruo.mutation.SubjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedSubjectsIDs(); len(nodes) > 0 && !ruo.mutation.SubjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.SubjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RelTypeconfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relation.RelTypeconfigsTable,
			Columns: []string{relation.RelTypeconfigsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typeconfig.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRelTypeconfigsIDs(); len(nodes) > 0 && !ruo.mutation.RelTypeconfigsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relation.RelTypeconfigsTable,
			Columns: []string{relation.RelTypeconfigsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RelTypeconfigsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   relation.RelTypeconfigsTable,
			Columns: []string{relation.RelTypeconfigsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.PermissionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !ruo.mutation.PermissionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PermissionsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.TypeconfigCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.TypeconfigIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Relation{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{relation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
