// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-farms/inventory/ent/predicate"
	"github.com/open-farms/inventory/ent/vehicle"
)

// VehicleUpdate is the builder for updating Vehicle entities.
type VehicleUpdate struct {
	config
	hooks    []Hook
	mutation *VehicleMutation
}

// Where appends a list predicates to the VehicleUpdate builder.
func (vu *VehicleUpdate) Where(ps ...predicate.Vehicle) *VehicleUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetMake sets the "make" field.
func (vu *VehicleUpdate) SetMake(s string) *VehicleUpdate {
	vu.mutation.SetMake(s)
	return vu
}

// SetModel sets the "model" field.
func (vu *VehicleUpdate) SetModel(s string) *VehicleUpdate {
	vu.mutation.SetModel(s)
	return vu
}

// SetYear sets the "year" field.
func (vu *VehicleUpdate) SetYear(s string) *VehicleUpdate {
	vu.mutation.SetYear(s)
	return vu
}

// SetTags sets the "tags" field.
func (vu *VehicleUpdate) SetTags(s []string) *VehicleUpdate {
	vu.mutation.SetTags(s)
	return vu
}

// SetCondition sets the "condition" field.
func (vu *VehicleUpdate) SetCondition(v vehicle.Condition) *VehicleUpdate {
	vu.mutation.SetCondition(v)
	return vu
}

// SetCreatedAt sets the "created_at" field.
func (vu *VehicleUpdate) SetCreatedAt(t time.Time) *VehicleUpdate {
	vu.mutation.SetCreatedAt(t)
	return vu
}

// SetUpdatedAt sets the "updated_at" field.
func (vu *VehicleUpdate) SetUpdatedAt(t time.Time) *VehicleUpdate {
	vu.mutation.SetUpdatedAt(t)
	return vu
}

// Mutation returns the VehicleMutation object of the builder.
func (vu *VehicleUpdate) Mutation() *VehicleMutation {
	return vu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VehicleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vu.hooks) == 0 {
		if err = vu.check(); err != nil {
			return 0, err
		}
		affected, err = vu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vu.check(); err != nil {
				return 0, err
			}
			vu.mutation = mutation
			affected, err = vu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vu.hooks) - 1; i >= 0; i-- {
			if vu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VehicleUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VehicleUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VehicleUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VehicleUpdate) check() error {
	if v, ok := vu.mutation.Condition(); ok {
		if err := vehicle.ConditionValidator(v); err != nil {
			return &ValidationError{Name: "condition", err: fmt.Errorf("ent: validator failed for field \"condition\": %w", err)}
		}
	}
	return nil
}

func (vu *VehicleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vehicle.Table,
			Columns: vehicle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: vehicle.FieldID,
			},
		},
	}
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Make(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldMake,
		})
	}
	if value, ok := vu.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldModel,
		})
	}
	if value, ok := vu.mutation.Year(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldYear,
		})
	}
	if value, ok := vu.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: vehicle.FieldTags,
		})
	}
	if value, ok := vu.mutation.Condition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: vehicle.FieldCondition,
		})
	}
	if value, ok := vu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicle.FieldCreatedAt,
		})
	}
	if value, ok := vu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicle.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vehicle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// VehicleUpdateOne is the builder for updating a single Vehicle entity.
type VehicleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VehicleMutation
}

// SetMake sets the "make" field.
func (vuo *VehicleUpdateOne) SetMake(s string) *VehicleUpdateOne {
	vuo.mutation.SetMake(s)
	return vuo
}

// SetModel sets the "model" field.
func (vuo *VehicleUpdateOne) SetModel(s string) *VehicleUpdateOne {
	vuo.mutation.SetModel(s)
	return vuo
}

// SetYear sets the "year" field.
func (vuo *VehicleUpdateOne) SetYear(s string) *VehicleUpdateOne {
	vuo.mutation.SetYear(s)
	return vuo
}

// SetTags sets the "tags" field.
func (vuo *VehicleUpdateOne) SetTags(s []string) *VehicleUpdateOne {
	vuo.mutation.SetTags(s)
	return vuo
}

// SetCondition sets the "condition" field.
func (vuo *VehicleUpdateOne) SetCondition(v vehicle.Condition) *VehicleUpdateOne {
	vuo.mutation.SetCondition(v)
	return vuo
}

// SetCreatedAt sets the "created_at" field.
func (vuo *VehicleUpdateOne) SetCreatedAt(t time.Time) *VehicleUpdateOne {
	vuo.mutation.SetCreatedAt(t)
	return vuo
}

// SetUpdatedAt sets the "updated_at" field.
func (vuo *VehicleUpdateOne) SetUpdatedAt(t time.Time) *VehicleUpdateOne {
	vuo.mutation.SetUpdatedAt(t)
	return vuo
}

// Mutation returns the VehicleMutation object of the builder.
func (vuo *VehicleUpdateOne) Mutation() *VehicleMutation {
	return vuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VehicleUpdateOne) Select(field string, fields ...string) *VehicleUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Vehicle entity.
func (vuo *VehicleUpdateOne) Save(ctx context.Context) (*Vehicle, error) {
	var (
		err  error
		node *Vehicle
	)
	if len(vuo.hooks) == 0 {
		if err = vuo.check(); err != nil {
			return nil, err
		}
		node, err = vuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuo.check(); err != nil {
				return nil, err
			}
			vuo.mutation = mutation
			node, err = vuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuo.hooks) - 1; i >= 0; i-- {
			if vuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VehicleUpdateOne) SaveX(ctx context.Context) *Vehicle {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VehicleUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VehicleUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VehicleUpdateOne) check() error {
	if v, ok := vuo.mutation.Condition(); ok {
		if err := vehicle.ConditionValidator(v); err != nil {
			return &ValidationError{Name: "condition", err: fmt.Errorf("ent: validator failed for field \"condition\": %w", err)}
		}
	}
	return nil
}

func (vuo *VehicleUpdateOne) sqlSave(ctx context.Context) (_node *Vehicle, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vehicle.Table,
			Columns: vehicle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: vehicle.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Vehicle.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vehicle.FieldID)
		for _, f := range fields {
			if !vehicle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vehicle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.Make(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldMake,
		})
	}
	if value, ok := vuo.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldModel,
		})
	}
	if value, ok := vuo.mutation.Year(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldYear,
		})
	}
	if value, ok := vuo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: vehicle.FieldTags,
		})
	}
	if value, ok := vuo.mutation.Condition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: vehicle.FieldCondition,
		})
	}
	if value, ok := vuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicle.FieldCreatedAt,
		})
	}
	if value, ok := vuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicle.FieldUpdatedAt,
		})
	}
	_node = &Vehicle{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vehicle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
