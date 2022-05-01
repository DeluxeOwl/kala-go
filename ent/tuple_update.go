// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DeluxeOwl/kala-go/ent/predicate"
	"github.com/DeluxeOwl/kala-go/ent/relation"
	"github.com/DeluxeOwl/kala-go/ent/subject"
	"github.com/DeluxeOwl/kala-go/ent/tuple"
)

// TupleUpdate is the builder for updating Tuple entities.
type TupleUpdate struct {
	config
	hooks    []Hook
	mutation *TupleMutation
}

// Where appends a list predicates to the TupleUpdate builder.
func (tu *TupleUpdate) Where(ps ...predicate.Tuple) *TupleUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetSubjectID sets the "subject_id" field.
func (tu *TupleUpdate) SetSubjectID(i int) *TupleUpdate {
	tu.mutation.SetSubjectID(i)
	return tu
}

// SetRelationID sets the "relation_id" field.
func (tu *TupleUpdate) SetRelationID(i int) *TupleUpdate {
	tu.mutation.SetRelationID(i)
	return tu
}

// SetResourceID sets the "resource_id" field.
func (tu *TupleUpdate) SetResourceID(i int) *TupleUpdate {
	tu.mutation.SetResourceID(i)
	return tu
}

// SetSubject sets the "subject" edge to the Subject entity.
func (tu *TupleUpdate) SetSubject(s *Subject) *TupleUpdate {
	return tu.SetSubjectID(s.ID)
}

// SetRelation sets the "relation" edge to the Relation entity.
func (tu *TupleUpdate) SetRelation(r *Relation) *TupleUpdate {
	return tu.SetRelationID(r.ID)
}

// SetResource sets the "resource" edge to the Subject entity.
func (tu *TupleUpdate) SetResource(s *Subject) *TupleUpdate {
	return tu.SetResourceID(s.ID)
}

// Mutation returns the TupleMutation object of the builder.
func (tu *TupleUpdate) Mutation() *TupleMutation {
	return tu.mutation
}

// ClearSubject clears the "subject" edge to the Subject entity.
func (tu *TupleUpdate) ClearSubject() *TupleUpdate {
	tu.mutation.ClearSubject()
	return tu
}

// ClearRelation clears the "relation" edge to the Relation entity.
func (tu *TupleUpdate) ClearRelation() *TupleUpdate {
	tu.mutation.ClearRelation()
	return tu
}

// ClearResource clears the "resource" edge to the Subject entity.
func (tu *TupleUpdate) ClearResource() *TupleUpdate {
	tu.mutation.ClearResource()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TupleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TupleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TupleUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TupleUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TupleUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TupleUpdate) check() error {
	if _, ok := tu.mutation.SubjectID(); tu.mutation.SubjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tuple.subject"`)
	}
	if _, ok := tu.mutation.RelationID(); tu.mutation.RelationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tuple.relation"`)
	}
	if _, ok := tu.mutation.ResourceID(); tu.mutation.ResourceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tuple.resource"`)
	}
	return nil
}

func (tu *TupleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tuple.Table,
			Columns: tuple.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tuple.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tu.mutation.SubjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.SubjectTable,
			Columns: []string{tuple.SubjectColumn},
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
	if nodes := tu.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.SubjectTable,
			Columns: []string{tuple.SubjectColumn},
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
	if tu.mutation.RelationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.RelationTable,
			Columns: []string{tuple.RelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RelationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.RelationTable,
			Columns: []string{tuple.RelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ResourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.ResourceTable,
			Columns: []string{tuple.ResourceColumn},
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
	if nodes := tu.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.ResourceTable,
			Columns: []string{tuple.ResourceColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tuple.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TupleUpdateOne is the builder for updating a single Tuple entity.
type TupleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TupleMutation
}

// SetSubjectID sets the "subject_id" field.
func (tuo *TupleUpdateOne) SetSubjectID(i int) *TupleUpdateOne {
	tuo.mutation.SetSubjectID(i)
	return tuo
}

// SetRelationID sets the "relation_id" field.
func (tuo *TupleUpdateOne) SetRelationID(i int) *TupleUpdateOne {
	tuo.mutation.SetRelationID(i)
	return tuo
}

// SetResourceID sets the "resource_id" field.
func (tuo *TupleUpdateOne) SetResourceID(i int) *TupleUpdateOne {
	tuo.mutation.SetResourceID(i)
	return tuo
}

// SetSubject sets the "subject" edge to the Subject entity.
func (tuo *TupleUpdateOne) SetSubject(s *Subject) *TupleUpdateOne {
	return tuo.SetSubjectID(s.ID)
}

// SetRelation sets the "relation" edge to the Relation entity.
func (tuo *TupleUpdateOne) SetRelation(r *Relation) *TupleUpdateOne {
	return tuo.SetRelationID(r.ID)
}

// SetResource sets the "resource" edge to the Subject entity.
func (tuo *TupleUpdateOne) SetResource(s *Subject) *TupleUpdateOne {
	return tuo.SetResourceID(s.ID)
}

// Mutation returns the TupleMutation object of the builder.
func (tuo *TupleUpdateOne) Mutation() *TupleMutation {
	return tuo.mutation
}

// ClearSubject clears the "subject" edge to the Subject entity.
func (tuo *TupleUpdateOne) ClearSubject() *TupleUpdateOne {
	tuo.mutation.ClearSubject()
	return tuo
}

// ClearRelation clears the "relation" edge to the Relation entity.
func (tuo *TupleUpdateOne) ClearRelation() *TupleUpdateOne {
	tuo.mutation.ClearRelation()
	return tuo
}

// ClearResource clears the "resource" edge to the Subject entity.
func (tuo *TupleUpdateOne) ClearResource() *TupleUpdateOne {
	tuo.mutation.ClearResource()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TupleUpdateOne) Select(field string, fields ...string) *TupleUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tuple entity.
func (tuo *TupleUpdateOne) Save(ctx context.Context) (*Tuple, error) {
	var (
		err  error
		node *Tuple
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TupleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TupleUpdateOne) SaveX(ctx context.Context) *Tuple {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TupleUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TupleUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TupleUpdateOne) check() error {
	if _, ok := tuo.mutation.SubjectID(); tuo.mutation.SubjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tuple.subject"`)
	}
	if _, ok := tuo.mutation.RelationID(); tuo.mutation.RelationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tuple.relation"`)
	}
	if _, ok := tuo.mutation.ResourceID(); tuo.mutation.ResourceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Tuple.resource"`)
	}
	return nil
}

func (tuo *TupleUpdateOne) sqlSave(ctx context.Context) (_node *Tuple, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tuple.Table,
			Columns: tuple.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tuple.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tuple.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tuple.FieldID)
		for _, f := range fields {
			if !tuple.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tuple.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tuo.mutation.SubjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.SubjectTable,
			Columns: []string{tuple.SubjectColumn},
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
	if nodes := tuo.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.SubjectTable,
			Columns: []string{tuple.SubjectColumn},
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
	if tuo.mutation.RelationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.RelationTable,
			Columns: []string{tuple.RelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RelationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.RelationTable,
			Columns: []string{tuple.RelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ResourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.ResourceTable,
			Columns: []string{tuple.ResourceColumn},
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
	if nodes := tuo.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tuple.ResourceTable,
			Columns: []string{tuple.ResourceColumn},
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
	_node = &Tuple{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tuple.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
