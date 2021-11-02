// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/open-farms/inventory/ent/equipment"
	"github.com/open-farms/inventory/ent/predicate"
	"github.com/open-farms/inventory/ent/vehicle"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEquipment = "Equipment"
	TypeVehicle   = "Vehicle"
)

// EquipmentMutation represents an operation that mutates the Equipment nodes in the graph.
type EquipmentMutation struct {
	config
	op            Op
	typ           string
	id            *uint64
	name          *string
	tags          *[]string
	condition     *equipment.Condition
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Equipment, error)
	predicates    []predicate.Equipment
}

var _ ent.Mutation = (*EquipmentMutation)(nil)

// equipmentOption allows management of the mutation configuration using functional options.
type equipmentOption func(*EquipmentMutation)

// newEquipmentMutation creates new mutation for the Equipment entity.
func newEquipmentMutation(c config, op Op, opts ...equipmentOption) *EquipmentMutation {
	m := &EquipmentMutation{
		config:        c,
		op:            op,
		typ:           TypeEquipment,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEquipmentID sets the ID field of the mutation.
func withEquipmentID(id uint64) equipmentOption {
	return func(m *EquipmentMutation) {
		var (
			err   error
			once  sync.Once
			value *Equipment
		)
		m.oldValue = func(ctx context.Context) (*Equipment, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Equipment.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEquipment sets the old Equipment of the mutation.
func withEquipment(node *Equipment) equipmentOption {
	return func(m *EquipmentMutation) {
		m.oldValue = func(context.Context) (*Equipment, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EquipmentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EquipmentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Equipment entities.
func (m *EquipmentMutation) SetID(id uint64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EquipmentMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *EquipmentMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *EquipmentMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Equipment entity.
// If the Equipment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EquipmentMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *EquipmentMutation) ResetName() {
	m.name = nil
}

// SetTags sets the "tags" field.
func (m *EquipmentMutation) SetTags(s []string) {
	m.tags = &s
}

// Tags returns the value of the "tags" field in the mutation.
func (m *EquipmentMutation) Tags() (r []string, exists bool) {
	v := m.tags
	if v == nil {
		return
	}
	return *v, true
}

// OldTags returns the old "tags" field's value of the Equipment entity.
// If the Equipment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EquipmentMutation) OldTags(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTags is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTags requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTags: %w", err)
	}
	return oldValue.Tags, nil
}

// ResetTags resets all changes to the "tags" field.
func (m *EquipmentMutation) ResetTags() {
	m.tags = nil
}

// SetCondition sets the "condition" field.
func (m *EquipmentMutation) SetCondition(e equipment.Condition) {
	m.condition = &e
}

// Condition returns the value of the "condition" field in the mutation.
func (m *EquipmentMutation) Condition() (r equipment.Condition, exists bool) {
	v := m.condition
	if v == nil {
		return
	}
	return *v, true
}

// OldCondition returns the old "condition" field's value of the Equipment entity.
// If the Equipment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EquipmentMutation) OldCondition(ctx context.Context) (v equipment.Condition, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCondition is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCondition requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCondition: %w", err)
	}
	return oldValue.Condition, nil
}

// ResetCondition resets all changes to the "condition" field.
func (m *EquipmentMutation) ResetCondition() {
	m.condition = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *EquipmentMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *EquipmentMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Equipment entity.
// If the Equipment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EquipmentMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *EquipmentMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *EquipmentMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *EquipmentMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Equipment entity.
// If the Equipment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EquipmentMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *EquipmentMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the EquipmentMutation builder.
func (m *EquipmentMutation) Where(ps ...predicate.Equipment) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *EquipmentMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Equipment).
func (m *EquipmentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EquipmentMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.name != nil {
		fields = append(fields, equipment.FieldName)
	}
	if m.tags != nil {
		fields = append(fields, equipment.FieldTags)
	}
	if m.condition != nil {
		fields = append(fields, equipment.FieldCondition)
	}
	if m.created_at != nil {
		fields = append(fields, equipment.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, equipment.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EquipmentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case equipment.FieldName:
		return m.Name()
	case equipment.FieldTags:
		return m.Tags()
	case equipment.FieldCondition:
		return m.Condition()
	case equipment.FieldCreatedAt:
		return m.CreatedAt()
	case equipment.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EquipmentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case equipment.FieldName:
		return m.OldName(ctx)
	case equipment.FieldTags:
		return m.OldTags(ctx)
	case equipment.FieldCondition:
		return m.OldCondition(ctx)
	case equipment.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case equipment.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Equipment field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EquipmentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case equipment.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case equipment.FieldTags:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTags(v)
		return nil
	case equipment.FieldCondition:
		v, ok := value.(equipment.Condition)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCondition(v)
		return nil
	case equipment.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case equipment.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Equipment field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EquipmentMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EquipmentMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EquipmentMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Equipment numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EquipmentMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EquipmentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EquipmentMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Equipment nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EquipmentMutation) ResetField(name string) error {
	switch name {
	case equipment.FieldName:
		m.ResetName()
		return nil
	case equipment.FieldTags:
		m.ResetTags()
		return nil
	case equipment.FieldCondition:
		m.ResetCondition()
		return nil
	case equipment.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case equipment.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Equipment field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EquipmentMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EquipmentMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EquipmentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EquipmentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EquipmentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EquipmentMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EquipmentMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Equipment unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EquipmentMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Equipment edge %s", name)
}

// VehicleMutation represents an operation that mutates the Vehicle nodes in the graph.
type VehicleMutation struct {
	config
	op            Op
	typ           string
	id            *uint64
	make          *string
	model         *string
	year          *string
	active        *bool
	tags          *[]string
	condition     *vehicle.Condition
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Vehicle, error)
	predicates    []predicate.Vehicle
}

var _ ent.Mutation = (*VehicleMutation)(nil)

// vehicleOption allows management of the mutation configuration using functional options.
type vehicleOption func(*VehicleMutation)

// newVehicleMutation creates new mutation for the Vehicle entity.
func newVehicleMutation(c config, op Op, opts ...vehicleOption) *VehicleMutation {
	m := &VehicleMutation{
		config:        c,
		op:            op,
		typ:           TypeVehicle,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withVehicleID sets the ID field of the mutation.
func withVehicleID(id uint64) vehicleOption {
	return func(m *VehicleMutation) {
		var (
			err   error
			once  sync.Once
			value *Vehicle
		)
		m.oldValue = func(ctx context.Context) (*Vehicle, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Vehicle.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withVehicle sets the old Vehicle of the mutation.
func withVehicle(node *Vehicle) vehicleOption {
	return func(m *VehicleMutation) {
		m.oldValue = func(context.Context) (*Vehicle, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m VehicleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m VehicleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Vehicle entities.
func (m *VehicleMutation) SetID(id uint64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *VehicleMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetMake sets the "make" field.
func (m *VehicleMutation) SetMake(s string) {
	m.make = &s
}

// Make returns the value of the "make" field in the mutation.
func (m *VehicleMutation) Make() (r string, exists bool) {
	v := m.make
	if v == nil {
		return
	}
	return *v, true
}

// OldMake returns the old "make" field's value of the Vehicle entity.
// If the Vehicle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VehicleMutation) OldMake(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMake is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMake requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMake: %w", err)
	}
	return oldValue.Make, nil
}

// ResetMake resets all changes to the "make" field.
func (m *VehicleMutation) ResetMake() {
	m.make = nil
}

// SetModel sets the "model" field.
func (m *VehicleMutation) SetModel(s string) {
	m.model = &s
}

// Model returns the value of the "model" field in the mutation.
func (m *VehicleMutation) Model() (r string, exists bool) {
	v := m.model
	if v == nil {
		return
	}
	return *v, true
}

// OldModel returns the old "model" field's value of the Vehicle entity.
// If the Vehicle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VehicleMutation) OldModel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldModel is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldModel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldModel: %w", err)
	}
	return oldValue.Model, nil
}

// ResetModel resets all changes to the "model" field.
func (m *VehicleMutation) ResetModel() {
	m.model = nil
}

// SetYear sets the "year" field.
func (m *VehicleMutation) SetYear(s string) {
	m.year = &s
}

// Year returns the value of the "year" field in the mutation.
func (m *VehicleMutation) Year() (r string, exists bool) {
	v := m.year
	if v == nil {
		return
	}
	return *v, true
}

// OldYear returns the old "year" field's value of the Vehicle entity.
// If the Vehicle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VehicleMutation) OldYear(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldYear is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldYear requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldYear: %w", err)
	}
	return oldValue.Year, nil
}

// ResetYear resets all changes to the "year" field.
func (m *VehicleMutation) ResetYear() {
	m.year = nil
}

// SetActive sets the "active" field.
func (m *VehicleMutation) SetActive(b bool) {
	m.active = &b
}

// Active returns the value of the "active" field in the mutation.
func (m *VehicleMutation) Active() (r bool, exists bool) {
	v := m.active
	if v == nil {
		return
	}
	return *v, true
}

// OldActive returns the old "active" field's value of the Vehicle entity.
// If the Vehicle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VehicleMutation) OldActive(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldActive is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldActive requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldActive: %w", err)
	}
	return oldValue.Active, nil
}

// ResetActive resets all changes to the "active" field.
func (m *VehicleMutation) ResetActive() {
	m.active = nil
}

// SetTags sets the "tags" field.
func (m *VehicleMutation) SetTags(s []string) {
	m.tags = &s
}

// Tags returns the value of the "tags" field in the mutation.
func (m *VehicleMutation) Tags() (r []string, exists bool) {
	v := m.tags
	if v == nil {
		return
	}
	return *v, true
}

// OldTags returns the old "tags" field's value of the Vehicle entity.
// If the Vehicle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VehicleMutation) OldTags(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTags is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTags requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTags: %w", err)
	}
	return oldValue.Tags, nil
}

// ResetTags resets all changes to the "tags" field.
func (m *VehicleMutation) ResetTags() {
	m.tags = nil
}

// SetCondition sets the "condition" field.
func (m *VehicleMutation) SetCondition(v vehicle.Condition) {
	m.condition = &v
}

// Condition returns the value of the "condition" field in the mutation.
func (m *VehicleMutation) Condition() (r vehicle.Condition, exists bool) {
	v := m.condition
	if v == nil {
		return
	}
	return *v, true
}

// OldCondition returns the old "condition" field's value of the Vehicle entity.
// If the Vehicle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VehicleMutation) OldCondition(ctx context.Context) (v vehicle.Condition, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCondition is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCondition requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCondition: %w", err)
	}
	return oldValue.Condition, nil
}

// ResetCondition resets all changes to the "condition" field.
func (m *VehicleMutation) ResetCondition() {
	m.condition = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *VehicleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *VehicleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Vehicle entity.
// If the Vehicle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VehicleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *VehicleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *VehicleMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *VehicleMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Vehicle entity.
// If the Vehicle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VehicleMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *VehicleMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the VehicleMutation builder.
func (m *VehicleMutation) Where(ps ...predicate.Vehicle) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *VehicleMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Vehicle).
func (m *VehicleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *VehicleMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.make != nil {
		fields = append(fields, vehicle.FieldMake)
	}
	if m.model != nil {
		fields = append(fields, vehicle.FieldModel)
	}
	if m.year != nil {
		fields = append(fields, vehicle.FieldYear)
	}
	if m.active != nil {
		fields = append(fields, vehicle.FieldActive)
	}
	if m.tags != nil {
		fields = append(fields, vehicle.FieldTags)
	}
	if m.condition != nil {
		fields = append(fields, vehicle.FieldCondition)
	}
	if m.created_at != nil {
		fields = append(fields, vehicle.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, vehicle.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *VehicleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case vehicle.FieldMake:
		return m.Make()
	case vehicle.FieldModel:
		return m.Model()
	case vehicle.FieldYear:
		return m.Year()
	case vehicle.FieldActive:
		return m.Active()
	case vehicle.FieldTags:
		return m.Tags()
	case vehicle.FieldCondition:
		return m.Condition()
	case vehicle.FieldCreatedAt:
		return m.CreatedAt()
	case vehicle.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *VehicleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case vehicle.FieldMake:
		return m.OldMake(ctx)
	case vehicle.FieldModel:
		return m.OldModel(ctx)
	case vehicle.FieldYear:
		return m.OldYear(ctx)
	case vehicle.FieldActive:
		return m.OldActive(ctx)
	case vehicle.FieldTags:
		return m.OldTags(ctx)
	case vehicle.FieldCondition:
		return m.OldCondition(ctx)
	case vehicle.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case vehicle.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Vehicle field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VehicleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case vehicle.FieldMake:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMake(v)
		return nil
	case vehicle.FieldModel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetModel(v)
		return nil
	case vehicle.FieldYear:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetYear(v)
		return nil
	case vehicle.FieldActive:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActive(v)
		return nil
	case vehicle.FieldTags:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTags(v)
		return nil
	case vehicle.FieldCondition:
		v, ok := value.(vehicle.Condition)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCondition(v)
		return nil
	case vehicle.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case vehicle.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Vehicle field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *VehicleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *VehicleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VehicleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Vehicle numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *VehicleMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *VehicleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *VehicleMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Vehicle nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *VehicleMutation) ResetField(name string) error {
	switch name {
	case vehicle.FieldMake:
		m.ResetMake()
		return nil
	case vehicle.FieldModel:
		m.ResetModel()
		return nil
	case vehicle.FieldYear:
		m.ResetYear()
		return nil
	case vehicle.FieldActive:
		m.ResetActive()
		return nil
	case vehicle.FieldTags:
		m.ResetTags()
		return nil
	case vehicle.FieldCondition:
		m.ResetCondition()
		return nil
	case vehicle.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case vehicle.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Vehicle field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *VehicleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *VehicleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *VehicleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *VehicleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *VehicleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *VehicleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *VehicleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Vehicle unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *VehicleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Vehicle edge %s", name)
}
