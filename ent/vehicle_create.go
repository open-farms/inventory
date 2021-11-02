// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-farms/inventory/ent/vehicle"
)

// VehicleCreate is the builder for creating a Vehicle entity.
type VehicleCreate struct {
	config
	mutation *VehicleMutation
	hooks    []Hook
}

// SetMake sets the "make" field.
func (vc *VehicleCreate) SetMake(s string) *VehicleCreate {
	vc.mutation.SetMake(s)
	return vc
}

// SetModel sets the "model" field.
func (vc *VehicleCreate) SetModel(s string) *VehicleCreate {
	vc.mutation.SetModel(s)
	return vc
}

// SetYear sets the "year" field.
func (vc *VehicleCreate) SetYear(s string) *VehicleCreate {
	vc.mutation.SetYear(s)
	return vc
}

// SetTags sets the "tags" field.
func (vc *VehicleCreate) SetTags(s []string) *VehicleCreate {
	vc.mutation.SetTags(s)
	return vc
}

// SetCondition sets the "condition" field.
func (vc *VehicleCreate) SetCondition(v vehicle.Condition) *VehicleCreate {
	vc.mutation.SetCondition(v)
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VehicleCreate) SetCreatedAt(t time.Time) *VehicleCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VehicleCreate) SetNillableCreatedAt(t *time.Time) *VehicleCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VehicleCreate) SetUpdatedAt(t time.Time) *VehicleCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VehicleCreate) SetNillableUpdatedAt(t *time.Time) *VehicleCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetID sets the "id" field.
func (vc *VehicleCreate) SetID(u uint64) *VehicleCreate {
	vc.mutation.SetID(u)
	return vc
}

// Mutation returns the VehicleMutation object of the builder.
func (vc *VehicleCreate) Mutation() *VehicleMutation {
	return vc.mutation
}

// Save creates the Vehicle in the database.
func (vc *VehicleCreate) Save(ctx context.Context) (*Vehicle, error) {
	var (
		err  error
		node *Vehicle
	)
	vc.defaults()
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VehicleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VehicleCreate) SaveX(ctx context.Context) *Vehicle {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VehicleCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VehicleCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VehicleCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := vehicle.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := vehicle.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VehicleCreate) check() error {
	if _, ok := vc.mutation.Make(); !ok {
		return &ValidationError{Name: "make", err: errors.New(`ent: missing required field "make"`)}
	}
	if _, ok := vc.mutation.Model(); !ok {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required field "model"`)}
	}
	if _, ok := vc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "year"`)}
	}
	if _, ok := vc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`ent: missing required field "tags"`)}
	}
	if _, ok := vc.mutation.Condition(); !ok {
		return &ValidationError{Name: "condition", err: errors.New(`ent: missing required field "condition"`)}
	}
	if v, ok := vc.mutation.Condition(); ok {
		if err := vehicle.ConditionValidator(v); err != nil {
			return &ValidationError{Name: "condition", err: fmt.Errorf(`ent: validator failed for field "condition": %w`, err)}
		}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if v, ok := vc.mutation.ID(); ok {
		if err := vehicle.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "id": %w`, err)}
		}
	}
	return nil
}

func (vc *VehicleCreate) sqlSave(ctx context.Context) (*Vehicle, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (vc *VehicleCreate) createSpec() (*Vehicle, *sqlgraph.CreateSpec) {
	var (
		_node = &Vehicle{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vehicle.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: vehicle.FieldID,
			},
		}
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.Make(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldMake,
		})
		_node.Make = value
	}
	if value, ok := vc.mutation.Model(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldModel,
		})
		_node.Model = value
	}
	if value, ok := vc.mutation.Year(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vehicle.FieldYear,
		})
		_node.Year = value
	}
	if value, ok := vc.mutation.Tags(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: vehicle.FieldTags,
		})
		_node.Tags = value
	}
	if value, ok := vc.mutation.Condition(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: vehicle.FieldCondition,
		})
		_node.Condition = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicle.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vehicle.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// VehicleCreateBulk is the builder for creating many Vehicle entities in bulk.
type VehicleCreateBulk struct {
	config
	builders []*VehicleCreate
}

// Save creates the Vehicle entities in the database.
func (vcb *VehicleCreateBulk) Save(ctx context.Context) ([]*Vehicle, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vehicle, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VehicleMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VehicleCreateBulk) SaveX(ctx context.Context) []*Vehicle {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VehicleCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VehicleCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
